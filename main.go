package main

import (
	"github.com/gonfire/fire"
	"github.com/gonfire/fire/components"
	"github.com/gonfire/fire/jsonapi"
)

func main() {
	// create pool
	pool := fire.NewPool("mongodb://localhost/archive")

	// create app
	app := fire.New()

	// TODO: Add basic authentication and protect resources.

	// create group
	group := jsonapi.New("api")

	// add controllers
	group.Add(&jsonapi.Controller{
		Model: &documentation{},
		Pool:  pool,
		Validator: jsonapi.Combine(
			madekDataValidator,
		),
	}, &jsonapi.Controller{
		Model: &person{},
		Pool:  pool,
	}, &jsonapi.Controller{
		Model: &tag{},
		Pool:  pool,
	})

	// mount group
	app.Mount(group)

	// mount hoster
	app.Mount(newHoster(pool))

	// mount inspector
	app.Mount(fire.DefaultInspector(app))

	// mount protector
	app.Mount(components.DefaultProtector())

	// start app
	app.Start("localhost:8080")
}
