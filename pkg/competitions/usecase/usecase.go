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

func (u *CompetitionUseCase) List(ctx context.Context) ([]*models.Competition, error) {
	return u.repo.List(ctx)
}
func (u *CompetitionUseCase) SortASCByData(ctx context.Context) ([]*models.Competition, error) {
	return u.repo.SortASCByData(ctx)
}
func (u *CompetitionUseCase) SortDESCByData(ctx context.Context) ([]*models.Competition, error) {
	return u.repo.SortDESCByData(ctx)
}
func (u *CompetitionUseCase) SortASCByNumOfPart(ctx context.Context) ([]*models.Competition, error) {
	return u.repo.SortASCByNumOfPart(ctx)
}
func (u *CompetitionUseCase) SortDESCByNumOfPart(ctx context.Context) ([]*models.Competition, error) {
	return u.repo.SortDESCByNumOfPart(ctx)
}
func (u *CompetitionUseCase) FindByWeaponType(ctx context.Context, weaponType string) ([]*models.Competition, error) {
	return u.repo.FindByWeaponType(ctx, weaponType)
}
func (u *CompetitionUseCase) FindByAgeCategory(ctx context.Context, ageCategory string) ([]*models.Competition, error) {
	return u.repo.FindByAgeCategory(ctx, ageCategory)
}
func (u *CompetitionUseCase) FindByStatus(ctx context.Context, status string) ([]*models.Competition, error) {
	return u.repo.FindByStatus(ctx, status)
}

func (u *CompetitionUseCase) FindByID(ctx context.Context, id int) (*models.Competition, error) {
	return u.repo.FindByID(ctx, id)
}

func (u *CompetitionUseCase) Analytic(ctx context.Context) ([]*models.Competition, error) {
	return nil, nil
}
