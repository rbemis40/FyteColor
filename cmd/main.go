package main

import (
	"fmt"

	fytecolor "github.com/TrashWithNoLid/FyteColor/pkg/color"
)

func main() {
	/* Construct a ColoredString manually using ColoredParts */
	colorString := fytecolor.NewColoredString([]fytecolor.ColoredPart{
		fytecolor.ColoredPart{
			PartString: "Hello ",
			PartColor:  fytecolor.BrightBlue,
			PartStyle:  fytecolor.Underlined,
		},
		fytecolor.ColoredPart{
			PartString: "World!",
			PartColor:  fytecolor.BrightRed,
		},
	})

	fmt.Println(colorString.String())
	fmt.Println(colorString.ColorString())

	/* Construct a ColoredString using printf-like syntax */
	colorString, err := fytecolor.NewFormattedStr(
		"$$$$Hello $$World, And other things can be added, like %d\n",
		fytecolor.Underlined,
		fytecolor.BrightBlue,
		fytecolor.BrightRed,
		10,
	)
	if err != nil {
		fmt.Printf("Error formatting new ColoredString: %v\n", err)
		return
	}

	fmt.Println(colorString.String())
	fmt.Println(colorString.ColorString())

	/* Or the printf syntax can be printed directly to the terminal */
	/* Keep in mind that when using the printf syntax, color arguments MUST come first */
	err = fytecolor.ColoredPrintf(
		"$$$$%s $$%s!\n",
		fytecolor.Underlined,
		fytecolor.BrightBlue,
		fytecolor.BrightRed,
		"Hello",
		"World",
	)
	if err != nil {
		fmt.Printf("Error printing ColoredString: %v\n", err)
		return
	}

	/* So this would cause an error */
	err = fytecolor.ColoredPrintf("%d$$,%d\n", 10, fytecolor.BrightCyan, 10)
	if err != nil {
		fmt.Printf("Error printing ColoredString: %v\n", err)
		return
	}
}
