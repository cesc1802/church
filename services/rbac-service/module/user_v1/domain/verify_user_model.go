package domain

type VerifyUserModel struct {
	LoginID  string `gorm:"column:login_id"`
	Password string `gorm:"column:password"`
}

func (VerifyUserModel) TableName() string {
	return UserModel{}.TableName()
}
