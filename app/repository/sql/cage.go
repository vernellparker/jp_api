package sql

var CreateCrateTable = `create table if not exists cage (
		cage_id serial primary key, 
		created_at timestamp with time zone, 
		updated_at timestamp with time zone, 
		deleted_at timestamp with time zone, 
		capacity bigint not null,
		power text not null, 
		species text[], 
		species_type text, 
		current_occupancy bigint GENERATED ALWAYS AS (array_length(dinosaurs,1) ) STORED not null, 
		dinosaurs  bigint [] 
                  );`

var CreateNewCreate = `insert into cage (
                  created_at, 
                  updated_at, 
                  capacity, 
                  power, 
                  species, 
                  species_type,
                  dinosaurs
                  ) values (
                            current_timestamp,
                            current_timestamp,
                            :capacity,
                            :power,
                            :species,
                            :species_type,
                            :dinosaurs
                  ) returning cage_id`

var GetAllCrates = `select cage.cage_id,
       capacity,
       power,
       cage.species,
       cage.species_type,
       current_occupancy,
       dino_id,
       name,
       d.species,
       d.species_type
from cage
         join public.dinosaur d on d.dino_id = ANY (cage.dinosaurs)`

var GetOneCrates = `select cage.cage_id,
      capacity,
      power,
      species,
      species_type,
      current_occupancy,
      dinosaurs
from cage where cage.cage_id=$1 `

var UpdateCage = `UPDATE cage
SET capacity          = COALESCE(NULLIF(:capacity,0), capacity),
    power             = COALESCE(NULLIF(:power,E''),power),
    species           = COALESCE(:species, species),
    species_type      = COALESCE(NULLIF(:species_type,E''), species_type),
    dinosaurs         = COALESCE(:dinosaurs, dinosaurs)

WHERE cage_id = :cage_id
  AND (cast(:capacity as bigint) IS NOT NULL AND cast(:capacity as bigint) IS DISTINCT FROM capacity OR
       cast(:power as text) IS NOT NULL AND cast(:power as text) IS DISTINCT FROM power OR
       cast(:species as text[])IS NOT NULL AND cast(:species as text[])IS DISTINCT FROM species OR
       cast(:species_type as text) IS NOT NULL AND  cast(:species_type as text) IS DISTINCT FROM species_type OR
       cast(:dinosaurs as bigint[]) IS NOT NULL AND cast(:dinosaurs as bigint[]) IS DISTINCT FROM dinosaurs
    );`

var UpdateDinoCageIDs = `update dinosaur set cage_id = (?) where dino_id in (?);`
