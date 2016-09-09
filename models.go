package main

import (
	"errors"

	"github.com/gonfire/fire"
	"gopkg.in/mgo.v2/bson"
)

type documentation struct {
	fire.Base    `bson:",inline" fire:"documentations"`
	Slug         string `json:"slug" bson:"slug" fire:"filterable,sortable"`
	MadekID      string `json:"madek-id" bson:"madek_id"`
	MadekCoverID string `json:"madek-cover-id" bson:"madek_cover_id"`
	Published    bool   `json:"published" fire:"filterable"`

	Title     string  `json:"title"`
	Subtitle  string  `json:"subtitle"`
	Abstract  string  `json:"abstract"`
	Year      string  `json:"year" fire:"filterable,sortable"`
	Cover     *image  `json:"cover"`
	Images    []image `json:"images"`
	Videos    []video `json:"videos"`
	Documents []file  `json:"documents"`
	Websites  []file  `json:"websites"`
	Files     []file  `json:"files"`

	TagIDs    []bson.ObjectId `json:"-" bson:"tag_ids" fire:"tags:tags"`
	PeopleIDs []bson.ObjectId `json:"-" bson:"people_ids" fire:"people:people"`
}

func (d *documentation) Validate(fresh bool) error {
	if d.Published && d.Slug == "" {
		return errors.New("Missing slug for published documentation")
	}

	return d.Base.Validate(fresh)
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
	fire.Base `bson:",inline" fire:"people"`
	Slug      string `json:"slug" fire:"filterable"`
	Name      string `json:"name"`

	Documentations fire.HasMany `json:"-" bson:"-" fire:"documentations:documentations:people"`
}

type tag struct {
	fire.Base `bson:",inline" fire:"tags"`
	Slug      string `json:"slug" fire:"filterable"`
	Name      string `json:"name"`

	Documentations fire.HasMany `json:"-" bson:"-" fire:"documentations:documentations:tags"`
}
