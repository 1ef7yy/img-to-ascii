package convert

import (
	"fmt"
	"image"
	"image/color"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"strings"

	"github.com/1ef7yy/img-to-ascii/types"
)

func ConvertImage(path string, opts types.Options) (string, error) {
	img, err := openImage(path)

	if err != nil {
		return "", err
	}

	if !opts.IsColored {

		grayscale := toGrayscale(img)
		return grayscaleToASCII(grayscale)

	} else {
		return coloredToASCII(img)
	}
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
		for x := 0; x < width; x++ {
			c := img.At(x, y)
			r, g, b, _ := c.RGBA()

			gray := uint8(0.299*float64(r>>8) + 0.587*float64(g>>8) + 0.114*float64(b>>8))
			grayImg.Set(x, y, color.Gray{Y: gray})
		}
	}
	return grayImg
}

func grayscaleToASCII(grayImg image.Image) (string, error) {
	asciiChars := []rune("$@B%8&WM#*oahkbdpqwmZO0QLCJUYXzcvunxrjft/\\|()1{}[]?-_+~<>i!lI;:,\"^`'. ")
	var asciiArt strings.Builder

	bounds := grayImg.Bounds()
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			c := grayImg.At(x, y)
			grayVal := color.GrayModel.Convert(c).(color.Gray).Y
			asciiChar := asciiChars[int(grayVal)*len(asciiChars)/256]
			_, err := asciiArt.WriteString(string(asciiChar))
			if err != nil {
				return "", err
			}
		}
		_, err := asciiArt.WriteString("\n")
		if err != nil {
			return "", err
		}
	}

	return asciiArt.String(), nil
}

func coloredToASCII(img image.Image) (string, error) {
	var asciiArt strings.Builder
	asciiChars := []rune("$@B%8&WM#*oahkbdpqwmZO0QLCJUYXzcvunxrjft/\\|()1{}[]?-_+~<>i!lI;:,\"^`'. ")

	bounds := img.Bounds()

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			c := img.At(x, y)
			r, g, b, _ := c.RGBA()

			r8 := uint8(r >> 8)
			g8 := uint8(g >> 8)
			b8 := uint8(b >> 8)

			brightness := (int(r8)*299 + int(g8)*587 + int(b8)*114) / 1000
			asciiChar := asciiChars[int(brightness)*len(asciiChars)/256]
			coloredChar := fmt.Sprintf("\033[38;2;%d;%d;%dm%c\033[0m", r/256, g/256, b/256, asciiChar)
			_, err := asciiArt.WriteString(coloredChar)
			if err != nil {
				return "", err
			}
		}
		_, err := asciiArt.WriteString("\n")
		if err != nil {
			return "", err
		}
	}

	return asciiArt.String(), nil
}
