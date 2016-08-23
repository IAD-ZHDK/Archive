package main

import (
	"github.com/256dpi/fire"
	"github.com/IAD-ZHDK/madek"
)

type documentation struct {
	fire.Base `bson:",inline" fire:"documentation:documentations"`
	Slug      string     `json:"slug" valid:"required" bson:"slug" fire:"filterable,sortable"`
	Title     string     `json:"title" valid:"required"`
	MadekSet  string     `json:"madek-set" valid:"required" bson:"madek_set"`
	MadekData *madek.Set `json:"madek-data" valid:"-" bson:"madek_data"`
}
