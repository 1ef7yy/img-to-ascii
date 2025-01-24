package convert

import (
    "github.com/1ef7yy/img-to-ascii/types"
    "image/color"
	"image"
    _ "image/jpeg"
    _ "image/png"
    "os"
)

func ConvertImage(path string, opts types.Options) (string, error) {
    img, err := openImage(path)

    if err != nil {
        return "", err
    }

    grayscale := toGrayscale(img)

    ascii := grayscalToASCII(grayscale)

    return ascii, nil
}


func openImage(path string) (image.Image, error) {
    file, err := os.Open(path)


    if err != nil {
        return nil, err
    }


    defer file.Close()

    img, _, err := image.Decode(file)


    if err != nil {
        return nil, err
    }

    return img, err
}

func toGrayscale(img image.Image) image.Image {
    bounds := img.Bounds()

    width, height := bounds.Max.X, bounds.Max.Y

    grayImg := image.NewGray(image.Rect(0, 0, width, height))

    for y := 0; y < height; y++ {
        for x := 0; x < width; x ++ {
            c := img.At(x, y)
            r, g, b, _ := c.RGBA()

            gray := uint8(0.299*float64(r>>8) + 0.587*float64(g>>8) + 0.114*float64(b>>8))
            grayImg.Set(x, y, color.Gray{Y: gray})
        }
    }
    return grayImg
}


func grayscalToASCII(grayImg image.Image) string {
    asciiChars := []rune("$@B%8&WM#*oahkbdpqwmZO0QLCJUYXzcvunxrjft/\\|()1{}[]?-_+~<>i!lI;:,\"^`'. ")
	var asciiArt string

	bounds := grayImg.Bounds()
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			c := grayImg.At(x, y)
			grayVal := color.GrayModel.Convert(c).(color.Gray).Y
			asciiChar := asciiChars[int(grayVal)*len(asciiChars)/256]
			asciiArt += string(asciiChar)
		}
		asciiArt += "\n"
	}

	return asciiArt
}
