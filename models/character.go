package models

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Chartacter struct {
	ID    primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Value string             `json:"value" bson:"value,omitempty"`
}

func GetCharacters() ([]Chartacter, error) {
	var chars []Chartacter
	c := DB.Collection(character_collection)
	curr, err := c.Find(context.TODO(), bson.M{})

	if err != nil {
		return []Chartacter{}, errors.New("error getting characters from db")
	}

	defer curr.Close(context.TODO())

	for curr.Next(context.TODO()) {
		var char Chartacter
		if err := curr.Decode(&char); err != nil {
			return []Chartacter{}, errors.New("error deserealizing Character entry")
		}
		chars = append(chars, char)
	}

	if err = curr.Err(); err != nil {
		return []Chartacter{}, errors.New("current connection presented some error")
	}

	return chars, nil
}

func AddCharacter(newChar Chartacter) error {

	c := DB.Collection(character_collection)
	_, err := c.InsertOne(context.TODO(), newChar)

	if err != nil {
		return errors.New(err.Error())
	}

	return nil
}

func DeleteCharacter(paramId string) error {
	id, err := primitive.ObjectIDFromHex(paramId)
	if err != nil {
		return errors.New(err.Error())
	}
	filter := bson.M{"_id": id}
	c := DB.Collection(character_collection)
	_, err = c.DeleteOne(context.TODO(), filter)

	if err != nil {
		return errors.New(err.Error())
	}

	return nil
}
