package models

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Number struct {
	ID    primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Value int                `json:"value" bson:"value,omitempty"`
}

func GetNumbers() ([]Number, error) {
	var numbers []Number
	c := DB.Collection(number_collection)
	curr, err := c.Find(context.TODO(), bson.M{})

	if err != nil {
		return []Number{}, errors.New("error getting numbers from db")
	}

	defer curr.Close(context.TODO())

	for curr.Next(context.TODO()) {
		var number Number
		if err := curr.Decode(&number); err != nil {
			return []Number{}, errors.New("error deserealizing number entry")
		}
		numbers = append(numbers, number)
	}

	if err = curr.Err(); err != nil {
		return []Number{}, errors.New("current connection presented some error")
	}

	return numbers, nil
}

func AddNumber(newNumber Number) error {

	c := DB.Collection(number_collection)
	_, err := c.InsertOne(context.TODO(), newNumber)

	if err != nil {
		return errors.New(err.Error())
	}

	return nil
}

func DeleteNumber(paramId string) error {
	id, err := primitive.ObjectIDFromHex(paramId)
	if err != nil {
		return errors.New(err.Error())
	}
	filter := bson.M{"_id": id}
	c := DB.Collection(number_collection)
	_, err = c.DeleteOne(context.TODO(), filter)

	if err != nil {
		return errors.New(err.Error())
	}

	return nil
}
