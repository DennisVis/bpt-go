package routes

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/DennisVis/bpt-go/persistence"
)

func NewRouter(daos map[string]persistence.DAO) *mux.Router {

	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {

		var handler http.HandlerFunc
		modelType := route.ModelType

		handler = route.HandlerFactory(daos[modelType])

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}
