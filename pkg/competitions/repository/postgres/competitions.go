package postgres

import (
	"context"
	"database/sql"

	"github.com/ArdentK/db-cp-final/models"
)

type CompetitionRepo struct {
	db *sql.DB
}

func NewCompRepository(db *sql.DB) *CompetitionRepo {
	return &CompetitionRepo{
		db: db,
	}
}

func (r CompetitionRepo) List(ctx context.Context) ([]*models.Competition, error) { return nil, nil }
func (r CompetitionRepo) SortASCByData(ctx context.Context) ([]*models.Competition, error) {
	return nil, nil
}
func (r CompetitionRepo) SortDESCByData(ctx context.Context) ([]*models.Competition, error) {
	return nil, nil
}
func (r CompetitionRepo) SortASCByNumOfPart(ctx context.Context) ([]*models.Competition, error) {
	return nil, nil
}
func (r CompetitionRepo) SortDESCByNumOfPart(ctx context.Context) ([]*models.Competition, error) {
	return nil, nil
}
func (r CompetitionRepo) FindByWeaponType(ctx context.Context, weaponType string) ([]*models.Competition, error) {
	return nil, nil
}
func (r CompetitionRepo) FindByAgeCategory(ctx context.Context, ageCategory string) ([]*models.Competition, error) {
	return nil, nil
}
func (r CompetitionRepo) FindByStatus(ctx context.Context, status string) ([]*models.Competition, error) {
	return nil, nil
}
