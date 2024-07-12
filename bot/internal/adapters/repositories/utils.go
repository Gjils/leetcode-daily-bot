package repositories

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

func toDoc(v interface{}) (doc *bson.D, err error) {
	data, err := bson.Marshal(v)
	if err != nil {
			return
	}

	err = bson.Unmarshal(data, &doc)
	fmt.Println(doc)
	return
}