package main

import (
	"fmt"

	fytecolor "github.com/TrashWithNoLid/FyteColor/pkg/color"
)

func main() {
	coloredString := fytecolor.NewColoredString([]fytecolor.ColoredPart{
		fytecolor.ColoredPart{
			PartColor:  fytecolor.Red,
			PartString: "Hello ",
		},
		fytecolor.ColoredPart{
			PartColor:  fytecolor.Blue,
			PartString: "World!",
		},
	})

	fmt.Println(coloredString.ColorString())
}
