package entity

import (
	"github.com/chriswp/api-rest-campeonato/pkg/entity"
	"time"
)

type FootballFan struct {
	ID        entity.ID `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Team      string    `json:"team"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
