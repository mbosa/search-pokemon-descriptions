# search-pokemon-descriptions

An app to search pokemons from keywords in their description.

It takes advantage of PostgreSQL [full text search](https://www.postgresql.org/docs/current/textsearch.html) feature.\
Each description is stored in a db as a list of [lexemes](https://en.wikipedia.org/wiki/Lexeme) (a [tsvector](https://www.postgresql.org/docs/current/datatype-textsearch.html#DATATYPE-TSVECTOR) in PostgreSQL).\
Similarly, each word in the input is converted to a lexeme and compared to the list of lexemes in each description.

The server serves static files and html pages made with the `html/template` package from Go.
