package service

import (
	"JurrassicParkAPI/app/repository"
	"JurrassicParkAPI/model"
	"golang.org/x/exp/slices"
)

type DinosaurService struct {
	dinosaurRepository *repository.DinosaurRepository
	cageRepository     *repository.CageRepository
}

func NewDinosaurService(dinosaurRepository *repository.DinosaurRepository, cageRepository *repository.CageRepository) *DinosaurService {
	return &DinosaurService{
		dinosaurRepository,
		cageRepository,
	}
}

func (d *DinosaurService) CreateDinosaur(dinosaur model.Dinosaur) error {
	return d.dinosaurRepository.CreateDinosaur(dinosaur)
}

func (d *DinosaurService) PreloadTestData() error {
	return d.dinosaurRepository.PreloadTestData()
}

func (d *DinosaurService) GetAllDinosaurs() (dinosaurs []model.Dinosaur, err error) {
	return d.dinosaurRepository.GetAllDinosaurs()
}

func (d *DinosaurService) GetOneDinosaur(id int) (dino model.Dinosaur, err error) {
	return d.dinosaurRepository.GetOneDinosaur(id)
}

// UpdateDinosaur also updates the cage if the dinosaur is moved
func (d *DinosaurService) UpdateDinosaur(dinosaur model.Dinosaur) (int, error) {
	updateDinosaur, err := d.dinosaurRepository.UpdateDinosaur(dinosaur)
	if err != nil {
		return 0, err
	}
	if *dinosaur.CageID != 0 {
		c := model.CageDAO{}
		c.CageID = dinosaur.CageID
		c.Dinosaurs = append(c.Dinosaurs, int64(dinosaur.DinoID))

		cage, err := d.cageRepository.GetOneCage(*c.CageID)
		if err != nil {
			return 0, err
		}
		if !slices.Contains(c.Dinosaurs, cage.Dinosaurs[0]) {
			c.Dinosaurs = append(c.Dinosaurs, cage.Dinosaurs...)
			_, err = d.cageRepository.UpdateCage(c)
			if err != nil {
				return 0, err
			}
		}

	}
	return updateDinosaur, err
}
