package user

import (
	"context"
	"errors"

	"github.com/lutfiharidha/google-trace/pkg/domain/repository"
	"github.com/lutfiharidha/google-trace/pkg/shared/tracing"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
)

type UserUsecase struct {
	userRepo repository.UserRepository
}

func NewUserUsecase(userRepo repository.UserRepository) *UserUsecase {
	return &UserUsecase{
		userRepo: userRepo,
	}
}

func (uc *UserUsecase) GetUser(ctx context.Context, data string) (string, error) {

	ctxNew, span := tracing.CreateSpan(ctx, "UserUsecase")
	defer span.End()

	span.SetAttributes(attribute.String("Request_Usecase", data))

	uc.userRepo.GetUser(ctxNew, data)

	span.RecordError(errors.New("test error11"))
	span.SetStatus(codes.Error, errors.New("test error11").Error())
	return data, nil
}
