package usecase

import (
	"context"
	"github.com/chriswp/api-rest-campeonato/internal/domain/entity"
	"github.com/chriswp/api-rest-campeonato/internal/domain/repository"
)

type UserUseCase struct {
	userRepository repository.UserRepository
}

func NewUserUseCase(userRepository repository.UserRepository) *UserUseCase {
	return &UserUseCase{
		userRepository: userRepository,
	}
}

func (u *UserUseCase) FindUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	user, err := u.userRepository.FindUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	return user, nil
}
