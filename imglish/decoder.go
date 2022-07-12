package imglish

import (
	"image"
	"image/color"
	"io"
)

func decodeHeader(r io.Reader) (int, int, error) {
	r.Read(make([]byte, len(Magic)))

	x, err := decodeInt16(r)
	if err != nil {
		return 0, 0, err
	}

	y, err := decodeInt16(r)
	if err != nil {
		return 0, 0, err
	}

	return int(x), int(y), nil
}

func Decode(r io.Reader) (image.Image, error) {
	x, y, err := decodeHeader(r)
	if err != nil {
		return nil, err
	}

	img := image.NewRGBA(image.Rect(0, 0, x, y))

	for y := 0; y < img.Bounds().Dy(); y++ {
		for x := 0; x < img.Bounds().Dx(); x++ {
			c, err := decodeRGBA(r)
			if err != nil {
				return nil, err
			}

			img.Set(x, y, c)
		}
	}

	return img, nil
}

func DecodeConfig(r io.Reader) (image.Config, error) {
	return image.Config{
		ColorModel: color.RGBAModel,
	}, nil
}
