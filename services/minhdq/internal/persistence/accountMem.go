package persistence

import (
	"context"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"sync"
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
	mu       *sync.RWMutex
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

func (a *AccountMem) Register(loginId string, password string, firstName string, lastName string) (err error) {
	a.mu.Lock()
	defer a.mu.Unlock()
	if loginId == "" {
		return errors.New("loginID can't be empty")
	}

	if a.FindById(loginId) {
		return errors.New("loginID already existed")
	}

	a.accounts = append(a.accounts, AccountModel{
		LoginID:   loginId,
		Password:  password,
		LastName:  &firstName,
		FirstName: &lastName,
	})

	return nil
}

func (a *AccountMem) Login(loginId string, password string) (jwt string, err error) {
	a.mu.Lock()
	defer a.mu.Unlock()
	fmt.Println(a.accounts)
	for _, account := range a.accounts {
		fmt.Println(account)
		if account.LoginID == loginId && account.Password == password {
			fmt.Printf("hello")
			return CreateJwtToken(loginId)
		}
	}
	return "", errors.New("login invalid")
}

func (a *AccountMem) FindById(id string) (existed bool) {
	for _, account := range a.accounts {
		if account.LoginID == id {
			return true
		}
	}
	return false
}

func (a *AccountMem) Authentization(jwtString string) (err error) {
	token, err := jwt.Parse(jwtString, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte("SECRET"), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if a.FindById(claims["ID"].(string)) {
			if claims["Issuer"].(string) == "minhdq" {
				return nil
			}
		}
	}

	return errors.New("token is invalid")
}

func newAccountRepoMem(ctx context.Context) (repo *AccountMem, err error) {
	return &AccountMem{accounts: []AccountModel{}, mu: &sync.RWMutex{}}, nil
}
