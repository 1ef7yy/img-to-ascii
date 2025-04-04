package convert

import (
	"fmt"
	"image"
	"image/color"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"strconv"
	"strings"

	"github.com/1ef7yy/img-to-ascii/types"
)

func parseColor(color string) ([]int, error) {
	// should be of format R,G,B (for example 123,34,5)
	colors := strings.Split(color, ",")
	if len(colors) != 3 {
		return nil, fmt.Errorf("error splitting single color value")
	}

	var RGBValues []int

	for _, color := range colors {
		val, err := strconv.Atoi(color)

		if err != nil {
			return nil, fmt.Errorf("error converting RGB values to numbers: %s", err.Error())
		}

		if val > 255 || val < 0 {
			return nil, fmt.Errorf("one of RGB values is over 255 or lower than 0")
		}

		RGBValues = append(RGBValues, val)
	}

	return RGBValues, nil
}

func ConvertImage(path string, opts types.Options) (string, error) {
	img, err := openImage(path)

	if err != nil {
		return "", err
	}

	if opts.IsColored {
		return coloredToASCII(img)
	} else if opts.SingleColor != "" {
		RGBValues, err := parseColor(opts.SingleColor)
		if err != nil {
			return "", fmt.Errorf("could not parse RGBValues: %s", err.Error())
		}
		return singleColorToASCII(img, RGBValues)
	} else {
		grayscale := toGrayscale(img)
		return grayscaleToASCII(grayscale)
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

// TODO
func singleColorToASCII(img image.Image, RGB []int) (string, error) {
	asciiChars := []rune("$@B%8&WM#*oahkbdpqwmZO0QLCJUYXzcvunxrjft/\\|()1{}[]?-_+~<>i!lI;:,\"^`'. ")

	var asciiArt strings.Builder

	fmt.Println("RGB Values: ", RGB)

	bounds := img.Bounds()
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			c := img.At(x, y)
			grayVal := color.GrayModel.Convert(c).(color.Gray).Y
			asciiChar := asciiChars[int(grayVal)*len(asciiChars)/256]
			_, err := asciiArt.WriteString(fmt.Sprintf("\033[38;2;%d;%d;%dm%c\033[0m", RGB[0], RGB[1], RGB[2], asciiChar))
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
