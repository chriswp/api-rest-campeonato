package usecase

import (
	"context"
	"errors"
	"github.com/chriswp/api-rest-campeonato/internal/domain/repository"
	"github.com/chriswp/api-rest-campeonato/internal/domain/usecase"
	"golang.org/x/crypto/bcrypt"
)

type AuthUseCaseImpl struct {
	userRepo repository.UserRepository
}

func NewAuthUseCase(userRepo repository.UserRepository) usecase.AuthUseCase {
	return &AuthUseCaseImpl{userRepo: userRepo}
}

func (u *AuthUseCaseImpl) Authenticate(ctx context.Context, email, password string) (map[string]interface{}, error) {
	user, err := u.userRepo.FindUserByEmail(ctx, email)
	if err != nil {
		return nil, errors.New("usuário não encontrado")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New("senha inválida")
	}

	return map[string]interface{}{
		"user_id": user.ID,
	}, nil
}
