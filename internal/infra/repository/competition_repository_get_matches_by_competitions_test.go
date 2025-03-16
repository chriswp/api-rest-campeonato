package repository_test

import (
	"errors"
	"net/http"
	"testing"

	"github.com/chriswp/api-rest-campeonato/internal/constants"
	"github.com/chriswp/api-rest-campeonato/internal/infra/repository"
	"github.com/chriswp/api-rest-campeonato/internal/mocks"
	"github.com/stretchr/testify/assert"
)

func TestGetMatchesByCompetition(t *testing.T) {
	t.Run("Success case", func(t *testing.T) {
		mockResponseBody := `{
			"matches": [
				{
					"homeTeam": {"name": "Team A"},
					"awayTeam": {"name": "Team B"},
					"score": {"fullTime": {"homeTeam": 2, "awayTeam": 1}}
				}
			]
		}`
		mockResponse := mocks.NewMockHTTPResponse(http.StatusOK, mockResponseBody)
		mockClient := &mocks.MockHTTPClient{MockResponse: mockResponse}
		repo := repository.NewCompetitionRepositoryImpl("http://api-test.com", "test-token", mockClient)

		matchday := 5
		matches, err := repo.GetMatchesByCompetition(1, &matchday, nil)
		assert.NoError(t, err)
		assert.NotNil(t, matches)
		assert.Len(t, *matches, 1)
		assert.Equal(t, "Team A", (*matches)[0].HomeTeam)
		assert.Equal(t, "Team B", (*matches)[0].AwayTeam)
		assert.Equal(t, "2-1", (*matches)[0].Score)
	})

	t.Run("Failure case - HTTP error", func(t *testing.T) {
		mockClient := &mocks.MockHTTPClient{MockError: errors.New("network error")}
		repo := repository.NewCompetitionRepositoryImpl("http://api-test.com", "test-token", mockClient)

		matches, err := repo.GetMatchesByCompetition(1, nil, nil)
		assert.Error(t, err)
		assert.Nil(t, matches)
	})

	t.Run("Failure case - Non-200 response", func(t *testing.T) {
		mockResponse := mocks.NewMockHTTPResponse(http.StatusInternalServerError, "")
		mockClient := &mocks.MockHTTPClient{MockResponse: mockResponse}
		repo := repository.NewCompetitionRepositoryImpl("http://api-test.com", "test-token", mockClient)

		matches, err := repo.GetMatchesByCompetition(1, nil, nil)
		assert.Error(t, err)
		assert.Nil(t, matches)
		assert.Contains(t, err.Error(), constants.FailedToFetch)
	})

	t.Run("Failure case - Invalid JSON", func(t *testing.T) {
		mockResponse := mocks.NewMockHTTPResponse(http.StatusOK, "invalid json")
		mockClient := &mocks.MockHTTPClient{MockResponse: mockResponse}
		repo := repository.NewCompetitionRepositoryImpl("http://api-test.com", "test-token", mockClient)

		matches, err := repo.GetMatchesByCompetition(1, nil, nil)
		assert.Error(t, err)
		assert.Nil(t, matches)
	})

	t.Run("Success case - Filter by team", func(t *testing.T) {
		mockResponseBody := `{
			"matches": [
				{"homeTeam": {"name": "Team A"}, "awayTeam": {"name": "Team B"}, "score": {"fullTime": {"homeTeam": 2, "awayTeam": 1}}},
				{"homeTeam": {"name": "Team C"}, "awayTeam": {"name": "Team D"}, "score": {"fullTime": {"homeTeam": 3, "awayTeam": 2}}}
			]
		}`
		mockResponse := mocks.NewMockHTTPResponse(http.StatusOK, mockResponseBody)
		mockClient := &mocks.MockHTTPClient{MockResponse: mockResponse}
		repo := repository.NewCompetitionRepositoryImpl("http://api-test.com", "test-token", mockClient)

		team := "Team A"
		matches, err := repo.GetMatchesByCompetition(1, nil, &team)
		assert.NoError(t, err)
		assert.NotNil(t, matches)
		assert.Len(t, *matches, 1)
		assert.Equal(t, "Team A", (*matches)[0].HomeTeam)
	})
}
