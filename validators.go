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
	if ctx.Action != fire.Create && ctx.Action != fire.Update {
		return nil
	}

	doc := ctx.Model.(*documentation)

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

	ok, err := regexp.MatchString(`^\d{4}$`, coll.MetaData.Year)
	if !ok || err != nil {
		return errors.New("Invalid year must be like '2016'.")
	}

	if len(coll.MetaData.Genres) != 1 || coll.MetaData.Genres[0] != "Design" {
		return errors.New("Collection genre must be 'Design'.")
	}

	// TODO: Affiliation must be Interaction Design (BDE_VIAD...)

	doc.Title = coll.MetaData.Title
	doc.Subtitle = coll.MetaData.Subtitle
	doc.Abstract = coll.MetaData.Description
	doc.Year = coll.MetaData.Year

	doc.PeopleIDs = nil
	doc.TagIDs = nil
	doc.Cover = nil
	doc.Videos = nil
	doc.Images = nil
	doc.Documents = nil
	doc.Files = nil

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

	for _, mediaEntry := range coll.MediaEntries {
		if len(mediaEntry.MetaData.Title) < 5 {
			return errors.New("Entry title must be longer than 5 characters.")
		}

		if mediaEntry.MetaData.Copyright.Holder != "Interaction Design" {
			return errors.New("Entry copyright holder must be 'Interaction Design'.")
		}

		if len(mediaEntry.MetaData.Copyright.Licenses) != 1 || mediaEntry.MetaData.Copyright.Licenses[0] != "Alle Rechte vorbehalten" {
			return errors.New("Entry copyright license must be 'Alle Rechte vorbehalten'.")
		}

		if mediaEntry.MetaData.Copyright.Usage != "Das Werk darf nur mit Einwilligung des Autors/Rechteinhabers weiter verwendet werden." {
			return errors.New("Entry copyright usage must be 'Das Werk darf nur mit Einwilligung des Autors/Rechteinhabers weiter verwendet werden.'.")
		}

		_file := file{
			Title:    mediaEntry.MetaData.Title,
			Stream:   mediaEntry.StreamURL,
			Download: mediaEntry.DownloadURL,
		}

		if strings.HasSuffix(mediaEntry.FileName, ".pdf") {
			doc.Documents = append(doc.Documents, _file)
			continue
		}

		var lowRes, highRes *madek.Preview
		var mp4Source, webmSource *madek.Preview

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

		if lowRes == nil || highRes == nil {
			doc.Files = append(doc.Files, _file)
			continue
		}

		_image := image{
			file:    _file,
			LowRes:  lowRes.URL,
			HighRes: highRes.URL,
		}

		if mediaEntry.ID == doc.MadekCoverID {
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
