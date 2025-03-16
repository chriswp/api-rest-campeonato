package usecase_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/chriswp/api-rest-campeonato/internal/domain/entity"
	"github.com/chriswp/api-rest-campeonato/internal/usecase"
)

type MockCompetitionRepository struct {
	mock.Mock
}

func (m *MockCompetitionRepository) GetCompetitions() (*[]entity.Competition, error) {
	args := m.Called()
	return args.Get(0).(*[]entity.Competition), args.Error(1)
}

func (m *MockCompetitionRepository) GetMatchesByCompetition(id int, matchday *int, team *string) (*[]entity.Match, error) {
	args := m.Called(id, matchday, team)
	return args.Get(0).(*[]entity.Match), args.Error(1)
}

func TestGetCompetitions_Success(t *testing.T) {
	repo := new(MockCompetitionRepository)
	usecase := usecase.NewCompetitionUsecase(repo)

	mockCompetitions := &[]entity.Competition{{ID: 1, Name: "Premier League", Season: 2023}}
	repo.On("GetCompetitions").Return(mockCompetitions, nil)

	competitions, err := usecase.GetCompetitions()

	assert.NoError(t, err)
	assert.Equal(t, mockCompetitions, competitions)
	repo.AssertExpectations(t)
}

func TestGetCompetitions_Error(t *testing.T) {
	repo := new(MockCompetitionRepository)
	usecase := usecase.NewCompetitionUsecase(repo)
	repo.On("GetCompetitions").Return((*[]entity.Competition)(nil), errors.New("database error"))

	competitions, err := usecase.GetCompetitions()

	assert.Error(t, err)
	assert.Nil(t, competitions)
	repo.AssertExpectations(t)
}

func TestGetMatches_Success(t *testing.T) {
	repo := new(MockCompetitionRepository)
	usecase := usecase.NewCompetitionUsecase(repo)

	mockMatches := &[]entity.Match{{HomeTeam: "Team A", AwayTeam: "Team B", Score: "2-1"}}
	matchday := 1
	team := "Team A"
	repo.On("GetMatchesByCompetition", 1, &matchday, &team).Return(mockMatches, nil)

	matches, err := usecase.GetMatches(1, &matchday, &team)

	assert.NoError(t, err)
	assert.Equal(t, mockMatches, matches)
	repo.AssertExpectations(t)
}

func TestGetMatches_Error(t *testing.T) {
	repo := new(MockCompetitionRepository)
	usecase := usecase.NewCompetitionUsecase(repo)
	matchday := 1
	team := "Team A"
	repo.On("GetMatchesByCompetition", 1, &matchday, &team).Return((*[]entity.Match)(nil), errors.New("network error"))

	matches, err := usecase.GetMatches(1, &matchday, &team)

	assert.Error(t, err)
	assert.Nil(t, matches)
	repo.AssertExpectations(t)
}
