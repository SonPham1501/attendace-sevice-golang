package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Account represents the details of single account
type Account struct {
	Id       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Role     int64              `json:"role"`
	Email    string             `json:"email"`
	Password string             `json:"password"`
}

// AccountListResponse
type AccountListResponse struct {
	Data  []Account `json:"data"`
	Total int       `json:"total"`
}

// AccountResponse
type AccountResponse struct {
	Data Account `json:"data"`
}
