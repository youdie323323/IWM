package main

import (
	"fmt"
	"github.com/alexeyco/simpletable"
	"strings"
)

var (
	// StyleUnicodeTwo - unicode table style:
	//
	// ╔═══╤══════════════════╤════════════╗
	// ║ # │       NAME       │    TAX     ║
	// ╠---┼------------------┼------------╣
	// ║ 1 │ Newton G. Goetz  │   $ 532.70 ║
	// ║ 2 │ Rebecca R. Edney │  $ 1423.25 ║
	// ║ 3 │ John R. Jackson  │  $ 7526.12 ║
	// ║ 4 │ Ron J. Gomes     │   $ 123.84 ║
	// ║ 5 │ Penny R. Lewis   │  $ 3221.11 ║
	// ╠---┼------------------┼------------╣
	// ║   │         Subtotal │ $ 12827.02 ║
	// ╚═══╧══════════════════╧════════════╝
	StyleUnicodeTwo = &simpletable.Style{
		Border: &simpletable.BorderStyle{
			TopLeft:            "╔",
			Top:                "═",
			TopRight:           "╗",
			Right:              "║",
			BottomRight:        "╝",
			Bottom:             "═",
			BottomLeft:         "╚",
			Left:               "║",
			TopIntersection:    "╤",
			BottomIntersection: "╧",
		},
		Divider: &simpletable.DividerStyle{
			Left:         "╟",
			Center:       "─",
			Right:        "╢",
			Intersection: "┼",
		},
		Cell: "│",
	}
)

func IWM_Path_test() {
	paths := []string{
		//2E01000
		//01000000
		"00000000000000000000040C",
		"00000000000000000000050C",
		"00000000000000000000058C",
		"00000000000000000000060C",
		"00000000000000000000064C",
		"00000000000000000000068C",
		"0000000000000000000006CC",
		"00000000000000000000070C",
		"00000000000000000000072C",
		"00000000000000000000074C",
		"00000000000000000000076C",
		"00000000000000000000078C",
		"0000000000000000000007AC",
		"0000000000000000000007CC",
		"0000000000000000000007EC",
		"00000000000000000000080C", /*"0",*/
	}
	table := simpletable.New()
	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "Pad?"},
			{Align: simpletable.AlignCenter, Text: "Float64"},
			{Align: simpletable.AlignCenter, Text: "Value"},
		},
	}
	table.SetStyle(StyleUnicodeTwo)
	for _, pa := range paths {
		val, err := HexadecimalToFloat64(pa[8:])
		if err != nil {
			continue
		}
		rk := []*simpletable.Cell{
			{Align: simpletable.AlignLeft, Text: pa[:8]},
			{Align: simpletable.AlignLeft, Text: strings.ToUpper(Float64ToHexadecimal(val))},
			{Align: simpletable.AlignLeft, Text: fmt.Sprintf("%g", val)},
		}
		table.Body.Cells = append(table.Body.Cells, rk)
	}
	fmt.Println(table.String())
}
