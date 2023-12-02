package main

import (
	"fmt"
	"github.com/crazy3lf/colorconv"
)

func Bright[T []*Object | ObjectsGroup](ImagePath string, width, height int) (T, int, error) {
	var total []*Object
	var numObject int
	img, err := loadImage(ImagePath, width, height)
	if err != nil {
		return nil, 0, err
	}
	bounds := img.Bounds()
	for y := bounds.Min.Y; y < bounds.Max.Y; y += 16 {
		for x := bounds.Min.X; x < bounds.Max.X; x += 16 {
			r, g, b, _ := img.At(x, y).RGBA()
			h, s, l := colorconv.RGBToHSL(uint8(r), uint8(g), uint8(b))
			l = l * 255
			s = s * 255
			switch {
			case h < 0:
				h = 0
			case h > 255:
				h = 0
			case s < 0:
				s = 0
			case s > 255:
				s = 255
			case l < 0:
				l = 0
			case l > 255:
				l = 255
			}
			newObject := NewObject(x, y, OBJECT_TYPE_LIGHT, []Param{
				{Key: "light_invert", Val: "0"},
				{Key: "scale", Val: fmt.Sprintf("0.2")},
				{Key: "layer", Val: "2"},
				{Key: "light_saturation", Val: fmt.Sprintf("%.4f", s)},
				{Key: "light_hue", Val: fmt.Sprintf("%.4f", h)},
				{Key: "light_lum", Val: fmt.Sprintf("%.4f", l)},
			})
			total = append(total, newObject)
			numObject++
		}
	}
	return total, numObject, nil
}
