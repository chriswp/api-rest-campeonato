package usecase_test

import (
	"context"
	"errors"
	"github.com/chriswp/api-rest-campeonato/internal/usecase"
	"testing"

	"github.com/chriswp/api-rest-campeonato/internal/domain/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockFootballFanRepository struct {
	mock.Mock
}

func (m *MockFootballFanRepository) SaveFootballFan(ctx context.Context, fan *entity.FootballFan) (*entity.FootballFan, error) {
	args := m.Called(ctx, fan)
	return args.Get(0).(*entity.FootballFan), args.Error(1)
}

func (m *MockFootballFanRepository) FindFootballFanByEmail(ctx context.Context, email string) (*entity.FootballFan, error) {
	args := m.Called(ctx, email)
	return args.Get(0).(*entity.FootballFan), args.Error(1)
}

type MockFootballFanValidator struct {
	mock.Mock
}

func (m *MockFootballFanValidator) Validate(ctx context.Context, fan *entity.FootballFan) error {
	args := m.Called(ctx, fan)
	return args.Error(0)
}

func TestCreateFootballFan(t *testing.T) {
	repo := new(MockFootballFanRepository)
	validator := new(MockFootballFanValidator)

	useCase := usecase.NewFootballFanUseCase(repo, validator)

	fan := &entity.FootballFan{
		Name:  "Torcedor 1",
		Email: "torcedor@example.com",
		Team:  "CR Flamengo",
	}

	validator.On("Validate", mock.Anything, fan).Return(errors.New("invalid fan"))
	createdFan, err := useCase.CreateFootballFan(context.Background(), fan)
	assert.Nil(t, createdFan)
	assert.EqualError(t, err, "invalid fan")
	validator.AssertExpectations(t)

	validator.ExpectedCalls = nil

	validator.On("Validate", mock.Anything, fan).Return(nil)
	repo.On("SaveFootballFan", mock.Anything, fan).Return(fan, nil)
	createdFan, err = useCase.CreateFootballFan(context.Background(), fan)
	assert.NotNil(t, createdFan)
	assert.NoError(t, err)
	assert.Equal(t, fan.Name, createdFan.Name)
	validator.AssertExpectations(t)
	repo.AssertExpectations(t)
}
