package imglish

import (
	"crypto/rand"
	"fmt"
	"image/color"
	"io"
	"math/big"
	"strings"
)

type RGBA struct {
	R, G, B, A uint8
}

func (r RGBA) RGBA() (uint32, uint32, uint32, uint32) {
	return uint32(r.R) * 256, uint32(r.G) * 256, uint32(r.B) * 256, uint32(r.A) * 256
}

const Magic = "Hello! "

func title(word string) string {
	return strings.ToUpper(word[:1]) + word[1:]
}

func choice(choices []string) string {
	n, err := rand.Int(rand.Reader, big.NewInt(int64(len(choices))))
	if err != nil {
		panic(err)
	}

	return choices[n.Int64()]
}

func getNamePrefix(name string) string {
	c := name[0]

	if c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' {
		return "an"
	}

	return "a"
}

func readSentence(r io.Reader) (string, error) {
	sentence := ""
	b := [1]byte{}

	for {
		_, err := r.Read(b[:])
		if err != nil {
			return "", err
		}

		if b[0] == '.' {
			break
		}

		sentence += string(b[0])
	}

	return sentence, nil
}

func encodeRGBA(value color.RGBA) string {
	subjectName := animals[value.R]
	subjectAdj := adjectives[value.G]
	objectName := animals[value.B]
	objectAdj := adjectives[value.A]

	action := choice(actions)
	place := choice(places)

	subjectPrefix := title(getNamePrefix(subjectAdj))
	objectPrefix := getNamePrefix(objectAdj)

	return fmt.Sprintf("%s %s %s in %s %s %s %s %s. ", subjectPrefix, subjectAdj, subjectName, place, action, objectPrefix, objectAdj, objectName)
}

func decodeRGBA(r io.Reader) (RGBA, error) {
	col := RGBA{}

	sentence, err := readSentence(r)
	if err != nil {
		return col, err
	}

	parts := strings.Split(strings.TrimSpace(sentence), " ")

	if len(parts) != 9 {
		return col, fmt.Errorf("Invalid sentence: %s", sentence)
	}

	sa, ok := reverseAdjectives[parts[1]]
	if !ok {
		return col, fmt.Errorf("Invalid adjective: %s", parts[1])
	}

	sn, ok := reverseAnimals[parts[2]]
	if !ok {
		return col, fmt.Errorf("Invalid animal: %s", parts[2])
	}

	oa, ok := reverseAdjectives[parts[7]]
	if !ok {
		return col, fmt.Errorf("Invalid adjective: %s", parts[7])
	}

	on, ok := reverseAnimals[parts[8]]
	if !ok {
		return col, fmt.Errorf("Invalid animal: %s", parts[8])
	}

	return RGBA{
		R: sn,
		G: sa,
		B: on,
		A: oa,
	}, nil
}

func encodeInt16(value uint16) string {
	subjectName := animals[uint8(value)]
	objectName := animals[uint8(value>>8)]

	action := choice(actions)
	place := choice(places)

	subjectPrefix := title(getNamePrefix(subjectName))
	objectPrefix := getNamePrefix(objectName)

	return fmt.Sprintf("%s %s in %s %s %s %s. ", subjectPrefix, subjectName, place, action, objectPrefix, objectName)
}


func decodeInt16(r io.Reader) (uint16, error) {
	sentence, err := readSentence(r)
	if err != nil {
		return 0, err
	}

	parts := strings.Split(strings.TrimSpace(sentence), " ")

	if len(parts) != 7 {
		return 0, fmt.Errorf("Invalid sentence: %s", sentence)
	}

	subject := parts[1]
	object := parts[6]

	n1, ok := reverseAnimals[subject]
	if !ok {
		return 0, fmt.Errorf("Invalid subject: %s", subject)
	}

	n2, ok := reverseAnimals[object]
	if !ok {
		return 0, fmt.Errorf("Invalid object: %s", object)
	}

	return uint16(n2) << 8 | uint16(n1), nil
}
