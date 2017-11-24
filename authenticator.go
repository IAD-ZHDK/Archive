package main

import (
	"context"
	"errors"
	"net/http"

	"github.com/256dpi/fire/coal"
	"github.com/256dpi/fire/flame"
	"github.com/256dpi/oauth2"
)

type contextKey int

const (
	accessTokenContextKey contextKey = iota
	applicationContextKey
	userContextKey
	candidateContextKey
)

func grantStrategy(scope oauth2.Scope, c flame.Client, ro flame.ResourceOwner) (oauth2.Scope, error) {
	// check scope
	if !scope.Empty() {
		return nil, flame.ErrInvalidScope
	}

	// grant empty scope to users
	if _, ok := ro.(*flame.User); ok {
		return scope, nil
	}

	// reject grant by default
	return nil, flame.ErrGrantRejected
}

func extendedAuthorizer(store *coal.Store, reporter func(error)) func(http.Handler) http.Handler {
	// TODO: Move to flame?

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// get current context
			ctx := r.Context()

			// get access token
			value := ctx.Value(flame.AccessTokenContextKey)
			if value == nil {
				// call next handler
				next.ServeHTTP(w, r.WithContext(ctx))
				return
			}

			// cast access token
			accessToken := value.(*flame.AccessToken)

			// save access token
			ctx = context.WithValue(ctx, accessTokenContextKey, accessToken)

			// copy store
			s := store.Copy()
			defer s.Close()

			// load application
			var application flame.Application
			err := s.C(&application).FindId(accessToken.ClientID).One(&application)
			if err != nil {
				reporter(err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			// store application
			ctx = context.WithValue(ctx, applicationContextKey, coal.Init(&application))
			ctx = context.WithValue(ctx, candidateContextKey, applicationContextKey)

			// report error if resource owner is missing
			if accessToken.ResourceOwnerID == nil {
				reporter(errors.New("missing resource owner id"))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			// attempt to load user
			var user flame.User
			err = s.C(&user).FindId(*accessToken.ResourceOwnerID).One(&user)
			if err != nil {
				reporter(err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			// store user
			ctx = context.WithValue(ctx, userContextKey, coal.Init(&user))
			ctx = context.WithValue(ctx, candidateContextKey, userContextKey)

			// call next handler
			next.ServeHTTP(w, r.WithContext(ctx))
			return
		})
	}
}
