package autoload

import "github.com/mbosa/search-pokemon-descriptions/dotenv"

func init() {
	dotenv.Load(".env")
}
