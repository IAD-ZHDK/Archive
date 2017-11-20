package main

import (
	"github.com/256dpi/fire"
	"github.com/256dpi/fire/coal"
)

func documentationController(store *coal.Store) *fire.Controller {
	return &fire.Controller{
		Model:       &Documentation{},
		Store:       store,
		Authorizers: nil,
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
		Authorizers: nil,
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
		Authorizers: nil,
		Validators: fire.L(
			fire.ModelValidator(),
			fire.RelationshipValidator(&Tag{}, group),
			slugAndNameValidator,
		),
	}
}
