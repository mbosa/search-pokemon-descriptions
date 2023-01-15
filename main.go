package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/mbosa/search-pokemon-descriptions/config"
	"github.com/mbosa/search-pokemon-descriptions/server"
)

func main() {
	port := config.PORT

	s := server.NewServer()
	defer s.Close()

	log.Printf("server starting at http://localhost:%s", port)
	log.Fatal(fmt.Sprint(http.ListenAndServe(":"+port, s)))
}
