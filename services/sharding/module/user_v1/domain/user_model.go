package domain

import "shard/common"

type UserModel struct {
	common.BaseModel
	ID       int64  `gorm:"primarykey"`
	UserID   string `gorm:"column:userid"`
	Password string `gorm:"column:password"`
}

func (UserModel) TableName() string {
	return "users"
}
