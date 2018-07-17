package main

import (
	"github.com/256dpi/fire/coal"
	"gopkg.in/mgo.v2/bson"
)

var catalog = coal.NewCatalog(&Collection{}, &Project{}, &Person{}, &Tag{})

var indexer = coal.NewIndexer()

func init() {
	indexer.Add(&Collection{}, true, false, "Slug")
	indexer.Add(&Project{}, true, false, "Slug")
	indexer.Add(&Project{}, false, false, "Published")
	indexer.Add(&Person{}, true, false, "Slug")
	indexer.Add(&Tag{}, true, false, "Slug")
}

// A Collection groups multiple projects.
type Collection struct {
	coal.Base `json:"-" bson:",inline" coal:"collections"`
	Slug      string          `json:"slug"`
	Name      string          `json:"name"`
	Projects  []bson.ObjectId `json:"-" bson:"project_ids" coal:"projects:projects"`
}

// A Project holds the full madek data of one project.
type Project struct {
	coal.Base    `json:"-" bson:",inline" coal:"projects"`
	Slug         string `json:"slug"`
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

	Tags        []bson.ObjectId `json:"-" bson:"tag_ids" coal:"tags:tags"`
	People      []bson.ObjectId `json:"-" bson:"people_ids" coal:"people:people"`
	Collections coal.HasMany    `json:"-" bson:"-" coal:"collections:collections:projects"`
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
