package main

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Newsletter struct {
	ID            primitive.ObjectID `json:"_id"  bson:"_id"`
	Name          string             `json:"name"  bson:"name"`
	Email         string             `json:"email" bson:"email"`
	TermsAccepted int                `json:"terms_accepted"  bson:"terms_accepted"`
	Created_At    time.Time          `json:"created_at"  bson:"created_at"`
}

type User struct {
	ID         primitive.ObjectID `json:"_id"  bson:"_id"`
	Name       string             `json:"name"  bson:"name"`
	Email      string             `json:"email" bson:"email"`
	ReCode     string             `json:"code"  bson:"code"`
	Valid      int                `json:"valid"  bson:"valid"`
	Role       int                `json:"role"  bson:"role"`
	Checked    int                `json:"checked"  bson:"checked"`
	Created_By string             `json:"created_by"  bson:"created_by"`
	Created_At time.Time          `json:"created_at"  bson:"created_at"`
}

type Order struct {
	Name    string    `json:"name"`
	Loc     string    `json:"loc,omitempty"`
	Payment string    `json:"payment"`
	Total   int       `json:"total"`
	Deliver string    `json:"deliver"`
	List    []Details `json:"order"`
}

type Details struct {
	Parent      string `json:"parent"`
	Item        string `json:"item"`
	MenuType    string `json:"menu_type"`
	Tortilla    string `json:"tortilla"`
	Quantity    string `json:"quantity"`
	Details     string `json:"details"`
	Total       int    `json:"total"`
	Ingredients string `json:"ingredients"`
	Action      string `json:"action"`
	Tipo        string `json:"tipo"`
	Topping     string `json:"topping"`
	Action2     string `json:"action_2"`
	Postres     string `json:"postres"`
	Toppings    string `json:"toppings"`
	Postre      string `json:"postre"`
}
