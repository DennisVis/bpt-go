package routes

import (
	"github.com/DennisVis/bpt-go/controllers"
	"github.com/DennisVis/bpt-go/persistence"
	"net/http"
)

type Route struct {
	Name           string
	Methods        []string
	Pattern        string
	ModelType      string
	HandlerFactory func(dao persistence.DAO) http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		[]string{"GET", "OPTIONS"},
		"/",
		"",
		func(dao persistence.DAO) http.HandlerFunc {
			return controllers.IndexController
		},
	},
	Route{
		"AllQuestions",
		[]string{"GET", "OPTIONS"},
		"/questions",
		"question",
		func(dao persistence.DAO) http.HandlerFunc {
			return controllers.AllQuestionsController(dao)
		},
	},
	Route{
		"CreateQuestion",
		[]string{"POST", "OPTIONS"},
		"/questions",
		"question",
		func(dao persistence.DAO) http.HandlerFunc {
			return controllers.CreateQuestionController(dao)
		},
	},
	Route{
		"ReadQuestion",
		[]string{"GET", "OPTIONS"},
		"/questions/{questionId}",
		"question",
		func(dao persistence.DAO) http.HandlerFunc {
			return controllers.ReadQuestionController(dao)
		},
	},
	Route{
		"UpdateQuestion",
		[]string{"PUT", "OPTIONS"},
		"/questions/{questionId}",
		"question",
		func(dao persistence.DAO) http.HandlerFunc {
			return controllers.UpdateQuestionController(dao)
		},
	},
	Route{
		"DeleteQuestion",
		[]string{"DELETE", "OPTIONS"},
		"/questions/{questionId}",
		"question",
		func(dao persistence.DAO) http.HandlerFunc {
			return controllers.DeleteQuestionController(dao)
		},
	},
}
