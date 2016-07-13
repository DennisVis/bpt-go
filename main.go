package main

import (
	"log"
	"net/http"
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/DennisVis/bpt-go/routes"
	"github.com/DennisVis/bpt-go/persistence"
)

func main() {

	db, err := sql.Open("postgres", "user=bpt dbname=bpt sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	router := routes.NewRouter(map[string]persistence.DAO{
		"question": persistence.QuestionsDAO{db},
	})

	log.Println("Listening...")
	http.ListenAndServe(":9000", router)
}
