package competitions

import (
	"context"

	"github.com/ArdentK/db-cp-final/models"
)

type CompRepo interface {
	Create(ctx context.Context, c *models.Competition) error
	Update(ctx context.Context, c *models.Competition) error
	Delete(ctx context.Context, id int) error
	FindByID(ctx context.Context, id int) (*models.Competition, error)
	List(ctx context.Context) ([]*models.Competition, error)
	SortASCByData(ctx context.Context) ([]*models.Competition, error)
	SortDESCByData(ctx context.Context) ([]*models.Competition, error)
	SortASCByNumOfPart(ctx context.Context) ([]*models.Competition, error)
	SortDESCByNumOfPart(ctx context.Context) ([]*models.Competition, error)
	FindByWeaponType(ctx context.Context, weaponType string) ([]*models.Competition, error)
	FindByAgeCategory(ctx context.Context, ageCategory string) ([]*models.Competition, error)
	FindByStatus(ctx context.Context, status string) ([]*models.Competition, error)
}
