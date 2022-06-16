package cmd

import (
	"context"

	migrate "github.com/rubenv/sql-migrate"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
	config "services.core-service/configs"
	"services.core-service/logger"
	"services.core-service/plugin/storage/sdkgorm"
	"services.rbac-service/constants"
)

func mustErr(err error) {
	if err != nil {
		logger.Fatal("ERR: ", err)
	}
}

var migrationCmd = &cobra.Command{
	Use:   "migration",
	Short: "This command used to migrate database",
	Long:  "This command used to migrate database",
}

func init() {
	migrationCmd.AddCommand(migrateUp)
	migrationCmd.AddCommand(migrateDown)
}

var migrateUp = &cobra.Command{
	Use:   "up",
	Short: "Run migration up",
	Long:  "Run migration up",
	RunE: func(cmd *cobra.Command, args []string) error {
		migrations := &migrate.FileMigrationSource{
			Dir: "./dbmigration",
		}

		baseCfg, err := config.LoadConfig()
		mustErr(err)
		baseDB := sdkgorm.NewGormDB(constants.KeyMainDb, constants.KeyMainDb, &baseCfg.SQLDBConfigs[0])

		baseDB.Start()
		dbGorm := baseDB.Get().(*gorm.DB)
		conn, err := dbGorm.DB()
		mustErr(err)

		applyNum, err := migrate.Exec(conn, "postgres", migrations, migrate.Up)
		mustErr(err)

		logger.Info("Apply number migration %d", applyNum)

		baseDB.Stop(context.Background())
		return nil
	},
}

var migrateDown = &cobra.Command{
	Use:   "down",
	Short: "down migration down",
	Long:  "down migration down",
	RunE: func(cmd *cobra.Command, args []string) error {
		migrations := &migrate.FileMigrationSource{
			Dir: "./dbmigration",
		}

		baseCfg, err := config.LoadConfig()
		mustErr(err)
		baseDB := sdkgorm.NewGormDB(constants.KeyMainDb, constants.KeyMainDb, &baseCfg.SQLDBConfigs[0])

		baseDB.Start()
		dbGorm := baseDB.Get().(*gorm.DB)
		conn, err := dbGorm.DB()
		mustErr(err)

		applyNum, err := migrate.Exec(conn, dialect, migrations, migrate.Down)
		mustErr(err)

		logger.Info("Apply number migration %d", applyNum)

		baseDB.Stop(context.Background())
		return nil
	},
}
