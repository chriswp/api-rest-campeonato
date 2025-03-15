//go:build wireinject
// +build wireinject

package infra

import (
	"github.com/chriswp/api-rest-campeonato/internal/infra/handler"
	"github.com/chriswp/api-rest-campeonato/internal/infra/repository"
	"github.com/chriswp/api-rest-campeonato/internal/usecase"
)
import "github.com/google/wire"

var CompetitionUseCaseSet = wire.NewSet(
	repository.NewCompetitionRepositoryImpl,
	usecase.NewCompetitionUsecase,
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
