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

func TestGetCompetitions(t *testing.T) {
	t.Run("Success case", func(t *testing.T) {
		mockResponseBody := `{
			"competitions": [
				{"id": 1, "name": "Campeonato Brasileiro Série A", "currentSeason": {"startDate": "2024-08-01"}}
			]
		}`
		mockResponse := mocks.NewMockHTTPResponse(http.StatusOK, mockResponseBody)

		mockClient := &mocks.MockHTTPClient{MockResponse: mockResponse}
		repo := repository.NewCompetitionRepositoryImpl("http://api-test.com", "test-token", mockClient)

		competitions, err := repo.GetCompetitions()
		assert.NoError(t, err)
		assert.NotNil(t, competitions)
		assert.Len(t, *competitions, 1)
		assert.Equal(t, 1, (*competitions)[0].ID)
		assert.Equal(t, "Campeonato Brasileiro Série A", (*competitions)[0].Name)
	})

	t.Run("Failure case - HTTP error", func(t *testing.T) {
		mockClient := &mocks.MockHTTPClient{MockError: errors.New("network error")}
		repo := repository.NewCompetitionRepositoryImpl("http://api-test.com", "test-token", mockClient)

		competitions, err := repo.GetCompetitions()
		assert.Error(t, err)
		assert.Nil(t, competitions)
	})

	t.Run("Failure case - Non-200 response", func(t *testing.T) {
		mockResponse := mocks.NewMockHTTPResponse(http.StatusInternalServerError, "")
		mockClient := &mocks.MockHTTPClient{MockResponse: mockResponse}
		repo := repository.NewCompetitionRepositoryImpl("http://api-test.com", "test-token", mockClient)

		competitions, err := repo.GetCompetitions()
		assert.Error(t, err)
		assert.Nil(t, competitions)
		assert.Contains(t, err.Error(), constants.FailedToFetch)
	})

	t.Run("Failure case - Invalid JSON", func(t *testing.T) {
		mockResponse := mocks.NewMockHTTPResponse(http.StatusOK, "invalid json")
		mockClient := &mocks.MockHTTPClient{MockResponse: mockResponse}
		repo := repository.NewCompetitionRepositoryImpl("http://api-test.com", "test-token", mockClient)

		competitions, err := repo.GetCompetitions()
		assert.Error(t, err)
		assert.Nil(t, competitions)
	})
}
