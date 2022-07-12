package domain

type CreateUserModel struct {
	UserID   string `gorm:"column:login_id"`
	Password string `gorm:"column:password"`
}

func (CreateUserModel) TableName() string {
	return UserModel{}.TableName()
}
