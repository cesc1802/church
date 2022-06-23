package app

import (
	"github.com/spf13/cobra"
	core "services.core-service"
	config "services.core-service/configs"
	"services.core-service/httpserver"
	"services.core-service/i18n"
	"services.core-service/logger"
)

var ChatCMD = &cobra.Command{
	Use:   "frontend",
	Short: "this command used to run frontend server",
	Long:  "this command used to run server",
	RunE: func(cmd *cobra.Command, args []string) error {
		coreCfg := config.Config{
			Env:          "",
			RedisConfig:  config.RedisConfig{},
			SQLDBConfigs: nil,
			NoSQLConfigs: nil,
			ServerConfig: config.ServerConfig{
				Host: "localhost",
				Port: "5000",
			},
			LoginServerConfig:    config.ServerConfig{},
			RegisterServerConfig: config.ServerConfig{},
			HttpClientConfig:     config.HttpClientConfig{},
			I18nConfig:           config.I18nConfig{},
			CORSConfig:           config.CORSConfig{},
			LogConfig:            config.LogConfig{},
		}
		i18n, _ := i18n.NewI18n(coreCfg.I18nConfig)
		serviceLogger := logger.NewLogger(logger.INFO)

		service := core.NewAppService(
			core.WithName("frontend"),
			core.WithVersion("v0.1"),
			core.WithHttpServer(httpserver.New(coreCfg, i18n)),
		)

		serviceLogger.Info("service Name: %s , service Version: %s", service.Name(),
			service.Version())

		service.HttpServer().AddHandler(ChatRouter(service, service.HttpServer().GetMelody()))

		if err := service.Run(); err != nil {
			serviceLogger.Fatal("Run RBAC service: %v", err)
		}
		return nil
	},
}

func GetChatCommand() *cobra.Command {
	return ChatCMD
}
