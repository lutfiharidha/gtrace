package user

import (
	"context"
)

// BranchCashInOutInputPort :
type UserInterface interface {
	GetUser(ctx context.Context, data string) (string, error)
}
