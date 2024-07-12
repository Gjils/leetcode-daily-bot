package repositories

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"leetcodebot/internal/domain/entities"
	"leetcodebot/internal/domain/repositories"
)

var ctx = context.TODO()

type GroupRepository struct {
	groups *mongo.Collection
}


func GetGroupsRepository(db *mongo.Database) repositoryInterfaces.GroupRepository {
	return GroupRepository{groups: db.Collection("groups")}
}

func (rep GroupRepository) GetAll() ([]entities.Group, error) {
	cur, err := rep.groups.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	var groupsList []entities.Group
	if err := cur.All(ctx, &groupsList); err != nil {
		return nil, err
	}
	fmt.Println("Get all operation")
	return groupsList, nil
}

func (rep GroupRepository) Add(group entities.GroupInfo) error {
	doc, err := toDoc(group)
	if err != nil {
		return err
	}
	if _, err := rep.groups.InsertOne(ctx, doc); err != nil {
		return err
	}
	fmt.Println("Added", group.Title)
	return nil
}

func (rep GroupRepository) Find(group entities.GroupQuery) (bool, error) {
	doc, err := toDoc(group)
	if err != nil {
		return false, err
	}
	var res entities.Group
	if err := rep.groups.FindOne(ctx, doc).Decode(&res); err != nil {
		return false, nil
	}
	return true, nil
}

func (rep GroupRepository) Delete(group entities.GroupQuery) error {
	if _, err := rep.groups.DeleteOne(ctx, group); err != nil {
		return err
	}
	return nil
}
