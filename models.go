package main

import "github.com/256dpi/fire"

type documentation struct {
	fire.Base `bson:",inline" fire:"documentation:documentations"`
	Slug      string  `json:"slug" valid:"required" bson:"slug" fire:"filterable,sortable"`
	Title     string  `json:"title" valid:"required"`
	MadekSet  string  `json:"madek-set" valid:"required" bson:"madek_set"`
	Images    []image `json:"images" valid:"-"`
	Videos    []video `json:"videos" valid:"-"`
	Files     []file  `json:"files" valid:"-"`
}

type image struct {
	Title   string `json:"title"`
	LowRes  string `json:"low-res" bson:"low_res"`
	HighRes string `json:"high-res" bson:"high_res"`
}

type video struct {
	image      `json:",inline" bson:",inline"`
	MP4Source  string `json:"mp4-source" bson:"mp4_source"`
	WebMSource string `json:"webm-source" bson:"webm_source"`
}

type file struct {
	Title    string `json:"title"`
	Download string `json:"download"`
}
