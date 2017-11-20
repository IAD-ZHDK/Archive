package main

import (
	"github.com/256dpi/fire/coal"
	"gopkg.in/mgo.v2/bson"
)

var group = coal.NewGroup(&Documentation{}, &Person{}, &Tag{})

var indexer = coal.NewIndexer()

// Documentation holds the full documentation info.
type Documentation struct {
	coal.Base    `json:"-" bson:",inline" coal:"documentations"`
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

// Person represents authors and collaborators on documentations.
type Person struct {
	coal.Base `json:"-" bson:",inline" coal:"people"`
	Slug      string `json:"slug"`
	Name      string `json:"name"`

	Documentations coal.HasMany `json:"-" bson:"-" coal:"documentations:documentations:people"`
}

// Tag is a single tag for categorizing documentations.
type Tag struct {
	coal.Base `json:"-" bson:",inline" coal:"tags"`
	Slug      string `json:"slug"`
	Name      string `json:"name" `

	Documentations coal.HasMany `json:"-" bson:"-" coal:"documentations:documentations:tags"`
}
