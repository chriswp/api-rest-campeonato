package registry

import (
	"context"
	"database/sql"
	"github.com/chriswp/api-rest-campeonato/internal/domain/repository"
	"github.com/chriswp/api-rest-campeonato/internal/infra/db"
	repositoryImpl "github.com/chriswp/api-rest-campeonato/internal/infra/repository"
	"log"
)

type Registry struct {
	Database        *sql.DB
	UserRepo        repository.UserRepository
	CompetitionRepo repository.CompetitionRepository
}

func NewRegistry(ctx context.Context) (*Registry, error) {
	connectionDB, err := db.NewPostgresConnection(ctx)
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco: %v", err)
		return nil, err
	}

	return &Registry{
		Database:        connectionDB,
		UserRepo:        repositoryImpl.NewUserRepositoryImpl(connectionDB),
		CompetitionRepo: repositoryImpl.NewCompetitionRepositoryImpl(),
	}, nil
}

func (r *Registry) Close() error {
	return r.Database.Close()
}
