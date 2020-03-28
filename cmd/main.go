package main

import (
	"fmt"

	"github.com/TrashWithNoLid/FyteColor/pkg/color"
	fytecolor "github.com/TrashWithNoLid/FyteColor/pkg/color"
)

func main() {
	coloredString, err := fytecolor.NewFormattedStr("%!Hello %!World", color.Red, color.Blue) //No Error
	if err != nil {
		fmt.Printf("Error building coloredString: %v\n", err)
		return
	}
	fmt.Println(coloredString.ColorString())

	coloredString, err = fytecolor.NewFormattedStr(
		"%!The number %!%d%! is pretty cool",
		color.BrightRed,
		color.Blue,
		color.BrightRed,
		77) //No Error, notice the %d arg is last

	if err != nil {
		fmt.Printf("Error building coloredString%v\n", err)
		return
	}

	fmt.Println(coloredString.ColorString())

	coloredString, err = fytecolor.NewFormattedStr("%!%d: %!Will be an error, as the printf-style args must come after colored ones!", color.Green, 10, color.ColorDef)
	if err != nil {
		fmt.Printf("Error building coloredString: %v\n", err)
		return
	}
}
