package queries

var Search = "SELECT species_id, name, description, url FROM pokemon, plainto_tsquery($1) tsquery WHERE description_tsvector @@ tsquery;"

var SearchAll = "SELECT species_id, name, description, url FROM pokemon;"
