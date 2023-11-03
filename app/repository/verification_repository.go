package repository

import (
	"JurrassicParkAPI/app/repository/sql"
	"JurrassicParkAPI/config"
	"JurrassicParkAPI/model"
	"github.com/jmoiron/sqlx"
)

type VerificationRepository struct {
	db *config.Database
}

func NewVerificationRepository(db *config.Database) *VerificationRepository {
	return &VerificationRepository{
		db,
	}
}

func (c *VerificationRepository) CheckIfDinoExitsInDB(ids []int64) (found []model.Dinosaur, err error) {
	query, args, err := sqlx.In(sql.VerifyDinoExist, ids)
	if err != nil {
		return nil, err
	}
	query = c.db.Rebind(query)
	rows, err := c.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var result model.Dinosaur
	for rows.Next() {
		err := rows.Scan(&result.DinoID, &result.Name, &result.Species, &result.SpeciesType)
		if err != nil {
			return nil, err
		}
		found = append(found, result)
	}

	return

}
