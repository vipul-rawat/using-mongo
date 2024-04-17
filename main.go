package main

import (
	mongo "github.com/vipul-rawat/gofr-mongo"
	"go.mongodb.org/mongo-driver/bson"
	"gofr.dev/pkg/gofr"
)

type Person struct {
	Name string `bson:"name" json:"name"`
	Age  int    `bson:"age" json:"age"`
	City string `bson:"city" json:"city"`
}

func main() {
	app := gofr.New()

	db := mongo.New(app.Config, app.Logger(), app.Metrics())
	app.UseMongo(db)

	app.POST("/mongo", Insert)
	app.GET("/mongo", Get)

	app.Run()
}

func Insert(ctx *gofr.Context) (interface{}, error) {
	var p Person
	err := ctx.Bind(&p)
	if err != nil {
		return nil, err
	}

	res, err := ctx.Mongo.InsertOne(ctx, "collection", p)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func Get(ctx *gofr.Context) (interface{}, error) {
	var result Person

	p := ctx.Param("name")

	err := ctx.Mongo.FindOne(ctx, "collection", bson.D{{"name", p}} /* valid filter */, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
