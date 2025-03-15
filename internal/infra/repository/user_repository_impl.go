package repository

import (
	"context"
	"database/sql"
	"github.com/chriswp/api-rest-campeonato/internal/domain/entity"
	"github.com/chriswp/api-rest-campeonato/internal/domain/repository"
)

type UserRepositoryImpl struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) repository.UserRepository {
	return &UserRepositoryImpl{db: db}
}

func (r *UserRepositoryImpl) FindUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	return nil, nil
}
