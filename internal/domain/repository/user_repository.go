package repository

import (
	"context"
	"github.com/chriswp/api-rest-campeonato/internal/domain/entity"
)

type UserRepository interface {
	FindUserByEmail(ctx context.Context, email string) (*entity.User, error)
}
