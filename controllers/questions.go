package controllers

import (
	"encoding/json"
	"github.com/DennisVis/bpt-go/models"
	"github.com/DennisVis/bpt-go/persistence"
	"github.com/gorilla/mux"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
)

func AllQuestionsController(dao persistence.DAO) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		questions, err := dao.All()

		if err != nil {
			serverError(w, err)
		} else {

			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			w.WriteHeader(http.StatusOK)

			json.NewEncoder(w).Encode(questions)
		}
	}
}

func CreateQuestionController(dao persistence.DAO) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		var question models.Question
		body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
		if err != nil {
			badRequest(w, err)
			return
		}
		if err := r.Body.Close(); err != nil {
			badRequest(w, err)
			return
		}
		if err := json.Unmarshal(body, &question); err != nil {
			badRequest(w, err)
			return
		}

		questionId, err := dao.Create(question)

		if err != nil {
			serverError(w, err)
		} else {

			w.WriteHeader(http.StatusCreated)
			w.Header().Set("Content-Type", "application/json; charset=utf-8")

			if err := json.NewEncoder(w).Encode(models.Question{questionId, question.Name, question.Labels}); err != nil {
				serverError(w, err)
			}
		}
	}
}

func ReadQuestionController(dao persistence.DAO) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		questionId, err := strconv.Atoi(mux.Vars(r)["questionId"])
		if err != nil {
			badRequest(w, err)
		}

		question, err := dao.Read(questionId)

		if err != nil {
			serverError(w, err)
		} else if question == nil {

		} else {

			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json; charset=utf-8")

			if err := json.NewEncoder(w).Encode(question); err != nil {
				serverError(w, err)
			}
		}
	}
}

func UpdateQuestionController(dao persistence.DAO) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		questionId, err := strconv.Atoi(mux.Vars(r)["questionId"])
		if err != nil {
			badRequest(w, err)
			return
		}

		var question models.Question
		body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
		if err != nil {
			badRequest(w, err)
			return
		}
		if err := r.Body.Close(); err != nil {
			badRequest(w, err)
			return
		}
		if err := json.Unmarshal(body, &question); err != nil {
			badRequest(w, err)
			return
		}

		ur, err := dao.Update(questionId, question)
		if err != nil {
			serverError(w, err)
			return
		} else {

			if q, ok := ur.(models.Question); ok {
				question = q
			}

			w.WriteHeader(http.StatusAccepted)
			w.Header().Set("Content-Type", "application/json; charset=utf-8")

			if err := json.NewEncoder(w).Encode(question); err != nil {
				serverError(w, err)
			}
		}
	}
}

func DeleteQuestionController(dao persistence.DAO) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		questionId, err := strconv.Atoi(mux.Vars(r)["questionId"])
		if err != nil {
			badRequest(w, err)
			return
		}

		rowsAffected, err := dao.Delete(questionId)

		if err != nil {
			serverError(w, err)
		} else if rowsAffected < 1 {
			notFound(w)
		} else {

			w.WriteHeader(http.StatusNoContent)
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
		}
	}
}
