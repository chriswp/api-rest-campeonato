package usecase

import (
	"github.com/chriswp/api-rest-campeonato/internal/domain/entity"
	"github.com/chriswp/api-rest-campeonato/internal/domain/repository"
)

type CompetitionUseCase struct {
	competitionRepository repository.CompetitionRepository
}

func NewCompetitionUsecase(competitionRepository repository.CompetitionRepository) *CompetitionUseCase {
	return &CompetitionUseCase{
		competitionRepository: competitionRepository,
	}
}

func (c *CompetitionUseCase) GetCompetitions() (*[]entity.Competition, error) {
	return c.competitionRepository.GetCompetitions()
}

func (c *CompetitionUseCase) GetMatches(id int, matchday *int, team *string) (*[]entity.Match, error) {
	return c.competitionRepository.GetMatchesByCompetition(id, matchday, team)
}
