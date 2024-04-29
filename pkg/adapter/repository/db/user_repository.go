package user

import (
	"context"
	"errors"

	"github.com/lutfiharidha/google-trace/pkg/shared/tracing"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
)

type UserRepository struct {
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (r *UserRepository) GetUser(ctx context.Context, data string) (string, error) {

	_, sp := tracing.CreateSpan(ctx, "UserRepository")
	defer sp.End()
	sp.SetAttributes(attribute.String("Request_Repo", data))

	sp.RecordError(errors.New("test error"))
	sp.SetStatus(codes.Error, errors.New("test error").Error())

	tracing.PrintError(errors.New("test error"))
	// Optionally add error attributes to provide additional context
	return data, nil
}
