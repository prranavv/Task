package models

import "time"

type User struct {
	ID          string    `json:"id" bson:"_id"`
	Name        string    `json:"name" bson:"name"`
	DOB         time.Time `json:"dob" bson:"dob"`
	Address     string    `json:"address" bson:"address"`
	Description string    `json:"description" bson:"description"`
	CreatedAt   time.Time `json:"createdAt" bson:"createdAt"`
}
