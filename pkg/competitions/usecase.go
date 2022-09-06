package competitions

import (
	"context"

	"github.com/ArdentK/db-cp-final/models"
)

type CompUseCase interface {
	List(ctx context.Context, c *models.Competition) (*models.Competition, error)
}
