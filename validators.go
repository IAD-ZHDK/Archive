package main

import (
	"errors"
	"os"
	"regexp"
	"strings"

	"github.com/256dpi/fire"
	"github.com/IAD-ZHDK/madek"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var client = madek.NewClient(
	os.Getenv("MADEK_ADDRESS"),
	os.Getenv("MADEK_USERNAME"),
	os.Getenv("MADEK_PASSWORD"),
)

func projectValidator(ctx *fire.Context) error {
	// only run on create and update
	if ctx.Action != fire.Create && ctx.Action != fire.Update {
		return nil
	}

	// get model
	project := ctx.Model.(*Project)

	// check slug on publishing
	if project.Published && len(project.Slug) < 5 {
		return errors.New("slug must be at least 5 characters")
	}

	// check madek id
	if len(project.MadekID) < 30 {
		return errors.New("invalid madek id")
	}

	// TODO: Enforce existence of a cover?

	// check madek cover id
	if project.MadekCoverID != "" && len(project.MadekID) < 30 {
		return errors.New("invalid madek cover id")
	}

	// force unpublished on create
	if ctx.Action == fire.Create {
		project.Published = false
	}

	// compile collection
	coll, err := client.CompileCollection(project.MadekID)
	if err != nil {
		switch err {
		case madek.ErrNotFound:
			return errors.New("collection cannot be found")
		case madek.ErrAccessForbidden:
			return errors.New("collection is not publicly accessible")
		default:
			return fire.Fatal(err)
		}
	}

	// validate title
	if len(coll.MetaData.Title) < 5 {
		return errors.New("collection title must be longer than 5 characters")
	}

	// TODO: What are the rules here?

	//if len(coll.MetaData.Subtitle) < 50 {
	//	return errors.New("Collection subtitle must be longer than 50 characters.")
	//}

	//if len(coll.MetaData.Description) < 200 {
	//	return errors.New("Collection description must be longer than 200 characters.")
	//}

	// validate year
	ok, err := regexp.MatchString(`^\d{4}$`, coll.MetaData.Year)
	if !ok || err != nil {
		return errors.New("invalid year must be like '2016'")
	}

	// validate genre
	if len(coll.MetaData.Genres) != 1 || coll.MetaData.Genres[0] != "Design" {
		return errors.New("collection genre must be 'Design'")
	}

	// TODO: Affiliation must be Interaction Design (BDE_VIAD...)

	// set data
	project.Title = coll.MetaData.Title
	project.Subtitle = coll.MetaData.Subtitle
	project.Abstract = coll.MetaData.Description
	project.Year = coll.MetaData.Year

	// reset lists
	project.People = nil
	project.Tags = nil
	project.Cover = nil
	project.Videos = nil
	project.Images = nil
	project.Documents = nil
	project.Websites = nil
	project.Files = nil

	// add authors
	for _, author := range coll.MetaData.Authors {
		authorName := author.FirstName + " " + author.LastName
		var p Person
		err := ctx.Store.DB().C("people").Find(bson.M{
			"name": authorName,
		}).One(&p)
		if err == mgo.ErrNotFound {
			return errors.New("person " + authorName + " has not yet been created.")
		}
		if err != nil {
			return err
		}

		project.People = append(project.People, p.ID())
	}

	// add tags
	for _, keyword := range coll.MetaData.Keywords {
		var t Tag
		err := ctx.Store.DB().C("tags").Find(bson.M{
			"name": keyword,
		}).One(&t)
		if err == mgo.ErrNotFound {
			return errors.New("tag " + keyword + " has not yet been created.")
		}
		if err != nil {
			return err
		}

		project.Tags = append(project.Tags, t.ID())
	}

	// process media entries
	for _, mediaEntry := range coll.MediaEntries {
		// validate title
		if len(mediaEntry.MetaData.Title) < 5 {
			return errors.New("entry title must be longer than 5 characters")
		}

		// validate copyright holder
		if mediaEntry.MetaData.Copyright.Holder != "Interaction Design" {
			return errors.New("entry copyright holder must be 'Interaction Design'")
		}

		// validate copyright license
		if len(mediaEntry.MetaData.Copyright.Licenses) != 1 || mediaEntry.MetaData.Copyright.Licenses[0] != "Alle Rechte vorbehalten" {
			return errors.New("entry copyright license must be 'Alle Rechte vorbehalten'")
		}

		// validate copyright usage
		if mediaEntry.MetaData.Copyright.Usage != "Das Werk darf nur mit Einwilligung des Autors/Rechteinhabers weiter verwendet werden." {
			return errors.New("entry copyright usage must be 'Das Werk darf nur mit Einwilligung des Autors/Rechteinhabers weiter verwendet werden.'")
		}

		// prepare basic file
		fl := File{
			Title:    mediaEntry.MetaData.Title,
			Stream:   mediaEntry.StreamURL,
			Download: mediaEntry.DownloadURL,
		}

		// add documents and continue
		if strings.HasSuffix(mediaEntry.FileName, ".pdf") {
			project.Documents = append(project.Documents, fl)
			continue
		}

		// add websites and continue
		if strings.HasSuffix(mediaEntry.FileName, ".web.zip") {
			project.Websites = append(project.Websites, fl)
			continue
		}

		// prepare previews
		var lowRes, highRes *madek.Preview
		var mp4Source, webmSource *madek.Preview

		// process previews
		for _, preview := range mediaEntry.Previews {
			if preview.Type == "image" {
				if preview.Size == "large" {
					lowRes = preview
				} else if preview.Size == "x_large" {
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

		// add ordinary file and continue when previews are missing
		if lowRes == nil || highRes == nil {
			project.Files = append(project.Files, fl)
			continue
		}

		// prepare image
		img := Image{
			File:    fl,
			LowRes:  lowRes.URL,
			HighRes: highRes.URL,
		}

		// add cover if ids match
		if mediaEntry.ID == project.MadekCoverID {
			project.Cover = &img
			continue
		}

		// add image if video sources are missing
		if mp4Source == nil || webmSource == nil {
			project.Images = append(project.Images, img)
			continue
		}

		// add video
		project.Videos = append(project.Videos, Video{
			Image:      img,
			MP4Source:  mp4Source.URL,
			WebMSource: webmSource.URL,
		})
	}

	return nil
}

func slugAndNameValidator(ctx *fire.Context) error {
	// only validate on create and update
	if ctx.Action != fire.Create && ctx.Action != fire.Update {
		return nil
	}

	if len(ctx.Model.MustGet("slug").(string)) < 5 {
		return errors.New("slug must be at least 5 characters")
	}

	if len(ctx.Model.MustGet("name").(string)) < 5 {
		return errors.New("mame must be at least 5 characters")
	}

	return nil
}
