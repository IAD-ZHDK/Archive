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
	policy.GrantStrategy = grantStrategy

	// set resource owner callback
	policy.ResourceOwners = func(c flame.Client) []flame.ResourceOwner {
		return []flame.ResourceOwner{&flame.User{}}
	}

	// set data for token callback
	policy.TokenData = func(c flame.Client, ro flame.ResourceOwner) map[string]interface{} {
		return nil
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
	g.Add(documentationController(store))
	g.Add(personController(store))
	g.Add(tagController(store))

	// register group
	mux.Handle("/api/", fire.Compose(
		authenticator.Authorizer("", false),
		extendedAuthorizer(store, reporter),
		g.Endpoint("/api/"),
	))

	// create protector
	protector := wood.NewProtector("8M", cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedHeaders: []string{"Origin", "Accept", "Content-Type",
			"Authorization", "Cache-Control", "X-Requested-With"},
		AllowedMethods: []string{"GET", "POST", "PATCH", "DELETE"},
	})

	// compose handler
	handler := fire.Compose(
		protector,
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
