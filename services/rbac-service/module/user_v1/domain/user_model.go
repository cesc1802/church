package domain

import "services.rbac-service/common"

type UserModel struct {
	common.BaseModel
	LoginID   string `gorm:"column:login_id"`
	Password  string `gorm:"column:password"`
	LastName  string `gorm:"column:last_name"`
	FirstName string `gorm:"column:first_name"`
}

func (UserModel) TableName() string {
	return "users"
}
