package main

import (
	"fmt"
	"github.com/nfnt/resize"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"strconv"
	"strings"
)

func RGBtoIWMBlend(r, g, b uint32) uint32 {
	return (b>>8)<<16 | (g>>8)<<8 | (r >> 8)
}

func mostFrequentColor(img image.Image, minX, minY, maxX, maxY int) string {
	colorMap := make(map[string]int)
	for y := minY; y < maxY; y++ {
		for x := minX; x < maxX; x++ {
			r, g, b, _ := img.At(x, y).RGBA()
			t := fmt.Sprintf("%d", RGBtoIWMBlend(r, g, b))
			if t != "0" {
				colorMap[t]++
			}
		}
	}
	var mostFrequent struct {
		color string
		count int
	}
	for color, count := range colorMap {
		if count > mostFrequent.count {
			mostFrequent.color = color
			mostFrequent.count = count
		}
	}
	return mostFrequent.color
}

// IgnoreColors format: B G R
// return: Objects, numobjects, error
// width height recommended: 794, 608(full map)
func Image[T []*Object | ObjectsGroup](ImagePath string, width, height int, IgnoreColors []int, Scale float64, EventOnObjects []Event) (T, int, error) {
	var total []*Object
	var numObject int
	ignoreColorsMap := make(map[string]bool)
	for _, v := range IgnoreColors {
		ignoreColorsMap[strconv.Itoa(v)] = true
	}
	img, err := loadImage(ImagePath, width, height)
	if err != nil {
		return nil, 0, err
	}
	bounds := img.Bounds()
	s, err := strconv.Atoi(strings.Split(fmt.Sprintf("%.1f", Scale), ".")[1])
	if err != nil {
		return nil, 0, err
	}
	for y := bounds.Min.Y; y < bounds.Max.Y; y += s {
		for x := bounds.Min.X; x < bounds.Max.X; x += s {
			t := mostFrequentColor(img, x-s, y-s, x, y)
			if t == "0" {
				continue
			} else if ignoreColorsMap[t] {
				continue
			}
			newObject := NewObject(x, y, OBJECT_TYPE_CHERRY, []Param{
				{Key: "blend_color", Val: t},
				{Key: "scale", Val: fmt.Sprintf("%.1f", Scale)},
				{Key: "layer", Val: "2"},
				{Key: "cherry_color", Val: strconv.Itoa(11)},
				{Key: "bounce", Val: "0"},
			})
			if EventOnObjects != nil {
				newObject.AppendEvent(EventOnObjects)
			}
			total = append(total, newObject)
			numObject++
		}
	}
	return total, numObject, nil
}

func loadImage(path string, width, height int) (image.Image, error) {
	f, err := os.Open(path) // 画像ファイルを開く
	if err != nil {
		return nil, err
	}
	defer f.Close()
	img, _, err := image.Decode(f)
	img = resize.Resize(uint(width), uint(height), img, resize.Bilinear)
	return img, err
}
