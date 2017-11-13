package main

import (
	"github.com/gonfire/fire/model"
	"gopkg.in/mgo.v2/bson"
)

type documentation struct {
	model.Base   `json:"-" bson:",inline" fire:"documentations"`
	Slug         string `json:"slug" bson:"slug" fire:"filterable,sortable"`
	MadekID      string `json:"madek-id" bson:"madek_id"`
	MadekCoverID string `json:"madek-cover-id" bson:"madek_cover_id"`
	Published    bool   `json:"published" fire:"filterable"`

	Title     string  `json:"title"`
	Subtitle  string  `json:"subtitle"`
	Abstract  string  `json:"abstract"`
	Year      string  `json:"year" fire:"filterable,sortable"`
	Cover     *Image  `json:"cover"`
	Images    []Image `json:"images"`
	Videos    []video `json:"videos"`
	Documents []File  `json:"documents"`
	Websites  []File  `json:"websites"`
	Files     []File  `json:"files"`

	TagIDs    []bson.ObjectId `json:"-" bson:"tag_ids" fire:"tags:tags"`
	PeopleIDs []bson.ObjectId `json:"-" bson:"people_ids" fire:"people:people"`
}

type File struct {
	Title    string `json:"title"`
	Stream   string `json:"stream"`
	Download string `json:"download"`
}

type Image struct {
	File           `json:",inline" bson:",inline"`
	LowRes  string `json:"low-res" bson:"low_res"`
	HighRes string `json:"high-res" bson:"high_res"`
}

type video struct {
	Image             `json:",inline" bson:",inline"`
	MP4Source  string `json:"mp4-source" bson:"mp4_source"`
	WebMSource string `json:"webm-source" bson:"webm_source"`
}

type person struct {
	model.Base `json:"-" bson:",inline" fire:"people"`
	Slug       string `json:"slug" fire:"filterable"`
	Name       string `json:"name"`

	Documentations model.HasMany `json:"-" bson:"-" fire:"documentations:documentations:people"`
}

type tag struct {
	model.Base `json:"-" bson:",inline" fire:"tags"`
	Slug       string `json:"slug" fire:"filterable"`
	Name       string `json:"name" `

	Documentations model.HasMany `json:"-" bson:"-" fire:"documentations:documentations:tags"`
}
