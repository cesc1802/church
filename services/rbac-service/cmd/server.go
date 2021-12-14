package cmd

import (
	"github.com/spf13/cobra"
	core "services.core-service"
	config "services.core-service/configs"
	"services.core-service/httpserver"
	"services.core-service/i18n"
	"services.core-service/logger"
	"services.core-service/plugin/storage/sdkgorm"
	"services.rbac-service/cmd/handlers"
	"services.rbac-service/constants"
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
		i18n, _ := i18n.NewI18n(coreCfg.I18nConfig)
		serviceLogger := logger.NewLogger(logger.INFO)

		rbacService := core.NewAppService(
			core.WithName(serviceName),
			core.WithVersion(version),
			core.WithHttpServer(httpserver.New(coreCfg, i18n)),
			core.WithInitRunnable(sdkgorm.NewGormDB(constants.KeyMainDb, constants.KeyMainDb, &coreCfg.SQLDBConfigs[0])),
		)

		serviceLogger.Info("RBAC Name: %s , RBAC Version: %s", rbacService.Name(),
			rbacService.Version())

		rbacService.HttpServer().AddHandler(handlers.EndUserRoutes(rbacService))

		if err := rbacService.Run(); err != nil {
			serviceLogger.Fatal("Run RBAC service: %v", err)
		}
		return nil
	},
}
