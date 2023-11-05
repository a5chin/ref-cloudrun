package entity

type Ingredient struct {
	ID         string       `json:"id"`
	Name       string       `json:"name"`
	Nutritions []*Nutrition `json:"nutritions"`
	Weight     uint         `json:"weight"`
}
