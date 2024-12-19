package model

type DuckFact struct {
	ID     int    `json:"id" db:"id"`
	Type   string `json:"type" db:"type"`
	Text   string `json:"text" db:"text"`
	Length int    `json:"length" db:"length"`
}

func (f DuckFact) TableName() string {
	return "facts"
}

func (f DuckFact) FactType() string {
	return "duck"
}
