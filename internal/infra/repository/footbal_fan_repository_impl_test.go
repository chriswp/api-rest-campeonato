package repository_test

import (
	"context"
	"database/sql"
	"errors"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/chriswp/api-rest-campeonato/internal/domain/entity"
	"github.com/chriswp/api-rest-campeonato/internal/infra/repository"
	"github.com/stretchr/testify/assert"
)

func TestFootballFanRepository_FindFootballFanByEmail(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	repo := repository.NewFootballFanRepositoryImpl(db)

	userID := uuid.New()
	email := "user@test.com"
	team := "flamengo"
	createdAt := time.Date(2025, 3, 11, 0, 0, 0, 0, time.UTC)
	updatedAt := time.Date(2025, 3, 11, 0, 0, 0, 0, time.UTC)
	mockRow := sqlmock.NewRows([]string{"id", "name", "team", "email", "created_at", "updated_at"}).
		AddRow(userID.String(), "Test Fan", email, team, createdAt, updatedAt)

	mock.ExpectQuery(`-- name: FindFootballFanByEmail :one SELECT id, name, email, team, created_at, updated_at FROM football_fans WHERE email = \$1`).
		WithArgs(email).
		WillReturnRows(mockRow)

	fan, err := repo.FindFootballFanByEmail(context.Background(), email)
	assert.NoError(t, err)
	assert.NotNil(t, fan)
	assert.Equal(t, userID, fan.ID)
	assert.Equal(t, "Test Fan", fan.Name)

	mock.ExpectQuery(`-- name: FindFootballFanByEmail :one SELECT id, name, email, team, created_at, updated_at FROM football_fans WHERE email = \$1`).
		WithArgs("notfound@example.com").
		WillReturnError(sql.ErrNoRows)

	fan, err = repo.FindFootballFanByEmail(context.Background(), "notfound@example.com")
	assert.NoError(t, err)
	assert.Nil(t, fan)

	mock.ExpectQuery(`-- name: FindFootballFanByEmail :one SELECT id, name, email, team, created_at, updated_at FROM football_fans WHERE email = \$1`).
		WithArgs("error@example.com").
		WillReturnError(errors.New("some error"))

	fan, err = repo.FindFootballFanByEmail(context.Background(), "error@example.com")
	assert.Error(t, err)
	assert.Nil(t, fan)
}

func TestFootballFanRepository_SaveFootballFan(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock database: %v", err)
	}
	defer db.Close()

	repo := repository.NewFootballFanRepositoryImpl(db)
	id := uuid.New()
	mock.ExpectQuery(`-- name: CreateFootballFan :one INSERT INTO football_fans \(id, name, email, team\) VALUES \(\$1, \$2, \$3, \$4\) RETURNING id, name, email, team, created_at, updated_at`).
		WithArgs(sqlmock.AnyArg(), "New Fan", "newfan@example.com", "Team A").
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "email", "team", "created_at", "updated_at"}).
			AddRow(id, "New Fan", "newfan@example.com", "Team A", time.Now(), time.Now()))

	fan := &entity.FootballFan{
		ID:    id,
		Name:  "New Fan",
		Email: "newfan@example.com",
		Team:  "Team A",
	}

	savedFan, err := repo.SaveFootballFan(context.Background(), fan)
	assert.NoError(t, err)
	assert.NotNil(t, savedFan)
	assert.Equal(t, id, savedFan.ID)
	assert.Equal(t, "New Fan", savedFan.Name)

	mock.ExpectExec(`-- name: CreateFootballFan :one INSERT INTO football_fans \(id, name, email, team\) VALUES \(\$1, \$2, \$3, \$4\) RETURNING id, name, email, team, created_at, updated_at`).
		WithArgs(id, "New Fan", "newfan@example.com", "Team A").
		WillReturnError(errors.New("some error"))

	savedFan, err = repo.SaveFootballFan(context.Background(), fan)
	assert.Error(t, err)
	assert.Nil(t, savedFan)
}
