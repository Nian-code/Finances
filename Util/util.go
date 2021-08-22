package Util

import (
	e "github.com/Nian-code/Finances/Entity"
	"go.mongodb.org/mongo-driver/bson"
)

func GetArrayByType(obj string) (array interface{}){
	switch obj{
		case "Account":
			return []e.Account{}
	default:
		return
	}
}

func ToDoc(v interface{}) (doc *bson.D, err error) {
	data, err := bson.Marshal(v)
	if err != nil {
	return
	}

	err = bson.Unmarshal(data, &doc)
	return
}
