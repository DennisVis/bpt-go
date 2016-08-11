package main

import (
	"database/sql"
	"github.com/DennisVis/bpt-go/persistence"
	"github.com/DennisVis/bpt-go/routes"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
)

func main() {

	db, err := sql.Open("postgres", "user=bpt dbname=bpt sslmode=disable")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	defer db.Close()
	db.SetMaxOpenConns(20)

	router := routes.NewRouter(map[string]persistence.DAO{
		"question": persistence.QuestionsDAO{db},
	})

	log.Println("Listening...")
	http.ListenAndServe(":3000", router)
}
