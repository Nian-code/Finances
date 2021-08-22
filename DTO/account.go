package DTO

type AccountDTO struct {
	Title string  `json:"title" bson:"title,omitempty"`
	Money float64 `json:"money" bson:"money,omitempty"`
}
