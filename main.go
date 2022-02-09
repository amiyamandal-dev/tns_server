package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"path/filepath"

	"os"
	"os/exec"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const TEMP_DIR = "temp_dir"

func upload(c echo.Context) error {
	// Read form fields
	name := c.FormValue("name")
	email := c.FormValue("email")

	//------------
	// Read files
	//------------

	// Multipart form
	form, err := c.MultipartForm()
	if err != nil {
		return err
	}
	files := form.File["files"]

	for _, file := range files {
		// Source
		src, err := file.Open()
		if err != nil {
			return err
		}
		defer src.Close()

		// Destination
		dst, err := os.Create(filepath.Join(TEMP_DIR, file.Filename))
		if err != nil {
			return err
		}
		defer dst.Close()

		// Copy
		if _, err = io.Copy(dst, src); err != nil {
			return err
		}

	}

	return c.HTML(http.StatusOK, fmt.Sprintf("<p>Uploaded successfully %d files with fields name=%s and email=%s.</p>", len(files), name, email))
}

func init() {
	log.Println("getting all model to work")
	// over here all model download option should be
	if _, err := os.Stat("deepspeech-0.9.3-models.scorer"); errors.Is(err, os.ErrNotExist) {
		log.Println("Downloading deepspeech-0.9.3-models.scorer")
		cmd := exec.Command("curl", "-LO", "https://github.com/mozilla/DeepSpeech/releases/download/v0.9.3/deepspeech-0.9.3-models.pbmm")
		_, err := cmd.Output()
		if err != nil {
			log.Fatalln(err)
			panic(err)
		}
	}
	if _, err := os.Stat("deepspeech-0.9.3-models.scorer"); errors.Is(err, os.ErrNotExist) {
		log.Println("downloading deepspeech-0.9.3-models.scorer")
		cmd := exec.Command("curl", "-LO", "https://github.com/mozilla/DeepSpeech/releases/download/v0.9.3/deepspeech-0.9.3-models.scorer")
		_, err := cmd.Output()
		if err != nil {
			log.Fatalln(err)
			panic(err)
		}
	}
	if _, err := os.Stat(TEMP_DIR); errors.Is(err, os.ErrNotExist) {
		os.Mkdir(TEMP_DIR, 0777)
	}
}

func main() {
	e := echo.New()
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.File("/favicon.ico", "images/favicon.ico")
	e.Static("/", "public")
	e.POST("/upload", upload)
	e.Logger.Fatal(e.Start(":1323"))
}
