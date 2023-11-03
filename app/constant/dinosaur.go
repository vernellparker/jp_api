package constant

var SpeciesToSpeciesType = map[Species]SpeciesType{
	Brachiosaurus: Herbivore,
	Stegosaurus:   Herbivore,
	Ankylosaurus:  Herbivore,
	Triceratops:   Herbivore,
	Protoceratops: Herbivore,
	Hadrosaur:     Herbivore,
	Tyrannosaurus: Carnivore,
	Velociraptor:  Carnivore,
	Spinosaurus:   Carnivore,
	Megalosaurus:  Carnivore,
}

type Species string

const (
	Brachiosaurus Species = "Brachiosaurus"
	Stegosaurus   Species = "Stegosaurus"
	Ankylosaurus  Species = "Ankylosaurus"
	Triceratops   Species = "Triceratops"
	Protoceratops Species = "Protoceratops"
	Hadrosaur     Species = "Hadrosaur"
	Tyrannosaurus Species = "Tyrannosaurus"
	Velociraptor  Species = "Velociraptor"
	Spinosaurus   Species = "Spinosaurus"
	Megalosaurus  Species = "Megalosaurus"
)

type SpeciesType string

const (
	Carnivore SpeciesType = "Carnivore"
	Herbivore SpeciesType = "Herbivore"
)
