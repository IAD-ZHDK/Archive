package main

import (
	"os"
	"strings"

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

	doc.Title = set.MetaData["title"]
	doc.Subtitle = set.MetaData["subtitle"]

	doc.Cover = nil
	doc.Videos = nil
	doc.Images = nil
	doc.Documents = nil
	doc.Files = nil

	// TODO: Check if madek copyright field is correct.

	for _, mediaEntry := range set.MediaEntries {
		_file := file{
			Title:    mediaEntry.MetaData["title"],
			Stream:   mediaEntry.StreamURL,
			Download: mediaEntry.DownloadURL,
		}

		if strings.HasSuffix(mediaEntry.FileName, ".pdf") {
			doc.Documents = append(doc.Documents, _file)
			continue
		}

		if mediaEntry.FileName == "Abstract.md" {
			res, err := client.Fetch(mediaEntry.StreamURL)
			if err != nil {
				return fire.Fatal(err)
			}

			doc.Abstract = res
			continue
		}

		var lowRes, highRes *madek.Preview
		var mp4Source, webmSource *madek.Preview

		for _, preview := range mediaEntry.Previews {
			if preview.Type == "image" {
				if preview.Size == "large" {
					lowRes = &preview
				} else if preview.Size == "x_large" {
					highRes = &preview
				}
			}

			if preview.Type == "video" && preview.Size == "large" {
				if preview.ContentType == "video/mp4" {
					mp4Source = &preview
				} else if preview.ContentType == "video/webm" {
					webmSource = &preview
				}
			}
		}

		if lowRes == nil || highRes == nil {
			doc.Files = append(doc.Files, _file)
			continue
		}

		_image := image{
			file:    _file,
			LowRes:  lowRes.URL,
			HighRes: highRes.URL,
		}

		if mediaEntry.ID == doc.MadekCover {
			doc.Cover = &_image
			continue
		}

		if mp4Source == nil || webmSource == nil {
			doc.Images = append(doc.Images, _image)
			continue
		}

		doc.Videos = append(doc.Videos, video{
			image:      _image,
			MP4Source:  mp4Source.URL,
			WebMSource: webmSource.URL,
		})
	}

	return nil
}
