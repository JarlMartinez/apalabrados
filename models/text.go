package models

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Text struct {
	ID    primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Value string             `json:"value" bson:"value,omitempty"`
}

func GetTexts() ([]Text, error) {
	var texts []Text
	c := DB.Collection(text_collection)
	curr, err := c.Find(context.TODO(), bson.M{})

	if err != nil {
		return []Text{}, errors.New("error getting texts from db")
	}

	defer curr.Close(context.TODO())

	for curr.Next(context.TODO()) {
		var text Text
		if err := curr.Decode(&text); err != nil {
			return []Text{}, errors.New("error deserealizing text entry")
		}
		texts = append(texts, text)
	}

	if err = curr.Err(); err != nil {
		return []Text{}, errors.New("current connection presented some error")
	}

	return texts, nil
}

func AddText(newText Text) error {

	c := DB.Collection(text_collection)
	_, err := c.InsertOne(context.TODO(), newText)

	if err != nil {
		return errors.New(err.Error())
	}

	return nil
}

func DeleteText(paramId string) error {
	id, err := primitive.ObjectIDFromHex(paramId)
	if err != nil {
		return errors.New(err.Error())
	}
	filter := bson.M{"_id": id}
	c := DB.Collection(text_collection)
	_, err = c.DeleteOne(context.TODO(), filter)

	if err != nil {
		return errors.New(err.Error())
	}

	return nil
}
