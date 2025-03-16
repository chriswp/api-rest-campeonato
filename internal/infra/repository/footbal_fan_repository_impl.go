package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/chriswp/api-rest-campeonato/internal/domain/entity"
	"github.com/chriswp/api-rest-campeonato/internal/infra/repository/sqlc"
	uuid "github.com/chriswp/api-rest-campeonato/pkg/entity"
)

type FootballFanRepositoryImpl struct {
	queries *sqlc.Queries
}

func NewFootballFanRepositoryImpl(db *sql.DB) *FootballFanRepositoryImpl {
	return &FootballFanRepositoryImpl{
		queries: sqlc.New(db),
	}
}

func (r *FootballFanRepositoryImpl) FindFootballFanByEmail(ctx context.Context, email string) (*entity.FootballFan, error) {
	fan, err := r.queries.FindFootballFanByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &entity.FootballFan{
		ID:        fan.ID,
		Name:      fan.Name,
		Email:     fan.Email,
		CreatedAt: fan.CreatedAt,
		UpdatedAt: fan.UpdatedAt,
	}, nil
}

func (r *FootballFanRepositoryImpl) SaveFootballFan(ctx context.Context, fan *entity.FootballFan) (*entity.FootballFan, error) {
	fanSQLC, err := r.queries.CreateFootballFan(ctx, sqlc.CreateFootballFanParams{
		ID:    uuid.NewID(),
		Name:  fan.Name,
		Email: fan.Email,
		Team:  fan.Team,
	})
	if err != nil {
		return nil, err
	}
	return &entity.FootballFan{
		ID:        fanSQLC.ID,
		Name:      fanSQLC.Name,
		Email:     fanSQLC.Email,
		Team:      fanSQLC.Team,
		CreatedAt: fanSQLC.CreatedAt,
		UpdatedAt: fanSQLC.UpdatedAt,
	}, nil
}
