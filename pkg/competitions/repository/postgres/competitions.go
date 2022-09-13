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

func (r *CompetitionRepo) Create(ctx context.Context, c *models.Competition) error {
	return nil
}

func (r *CompetitionRepo) Update(ctx context.Context, c *models.Competition) error {
	return nil
}

func (r *CompetitionRepo) Delete(ctx context.Context, id int) error {
	return nil
}

func (r *CompetitionRepo) FindByID(ctx context.Context, id int) (*models.Competition, error) {
	newComp := &models.Competition{}
	err := r.db.QueryRow(
		"SELECT * FROM competitions WHERE id = $1;",
		id,
	).Scan(
		&newComp.ID,
		&newComp.Name,
		&newComp.Date,
		&newComp.AgeCategory,
		&newComp.WeaponType,
		&newComp.IsTeam,
		&newComp.Status,
		&newComp.Sex,
		&newComp.NumOfAthlets,
	)
	if err != nil {
		return nil, err
	}

	return newComp, nil
}

func (r *CompetitionRepo) List(ctx context.Context) ([]*models.Competition, error) {
	items := []*models.Competition{}

	rows, err := r.db.Query(
		"SELECT * FROM competitions;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		post := &models.Competition{}
		err = rows.Scan(
			&post.ID,
			&post.Name,
			&post.Date,
			&post.AgeCategory,
			&post.WeaponType,
			&post.IsTeam,
			&post.Status,
			&post.Sex,
			&post.NumOfAthlets,
		)
		if err != nil {
			return nil, err
		}
		items = append(items, post)
	}

	return items, nil
}
func (r *CompetitionRepo) SortASCByData(ctx context.Context) ([]*models.Competition, error) {
	return nil, nil
}
func (r *CompetitionRepo) SortDESCByData(ctx context.Context) ([]*models.Competition, error) {
	return nil, nil
}
func (r *CompetitionRepo) SortASCByNumOfPart(ctx context.Context) ([]*models.Competition, error) {
	return nil, nil
}
func (r *CompetitionRepo) SortDESCByNumOfPart(ctx context.Context) ([]*models.Competition, error) {
	return nil, nil
}
func (r *CompetitionRepo) FindByWeaponType(ctx context.Context, weaponType string) ([]*models.Competition, error) {
	return nil, nil
}
func (r *CompetitionRepo) FindByAgeCategory(ctx context.Context, ageCategory string) ([]*models.Competition, error) {
	return nil, nil
}
func (r *CompetitionRepo) FindByStatus(ctx context.Context, status string) ([]*models.Competition, error) {
	return nil, nil
}
