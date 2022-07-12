package imglish

import (
	"fmt"
	"image"
	"image/color"
	"io"
)


func Encode(w io.Writer, m image.Image) error {
	_, err := w.Write([]byte(Magic))
	if err != nil {
		return err
	}

	b := m.Bounds()

	if b.Dx() > 0xFFFF || b.Dy() > 0xFFFF {
		return fmt.Errorf("Image bounds are too large to encode.")
	}

	x, y := uint16(b.Dx()), uint16(b.Dy())

	_, err = w.Write([]byte(encodeInt16(x)))
	if err != nil {
		return err
	}
	_, err = w.Write([]byte(encodeInt16(y)))
	if err != nil {
		return err
	}

	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			c := m.At(x, y)

			r, g, b, a := c.RGBA()

			_, err := w.Write([]byte(encodeRGBA(color.RGBA{
				R: uint8(r / 256),
				G: uint8(g / 256),
				B: uint8(b / 256),
				A: uint8(a / 256),
			})))
			if err != nil {
				return err
			}
		}
	}

	return nil
}
