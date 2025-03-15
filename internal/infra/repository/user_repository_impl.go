package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/chriswp/api-rest-campeonato/internal/domain/entity"
	"github.com/chriswp/api-rest-campeonato/internal/domain/repository"
	"github.com/chriswp/api-rest-campeonato/internal/infra/repository/sqlc"
)

type UserRepositoryImpl struct {
	queries *sqlc.Queries
}

func NewUserRepositoryImpl(db *sql.DB) repository.UserRepository {
	return &UserRepositoryImpl{
		queries: sqlc.New(db),
	}
}

func (r *UserRepositoryImpl) FindUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	user, err := r.queries.FindUserByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &entity.User{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}
