package main

import (
	"github.com/256dpi/fire/coal"
	"gopkg.in/mgo.v2/bson"
)

var group = coal.NewGroup(&Project{}, &Person{}, &Tag{})

var indexer = coal.NewIndexer()

// Project holds the full madek data of one project.
type Project struct {
	coal.Base    `json:"-" bson:",inline" coal:"projects"`
	Slug         string `json:"slug" bson:"slug"`
	MadekID      string `json:"madek-id" bson:"madek_id"`
	MadekCoverID string `json:"madek-cover-id" bson:"madek_cover_id"`
	Published    bool   `json:"published"`

	Title     string  `json:"title"`
	Subtitle  string  `json:"subtitle"`
	Abstract  string  `json:"abstract"`
	Year      string  `json:"year"`
	Cover     *Image  `json:"cover"`
	Images    []Image `json:"images"`
	Videos    []Video `json:"videos"`
	Documents []File  `json:"documents"`
	Websites  []File  `json:"websites"`
	Files     []File  `json:"files"`

	Tags   []bson.ObjectId `json:"-" bson:"tag_ids" coal:"tags:tags"`
	People []bson.ObjectId `json:"-" bson:"people_ids" coal:"people:people"`
}

// File holds information about a downloadable file.
type File struct {
	Title    string `json:"title"`
	Stream   string `json:"stream"`
	Download string `json:"download"`
}

// Image holds information about a viewable image.
type Image struct {
	File    `json:",inline" bson:",inline"`
	LowRes  string `json:"low-res" bson:"low_res"`
	HighRes string `json:"high-res" bson:"high_res"`
}

// Video holds information about an viewable video.
type Video struct {
	Image      `json:",inline" bson:",inline"`
	MP4Source  string `json:"mp4-source" bson:"mp4_source"`
	WebMSource string `json:"webm-source" bson:"webm_source"`
}

// Person represents authors and collaborators on projects.
type Person struct {
	coal.Base `json:"-" bson:",inline" coal:"people"`
	Slug      string `json:"slug"`
	Name      string `json:"name"`

	Projects coal.HasMany `json:"-" bson:"-" coal:"projects:projects:people"`
}

// Tag is a single tag for categorizing projects.
type Tag struct {
	coal.Base `json:"-" bson:",inline" coal:"tags"`
	Slug      string `json:"slug"`
	Name      string `json:"name" `

	Projects coal.HasMany `json:"-" bson:"-" coal:"projects:projects:tags"`
}
