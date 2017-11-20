package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/256dpi/fire/coal"
	"github.com/256dpi/fire/flame"
	"github.com/pborman/uuid"
	"gopkg.in/mgo.v2/bson"
)

var debug = os.Getenv("DEBUG") == "yes"

func main() {
	// create store
	store := coal.MustCreateStore(os.Getenv("MONGODB_URI"))

	// prepare database
	err := prepareDatabase(store)
	if err != nil {
		panic(err)
	}

	// create main mux
	mux := http.NewServeMux()

	// build v1 api handler
	mux.Handle("/v1/", handler(store, os.Getenv("SECRET"), debug))

	// get port
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	if port == 0 {
		port = 8000
	}

	// run plain server
	fmt.Printf("Running on http://0.0.0.0:%d\n", port)
	err = http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", port), mux)
	if err != nil {
		panic(err)
	}
}

func prepareDatabase(store *coal.Store) error {
	// ensure indexes
	err := indexer.Ensure(store)
	if err != nil {
		return err
	}

	// ensure first user
	err = ensureFirstUser(store)
	if err != nil {
		return err
	}

	// ensure admin application
	err = ensureApplication(store, "Admin")
	if err != nil {
		return err
	}

	// admin application key
	adminAppKey, err := getApplicationKey(store, "Admin")
	if err != nil {
		return err
	}

	// print main application keys
	fmt.Printf("Admin Application Key: %s\n", adminAppKey)

	return nil
}

func ensureFirstUser(store *coal.Store) error {
	// TODO: Move to flame.

	// copy store
	s := store.Copy()
	defer s.Close()

	// count super users
	n, err := s.C(&flame.User{}).Count()
	if err != nil {
		return err
	}

	// check existence
	if n > 0 {
		return nil
	}

	// create super user
	user := coal.Init(&flame.User{}).(*flame.User)
	user.Name = "Root"
	user.Email = "root@archive.iad.zhdk.ch"
	user.Password = "root"

	// set key and secret
	err = user.Validate()
	if err != nil {
		return err
	}

	// save super user
	return s.C(user).Insert(user)
}

func ensureApplication(store *coal.Store, name string) error {
	// TODO: Move to flame.

	// copy store
	s := store.Copy()
	defer s.Close()

	// count main applications
	var apps []flame.Application
	err := s.C(&flame.Application{}).Find(bson.M{
		coal.F(&flame.Application{}, "Name"): name,
	}).All(&apps)
	if err != nil {
		return err
	}

	// check count
	if len(apps) > 1 {
		return errors.New("to many applications with that name")
	}

	// application is missing

	// create application
	app := coal.Init(&flame.Application{}).(*flame.Application)
	app.Name = name
	app.Key = uuid.New()
	app.Secret = uuid.New()

	// validate model
	err = app.Validate()
	if err != nil {
		return err
	}

	// save application
	err = s.C(app).Insert(app)
	if err != nil {
		return err
	}

	return nil
}

func getApplicationKey(store *coal.Store, name string) (string, error) {
	// TODO: Move to flame.

	// copy store
	s := store.Copy()
	defer s.Close()

	// get application
	var app flame.Application
	err := s.C(&app).Find(bson.M{
		"name": name,
	}).One(&app)
	if err != nil {
		return "", err
	}

	return app.Key, nil
}
