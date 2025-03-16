//go:build wireinject
// +build wireinject

package infra

import (
	"database/sql"
	domainRepository "github.com/chriswp/api-rest-campeonato/internal/domain/repository"
	domainUseCase "github.com/chriswp/api-rest-campeonato/internal/domain/usecase"
	"github.com/chriswp/api-rest-campeonato/internal/infra/handler"
	"github.com/chriswp/api-rest-campeonato/internal/infra/http"
	"github.com/chriswp/api-rest-campeonato/internal/infra/registry"
	"github.com/chriswp/api-rest-campeonato/internal/infra/repository"
	"github.com/chriswp/api-rest-campeonato/internal/usecase"
	"github.com/chriswp/api-rest-campeonato/internal/usecase/validators"
	"os"
	"time"
)
import "github.com/google/wire"

func ProvideRegistry(db *sql.DB) *registry.Registry {
	return &registry.Registry{
		Database:        db,
		UserRepo:        repository.NewUserRepositoryImpl(db),
		CompetitionRepo: ProvideCompetitionRepository(),
	}
}

func ProvideCompetitionRepository() domainRepository.CompetitionRepository {
	apiURL := os.Getenv("FOOTBALL_API_URL")
	token := os.Getenv("FOOTBALL_API_TOKEN")
	httpClient := http.NewHTTPClient(5 * time.Second)

	return repository.NewCompetitionRepositoryImpl(apiURL, token, httpClient)
}

var RegistrySet = wire.NewSet(
	ProvideRegistry,
	UserUseCaseSet,
	CompetitionUseCaseSet,
)

func NewRegistry(db *sql.DB) (*registry.Registry, error) {
	wire.Build(RegistrySet)
	return nil, nil
}

var AuthUseCaseSet = wire.NewSet(
	usecase.NewAuthUseCase,
	wire.Bind(new(domainUseCase.AuthUseCase), new(*usecase.AuthUseCaseImpl)),
)

var CompetitionUseCaseSet = wire.NewSet(
	ProvideCompetitionRepository,
	usecase.NewCompetitionUsecase,
)

var UserUseCaseSet = wire.NewSet(
	repository.NewUserRepositoryImpl,
	usecase.NewUserUseCase,
)

var FootballFanUseCaseSet = wire.NewSet(
	ProvideFootballFanRepository,
	ProvideFootballFanValidator,
	ProvideFootballFanUseCase,
)

func ProvideFootballFanRepository(db *sql.DB) domainRepository.FootballFanRepository {
	return repository.NewFootballFanRepositoryImpl(db)
}

func ProvideFootballFanValidator(footballFanRepository domainRepository.FootballFanRepository) *validators.FootballFanValidator {
	return validators.NewFootballFanValidator(footballFanRepository)
}

func ProvideFootballFanUseCase(footballFanRepository domainRepository.FootballFanRepository, validator *validators.FootballFanValidator) *usecase.FootballFanUseCase {
	return usecase.NewFootballFanUseCase(footballFanRepository, validator)
}

func NewCompetitionUseCase() *usecase.CompetitionUseCase {
	wire.Build(CompetitionUseCaseSet)
	return nil
}

func NewFootballFanUseCase(db *sql.DB) *usecase.FootballFanUseCase {
	wire.Build(FootballFanUseCaseSet)
	return nil
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

func NewFootballFanHandler(db *sql.DB) *handler.FootballFanHandler {
	wire.Build(
		FootballFanUseCaseSet,
		handler.NewFootballFanHandler,
	)
	return &handler.FootballFanHandler{}
}
