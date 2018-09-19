package main

import (
	"bytes"
	"context"
	"image"
	"image/png"
	"io/ioutil"
	"log"
	"os"

	"github.com/SUN-XIN/google-cloud-helper/storage"
)

const (
	BUCKET_NAME = "local-tests"
)

func main() {
	ctx := context.Background()

	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Printf("Failed NewClient: %+v", err)
		return
	}

	////////////////////////////////////////////////////////////
	////////////////////     Upload file      //////////////////
	////////////////////////////////////////////////////////////
	f, err := os.Open("data_file")
	if err != nil {
		log.Printf("Failed Open: %+v", err)
		return
	}
	defer f.Close()

	b, err := ioutil.ReadAll(f)
	if err != nil {
		log.Printf("Failed ReadAll: %+v", err)
		return
	}

	err, _ = storage.UploadAsPublic(ctx, client, BUCKET_NAME, "test09192018", b)
	if err != nil {
		log.Printf("Failed UploadAsPublic: %+v", err)
		return
	}
	log.Printf("UploadAsPublic file ok")

	////////////////////////////////////////////////////////////
	////////////////////     Upload image      /////////////////
	////////////////////////////////////////////////////////////
	imgFile, err := os.Open("hello.png")
	if err != nil {
		log.Printf("Failed Open img: %+v", err)
		return
	}

	buf := new(bytes.Buffer)
	img, _, err := image.Decode(imgFile)
	if err != nil {
		log.Printf("Failed image.Decode: %+v", err)
		return
	}

	err = png.Encode(buf, img)
	if err != nil {
		log.Printf("Failed png.Encode: %+v", err)
		return
	}

	err, link := storage.UploadAsPublic(ctx, client, BUCKET_NAME, "test09192018img", buf.Bytes())
	if err != nil {
		log.Printf("Failed UploadAsPublic: %+v", err)
		return
	}
	log.Printf("UploadAsPublic img ok: link %s", link)
}
