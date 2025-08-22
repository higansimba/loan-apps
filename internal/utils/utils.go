package utils

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func StringToObjectID(str string) primitive.ObjectID {
	id, _ := primitive.ObjectIDFromHex(str)
	return id
}

func IsValidObjectID(str string) bool {
	_, err := primitive.ObjectIDFromHex(str)
	return err == nil
}
