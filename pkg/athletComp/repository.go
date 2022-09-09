package athletcomp

import (
	"context"

	"github.com/ArdentK/db-cp-final/models"
)

type ACRepo interface {
	AddRow(ctx context.Context, r *models.AthletComp) error
	DelRow(ctx context.Context, r *models.AthletComp) error
}
