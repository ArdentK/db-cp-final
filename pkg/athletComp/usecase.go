package athletcomp

import (
	"context"

	"github.com/ArdentK/db-cp-final/models"
)

type ACUseCase interface {
	Create(ctx context.Context, r *models.AthletComp) error
	Delete(ctx context.Context, id int) error
}
