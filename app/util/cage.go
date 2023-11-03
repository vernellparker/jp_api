package util

import (
	"JurrassicParkAPI/app/constant"
	"JurrassicParkAPI/model"
	"github.com/lib/pq"
)

// ConvertCageDao converts DAO to model.Cage
func ConvertCageDao(c model.CageDAO) (cage model.Cage) {
	cage = model.Cage{
		CageID:      *c.CageID,
		Capacity:    *c.Capacity,
		Power:       *c.Power,
		SpeciesType: *c.SpeciesType,
		Species:     []constant.Species{},
	}

	if c.CurrentOccupancy == nil {
		cage.Capacity = 0
	} else {
		cage.Capacity = *c.CurrentOccupancy
	}

	for _, v := range c.Species {
		cage.Species = append(cage.Species, constant.Species(v))
	}

	return
}

// ConvertIntakeCage converts model.IntakeCage to CageDAO
func ConvertIntakeCage(c model.IntakeCage) (cage model.CageDAO) {
	cage = model.CageDAO{
		CageID:           &c.CageID,
		Capacity:         &c.Capacity,
		CurrentOccupancy: &c.CurrentOccupancy,
		Power:            &c.Power,
		SpeciesType:      &c.SpeciesType,
		Dinosaurs:        pq.Int64Array(c.Dinosaurs),
	}

	for _, v := range c.Species {
		cage.Species = append(cage.Species, string(v))
	}

	return
}
