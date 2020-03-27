package main

import (
	"fmt"

	fytecolor "github.com/TrashWithNoLid/FyteColor/pkg/color"
)

func main() {
	coloredString := fytecolor.NewColoredString([]fytecolor.ColoredPart{
		fytecolor.ColoredPart{
			PartColor:  fytecolor.BrightRed,
			PartStyle:  fytecolor.Underlined,
			PartString: "Hello ",
		},
		fytecolor.ColoredPart{
			PartColor:  fytecolor.BrightBlue,
			PartString: "World!",
		},
	})

	fmt.Println(coloredString.ColorString())
}
