package repository_test

import (
	"context"
	"database/sql"
	"errors"
	"github.com/google/uuid"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/chriswp/api-rest-campeonato/internal/infra/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFindUserByEmail_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	userID := uuid.New()
	email := "user@test.com"
	createdAt := time.Date(2025, 3, 11, 0, 0, 0, 0, time.UTC)
	updatedAt := time.Date(2025, 3, 11, 0, 0, 0, 0, time.UTC)

	expectedUser := sqlmock.NewRows([]string{"id", "name", "email", "password", "created_at", "updated_at"}).
		AddRow(userID.String(), "User", email, "password", createdAt, updatedAt)

	mock.ExpectQuery(`-- name: FindUserByEmail :one SELECT id, name, email, password, created_at, updated_at FROM users WHERE email = \$1`).
		WithArgs(email).
		WillReturnRows(expectedUser)

	repo := repository.NewUserRepositoryImpl(db)
	user, err := repo.FindUserByEmail(context.Background(), email)

	require.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, "User", user.Name)
	assert.Equal(t, email, user.Email)
	assert.Equal(t, userID.String(), user.ID.String())
}

func TestFindUserByEmail_NotFound(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	email := "user@test.com.br"
	mock.ExpectQuery(`-- name: FindUserByEmail :one SELECT id, name, email, password, created_at, updated_at FROM users WHERE email = \$1`).
		WithArgs(email).
		WillReturnError(sql.ErrNoRows)

	repo := repository.NewUserRepositoryImpl(db)
	user, err := repo.FindUserByEmail(context.Background(), email)

	require.NoError(t, err)
	assert.Nil(t, user)
}

func TestFindUserByEmail_DBError(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	email := "user@test.com"
	mock.ExpectQuery(`-- name: FindUserByEmail :one SELECT id, name, email, password, created_at, updated_at FROM users WHERE email = \$1`).
		WithArgs(email).
		WillReturnError(errors.New("database error"))

	repo := repository.NewUserRepositoryImpl(db)
	user, err := repo.FindUserByEmail(context.Background(), email)

	require.Error(t, err)
	assert.Nil(t, user)
}
