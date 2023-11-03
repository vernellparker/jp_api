package model

import (
	"JurrassicParkAPI/app/constant"
	"github.com/lib/pq"
)

type Cage struct {
	CageID           int                  `json:"cage_id"`
	Capacity         int                  `json:"capacity"`
	CurrentOccupancy int                  `json:"current_occupancy"`
	Power            constant.PowerStatus `json:"power"`
	Species          []constant.Species   `json:"species"`
	SpeciesType      constant.SpeciesType `json:"species_type"`
	Dinosaurs        []Dinosaur           `json:"dinosaurs"`
}

type CageDAO struct {
	CageID           *int                  `db:"cage_id"`
	Capacity         *int                  `json:"capacity"`
	CurrentOccupancy *int                  `json:"-" db:"current_occupancy"`
	Power            *constant.PowerStatus `json:"power"`
	Species          pq.StringArray        `json:"species"`
	SpeciesType      *constant.SpeciesType `json:"species_type" db:"species_type"`
	Dinosaurs        pq.Int64Array         `json:"dinosaurs"`
}
type IntakeCage struct {
	CageID           int                  `json:"-"`
	Capacity         int                  `json:"capacity"`
	CurrentOccupancy int                  `json:"-"`
	Power            constant.PowerStatus `json:"power"`
	Species          []constant.Species   `json:"species"`
	SpeciesType      constant.SpeciesType `json:"species_type"`
	Dinosaurs        []int64              `json:"dinosaurs"`
}
