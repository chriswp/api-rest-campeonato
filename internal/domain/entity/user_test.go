package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("Usuario teste", "email@golang.com", "102030")
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.ID)
	assert.NotEmpty(t, user.Password)
	assert.Equal(t, "Usuario teste", user.Name)
	assert.Equal(t, "email@golang.com", user.Email)
}

func TestUser_CheckPassword(t *testing.T) {
	user, err := NewUser("Usuario teste", "email@golang.com", "102030")
	assert.Nil(t, err)
	assert.True(t, user.CheckPassword("102030"))
	assert.False(t, user.CheckPassword("102031"))
	assert.NotEqual(t, "102030", user.Password)
}
