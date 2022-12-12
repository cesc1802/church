package sdkgorm

import (
	"context"
	"errors"
	"math"
	"strings"
	"sync"
	"time"

	"gorm.io/gorm"

	config "services.core-service/configs"
	"services.core-service/logger"
	"services.core-service/plugin/storage/sdkgorm/gormdialects"
)

type GormDBType int

const (
	GormDBTypeMySQL GormDBType = iota + 1
	GormDBTypePostgres
	GormDBTypeSQLite
	GormDBTypeMsSQL
	GormDBTypeNotSupported
)

const (
	retryCount = 10
)

type gormDB struct {
	prefix       string
	name         string
	db           *gorm.DB
	isRunning    bool
	once         *sync.Once
	cfg          *config.SQLDBConfig
	plugins      []gorm.Plugin
	migrateTable []interface{}
}

func NewGormDB(name, prefix string, cfg *config.SQLDBConfig, plugins ...gorm.Plugin) *gormDB {
	return &gormDB{
		name:      name,
		prefix:    prefix,
		isRunning: false,
		once:      new(sync.Once),
		cfg:       cfg,
		plugins:   append([]gorm.Plugin{}, plugins...),
	}
}

func (gdb *gormDB) SetMigration(model ...interface{}) *gormDB {
	gdb.migrateTable = append(gdb.migrateTable, model...)
	return gdb
}

func (gdb *gormDB) GetPrefix() string {
	return gdb.prefix
}

func (gdb *gormDB) Name() string {
	return gdb.name
}

func getDBType(dbType string) GormDBType {
	switch strings.ToLower(dbType) {
	case "mysql":
		return GormDBTypeMySQL
	case "postgres":
		return GormDBTypePostgres
	case "mssql":
		return GormDBTypeMsSQL
	case "sqlite":
		return GormDBTypeSQLite
	default:
		return GormDBTypeNotSupported
	}
}

func (gdb *gormDB) getDBConn(t GormDBType) (*gorm.DB, error) {
	switch t {
	case GormDBTypeMsSQL:
		return gormdialects.MssqlDB(gdb.cfg)
	case GormDBTypeSQLite:
		return gormdialects.SqliteDB(gdb.cfg)
	case GormDBTypePostgres:
		return gormdialects.PostgresDB(gdb.cfg)
	case GormDBTypeMySQL:
		return gormdialects.MySqlDB(gdb.cfg)
	}
	return nil, nil
}

func (gdb *gormDB) reconnectIfNeed() {
	for {
		conn, err := gdb.db.DB()
		if err = conn.Ping(); err != nil {
			_ = conn.Close()
			logger.Info("connect is gone, try to connect %s\n", gdb.name)
			gdb.isRunning = false
			gdb.once = new(sync.Once)
			_ = gdb.Get()
			return
		}
		time.Sleep(time.Second * time.Duration(5))
	}
}

func (gdb *gormDB) Get() interface{} {
	gdb.once.Do(func() {
		if !gdb.isRunning {
			if db, err := gdb.getConnWithRetry(getDBType(gdb.cfg.DBType), math.MaxInt32); err == nil {
				gdb.db = db
				gdb.isRunning = true
			} else {
				logger.Fatal("connection cannot reconnect\n", gdb.name, err)
			}
		}
	})

	if gdb.db == nil {
		return nil
	}

	// TODO: need setup logger
	// gdb.db.Logger =

	return gdb.db
}

func (gdb *gormDB) getConnWithRetry(dbType GormDBType, retry int) (*gorm.DB, error) {
	db, err := gdb.getDBConn(dbType)

	if err != nil {
		for {
			time.Sleep(time.Second * 1)
			db, err = gdb.getDBConn(dbType)

			if err == nil {
				go gdb.reconnectIfNeed()
				break
			}
		}
	} else {
		go gdb.reconnectIfNeed()
	}
	return db, err
}

func (gdb *gormDB) Configure() error {
	if gdb.isRunning {
		return nil
	}

	dbType := getDBType(gdb.cfg.DBType)
	if dbType == GormDBTypeNotSupported {
		return errors.New("gorm database type is not supported")
	}

	var err error
	gdb.db, err = gdb.getConnWithRetry(dbType, retryCount)
	if err != nil {
		return nil
	}

	err = gdb.setPlugin(context.Background())
	if err != nil {
		return nil
	}

	err = gdb.migrate()
	if err != nil {
		return nil
	}

	gdb.isRunning = true
	return nil
}

func (gdb *gormDB) Start() error {
	if err := gdb.Configure(); err != nil {
		return nil
	}
	return nil
}

func (gdb *gormDB) migrate() error {
	if gdb.db == nil {
		return errors.New("gorm database is not initialized")
	}
	for _, model := range gdb.migrateTable {
		err := gdb.db.Migrator().AutoMigrate(model)
		if err != nil {
			return err
		}
	}
	return nil
}

func (gdb *gormDB) setPlugin(ctx context.Context) error {
	if gdb.db == nil {
		return errors.New("gorm database is not initialized")
	}
	for _, plugin := range gdb.plugins {
		err := gdb.db.Use(plugin)
		if err != nil {
			return errors.New("gorm plugin is not supported")
		}
	}

	return nil
}

func (gdb *gormDB) Stop(ctx context.Context) error {
	if gdb.db != nil {
		if conn, err := gdb.db.DB(); err != nil {
			conn.Close()
		}
	}
	return nil
}
