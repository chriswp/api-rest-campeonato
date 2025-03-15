package repository

import (
	"encoding/json"
	"fmt"
	config "github.com/chriswp/api-rest-campeonato/configs"
	"github.com/chriswp/api-rest-campeonato/internal/constants"
	"github.com/chriswp/api-rest-campeonato/internal/domain/entity"
	"github.com/chriswp/api-rest-campeonato/internal/domain/repository"
	"github.com/chriswp/api-rest-campeonato/internal/dto"
	"github.com/chriswp/api-rest-campeonato/internal/infra/http"
	"github.com/chriswp/api-rest-campeonato/internal/utils"
	httpResponse "net/http"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"time"
)

type CompetitionRepositoryImpl struct {
	apiURL     string
	token      string
	httpClient http.HTTPClient
}

func NewCompetitionRepositoryImpl() repository.CompetitionRepository {
	return &CompetitionRepositoryImpl{
		apiURL:     config.Envs.FootballAPIURL,
		token:      config.Envs.FootballAPIToken,
		httpClient: http.NewHTTPClient(5 * time.Second),
	}
}

func (r *CompetitionRepositoryImpl) fetchData(endpoint string, queryParams map[string]string) (*httpResponse.Response, error) {
	headers := map[string]string{
		"Content-Type": "application/json",
		"X-Auth-Token": r.token,
	}

	baseUrl := fmt.Sprintf("%s%s", r.apiURL, endpoint)
	if len(queryParams) > 0 {
		params := url.Values{}
		for key, value := range queryParams {
			params.Add(key, value)
		}
		baseUrl += "?" + params.Encode()
	}

	resp, err := r.httpClient.DoRequest("GET", baseUrl, headers, nil)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != httpResponse.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("%s: %s", constants.FailedToFetch, endpoint)
	}

	return resp, nil
}

func (r *CompetitionRepositoryImpl) GetCompetitions() (*[]entity.Competition, error) {
	resp, err := r.fetchData("/competitions", nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != httpResponse.StatusOK {
		return nil, fmt.Errorf("%s: competitions", constants.FailedToFetch)
	}

	var result struct {
		Competitions []struct {
			ID            int    `json:"id"`
			Name          string `json:"name"`
			CurrentSeason struct {
				StartDate string `json:"startDate"`
			} `json:"currentSeason"`
		} `json:"competitions"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	var competitions []entity.Competition
	for _, c := range result.Competitions {
		year, _ := utils.ExtractYear(c.CurrentSeason.StartDate)
		competitions = append(competitions, entity.Competition{
			ID:     c.ID,
			Name:   c.Name,
			Season: year,
		})
	}

	return &competitions, nil
}

func (r *CompetitionRepositoryImpl) GetMatchesByCompetition(id int, matchday *int, team *string) (*[]entity.Match, error) {
	queryParams := map[string]string{}
	if matchday != nil {
		queryParams["matchday"] = strconv.Itoa(*matchday)
	}

	resp, err := r.fetchData(fmt.Sprintf("/competitions/%d/matches", id), queryParams)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != httpResponse.StatusOK {
		return nil, fmt.Errorf("%s: competitions", constants.FailedToFetch)
	}

	var apiResponse struct {
		Matches []dto.MatchDTO `json:"matches"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
		return nil, err
	}

	var mu sync.Mutex
	var wg sync.WaitGroup
	var filteredMatches []entity.Match

	matchChan := make(chan entity.Match, len(apiResponse.Matches))
	numWorkers := 5
	jobs := make(chan dto.MatchDTO, len(apiResponse.Matches))

	for w := 0; w < numWorkers; w++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for m := range jobs {
				homeScore, awayScore := 0, 0
				if m.Score.FullTime.HomeTeam != nil {
					homeScore = *m.Score.FullTime.HomeTeam
				}
				if m.Score.FullTime.AwayTeam != nil {
					awayScore = *m.Score.FullTime.AwayTeam
				}

				match := entity.Match{
					HomeTeam: m.HomeTeam.Name,
					AwayTeam: m.AwayTeam.Name,
					Score:    fmt.Sprintf("%d-%d", homeScore, awayScore),
				}

				if team == nil || strings.Contains(strings.ToLower(match.HomeTeam), strings.ToLower(*team)) ||
					strings.Contains(strings.ToLower(match.AwayTeam), strings.ToLower(*team)) {
					matchChan <- match
				}
			}
		}()
	}
	go func() {
		for _, match := range apiResponse.Matches {
			jobs <- match
		}
		close(jobs)
	}()

	go func() {
		wg.Wait()
		close(matchChan)
	}()

	for match := range matchChan {
		mu.Lock()
		filteredMatches = append(filteredMatches, match)
		mu.Unlock()
	}

	return &filteredMatches, nil
}
