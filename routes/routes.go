package routes

import (
	"net/http"
	"github.com/DennisVis/bpt-go/controllers"
	"github.com/DennisVis/bpt-go/persistence"
)

type Route struct {
	Name           string
	Method         string
	Pattern        string
	ModelType      string
	HandlerFactory func (dao persistence.DAO) http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		"",
		func (dao persistence.DAO) http.HandlerFunc {
			return controllers.IndexController
		},
	},
	Route{
		"AllQuestions",
		"GET",
		"/questions",
		"question",
		func (dao persistence.DAO) http.HandlerFunc {
			return controllers.AllQuestionsController(dao)
		},
	},
	Route{
		"CreateQuestion",
		"POST",
		"/questions",
		"question",
		func (dao persistence.DAO) http.HandlerFunc {
			return controllers.CreateQuestionController(dao)
		},
	},
	Route{
		"ReadQuestion",
		"GET",
		"/questions/{questionId}",
		"question",
		func (dao persistence.DAO) http.HandlerFunc {
			return controllers.ReadQuestionController(dao)
		},
	},
	Route{
		"UpdateQuestion",
		"PUT",
		"/questions/{questionId}",
		"question",
		func (dao persistence.DAO) http.HandlerFunc {
			return controllers.UpdateQuestionController(dao)
		},
	},
	Route{
		"DeleteQuestion",
		"DELETE",
		"/questions/{questionId}",
		"question",
		func (dao persistence.DAO) http.HandlerFunc {
			return controllers.DeleteQuestionController(dao)
		},
	},
}
