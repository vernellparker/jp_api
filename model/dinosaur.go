package model

import (
	"JurrassicParkAPI/app/constant"
)

type Dinosaur struct {
	DinoID      int                  `db:"dino_id"`
	Name        string               `json:"name"`
	Species     constant.Species     `json:"species"`
	SpeciesType constant.SpeciesType `db:"species_type" json:"species_type"`
	CageID      *int                 `db:"cage_id" json:"cage_id"`
}
