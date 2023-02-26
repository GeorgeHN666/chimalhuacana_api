package main

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (app *application) ConnectDB() *mongo.Client {

	var opt = options.Client().ApplyURI(app.config.db.URI)

	client, err := mongo.Connect(context.TODO(), opt)
	if err != nil {
		app.ErrorL.Fatalln("Couldn´t connect with db")
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		app.ErrorL.Fatal("No connection with database")
	}

	return client

}

func (app *application) AddUser(u User) error {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := app.ConnectDB().Database(app.config.db.Database).Collection("users")

	u.ID = primitive.NewObjectID()
	u.Created_At = time.Now()
	u.ReCode, _ = app.GenerateCode()
	u.Valid = 0
	u.Checked = 0

	_, err := db.InsertOne(ctx, u)
	if err != nil {
		return err
	}

	return nil
}

func (app *application) DeleteUser(u User, A User) error {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := app.ConnectDB().Database(app.config.db.Database).Collection("users")

	if A.Role < 3 {
		return errors.New("error")
	}

	filter := bson.M{
		"_id": u.ID,
	}

	_, err := db.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return nil
}

func (app *application) CheckUser(u User) (User, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := app.ConnectDB().Database(app.config.db.Database).Collection("users")

	filter := bson.M{
		"email": bson.M{"$eq": u.Email},
	}

	var result User

	err := db.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return result, err
	}

	return result, nil
}

func (app *application) AddNewsletter(n Newsletter) error {

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	db := app.ConnectDB().Database(app.config.db.Database).Collection("boletin")

	n.ID = primitive.NewObjectID()
	n.TermsAccepted = 1

	_, err := db.InsertOne(ctx, n)
	if err != nil {
		return err
	}
	return nil

}

func (app *application) FindNewsletterUser(e string) (Newsletter, bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	db := app.ConnectDB().Database(app.config.db.Database).Collection("boletin")

	filter := bson.M{
		"email": bson.M{"$eq": e},
	}

	var result Newsletter

	err := db.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return result, false, err
	}
	return result, true, nil
}

func (app *application) DeleteNewsletterUser(u User) error {

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	db := app.ConnectDB().Database(app.config.db.Database).Collection("newsletter")

	filter := bson.M{
		"_id": u.ID,
	}

	_, err := db.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	return nil

}

func (app *application) GetOrderStr(o Order) (string, error) {

	orderN, _ := app.GenerateOrderNumber()

	var tempString []string

	for _, v := range o.List {

		if v.MenuType == "Tacos de muerte lenta" {
			st := fmt.Sprintf("- %v <%v> %v %v (%v)", v.Quantity, v.Tortilla, v.Parent, v.Item, v.Details)

			tempString = append(tempString, st)
			st = ""
		}
		if v.MenuType == "El Patio" {
			st := fmt.Sprintf("- %v <%v:%v> %v %v (%v)", v.Quantity, v.Action, v.Ingredients, v.Parent, v.Item, v.Details)
			tempString = append(tempString, st)
			st = ""
		}
		if v.MenuType == "Tacos de guisado" {
			st := fmt.Sprintf("- %v <%v:%v> %v %v (%v)", v.Quantity, v.Action, v.Tipo, v.Parent, v.Item, v.Details)
			tempString = append(tempString, st)
			st = ""
		}
		if v.MenuType == "Bebidas" {
			st := fmt.Sprintf("- %v <%v:%v> %v %v (%v)", v.Quantity, v.Action, v.Tipo, v.Parent, v.Item, v.Details)
			tempString = append(tempString, st)
			st = ""

		}
		if v.MenuType == "Postres" {
			st := fmt.Sprintf("-$%v <Tipo:%v Postre:%v %v toppings:%v %v > %v %v (%v)", v.Quantity, v.Tipo, v.Postre, v.Postres, v.Toppings, v.Topping, v.Parent, v.Item, v.Details)
			tempString = append(tempString, st)
			st = ""
		}

	}

	OR := strings.Join(tempString, "%0a")

	final := fmt.Sprintf("%v %0a Numero de orden: %v %0a <%v> %0a *%v %0a -%v %0a Total: %v %0a %0a ---ORDEN---%0a %v", o.Name, orderN, o.Deliver, o.Loc, o.Payment, o.Total, OR)

	if len(final) < 1 {
		return final, errors.New("bad string")
	}

	return final, nil

}
