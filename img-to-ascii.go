package main

import (
	"flag"
	"fmt"

	"github.com/1ef7yy/img-to-ascii/src/convert"
	"github.com/1ef7yy/img-to-ascii/src/save"
	"github.com/1ef7yy/img-to-ascii/types"
)


var opts = types.Options{}


func init() {
    flag.BoolVar(&opts.IsColored, "color", false, "Toggles the color for output")
    flag.StringVar(&opts.SaveToFile, "save", "", "Saves the output to a file")
    flag.BoolVar(&opts.NoOutput, "no-output", false, "Doesn't output the image to console")
    flag.Parse()
}

func main() {

	vals, err := convert.ConvertImage("static/test.jpg", opts)

    if err != nil {
        fmt.Println(err.Error())
        return
    }


    if !opts.NoOutput{
        fmt.Println(vals)
    }

    if opts.SaveToFile != ""{
        err := save.SaveToFile(opts.SaveToFile, vals)

        if err != nil {
            fmt.Println("error writing to file: ", err.Error())
        }
    }
}
