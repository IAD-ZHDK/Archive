package main

import (
	"time"

	"github.com/256dpi/fire"
	"github.com/gin-gonic/gin"
	"github.com/itsjamie/gin-cors"
	"gopkg.in/mgo.v2"
)

func main() {
	// connect to mongodb
	sess, err := mgo.Dial("mongodb://localhost/archive")
	if err != nil {
		panic(err)
	}

	// close session when finished
	defer sess.Close()

	// get database
	db := sess.DB("")

	// get a router
	router := gin.Default()

	// prepare endpoint
	endpoint := fire.NewEndpoint(db)

	// register documentation resource
	endpoint.AddResource(&fire.Resource{
		Model: &documentation{},
	})

	// setup cors
	router.Use(cors.Middleware(cors.Config{
		Origins:        "*",
		Methods:        "GET, POST, PUT, PATCH, DELETE",
		RequestHeaders: "Origin, Authorization, Content-Type, Accept, Cache-Control, X-Requested-With",
		ExposedHeaders: "",
		MaxAge:         time.Minute,
		Credentials:    true,
	}))

	// register endpoint on router
	endpoint.Register("", router)

	// run server
	err = router.Run("localhost:8080")
	if err != nil {
		panic(err)
	}
}
