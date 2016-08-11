package routes

import (
	"github.com/DennisVis/bpt-go/persistence"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"net/http"
)

func NewRouter(daos map[string]persistence.DAO) *mux.Router {

	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		modelType := route.ModelType

		c := cors.New(cors.Options{
			AllowedMethods: []string{"GET", "OPTIONS", "PATCH", "POST", "PUT"},
		})

		handler := route.HandlerFactory(daos[modelType])

		corsHhandler := c.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method == "OPTIONS" {
				println("OPTIONS!!!!")
				w.Header().Set("Access-Control-Allow-Origin", "*")
				w.Header().Set("Access-Control-Allow-Methods", "*")
				w.Header().Set("Access-Control-Allow-Headers", "*")
				w.WriteHeader(http.StatusOK)
			} else {
				handler.ServeHTTP(w, r)
			}
		}))

		router.
			Methods(route.Methods...).
			Path(route.Pattern).
			Name(route.Name).
			Handler(corsHhandler)
	}

	return router
}
