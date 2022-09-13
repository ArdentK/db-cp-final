package usecase

import (
	"context"

	"github.com/ArdentK/db-cp-final/models"
	athletcomp "github.com/ArdentK/db-cp-final/pkg/athletComp"
)

type AthletCompUseCase struct {
	repo athletcomp.ACRepo
}

func NewACUseCase(r athletcomp.ACRepo) *AthletCompUseCase {
	return &AthletCompUseCase{
		repo: r,
	}
}

func (ac AthletCompUseCase) Create(ctx context.Context, user *models.User, idComp int) error {
	newAC := &models.AthletComp{
		Email:  user.Email,
		IDComp: idComp,
	}
	return ac.repo.AddRow(ctx, newAC)
}
func (ac AthletCompUseCase) Delete(ctx context.Context, user *models.User, idComp int) error {
	newAC := &models.AthletComp{
		Email:  user.Email,
		IDComp: idComp,
	}
	return ac.repo.DelRow(ctx, newAC)
}
