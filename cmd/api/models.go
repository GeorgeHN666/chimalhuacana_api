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

/*
{
	"name" : "George Hernandez",
	"loc" : "av union 234",
	"payment : "cash",
	"total" : "2001",
	"deliver" : "domicilio"
	"order" :  [{"parent":"Tacos","order_id":"ad51Pastor","item":"Pastor","menu_type":"Tacos de muerte lenta","price":12,"tortilla":"","quantity":"1","details":"","total":12},{"parent":"Tacos","order_id":"ad53Sirloin","item":"Sirloin","menu_type":"Tacos de muerte lenta","price":"35","action":"Tortilla","tortilla":"Queso","quantity":"1","details":"","total":35},{"parent":"Volcanes","order_id":"fd96Orden de Chuleta","item":"Orden de Chuleta","menu_type":"Tacos de muerte lenta","price":50,"tortilla":"","quantity":"1","details":"","total":50},{"parent":"Tortas","order_id":"gr61Pastor","item":"Pastor","menu_type":"Tacos de muerte lenta","price":"45","action":"Tipo","tortilla":"C/queso","quantity":"1","details":"","total":45},{"parent":"Gringas","order_id":"undefined2Suadero","item":"Suadero","menu_type":"Tacos de muerte lenta","price":60,"tortilla":"","quantity":"1","details":"","total":60},{"parent":"Gringas","order_id":"undefined3Bistec","item":"Bistec","menu_type":"Tacos de muerte lenta","price":60,"tortilla":"","quantity":"6","details":"","total":360},{"parent":"Suadero","order_id":"od51Orden de gorditas","item":"Orden de gorditas","menu_type":"Tacos de muerte lenta","price":"85","action":"Tipo","tortilla":"C/queso","quantity":"1","details":"","total":85},{"parent":"Especialidad de la casa","order_id":"pg81Parrillada","item":"Parrillada","menu_type":"Tacos de muerte lenta","price":180,"tortilla":"","quantity":"1","details":"","total":180},{"parent":"Tlayudas","order_id":"5s11Del Patio","item":"Del Patio","menu_type":"El Patio","price":150,"quantity":"1","details":"","ingredients":"","total":150},{"parent":"Tlayudas","order_id":"5s12Sencilla 1 ingrediente","item":"Sencilla 1 ingrediente","menu_type":"El Patio","price":100,"action":"Ingrediente","quantity":"1","details":"","ingredients":"Cecina","total":100},{"parent":"Tlayudas","order_id":"5s142 Ingredientes","item":"2 Ingredientes","menu_type":"El Patio","price":110,"action":"Ingredientes","quantity":"1","details":"","ingredients":"Tasajo,Cecina","total":110},{"parent":"Hamburguesas","order_id":"94f7Hawaiana mini","item":"Hawaiana mini","menu_type":"El Patio","price":50,"quantity":"1","details":"","ingredients":"","total":50},{"parent":"Pa' Botanear","order_id":"ds51Costillas","item":"Costillas","menu_type":"El Patio","price":"90","action":"tipo","quantity":"1","details":"","ingredients":"Mango,habanero","total":90},{"parent":"Pa' Botanear","order_id":"ds52Alitas","item":"Alitas","menu_type":"El Patio","price":70,"action":"tipo","quantity":"1","details":"","ingredients":"Mango,habanero","total":70},{"parent":"Guisados","order_id":"s8f1Chicharron en salsa verde","item":"Chicharron en salsa verde","menu_type":"Tacos de guisado","price":"50","action":"Tipo","tipo":"Plato","quantity":"1","details":"","total":50},{"parent":"Aguas","order_id":"ogf1Naranjada","item":"Naranjada","menu_type":"Bebidas","price":40,"tipo":"","quantity":"1","details":"","total":40},{"parent":"Aguas","order_id":"ogf5Jamaica","item":"Jamaica","menu_type":"Bebidas","price":"50","action":"Tamaño","tipo":"1 litro","quantity":"1","details":"","total":50},{"parent":"Del Barrio","order_id":"6f3z1Piña colada 1lt","item":"Piña colada 1lt","menu_type":"Bebidas","price":90,"tipo":"","quantity":"1","details":"","total":90},{"parent":"Postres","order_id":"6sg31Fresas con crema","item":"Fresas con crema","menu_type":"Postres","price":"45","action":"Tamano","postre":"","topping":"Chocolate liqudo","tipo":"Charola de 3","action_2":"topping","quantity":"1","details":"","postres":"","toppings":"","total":45},{"parent":"Rebanadas","order_id":"hy33Pay de limon","item":"Pay de limon","menu_type":"Postres","price":"60","action":"Tipo","postre":"Manzana con zanahoria","topping":"","tipo":"Con Postre","action_2":"Postre","quantity":"1","details":"","postres":"","toppings":"","total":60},{"parent":"Especiales","order_id":"av51Postre especial","item":"Postre especial","menu_type":"Postres","price":130,"action":"","postre":"","topping":"","tipo":"","quantity":"1","details":"","postres":"Fresas,con crema,Manzana con zanahoria,Durazno","toppings":"Lechera,Nuez","total":130},{"parent":"Especiales","order_id":"av52Platanos capeados","item":"Platanos capeados","menu_type":"Postres","price":40,"action":"Tipo","postre":"","topping":"","tipo":"","quantity":"1","details":"","postres":"","toppings":"","total":40},{"parent":"Hamburguesas","order_id":"94f3Del Gasper","item":"Del Gasper","menu_type":"El Patio","price":95,"quantity":"1","details":"","ingredients":"","total":95},{"parent":"Hamburguesas","order_id":"94f1Sencilla","item":"Sencilla","menu_type":"El Patio","price":70,"quantity":"1","details":"","ingredients":"","total":70}]
}


*/
