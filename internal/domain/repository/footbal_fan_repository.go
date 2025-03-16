package repository

import (
	"context"
	"github.com/chriswp/api-rest-campeonato/internal/domain/entity"
)

type FootballFanRepository interface {
	FindFootballFanByEmail(ctx context.Context, email string) (*entity.FootballFan, error)
	SaveFootballFan(ctx context.Context, fan *entity.FootballFan) (*entity.FootballFan, error)
}
