package main_test

import (
	"fmt"
	"testing"

	"github.com/1ef7yy/img-to-ascii/src/convert"
	"github.com/1ef7yy/img-to-ascii/types"
)

var (
	testJPGPath = "static/sashaGC.PNG"
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

func imgToAsciiColoredBenchmark(b *testing.B) {
	val, err := imgToAsciiColored()

	if err != nil {
		b.Error(err)
	}
	fmt.Println(val)
}
