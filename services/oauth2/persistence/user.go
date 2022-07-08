package persistence

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"oauth/config"
	"oauth/model"
	"sync"
)

var (
	_userRepo        UserRepository
	loadUserRepoOnce sync.Once
)

type UserRepository interface {
	FindAll(ctx context.Context) (users []*model.User, err error)
	FindOneByID(ctx context.Context, id string) (user *model.User, err error)
	CheckUser(ctx context.Context, username string, password string) (err error)
	CreateUser(ctx context.Context, username string, password string) (err error)
}

func User() UserRepository {
	if _userRepo == nil {
		panic("persistence: _userRepo Repository not initiated")
	}

	return _userRepo
}

func LoadUserRepoRespositoryMem(ctx context.Context) (err error) {
	cfg, err := pgxpool.ParseConfig(config.Get().PostgresURL)
	if err != nil {
		return err
	}
	cfg.MaxConns = 10
	conn, err := pgxpool.ConnectConfig(ctx, cfg)
	if err != nil {
		log.Fatalln(err)
	}
	loadUserRepoOnce.Do(func() {
		_userRepo, err = newUserRepoPSQL(ctx, conn)
	})
	return
}
