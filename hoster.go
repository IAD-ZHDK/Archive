package main

import (
	"archive/zip"
	"bytes"
	"errors"
	"io"
	"mime"
	"net/http"
	"path"
	"strconv"

	"github.com/gonfire/fire"
	"github.com/labstack/echo"
	"gopkg.in/mgo.v2/bson"
)

type hoster struct {
	pool fire.Pool
}

func newHoster(pool fire.Pool) *hoster {
	return &hoster{
		pool: pool,
	}
}

func (h *hoster) Register(router *echo.Echo) {
	router.GET("web/:id/:num/:file", h.serveFile)
}

func (h *hoster) Inspect() fire.ComponentInfo {
	return fire.ComponentInfo{
		Name: "Hoster",
	}
}

func (h *hoster) serveFile(ctx echo.Context) error {
	// validate id
	if !bson.IsObjectIdHex(ctx.Param("id")) {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid id")
	}

	// get id
	id := bson.ObjectIdHex(ctx.Param("id"))

	// get connection from pool
	sess, db, err := h.pool.Get()
	if err != nil {
		return err
	}

	// ensure session gets closed
	defer sess.Close()

	// get documentation
	var doc documentation
	err = db.C("documentations").FindId(id).One(&doc)
	if err != nil {
		return err
	}

	// read website number
	num, err := strconv.Atoi(ctx.Param("num"))
	if err != nil {
		return err
	}

	// validate number
	if num >= len(doc.Websites) {
		return errors.New("invalid website number")
	}

	// get website
	website := doc.Websites[num]

	// load website container
	res, err := http.Get(website.Stream)
	if err != nil {
		return err
	}

	// make sure the body will be properly closed
	defer res.Body.Close()

	// read full response
	buf := make([]byte, res.ContentLength)
	_, err = io.ReadFull(res.Body, buf)
	if err != nil {
		return err
	}

	// make reader
	reader := bytes.NewReader(buf)

	// create zip reader
	archive, err := zip.NewReader(reader, res.ContentLength)
	if err != nil {
		return err
	}

	// prepare zipped file
	var zippedFile *zip.File

	// iterate over files
	for _, file := range archive.File {
		if file.Name == ctx.Param("file") {
			zippedFile = file
			break
		}
	}

	// check existence
	if zippedFile == nil {
		return echo.NewHTTPError(http.StatusNotFound, "file not found")
	}

	// open file for reading
	f, err := zippedFile.Open()
	if err != nil {
		return err
	}

	// make sure the file will be properly closed
	defer f.Close()

	// get content type
	contentType := mime.TypeByExtension(path.Ext(zippedFile.Name))

	// set content type header
	ctx.Response().Header().Set(echo.HeaderContentType, contentType)

	// set header
	ctx.Response().WriteHeader(http.StatusOK)

	// write file to request
	_, err = io.Copy(ctx.Response(), f)
	if err != nil {
		return err
	}

	return nil
}
