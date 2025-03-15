//go:build wireinject
// +build wireinject

package infra

import (
	"database/sql"
	"github.com/chriswp/api-rest-campeonato/internal/infra/handler"
	"github.com/chriswp/api-rest-campeonato/internal/infra/repository"
	"github.com/chriswp/api-rest-campeonato/internal/usecase"
)
import "github.com/google/wire"

var CompetitionUseCaseSet = wire.NewSet(
	repository.NewCompetitionRepositoryImpl,
	usecase.NewCompetitionUsecase,
)

var UserUseCaseSet = wire.NewSet(
	repository.NewUserRepositoryImpl,
	usecase.NewUserUseCase,
)

func NewCompetitionUseCase() *usecase.CompetitionUseCase {
	wire.Build(
		repository.NewCompetitionRepositoryImpl,
		usecase.NewCompetitionUsecase,
	)
	return &usecase.CompetitionUseCase{}
}

func NewCompetitionHandler() *handler.CompetitionHandler {
	wire.Build(
		CompetitionUseCaseSet,
		handler.NewCompetitionHandler,
	)
	return &handler.CompetitionHandler{}
}

func NewUserHandler(db *sql.DB) *handler.UserHandler {
	wire.Build(
		UserUseCaseSet,
		handler.NewUserHandler,
	)
	return &handler.UserHandler{}
}
