package auth

import (
	"context"
	"db-cp-last/models"
)

const CtxUserKey = "user"

type UseCase interface {
	SignUp(ctx context.Context, user *models.User) error
	SignIn(ctx context.Context, user *models.User) (string, error)
}
