package validators

import (
	"context"
	"github.com/chriswp/api-rest-campeonato/internal/domain/entity"
)

type FootballFanValidator interface {
	Validate(ctx context.Context, fan *entity.FootballFan) error
}
