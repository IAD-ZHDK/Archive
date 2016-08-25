package main

import (
	"github.com/256dpi/fire"
	"gopkg.in/mgo.v2/bson"
)

type documentation struct {
	fire.Base    `bson:",inline" fire:"documentation:documentations"`
	Slug         string  `json:"slug" valid:"required" bson:"slug" fire:"filterable,sortable"`
	MadekID      string  `json:"madek-id" valid:"required" bson:"madek_id"`
	MadekCoverID string  `json:"madek-cover-id" valid:"required" bson:"madek_cover_id"`
	Title        string  `json:"title" valid:"-"`
	Subtitle     string  `json:"subtitle" valid:"-"`
	Abstract     string  `json:"abstract" valid:"-"`
	Cover        *image  `json:"cover" valid:"-"`
	Images       []image `json:"images" valid:"-"`
	Videos       []video `json:"videos" valid:"-"`
	Documents    []file  `json:"documents" valid:"-"`
	Files        []file  `json:"files" valid:"-"`

	PeopleIds []bson.ObjectId `json:"-" bson:"people_ids" fire:"people:people"`
}

type file struct {
	Title    string `json:"title"`
	Stream   string `json:"stream"`
	Download string `json:"download"`
}

type image struct {
	file    `json:",inline" bson:",inline"`
	LowRes  string `json:"low-res" bson:"low_res"`
	HighRes string `json:"high-res" bson:"high_res"`
}

type video struct {
	image      `json:",inline" bson:",inline"`
	MP4Source  string `json:"mp4-source" bson:"mp4_source"`
	WebMSource string `json:"webm-source" bson:"webm_source"`
}

type person struct {
	fire.Base `bson:",inline" fire:"person:people"`
	Slug      string `json:"slug" valid:"-"`
	Name      string `json:"name" valid:"-"`

	Documentations fire.HasMany `json:"-" bson:"-" fire:"documentations:documentations"`
}
