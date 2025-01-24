package main

import (
    "github.com/1ef7yy/img-to-ascii/types"
    "flag"
	"fmt"
	"github.com/1ef7yy/img-to-ascii/src/convert"
)


var opts = types.Options{}


func init() {
    flag.BoolVar(&opts.IsColored, "color", false, "Toggles the color for output")
    flag.StringVar(&opts.SaveToFile, "save", "", "Saves the output to a file")
    flag.Parse()
}

func main() {

	vals, err := convert.ConvertImage("static/test.jpg", opts)

    fmt.Println(opts)

    if err != nil {
        fmt.Println(err.Error())
        return
    }
	fmt.Println(vals)
}
