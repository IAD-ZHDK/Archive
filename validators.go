package main

import (
	"os"

	"github.com/256dpi/fire"
	"github.com/IAD-ZHDK/madek"
)

var client *madek.Client

func init() {
	client = madek.NewClient(
		os.Getenv("MADEK_ADDRESS"),
		os.Getenv("MADEK_USERNAME"),
		os.Getenv("MADEK_PASSWORD"),
	)
}

func madekDataValidator(ctx *fire.Context) error {
	if ctx.Action != fire.Create && ctx.Action != fire.Update {
		return nil
	}

	doc := ctx.Model.(*documentation)

	set, err := client.CompileSet(doc.MadekSet)
	if err != nil {
		return fire.Fatal(err)
	}

	// TODO: We should strip the madek dump and only save necessary data.

	doc.MadekData = set

	return nil
}
