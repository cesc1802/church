package persistence

import (
	"context"
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type AccountModel struct {
	LoginID   string  `json:"login_id" binding:"required"`
	Password  string  `json:"password" binding:"required"`
	LastName  *string `json:"last_name" binding:"omitempty"`
	FirstName *string `json:"first_name" binding:"omitempty"`
}

type AccountMem struct {
	accounts []AccountModel
}

func CreateJwtToken(LoginID string) (token string, err error) {
	expr := 24 * 2 * 60 * 60

	j := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"ID":        LoginID,
		"IssuedAt":  time.Now().Unix(),
		"Issuer":    "minhdq",
		"ExpiresAt": expr,
	})

	token, err = j.SignedString([]byte("SECRET"))

	return
}

func (a AccountMem) Register() {
	//TODO implement me
	panic("implement me")
}

func (a AccountMem) Login() (jwt string, err error) {
	//TODO implement me
	panic("implement me")
}

func (a AccountMem) FindById(id string) (err error) {
	for _, account := range a.accounts {
		if account.LoginID == id {
			return errors.New("this id already existed")
		}
	}
	return err
}

func (a AccountMem) Authentization(jwt string) (err error) {
	//TODO implement me
	panic("implement me")
}

func newAccountRepoMem(ctx context.Context) (repo *AccountMem, err error) {
	return &AccountMem{accounts: []AccountModel{}}, nil
}
