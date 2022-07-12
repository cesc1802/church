package domain

import "shard/common"

type UserModel struct {
	common.BaseModel
	ID       int64  `gorm:"primarykey"`
	UserID   string `gorm:"column:login_id"`
	Password string `gorm:"column:password"`
}

func (UserModel) TableName() string {
	return "users"
}
