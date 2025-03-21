package main

import (
	"testing"

	"github.com/1ef7yy/img-to-ascii/src/convert"
	"github.com/1ef7yy/img-to-ascii/types"
)

var (
	testJPGPath = "static/test.jpg"
)

func imgToAsciiColored() (string, error) {
	opts := types.Options{
		IsColored:  true,
		NoOutput:   false,
		SaveToFile: "",
	}
	val, err := convert.ConvertImage(testJPGPath, opts)

	if err != nil {
		return "", err
	}

	return val, err
}

func BenchmarkImgToColor(b *testing.B) {
	_, err := imgToAsciiColored()

	if err != nil {
		b.Error(err)
	}
}
