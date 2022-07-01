package broker

import (
	"context"
	"github.com/RichardKnop/machinery/v1"
	"github.com/RichardKnop/machinery/v1/log"
	config "services.core-service/configs"
	"services.core-service/logger"
)

type brokerConfig struct {
	name string
}

type brokerService struct {
	server *machinery.Server
	tasks  map[string]interface{}
	worker *machinery.Worker
	*brokerConfig
}

func (b *brokerService) Name() string {
	return b.name
}

func (b *brokerService) Configure() error {
	//TODO implement me
	panic("implement me")
}

func (b *brokerService) Start() error {
	err := b.server.RegisterTasks(b.tasks)
	if err != nil {
		return err
	}

	if b.worker != nil {
		return b.worker.Launch()
	}

	return nil
}

func (b *brokerService) Stop(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func New(c config.Config, mlogger MachineryLogger) *brokerService {
	server, err := machinery.NewServer(&c.Broker)
	if err != nil {
		logger.Fatal("Cant create new broker server, err: ", err)
	}
	log.Set(mlogger)
	return &brokerService{
		brokerConfig: &brokerConfig{name: "Broker"},
		server:       server,
		tasks:        map[string]interface{}{},
	}
}

func (b *brokerService) NewWorker(consumerTag string, concurency int) *machinery.Worker {
	if b.server == nil {
		return nil
	}
	b.worker = b.server.NewWorker(consumerTag, concurency)
	return b.worker
}

func (b *brokerService) GetServer() *machinery.Server {
	return b.server
}

func (b *brokerService) SetTasks(task map[string]interface{}) {
	for name, funcc := range task {
		if funcc != nil {
			b.tasks[name] = funcc
			continue
		}
		if _, ok := b.tasks[name]; ok {
			delete(b.tasks, name)
		}
	}
}
