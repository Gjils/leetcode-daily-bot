package groupsService

import (
	"context"
	"fmt"

	dbLayer "leetcodebot/data/db"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var ctx = context.TODO()

type GroupsService struct {
	groups *mongo.Collection
}

type Group struct {
	Id int `bson:"id"`
	Enabled bool `bson:"enabled"`
}

func (service *GroupsService) GetAll() ([]Group, error) {
	cur, err := service.groups.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	var groupsList []Group
	if err := cur.All(ctx, &groupsList); err != nil {
		return nil, err
	}
	fmt.Println("Get all operation")
	return groupsList, nil
}

func (service *GroupsService) Add(id int) error {
	if _, err := service.groups.InsertOne(ctx, bson.D{{Key: "id", Value: id}, {Key: "enabled", Value: true}}); err != nil {
		return err
	}
	fmt.Println("Added", id)
	return nil
}

func (service *GroupsService) Check(id int) (bool, error) {
	var res Group
	if err := service.groups.FindOne(ctx, bson.D{{Key: "id", Value: id}}).Decode(&res); err != nil {
		return false, nil
	}
	return true, nil
}

func (service *GroupsService) Delete(id int) error {
	if _, err := service.groups.DeleteOne(ctx, bson.D{{Key: "id", Value: id}}); err != nil {
		return err
	}
	return nil
}

func GetApi() GroupsService {
	return GroupsService{groups: dbLayer.DB.Collection("groups")}
}
