package service

import "math/rand"

type KeyRange struct {
	MinID int `db:"min_id"`
	MaxID int `db:"max_id"`
}

func (kr KeyRange) GetRandomKey() int {
	return rand.Intn(kr.MaxID-kr.MinID+1) + kr.MinID
}
