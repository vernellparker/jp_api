package service

import (
	"JurrassicParkAPI/app/constant"
	"JurrassicParkAPI/app/repository"
	"JurrassicParkAPI/app/util"
	"JurrassicParkAPI/model"
	"fmt"
	"golang.org/x/exp/slices"
)

type VerificationService struct {
	verificationRepository *repository.VerificationRepository
	cageRepository         *repository.CageRepository
}

func NewVerificationService(verificationRepository *repository.VerificationRepository, cageRepository *repository.CageRepository) *VerificationService {
	return &VerificationService{
		verificationRepository,
		cageRepository,
	}
}

// IsNotOverCapacity checks if cage is over capacity
func (c *VerificationService) IsNotOverCapacity(cage model.IntakeCage) (bool, string) {
	if len(cage.Dinosaurs) > cage.Capacity {
		return false, "The number dinosaurs you have listed is greater than the capacity you have set for this cage."
	}
	return true, ""
}

// CageVerificationChecks handles several checks for verification
func (c *VerificationService) CageVerificationChecks(cage model.IntakeCage, dinoId int64) (bool, string, error) {
	var dinos []model.Dinosaur
	if cage.Dinosaurs != nil && len(cage.Dinosaurs) != 0 {

		if ok, msg := c.IsNotOverCapacity(cage); !ok {
			return ok, msg, nil
		}

		var err error

		housingDinos := cage.Dinosaurs
		if dinoId != 0 {
			housingDinos = []int64{dinoId}
		}

		if cage.Power == constant.Down {
			return false, "Dinosaurs cannot be moved into cages that are powered off", err
		}

		dinos, err = c.verificationRepository.CheckIfDinoExitsInDB(housingDinos)
		if err != nil {
			return false, "", err
		}

		if ok, msg, err := doesExistInDb(cage, dinos); !ok {
			return ok, msg, err
		}

		if ok, msg, err := dinoSpeciesCheck(cage, dinos); !ok {
			return ok, msg, err
		}

		if ok, msg, err := SpeciesTypeCheck(cage, dinos); !ok {
			return ok, msg, err
		}

		ok, msg := areDinoCompatible(dinos)
		return ok, msg, err
	}

	return true, "", nil
}

// CagePowerChecks Check if cage has power
func (c *VerificationService) CagePowerChecks(cage model.IntakeCage) (bool, string, error) {
	if cage.Power != "" {
		currentCage, err := c.cageRepository.GetOneCage(cage.CageID)

		if len(currentCage.Dinosaurs) != 0 && cage.Power == constant.Down {
			return false, "Cage still contains Dinosaurs and cannot be powered off", err
		}
	}

	return true, "", nil
}

// DinoMoveUpdateChecks checks if dinos should be moved
func (c *VerificationService) DinoMoveUpdateChecks(dino model.Dinosaur) (bool, string, error) {
	if *dino.CageID != 0 {
		currentCage, err := c.cageRepository.GetOneCage(*dino.CageID)

		if *currentCage.Power != constant.Active {
			return false, "Dinosaurs cannot be moved into cages that are powered off", err
		}

		if currentCage.CurrentOccupancy != nil && *currentCage.CurrentOccupancy+1 > *currentCage.Capacity {
			return false, "Dinosaurs cannot be moved into this cage as it would be fulled pass its maximumCapacity ", err
		}
		temp := util.ConvertCageDao(currentCage)

		if ok, msg, err := c.CageVerificationChecks(model.IntakeCage{
			CageID:           temp.CageID,
			Capacity:         temp.Capacity,
			CurrentOccupancy: temp.CurrentOccupancy,
			Power:            temp.Power,
			Species:          temp.Species,
			SpeciesType:      temp.SpeciesType,
			Dinosaurs:        currentCage.Dinosaurs,
		}, int64(dino.DinoID)); !ok {
			return ok, msg, err
		}
	}

	return true, "", nil
}

// SpeciesTypeCheck checks if Species can be assigned to cage by the SpeciesType
func SpeciesTypeCheck(cage model.IntakeCage, dinos []model.Dinosaur) (bool, string, error) {
	for _, v := range dinos {
		if constant.SpeciesToSpeciesType[v.Species] != cage.SpeciesType && cage.SpeciesType != "" {
			return false, fmt.Sprintf("Error: Species assigned to cage do not match the correct Species Type labed for cage."), nil
		}
	}
	return true, "", nil
}

// dinoSpeciesCheck ensure Species assigned to cage matches the correct Species labeled for cage
func dinoSpeciesCheck(cage model.IntakeCage, dinos []model.Dinosaur) (bool, string, error) {
	flatDinos := map[int]constant.Species{}
	for _, v := range dinos {
		flatDinos[v.DinoID] = v.Species
	}

	for _, v := range flatDinos {
		if cage.Species != nil && !slices.Contains(cage.Species, v) {
			return false, "Error: Species assigned to cage do not match the correct Species labeled for cage.", nil
		}
	}
	return true, "", nil
}

// areDinoCompatible handles dino housing compatible
func areDinoCompatible(dinos []model.Dinosaur) (bool, string) {

	if ok, msg := speciesTypeCheck(dinos); !ok {
		return false, msg
	}

	if ok, msg := carnivoreCheck(dinos); !ok {
		return false, msg
	}

	return true, ""
}

// speciesTypeCheck ensures users are not trying to house Herbivores with Carnivore
func speciesTypeCheck(dinos []model.Dinosaur) (bool, string) {
	var speciesType constant.SpeciesType
	if len(dinos) < 2 {
		return true, ""
	}
	for i, v := range dinos {
		if i == 0 {
			speciesType = v.SpeciesType
			continue
		}
		if speciesType != v.SpeciesType {
			return false, fmt.Sprintf("Error: You are trying to house Herbivores with Carnivore, which is not allowed")
		}
	}
	return true, ""
}

// carnivoreCheck checks if dino is a carnivore
func carnivoreCheck(dinos []model.Dinosaur) (bool, string) {
	var species constant.Species
	for i, v := range dinos {
		if v.SpeciesType == constant.Carnivore && i == 0 {
			species = v.Species
			continue
		}

		if species != v.Species && species != "" {
			return false, fmt.Sprintf("Error: Only Carnivores,of the same species are allowed to be housed together.")
		}
	}
	return true, ""
}

// doesExistInDb ensures dino is in db first before trying ro add them to a cage
func doesExistInDb(cage model.IntakeCage, dinos []model.Dinosaur) (bool, string, error) {
	var missingDinos []int

	var foundIDs []int
	for _, v := range dinos {
		foundIDs = append(foundIDs, v.DinoID)
	}

	if len(cage.Dinosaurs) != len(dinos) && len(cage.Dinosaurs) != 0 {

		for _, v := range cage.Dinosaurs {
			if !slices.Contains(foundIDs, int(v)) {
				missingDinos = append(missingDinos, int(v))

			}
		}
		message := fmt.Sprintf("Dinosaur(s) with the id(s) of %d cannot be found in the database and must be added first before they are added to a cage", missingDinos)
		return false, message, nil

	}

	return true, "", nil
}
