package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/1ef7yy/img-to-ascii/src/convert"
	"github.com/1ef7yy/img-to-ascii/src/save"
	"github.com/1ef7yy/img-to-ascii/types"
)

var opts = types.Options{}

func initFlags() {
	flag.BoolVar(&opts.IsColored, "is-colored", false, "Toggles the color for output")
	flag.StringVar(&opts.SaveToFile, "save", "", "Saves the output to a file")
	flag.BoolVar(&opts.NoOutput, "no-output", false, "Doesn't output the image to console")
	flag.StringVar(&opts.Src, "source", "", "Path to source image")
	flag.StringVar(&opts.Recursive, "recursive", "", "Path to a directory of images") // TODO
	flag.StringVar(&opts.SingleColor, "single-color", "", "Select color to make the whole output image this color")
	flag.Parse()
}

func main() {

	initFlags()

	start := time.Now()

	vals, err := convert.ConvertImage(opts.Src, opts)

	convert := time.Now()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if !opts.NoOutput {
		fmt.Println(vals)
	}

	if opts.SaveToFile != "" {
		err := save.SaveToFile(opts.SaveToFile, vals)

		if err != nil {
			fmt.Println("error writing to file: ", err.Error())
		}
	}

	finish := time.Now()

	fmt.Printf("conversion took %s\n", convert.Sub(start).String())
	fmt.Printf("program took %s\n", finish.Sub(start).String())
}
