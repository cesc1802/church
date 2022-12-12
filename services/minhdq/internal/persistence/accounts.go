package persistence

import (
	"context"
	"sync"
)

var (
	_accountRepo        AccountRepository
	loadAccountRepoOnce sync.Once
)

type AccountRepository interface {
	Register(loginId string, password string, firstName string, lastName string) (err error)
	Login(loginId string, password string) (jwt string, err error)
	FindById(id string) (existed bool)
	Authentization(jwt string) (err error)
}

func Account() AccountRepository {
	if _accountRepo == nil {
		panic("persistence: accountRepo Repository not initiated")
	}

	return _accountRepo
}

func LoadAccountRespositoryMem(ctx context.Context) (err error) {
	loadAccountRepoOnce.Do(func() {
		_accountRepo, err = newAccountRepoMem(ctx)
	})
	return
}
