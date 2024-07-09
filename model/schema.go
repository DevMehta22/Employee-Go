package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Employee struct {
	EmpID        primitive.ObjectID    `json:"_id,omitempty" bson:"_id,omitempty"`
	EmpName      string `json:"name,omitempty"`
	EmpDept     string `json:"dept,omitempty"`
}