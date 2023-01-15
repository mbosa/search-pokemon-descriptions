CREATE TABLE IF NOT EXISTS pokemon (
  species_id text,
  name text,
  description text,
  description_tsvector tsvector,
  url text
);

COPY pokemon(species_id, name, description, url) FROM '/etc/pokemon.csv' DELIMITER ',' CSV HEADER;
UPDATE pokemon SET description_tsvector=to_tsvector(pokemon.description);
CREATE INDEX idx_pokemon_description_tsvector ON pokemon USING GIN(description_tsvector);
