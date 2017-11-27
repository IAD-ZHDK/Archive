package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/256dpi/fire/coal"
	"github.com/256dpi/fire/flame"
)

var debug = os.Getenv("DEBUG") == "yes"

// TODO: Resize images on the fly.

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
	mux.Handle("/", handler(store, os.Getenv("SECRET"), debug))

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
	err = flame.EnsureFirstUser(store, "Root", "root@archive.iad.zhdk.ch", "root")
	if err != nil {
		return err
	}

	// ensure admin application
	err = flame.EnsureApplication(store, "Admin")
	if err != nil {
		return err
	}

	// admin application key
	adminAppKey, err := flame.GetApplicationKey(store, "Admin")
	if err != nil {
		return err
	}

	// print main application keys
	fmt.Printf("Admin Application Key: %s\n", adminAppKey)

	return nil
}
