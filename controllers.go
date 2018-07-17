package main

import (
	"github.com/256dpi/fire"
	"github.com/256dpi/fire/coal"
	"github.com/256dpi/fire/flame"
)

// TODO: Require Admin privileges to change stuff.

func userController(store *coal.Store) *fire.Controller {
	return &fire.Controller{
		Model:       &flame.User{},
		Store:       store,
		Authorizers: fire.L{},
	}
}

func collectionController(store *coal.Store) *fire.Controller {
	return &fire.Controller{
		Model:       &Collection{},
		Store:       store,
		Authorizers: fire.L{},
		Validators: fire.L{
			fire.RelationshipValidator(&Collection{}, catalog),
		},
	}
}

func projectController(store *coal.Store) *fire.Controller {
	return &fire.Controller{
		Model:       &Project{},
		Store:       store,
		Authorizers: fire.L{},
		Filters:     []string{"Published", "Year", "Slug"},
		Validators: fire.L{
			fire.RelationshipValidator(&Project{}, catalog),
			projectValidator(),
		},
	}
}

func personController(store *coal.Store) *fire.Controller {
	return &fire.Controller{
		Model:       &Person{},
		Store:       store,
		Authorizers: fire.L{},
		Validators: fire.L{
			fire.RelationshipValidator(&Person{}, catalog),
			slugAndNameValidator(),
		},
	}
}

func tagController(store *coal.Store) *fire.Controller {
	return &fire.Controller{
		Model:       &Tag{},
		Store:       store,
		Authorizers: fire.L{},
		Validators: fire.L{
			fire.RelationshipValidator(&Tag{}, catalog),
			slugAndNameValidator(),
		},
	}
}
