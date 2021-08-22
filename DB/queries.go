package DB

import (
	"github/nian-code/finances/Util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

func FindAll(nameCollection, obj string) (interface{}, error) {

	documents := Util.GetArrayByType(obj)
	client, ctx, err := connect()

	if err != nil{
		return documents, err
	}

	defer disconnect(client, ctx)

	collection := client.Database("finances").Collection(nameCollection)
	cur, currErr := collection.Find(ctx, bson.D{})

	if currErr != nil { panic(currErr) }
	defer cur.Close(ctx)

	if err := cur.All(ctx, &documents); err != nil {
		panic(err)
	}
	return documents, err
}

func InsertDocuments(nameCollection string, objs []interface{}) (*mongo.InsertManyResult, error) {

	client, ctx, err := connect()

	if err != nil{
		return nil, err
	}

	defer disconnect(client, ctx)

	collection := client.Database("finances").Collection(nameCollection)

	/*
	   Insert documents
	*/
	var docs []interface{}

	for _, elem := range objs {
		doc , _ := Util.ToDoc(elem)
		docs = append(docs, doc)
	}

	res, insertErr := collection.InsertMany(ctx, docs)

	if insertErr != nil {
		log.Fatal(insertErr)
	}

	return res, nil
}
