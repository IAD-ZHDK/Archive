package main

import (
	"os"
	"regexp"
	"strings"

	"github.com/IAD-ZHDK/madek"
	"github.com/gonfire/fire"
	"github.com/pkg/errors"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
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
	// only run on create and update
	if ctx.Action != fire.Create && ctx.Action != fire.Update {
		return nil
	}

	// get model
	doc := ctx.Model.(*documentation)

	// check madek id
	if len(doc.MadekID) < 30 {
		return errors.New("Invalid Madek ID. Did you copy the whole id?")
	}

	// TODO: Enforce existence of a cover?

	// check madek cover id
	if doc.MadekCoverID != "" && len(doc.MadekID) < 30 {
		return errors.New("Invalid Madek Cover ID. Did you copy the whole id?")
	}

	// compile collection
	coll, err := client.CompileCollection(doc.MadekID)
	if err != nil {
		switch errors.Cause(err) {
		case madek.ErrNotFound:
			return errors.New("Collection cannot be found. Did you pass the right id?")
		case madek.ErrAccessForbidden:
			return errors.New("Collection is not publicly accessible. Did you set the necessary permissions?")
		default:
			return fire.Fatal(err)
		}
	}

	// validate title
	if len(coll.MetaData.Title) < 5 {
		return errors.New("Collection title must be longer than 5 characters.")
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
		return errors.New("Invalid year must be like '2016'.")
	}

	// validate genre
	if len(coll.MetaData.Genres) != 1 || coll.MetaData.Genres[0] != "Design" {
		return errors.New("Collection genre must be 'Design'.")
	}

	// TODO: Affiliation must be Interaction Design (BDE_VIAD...)

	// set data
	doc.Title = coll.MetaData.Title
	doc.Subtitle = coll.MetaData.Subtitle
	doc.Abstract = coll.MetaData.Description
	doc.Year = coll.MetaData.Year

	// reset lists
	doc.PeopleIDs = nil
	doc.TagIDs = nil
	doc.Cover = nil
	doc.Videos = nil
	doc.Images = nil
	doc.Documents = nil
	doc.Files = nil

	// add authors
	for _, author := range coll.MetaData.Authors {
		var p person
		err := ctx.DB.C("people").Find(bson.M{
			"name": author,
		}).One(&p)
		if err == mgo.ErrNotFound {
			continue
		}
		if err != nil {
			return err
		}

		doc.PeopleIDs = append(doc.PeopleIDs, p.ID())
	}

	// add tags
	for _, keyword := range coll.MetaData.Keywords {
		var t tag
		err := ctx.DB.C("tags").Find(bson.M{
			"name": keyword,
		}).One(&t)
		if err == mgo.ErrNotFound {
			continue
		}
		if err != nil {
			return err
		}

		doc.TagIDs = append(doc.TagIDs, t.ID())
	}

	// process media entries
	for _, mediaEntry := range coll.MediaEntries {
		// validate title
		if len(mediaEntry.MetaData.Title) < 5 {
			return errors.New("Entry title must be longer than 5 characters.")
		}

		// validate copyright holder
		if mediaEntry.MetaData.Copyright.Holder != "Interaction Design" {
			return errors.New("Entry copyright holder must be 'Interaction Design'.")
		}

		// validate copyright license
		if len(mediaEntry.MetaData.Copyright.Licenses) != 1 || mediaEntry.MetaData.Copyright.Licenses[0] != "Alle Rechte vorbehalten" {
			return errors.New("Entry copyright license must be 'Alle Rechte vorbehalten'.")
		}

		// validate copyright usage
		if mediaEntry.MetaData.Copyright.Usage != "Das Werk darf nur mit Einwilligung des Autors/Rechteinhabers weiter verwendet werden." {
			return errors.New("Entry copyright usage must be 'Das Werk darf nur mit Einwilligung des Autors/Rechteinhabers weiter verwendet werden.'.")
		}

		// prepare basic file
		_file := file{
			Title:    mediaEntry.MetaData.Title,
			Stream:   mediaEntry.StreamURL,
			Download: mediaEntry.DownloadURL,
		}

		// add documents and continue
		if strings.HasSuffix(mediaEntry.FileName, ".pdf") {
			doc.Documents = append(doc.Documents, _file)
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
			doc.Files = append(doc.Files, _file)
			continue
		}

		// prepare image
		_image := image{
			file:    _file,
			LowRes:  lowRes.URL,
			HighRes: highRes.URL,
		}

		// add cover if ids match
		if mediaEntry.ID == doc.MadekCoverID {
			doc.Cover = &_image
			continue
		}

		// add image if video sources are missing
		if mp4Source == nil || webmSource == nil {
			doc.Images = append(doc.Images, _image)
			continue
		}

		// add video
		doc.Videos = append(doc.Videos, video{
			image:      _image,
			MP4Source:  mp4Source.URL,
			WebMSource: webmSource.URL,
		})
	}

	return nil
}
