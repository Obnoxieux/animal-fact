package service

import (
	"github.com/animal-fact/model"
	"github.com/pocketbase/dbx"
)

type DuckService struct{}

func (d DuckService) Type() string {
	return "duck"
}

func (d DuckService) GetByID(db *dbx.DB, id int) (model.Fact, error) {
	var fact model.DuckFact
	err := db.Select().Model(id, &fact)
	if err != nil {
		return model.DuckFact{}, err
	}

	return fact, nil
}
