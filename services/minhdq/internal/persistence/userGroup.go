package persistence

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"minhdq/internal/config"
	"minhdq/internal/model"
	"sync"
)

var (
	_userGroupRepo        UserGroupRepository
	loadUserGroupRepoOnce sync.Once
)

type UserGroupRepository interface {
	FindAll(ctx context.Context) (usergroup []*model.UserGroup, err error)
	Save(ctx context.Context, userGroup model.UserGroup) (err error)
	Update(ctx context.Context, userGroup model.UserGroup) (record model.UserGroup, err error)
	Delete(ctx context.Context, userID int, groupID int) (err error)
}

func UserGroup() UserGroupRepository {
	if _userGroupRepo == nil {
		panic("persistence: userGroup Repository not initiated")
	}

	return _userGroupRepo
}

func LoadUserGrouprRespositorySql(ctx context.Context) (err error) {
	cfg, err := pgxpool.ParseConfig(config.Get().PostgresURL)

	if err != nil {
		return err
	}
	cfg.MaxConns = 10
	conn, err := pgxpool.ConnectConfig(ctx, cfg)
	if err != nil {
		log.Fatalln(err)
	}
	loadUserGroupRepoOnce.Do(func() {
		_userGroupRepo, err = newUserGroupRepoPSQL(ctx, conn)
	})
	return
}
