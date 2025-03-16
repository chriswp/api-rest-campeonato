package registry

import (
	"database/sql"
	"github.com/chriswp/api-rest-campeonato/internal/domain/repository"
)

type Registry struct {
	Database        *sql.DB
	UserRepo        repository.UserRepository
	CompetitionRepo repository.CompetitionRepository
}

func (r *Registry) Close() error {
	return r.Database.Close()
}
