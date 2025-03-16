package usecase_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"

	"github.com/chriswp/api-rest-campeonato/internal/domain/entity"
	"github.com/chriswp/api-rest-campeonato/internal/usecase"
	uuid "github.com/chriswp/api-rest-campeonato/pkg/entity"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) FindUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	args := m.Called(ctx, email)
	if args.Get(0) != nil {
		return args.Get(0).(*entity.User), args.Error(1)
	}
	return nil, args.Error(1)
}

func TestAuthenticate_Success(t *testing.T) {
	mockRepo := new(MockUserRepository)
	usecase := usecase.NewAuthUseCase(mockRepo)
	ctx := context.Background()
	email := "test@example.com"
	password := "password123"
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	user := &entity.User{ID: uuid.NewID(), Email: email, Password: string(hashedPassword)}

	mockRepo.On("FindUserByEmail", ctx, email).Return(user, nil)

	result, err := usecase.Authenticate(ctx, email, password)

	assert.NoError(t, err)
	assert.Equal(t, user.ID, result["user_id"])
	mockRepo.AssertExpectations(t)
}

func TestAuthenticate_UserNotFound(t *testing.T) {
	mockRepo := new(MockUserRepository)
	usecase := usecase.NewAuthUseCase(mockRepo)
	ctx := context.Background()
	email := "notfound@example.com"
	password := "password123"

	mockRepo.On("FindUserByEmail", ctx, email).Return(nil, errors.New("usuário não encontrado"))

	result, err := usecase.Authenticate(ctx, email, password)

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, "usuário não encontrado", err.Error())
	mockRepo.AssertExpectations(t)
}

func TestAuthenticate_InvalidPassword(t *testing.T) {
	mockRepo := new(MockUserRepository)
	usecase := usecase.NewAuthUseCase(mockRepo)
	ctx := context.Background()
	email := "test@example.com"
	password := "wrongpassword"
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
	user := &entity.User{ID: uuid.NewID(), Email: email, Password: string(hashedPassword)}

	mockRepo.On("FindUserByEmail", ctx, email).Return(user, nil)

	result, err := usecase.Authenticate(ctx, email, password)

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, "senha inválida", err.Error())
	mockRepo.AssertExpectations(t)
}
