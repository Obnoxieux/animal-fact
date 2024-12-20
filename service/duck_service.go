package service

import (
	"github.com/animal-fact/model"
	"github.com/pocketbase/dbx"
)

const MaxTries = 10

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

func (d DuckService) GetRandom(db *dbx.DB) (model.Fact, error) {
	var keyRange KeyRange
	query := db.NewQuery("SELECT MIN(id) AS min_id, MAX(id) AS max_id FROM facts;")
	err := query.One(&keyRange)
	if err != nil || keyRange.MinID > keyRange.MaxID || keyRange.MinID == 0 {
		return nil, err
	}

	randomNumber := keyRange.GetRandomKey()
	tries := 0
	for {
		if tries > MaxTries {
			return nil, err
		}
		fact, err := d.GetByID(db, randomNumber)
		if err != nil {
			randomNumber = keyRange.GetRandomKey()
			tries++
		} else if fact.FactType() != d.Type() {
			randomNumber = keyRange.GetRandomKey()
			tries++
		} else {
			return fact, nil
		}
	}
}
