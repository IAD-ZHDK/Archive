package main

import (
	"github.com/gonfire/fire"
	"github.com/gonfire/fire/components"
	"github.com/gonfire/fire/jsonapi"
	"github.com/gonfire/fire/model"
)

func main() {
	// create store
	store := model.CreateStore("mongodb://localhost/archive")

	// create app
	app := fire.New()

	// create group
	group := jsonapi.New("api")

	// add controllers
	group.Add(&jsonapi.Controller{
		Model:      &documentation{},
		Store:      store,
		Authorizer: passwordAuthorizer(true),
		Validator:  documentationValidator,
	}, &jsonapi.Controller{
		Model:      &person{},
		Store:      store,
		Authorizer: passwordAuthorizer(false),
		Validator:  slugAndNameValidator,
	}, &jsonapi.Controller{
		Model:      &tag{},
		Store:      store,
		Authorizer: passwordAuthorizer(false),
		Validator:  slugAndNameValidator,
	})

	// mount group
	app.Mount(group)

	// mount hoster
	app.Mount(newHoster(store))

	// mount inspector
	app.Mount(fire.DefaultInspector(app))

	// mount protector
	app.Mount(components.DefaultProtector())

	// start app
	app.Start("localhost:8080")
}
