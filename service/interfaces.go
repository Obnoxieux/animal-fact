package service

import (
	"github.com/animal-fact/model"
	"github.com/pocketbase/dbx"
)

type Service interface {
	Type() string
	GetByID(db *dbx.DB, id int) (model.Fact, error)
	GetRandom(db *dbx.DB) (model.Fact, error)
}
