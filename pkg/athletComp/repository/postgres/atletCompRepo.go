package postgres

import (
	"context"
	"database/sql"

	"github.com/ArdentK/db-cp-final/models"
)

type ACRepo struct {
	db *sql.DB
}

func NewAcRepository(db *sql.DB) *ACRepo {
	return &ACRepo{
		db: db,
	}
}

func (r *ACRepo) AddRow(ctx context.Context, ac *models.AthletComp) error {
	return r.db.QueryRow(
		"INSERT INTO AthletComp (email, id_competition) VALUES ($1, $2) RETURNING id",
		ac.Email,
		ac.IDComp,
	).Scan(&ac.ID)
}
func (r *ACRepo) DelRow(ctx context.Context, ac *models.AthletComp) error {
	_, err := r.db.Exec("DELETE FROM AthletComp WHERE email = $1 AND id_competition = $2",
		ac.Email,
		ac.IDComp,
	)
	return err
}
