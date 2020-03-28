package main

import (
	"fmt"

	"github.com/TrashWithNoLid/FyteColor/pkg/color"
	fytecolor "github.com/TrashWithNoLid/FyteColor/pkg/color"
)

func main() {
	/*
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
	*/

	//fmt.Println(coloredString.ColorString())

	coloredString, err := fytecolor.NewFormattedStr("%!Hello %!World!", color.Red, color.Blue)
	if err != nil {
		fmt.Printf("Error creating format color: %v\n", err)
		return
	}

	fmt.Println(coloredString.ColorString())
}
