package sql

var CreateDinosaurTable = `create table if not exists dinosaur (
		dino_id serial primary key , 
		created_at timestamp with time zone, 
		updated_at timestamp with time zone, 
		deleted_at timestamp with time zone, 
		name text not null, 
		species text not null, 
		species_type text not null, 
		cage_id bigint 
                  );`

var InsertPreLoadDinosaurs = `INSERT INTO dinosaur (
                      created_at, 
                      updated_at,  
                      name,
                      species,
                      species_type,
                      cage_id
                      ) values 
                            (current_timestamp, current_timestamp, 'Barney','Tyrannosaurus', 'Carnivore',0),
                            (current_timestamp, current_timestamp, 'Baby Bop','Triceratops','Herbivore',0),
                            (current_timestamp, current_timestamp, 'BJ','Protoceratops','Herbivore',0),
                            (current_timestamp, current_timestamp, 'Riff','Hadrosaur','Herbivore',0),
                            (current_timestamp, current_timestamp, 'Moe','Velociraptor','Carnivore',0),
                        	(current_timestamp, current_timestamp, 'Larry','Velociraptor','Carnivore',0),
                        	(current_timestamp, current_timestamp, 'Curly','Velociraptor','Carnivore',0)
                  	ON CONFLICT (dino_id) DO NOTHING;`

var CreateDinosaur = `INSERT INTO dinosaur (
                      created_at, 
                      updated_at,  
                      name,
                      species,
                      species_type 
                      ) values (
                                current_timestamp, 
                                current_timestamp, 
                                :name,
                                :species, 
                                :species_type
                            );`

var GetAllDinos = `select dino_id,name,species, species_type, cage_id from dinosaur`

var UpdateDino = `UPDATE dinosaur
SET name         = COALESCE(NULLIF(:name, E''), name),
    species      = COALESCE(NULLIF(:species, E''), species),
    species_type = COALESCE(NULLIF(:species_type, E''), species_type),
    cage_id      = COALESCE(NULLIF(:cage_id, 0), cage_id)
WHERE dino_id = :dino_id
  AND (cast(:name as text) IS NOT NULL AND cast(:name as text) IS DISTINCT FROM name OR
       cast(:species as text) IS NOT NULL AND cast(:species as text) IS DISTINCT FROM species OR
       cast(:species_type as text) IS NOT NULL AND cast(:species_type as text) IS DISTINCT FROM species_type OR
       cast(:cage_id as bigint) IS NOT NULL AND cast(:cage_id as bigint) IS DISTINCT FROM cage_id);`

var GetOneDino = `select dinosaur.dino_id,
       dinosaur.name,
       dinosaur.species,
       dinosaur.species_type,
       dinosaur.cage_id
from dinosaur where dino_id = $1`
