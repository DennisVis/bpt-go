package controllers

import (
	"net/http"
	"fmt"
	"strings"
	"log"
)

func serverError(w http.ResponseWriter, err error) {

	log.Println(err)

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusInternalServerError)

	r := strings.NewReplacer("\"", "'")
	w.Write([]byte(fmt.Sprintf(`{"message": "%s"}`, r.Replace(err.Error()))))
}

func badRequest(w http.ResponseWriter, err error) {

	log.Println(err)

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusBadRequest)

	r := strings.NewReplacer("\"", "'")
	w.Write([]byte(fmt.Sprintf(`{"message": "%s"}`, r.Replace(err.Error()))))
}


func notFound(w http.ResponseWriter) {

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusNotFound)

	w.Write([]byte(`{"message": "Not found"}`))
}

func IndexController(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	w.Write([]byte(`{"version": "1.0.0"}`))
}
