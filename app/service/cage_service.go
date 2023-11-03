package service

import (
	"JurrassicParkAPI/app/repository"
	"JurrassicParkAPI/app/util"
	"JurrassicParkAPI/model"
)

type CageService struct {
	cageRepository     *repository.CageRepository
	dinosaurRepository *repository.DinosaurRepository
}

func NewCrateService(cageRepository *repository.CageRepository, dinosaurRepository *repository.DinosaurRepository) *CageService {
	return &CageService{
		cageRepository,
		dinosaurRepository,
	}
}

func (c *CageService) CreateCage(cage model.IntakeCage) error {
	cageId, err := c.cageRepository.CreateCage(util.ConvertIntakeCage(cage))
	if err != nil {
		return err
	}
	if len(cage.Dinosaurs) > 0 {
		err = c.dinosaurRepository.UpdateCageID(cage.Dinosaurs, cageId)
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *CageService) GetAllCages() (cage []model.Cage, err error) {
	return c.cageRepository.GetAllCages()

}
func (c *CageService) GetOneCages(id int) (cage model.CageDAO, err error) {
	return c.cageRepository.GetOneCage(id)

}

func (c *CageService) UpdateCage(cage model.IntakeCage) (int, error) {
	return c.cageRepository.UpdateCage(util.ConvertIntakeCage(cage))
}
