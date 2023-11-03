package repository

import (
	"JurrassicParkAPI/app/repository/sql"
	"JurrassicParkAPI/app/util"
	"JurrassicParkAPI/config"
	"JurrassicParkAPI/model"
)

type CageRepository struct {
	db *config.Database
}

func NewCageRepository(db *config.Database) *CageRepository {
	return &CageRepository{
		db,
	}
}

func (c *CageRepository) CreateCage(cage model.CageDAO) (id int, err error) {

	rows, err := c.db.NamedQuery(sql.CreateNewCreate, cage)
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	if rows.Next() {
		err := rows.Scan(&id)
		if err != nil {
			return 0, err
		}
	}

	return id, rows.Err()

}

func (c *CageRepository) GetOneCage(id int) (cage model.CageDAO, err error) {
	rows, err := c.db.Query(sql.GetOneCrates, id)
	if err != nil {
		return cage, err
	}
	defer rows.Close()

	if rows.Next() {

		err := rows.Scan(&cage.CageID, &cage.Capacity, &cage.Power, &cage.Species, &cage.SpeciesType, &cage.CurrentOccupancy, &cage.Dinosaurs)
		if err != nil {
			return cage, err
		}
	}

	return

}

func (c *CageRepository) GetAllCages() (cages []model.Cage, err error) {
	rows, err := c.db.Query(sql.GetAllCrates)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tempCage model.CageDAO
	var tempDino model.Dinosaur
	var tempCageId = 0

	for rows.Next() {

		err := rows.Scan(&tempCage.CageID, &tempCage.Capacity, &tempCage.Power, &tempCage.Species, &tempCage.SpeciesType, &tempCage.CurrentOccupancy, &tempDino.DinoID, &tempDino.Name, &tempDino.Species, &tempDino.SpeciesType)
		if err != nil {
			return nil, err
		}

		if tempCageId == *tempCage.CageID {
			cages[len(cages)-1].Dinosaurs = append(cages[len(cages)-1].Dinosaurs, tempDino)
		} else {
			cageDTO := util.ConvertCageDao(tempCage)
			cageDTO.Dinosaurs = append(cageDTO.Dinosaurs, tempDino)
			cages = append(cages, cageDTO)
			tempCageId = cageDTO.CageID
		}

	}

	return

}

func (c *CageRepository) UpdateCage(cage model.CageDAO) (id int, err error) {
	rows, err := c.db.NamedQuery(sql.UpdateCage, cage)
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	return *cage.CageID, rows.Err()

}
