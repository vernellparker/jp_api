package repository

import (
	"JurrassicParkAPI/app/repository/sql"
	"JurrassicParkAPI/config"
	"JurrassicParkAPI/model"
	sql2 "database/sql"
	"github.com/jmoiron/sqlx"
)

type DinosaurRepository struct {
	db *config.Database
}

func NewDinosaurRepository(db *config.Database) *DinosaurRepository {
	return &DinosaurRepository{
		db,
	}
}

func (d *DinosaurRepository) CreateDinosaur(dinosaur model.Dinosaur) (err error) {
	rows, err := d.db.NamedQuery(sql.CreateDinosaur, dinosaur)
	if err != nil {
		return err
	}
	defer rows.Close()
	var id int
	if rows.Next() {
		err := rows.Scan(&id)
		if err != nil {
			return err
		}
	}

	return rows.Err()

}

func (d *DinosaurRepository) UpdateDinosaur(dinosaur model.Dinosaur) (id int, err error) {
	rows, err := d.db.NamedQuery(sql.UpdateDino, dinosaur)
	if err != nil {
		return id, err
	}
	defer rows.Close()
	if rows.Next() {
		err := rows.Scan(&id)
		if err != nil {
			return id, err
		}
	}

	return id, rows.Err()

}

func (d *DinosaurRepository) PreloadTestData() (err error) {
	rows, err := d.db.Query(sql.InsertPreLoadDinosaurs)
	if err != nil {
		return err
	}
	defer rows.Close()

	return rows.Err()

}

func (d *DinosaurRepository) GetAllDinosaurs() (dinos []model.Dinosaur, err error) {
	rows, err := d.db.Query(sql.GetAllDinos)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var tempDino model.Dinosaur
	for rows.Next() {
		err := rows.Scan(&tempDino.DinoID, &tempDino.Name, &tempDino.Species, &tempDino.SpeciesType, &tempDino.CageID)
		if err != nil {
			return nil, err
		}
		dinos = append(dinos, tempDino)
	}

	return

}

func (d *DinosaurRepository) GetOneDinosaur(id int) (dino model.Dinosaur, err error) {
	rows, err := d.db.Query(sql.GetOneDino, id)
	if err != nil {
		return dino, err
	}

	if rows.Next() {
		err := rows.Scan(&dino.DinoID, &dino.Name, &dino.Species, &dino.SpeciesType, &dino.CageID)
		if err != nil {
			return dino, err
		}
	}

	return

}

func (d *DinosaurRepository) UpdateCageID(dinoID []int64, cageID int) (err error) {
	query, args, err := sqlx.In(sql.UpdateDinoCageIDs, cageID, dinoID)
	if err != nil {
		return err
	}

	// sqlx.In returns queries with the `?` bindvar, we can rebind it for our backend
	query = d.db.Rebind(query)
	rows, err := d.db.Query(query, args...)
	if err != nil {
		return err
	}
	defer func(rows *sql2.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)
	return rows.Err()
}
