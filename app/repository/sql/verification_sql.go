package sql

var VerifyDinoExist = `select dino_id, name, species, species_type from dinosaur where dino_id in (?)`
