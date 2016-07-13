package persistence

import (
	"github.com/DennisVis/bpt-go/models"
	"database/sql"
)

type DB interface {
	Begin() (*sql.Tx, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	Exec(query string, args ...interface{}) (sql.Result, error)
}

type DAO interface {
	All() ([]models.Model, error)
	Create(item models.Model) (int, error)
	Read(id int) (*models.Model, error)
	Update(id int, item models.Model) (models.Model, error)
	Delete(id int) (int, error)
}

func mapToSlice(theMap map[int]models.Model) []models.Model {

	slice := make([]models.Model, 0, len(theMap))

	for _, value := range theMap {
		slice = append(slice, value)
	}

	return slice
}
