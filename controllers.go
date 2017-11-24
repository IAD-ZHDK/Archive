package main

import (
	"github.com/256dpi/fire"
	"github.com/256dpi/fire/coal"
	"github.com/256dpi/fire/flame"
)

// TODO: Require Admin privileges to change stuff.

func userController(store *coal.Store) *fire.Controller {
	return &fire.Controller{
		Model: &flame.User{},
		Store: store,
		Authorizers: fire.L(),
		Validators: fire.L(
			fire.ModelValidator(),
		),
	}
}

func documentationController(store *coal.Store) *fire.Controller {
	return &fire.Controller{
		Model:       &Documentation{},
		Store:       store,
		Authorizers: fire.L(),
		Validators: fire.L(
			fire.ModelValidator(),
			fire.RelationshipValidator(&Tag{}, group),
			documentationValidator,
		),
	}
}

func personController(store *coal.Store) *fire.Controller {
	return &fire.Controller{
		Model:       &Person{},
		Store:       store,
		Authorizers: fire.L(),
		Validators: fire.L(
			fire.ModelValidator(),
			fire.RelationshipValidator(&Tag{}, group),
			slugAndNameValidator,
		),
	}
}

func tagController(store *coal.Store) *fire.Controller {
	return &fire.Controller{
		Model:       &Tag{},
		Store:       store,
		Authorizers: fire.L(),
		Validators: fire.L(
			fire.ModelValidator(),
			fire.RelationshipValidator(&Tag{}, group),
			slugAndNameValidator,
		),
	}
}
