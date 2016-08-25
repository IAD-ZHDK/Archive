package main

import (
	"strings"

	"github.com/metal3d/go-slugify"
)

var replacer = strings.NewReplacer(
	"ä", "ae",
	"ö", "oe",
	"ü", "ue",
)

func makeSlug(str string) string {
	str = replacer.Replace(str)
	return slugify.Marshal(str, true)
}
