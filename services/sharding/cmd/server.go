package cmd

import (
	"github.com/spf13/cobra"
	"gorm.io/gorm"
	"gorm.io/sharding"
	core "services.core-service"
	config "services.core-service/configs"
	"services.core-service/httpserver"
	"services.core-service/i18n"
	"services.core-service/logger"
	"services.core-service/plugin/storage/sdkgorm"
	"shard/cmd/handlers"
	"shard/constants"
	"shard/module/user_v1/domain"
)

const (
	serviceName = "rbac-service"
	version     = "1.0.0"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "this command used to run server",
	Long:  "this command used to run server",
	RunE: func(cmd *cobra.Command, args []string) error {
		coreCfg, _ := config.LoadConfig()
		plugins := []gorm.Plugin{}
		plugins = append(plugins, sharding.Register(sharding.Config{DoubleWrite: true, ShardingKey: "userid", NumberOfShards: 64, PrimaryKeyGenerator: sharding.PKSnowflake}, domain.UserModel{}))

		i18n, _ := i18n.NewI18n(coreCfg.I18nConfig)
		serviceLogger := logger.NewLogger(logger.INFO)

		Service := core.NewAppService(
			core.WithName(serviceName),
			core.WithVersion(version),
			core.WithHttpServer(httpserver.New(coreCfg, i18n)),
			core.WithInitRunnable(sdkgorm.NewGormDB(constants.KeyMainDb, constants.KeyMainDb, &coreCfg.SQLDBConfigs[0], plugins...).SetMigration(domain.UserModel{})),
		)

		serviceLogger.Info("RBAC Name: %s , RBAC Version: %s", Service.Name(),
			Service.Version())

		Service.HttpServer().AddHandler(handlers.EndUserRoutes(Service))
		if err := Service.Run(); err != nil {
			serviceLogger.Fatal("Run RBAC service: %v", err)
		}
		return nil
	},
}
