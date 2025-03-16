package usecase

import (
	"context"
	"github.com/chriswp/api-rest-campeonato/internal/domain/entity"
	"github.com/chriswp/api-rest-campeonato/internal/domain/repository"
	"github.com/chriswp/api-rest-campeonato/internal/domain/validators"
)

type FootballFanUseCase struct {
	footballFanRepository repository.FootballFanRepository
	validator             validators.FootballFanValidator
}

func NewFootballFanUseCase(footballFanRepository repository.FootballFanRepository, validator validators.FootballFanValidator) *FootballFanUseCase {
	return &FootballFanUseCase{
		footballFanRepository: footballFanRepository,
		validator:             validator,
	}
}

func (u *FootballFanUseCase) CreateFootballFan(ctx context.Context, fan *entity.FootballFan) (*entity.FootballFan, error) {
	if err := u.validator.Validate(ctx, fan); err != nil {
		return nil, err
	}

	return u.footballFanRepository.SaveFootballFan(ctx, fan)
}
