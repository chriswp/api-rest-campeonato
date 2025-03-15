package entity

import (
	"github.com/chriswp/api-rest-campeonato/pkg/entity"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       entity.ID `json:"id"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Password string    `json:"-"`
}

func NewUser(name, email, password string) (*User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		panic(err)
	}

	return &User{
		ID:       entity.NewID(),
		Name:     name,
		Email:    email,
		Password: string(hash),
	}, nil
}

func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
