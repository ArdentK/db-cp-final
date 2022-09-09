package athletcomp

import (
	"context"

	"github.com/ArdentK/db-cp-final/models"
)

type ACUseCase interface {
	Create(ctx context.Context, user *models.User, idComp int) error
	Delete(ctx context.Context, user *models.User, idComp int) error
}
