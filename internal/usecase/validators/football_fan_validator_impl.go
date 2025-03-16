package validators

import (
	"context"
	"errors"
	"github.com/chriswp/api-rest-campeonato/internal/constants"
	"github.com/chriswp/api-rest-campeonato/internal/domain/entity"
	"github.com/chriswp/api-rest-campeonato/internal/domain/repository"
	"regexp"
)

type FootballFanValidator struct {
	footballFanRepository repository.FootballFanRepository
}

func NewFootballFanValidator(footballFanRepository repository.FootballFanRepository) *FootballFanValidator {
	return &FootballFanValidator{
		footballFanRepository: footballFanRepository,
	}
}

func (v *FootballFanValidator) Validate(ctx context.Context, fan *entity.FootballFan) error {
	if fan.Name == "" {
		return errors.New(constants.RequiredField("name"))
	}

	if fan.Email == "" {
		return errors.New(constants.RequiredField("email"))
	}

	if !isValidEmail(fan.Email) {
		return errors.New(constants.InvalidFieldError("email"))
	}

	existingFan, _ := v.footballFanRepository.FindFootballFanByEmail(ctx, fan.Email)
	if existingFan != nil {
		return errors.New(constants.IsAlreadyExists("email"))
	}

	return nil
}

func isValidEmail(email string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}
