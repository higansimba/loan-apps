package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserRole string

const (
	UserRoleAdmin   UserRole = "admin"
	UserRoleNasabah UserRole = "nasabah"
	UserRoleOfficer UserRole = "officer"
)

type User struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	Nama      string             `json:"nama" bson:"nama"`
	NoKTP     string             `json:"no_ktp" bson:"no_ktp"`
	NoHP      string             `json:"no_hp" bson:"no_hp"`
	Email     string             `json:"email" bson:"email"`
	Alamat    string             `json:"alamat" bson:"alamat"`
	Password  string             `json:"password" bson:"password"`
	Role      UserRole           `json:"role" bson:"role"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
}

func (n *User) CollectionName() string {
	return CollectionUsers.String()
}
