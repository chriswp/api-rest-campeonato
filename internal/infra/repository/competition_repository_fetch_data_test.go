package repository_test

import (
	"errors"
	"net/http"
	"testing"

	"github.com/chriswp/api-rest-campeonato/internal/infra/repository"
	"github.com/chriswp/api-rest-campeonato/internal/mocks"
	"github.com/stretchr/testify/assert"
)

func TestFetchData_Success(t *testing.T) {
	mockClient := &mocks.MockHTTPClient{
		MockResponse: mocks.NewMockHTTPResponse(http.StatusOK, `{"message": "success"}`),
	}
	repo := repository.NewCompetitionRepositoryImpl("https://api-test.com", "token", mockClient)

	resp, err := repo.FetchData("/test", nil)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func TestFetchData_Failure(t *testing.T) {
	mockClient := &mocks.MockHTTPClient{
		MockError: errors.New("request failed"),
	}
	repo := repository.NewCompetitionRepositoryImpl("https://api-test.com", "token", mockClient)

	resp, err := repo.FetchData("/test", nil)
	assert.Error(t, err)
	assert.Nil(t, resp)
}
