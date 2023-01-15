package config

import (
	_ "github.com/mbosa/search-pokemon-descriptions/dotenv/autoload"

	"os"
)

var (
	PORT         = os.Getenv("PORT")
	DATABASE_URL = os.Getenv("DATABASE_URL")
)
