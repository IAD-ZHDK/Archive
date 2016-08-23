package main

import (
	"os"

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

	doc.Videos = nil
	doc.Images = nil
	doc.Files = nil

	set, err := client.CompileSet(doc.MadekSet)
	if err != nil {
		return fire.Fatal(err)
	}

	for _, mediaEntry := range set.MediaEntries {
		var lowRes, highRes *madek.Preview
		var mp4Source, webmSource *madek.Preview

		for _, preview := range mediaEntry.Previews {
			if preview.Type == "image" {
				if preview.Size == "x_large" {
					lowRes = preview
				} else if preview.Size == "maximum" {
					highRes = preview
				}
			}

			if preview.Type == "video" && preview.Size == "large" {
				if preview.ContentType == "video/mp4" {
					mp4Source = preview
				} else if preview.ContentType == "video/webm" {
					webmSource = preview
				}
			}
		}

		if lowRes != nil && highRes != nil {
			if mp4Source != nil && webmSource != nil {
				doc.Videos = append(doc.Videos, video{
					image: image{
						Title: mediaEntry.Title,
						LowRes:  lowRes.URL,
						HighRes: highRes.URL,
					},
					MP4Source:  mp4Source.URL,
					WebMSource: webmSource.URL,
				})

				continue
			}

			doc.Images = append(doc.Images, image{
				Title: mediaEntry.Title,
				LowRes:  lowRes.URL,
				HighRes: highRes.URL,
			})

			continue
		}

		doc.Files = append(doc.Files, file{
			Title:    mediaEntry.Title,
			Download: mediaEntry.DownloadURL,
		})
	}

	return nil
}
