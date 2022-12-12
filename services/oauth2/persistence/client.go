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
	_clientRepo        ClientRepository
	loadClientRepoOnce sync.Once
)

type ClientRepository interface {
	FindAll(ctx context.Context) (clients []*model.Client, err error)
	FindOneByID(ctx context.Context, id string) (client *model.Client, err error)
	CreateByID(ctx context.Context, id string, redirectURIs []string, responeType []string, scopes []string, grantTypes []string, public bool, audiens []string) (secret string, err error)
}

func Client() ClientRepository {
	if _clientRepo == nil {
		panic("persistence: _clientRepo Repository not initiated")
	}

	return _clientRepo
}

func LoadClientRepoRespositoryMem(ctx context.Context) (err error) {
	cfg, err := pgxpool.ParseConfig(config.Get().PostgresURL)
	if err != nil {
		return err
	}
	cfg.MaxConns = 10
	conn, err := pgxpool.ConnectConfig(ctx, cfg)
	if err != nil {
		log.Fatalln(err)
	}
	loadClientRepoOnce.Do(func() {
		_clientRepo, err = newClientGroupRepoPSQL(ctx, conn)
	})
	return
}
