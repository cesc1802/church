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
	Register()
	Login() (jwt string, err error)
	FindById(id string) (err error)
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
