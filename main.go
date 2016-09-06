package main

import (
	"time"

	"github.com/gonfire/fire"
	"github.com/gin-gonic/gin"
	"github.com/itsjamie/gin-cors"
	"gopkg.in/mgo.v2"
)

func main() {
	sess, err := mgo.Dial("mongodb://localhost/archive")
	if err != nil {
		panic(err)
	}

	defer sess.Close()

	db := sess.DB("")

	router := gin.Default()

	app := fire.New(db, "")

	// TODO: Add authentication and protect resources.

	app.Mount(&fire.Controller{
		Model: &documentation{},
		Validator: fire.Combine(
			madekDataValidator,
		),
	}, &fire.Controller{
		Model: &person{},
	}, &fire.Controller{
		Model: &tag{},
	})

	router.Use(cors.Middleware(cors.Config{
		Origins:        "*",
		Methods:        "GET, POST, PUT, PATCH, DELETE",
		RequestHeaders: "Origin, Authorization, Content-Type, Accept, Cache-Control, X-Requested-With",
		ExposedHeaders: "",
		MaxAge:         time.Minute,
		Credentials:    true,
	}))

	app.Register(router)

	err = router.Run("localhost:8080")
	if err != nil {
		panic(err)
	}
}
