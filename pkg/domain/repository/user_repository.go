package repository

import "context"

type UserRepository interface {
	GetUser(ctx context.Context, data string) (string, error)
}
