package app

import (
	"github.com/RichardKnop/machinery/v1"
	machineryC "github.com/RichardKnop/machinery/v1/config"
)

var (
	server *machinery.Server
)

func NewServer() error {
	cnf := &machineryC.Config{
		DefaultQueue:    "chat_lmao",
		ResultsExpireIn: 3600,
		Broker:          "redis://localhost:6379",
		ResultBackend:   "redis://localhost:6379",
		Redis: &machineryC.RedisConfig{
			MaxIdle:                3,
			IdleTimeout:            240,
			ReadTimeout:            15,
			WriteTimeout:           15,
			ConnectTimeout:         15,
			NormalTasksPollPeriod:  1000,
			DelayedTasksPollPeriod: 500,
		},
	}

	se, err := machinery.NewServer(cnf)
	if err != nil {
		return err
	}
	server = se
	// Register tasks

	return nil
}

func GetServer() *machinery.Server {
	return server
}
