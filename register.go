package main

import (
	"net/http"

	"github.com/256dpi/fire"
	"github.com/256dpi/fire/coal"
	"github.com/256dpi/fire/flame"
	"github.com/256dpi/fire/wood"
	"github.com/goware/cors"
)

func handler(store *coal.Store, secret string, debug bool) http.Handler {
	// create mux
	mux := http.NewServeMux()

	// create reporter
	reporter := wood.DefaultErrorReporter()

	// create policy
	policy := flame.DefaultPolicy(secret)
	policy.PasswordGrant = true
	policy.AccessToken = &flame.AccessToken{}
	policy.RefreshToken = &flame.RefreshToken{}
	policy.Clients = []flame.Client{&flame.Application{}}
	policy.GrantStrategy = flame.DefaultGrantStrategy
	policy.ResourceOwners = func(c flame.Client) []flame.ResourceOwner {
		return []flame.ResourceOwner{&flame.User{}}
	}

	// create authenticator
	authenticator := flame.NewAuthenticator(store, policy)
	authenticator.Reporter = reporter

	// register authenticator
	mux.Handle("/auth/", authenticator.Endpoint("/auth/"))

	// create controller group
	g := fire.NewGroup()

	// set reporter
	g.Reporter = reporter

	// add all controllers
	g.Add(userController(store))
	g.Add(collectionController(store))
	g.Add(projectController(store))
	g.Add(personController(store))
	g.Add(tagController(store))

	// register group
	mux.Handle("/api/", fire.Compose(
		authenticator.Authorizer("", false, true, true),
		g.Endpoint("/api/"),
	))

	// compose handler
	handler := fire.Compose(
		wood.NewProtector("4M", cors.Options{
			AllowedOrigins: []string{"*"},
			AllowedHeaders: []string{"Origin", "Accept", "Content-Type", "Authorization"},
			AllowedMethods: []string{"GET", "POST", "PATCH", "DELETE"},
		}),
		mux,
	)

	// inject request logger in debug mode
	if debug {
		handler = fire.Compose(
			wood.DefaultRequestLogger(),
			handler,
		)
	}

	return handler
}
