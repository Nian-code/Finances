package Entity

type Account struct {
	Id    string   `json:"id" bson:"_id,omitempty"`
	Title string  `json:"title" bson:"title,omitempty"`
	Money float64 `json:"money" bson:"money,omitempty"`
}

