package repository

import (
	"github.com/chriswp/api-rest-campeonato/internal/domain/entity"
)

type CompetitionRepository interface {
	GetCompetitions() (*[]entity.Competition, error)
	GetMatchesByCompetition(id int, matchday *int, team *string) (*[]entity.Match, error)
}
