package main

import (
	"image"
	"image/png"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/vcokltfre/imglish/imglish"
)

func init() {
	rand.Seed(time.Now().UnixNano())

	image.RegisterFormat("imglish", imglish.Magic, imglish.Decode, imglish.DecodeConfig)
	image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)
}

func encode(inFile, outFile string) error {
	inf, err := os.Open(inFile)
	if err != nil {
		return err
	}
	defer inf.Close()

	outf, err := os.Create(outFile)
	if err != nil {
		return err
	}
	defer outf.Close()

	img, _, err:= image.Decode(inf)
	if err != nil {
		return err
	}

	return imglish.Encode(outf, img)
}

func decode(inFile, outFile string) error {
	inf, err := os.Open(inFile)
	if err != nil {
		return err
	}
	defer inf.Close()

	outf, err := os.Create(outFile)
	if err != nil {
		return err
	}
	defer outf.Close()

	img, _, err := image.Decode(inf)
	if err != nil {
		return err
	}

	return png.Encode(outf, img)
}

func main() {
	if len(os.Args) != 4 {
		log.Fatal("Expected exactly three arguments: <operation> <input> <output>")
	}

	switch os.Args[1] {
	case "encode":
		if err := encode(os.Args[2], os.Args[3]); err != nil {
			log.Fatal(err)
		}
	case "decode":
		if err := decode(os.Args[2], os.Args[3]); err != nil {
			log.Fatal(err)
		}
	default:
		log.Fatal("Unknown operation:", os.Args[1])
	}
}
