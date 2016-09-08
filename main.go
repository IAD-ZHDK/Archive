package main

import "github.com/gonfire/fire"

func main() {
	app := fire.New("mongodb://localhost/archive", "")

	// TODO: Add basic authentication and protect resources.

	app.EnableDevMode()

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

	app.Start("localhost:8080")
}
