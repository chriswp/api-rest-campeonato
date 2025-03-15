package usecase

import "context"

type AuthUseCase interface {
	Authenticate(ctx context.Context, username, password string) (map[string]interface{}, error)
}
