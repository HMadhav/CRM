package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Customer struct {
	Id        primitive.ObjectID `json:"id"`
	Name      string             `json:"name"`
	Role      string             `json:"role"`
	Email     string             `json:"email"`
	Phone     string             `json:"Phone"`
	Contacted bool               `json:"contacted"`
}

type CustomerResponse struct {
	Status  int                    `json:"status"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}
