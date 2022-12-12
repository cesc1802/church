package asynq

import (
	"context"
	"github.com/hibiken/asynq"
	"net/http"
	config "services.core-service/configs"
)

type AsynqWorker struct {
	config       config.Config
	workerConfig asynq.Config
	mux          *asynq.ServeMux
	handler      map[string]interface{}
	middlerware  []func(request *http.Request, w http.Response)
}

func NewAsynqWorket(Workerconfig asynq.Config, middlerwares []func(request *http.Request, w http.Response)) *AsynqWorker {

}

func (a *AsynqWorker) Name() string {
	//TODO implement me
	panic("implement me")
}

func (a *AsynqWorker) Configure() error {
	//TODO implement me
	panic("implement me")
}

func (a *AsynqWorker) Start() error {
	//TODO implement me
	panic("implement me")
}

func (a *AsynqWorker) Stop(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (a *AsynqWorker) SetTasks(task map[string]interface{}) {

}
