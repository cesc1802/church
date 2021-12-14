package domain

type CreateUserModel struct {
	LoginID   string  `gorm:"column:login_id"`
	Password  string  `gorm:"column:password"`
	LastName  *string `gorm:"column:last_name"`
	FirstName *string `gorm:"column:first_name"`
}

func (CreateUserModel) TableName() string {
	return UserModel{}.TableName()
}
