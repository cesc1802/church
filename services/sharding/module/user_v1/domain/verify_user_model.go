package domain

type VerifyUserModel struct {
	LoginID  string `gorm:"column:userid"`
	Password string `gorm:"column:password"`
}

func (VerifyUserModel) TableName() string {
	return UserModel{}.TableName()
}
