package usecase

import (
	"context"

	"github.com/ArdentK/db-cp-final/models"
	"github.com/ArdentK/db-cp-final/pkg/competitions"
)

type CompetitionUseCase struct {
	repo competitions.CompRepo
}

func NewCompUseCase(r competitions.CompRepo) *CompetitionUseCase {
	return &CompetitionUseCase{
		repo: r,
	}
}

func (u CompetitionUseCase) List(ctx context.Context) ([]*models.Competition, error) { return nil, nil }
func (u CompetitionUseCase) SortASCByData(ctx context.Context) ([]*models.Competition, error) {
	return nil, nil
}
func (u CompetitionUseCase) SortDESCByData(ctx context.Context) ([]*models.Competition, error) {
	return nil, nil
}
func (u CompetitionUseCase) SortASCByNumOfPart(ctx context.Context) ([]*models.Competition, error) {
	return nil, nil
}
func (u CompetitionUseCase) SortDESCByNumOfPart(ctx context.Context) ([]*models.Competition, error) {
	return nil, nil
}
func (u CompetitionUseCase) FindByWeaponType(ctx context.Context, weaponType string) ([]*models.Competition, error) {
	return nil, nil
}
func (u CompetitionUseCase) FindByAgeCategory(ctx context.Context, ageCategory string) ([]*models.Competition, error) {
	return nil, nil
}
func (u CompetitionUseCase) FindByStatus(ctx context.Context, status string) ([]*models.Competition, error) {
	return nil, nil
}
