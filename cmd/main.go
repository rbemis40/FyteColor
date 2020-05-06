package main

import (
	"fmt"

	fytecolor "github.com/TrashWithNoLid/FyteColor/color"
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

	fmt.Println("")

	/* Construct a ColoredString using printf-like syntax */
	colorString, err := fytecolor.NewFormattedStr(
		"$$$$Using $$Printf-syntax, other things can be added, like %d\n",
		fytecolor.Underlined,
		fytecolor.BrightBlue,
		fytecolor.BrightRed,
		10,
	)
	if err != nil {
		fmt.Printf("Error formatting new ColoredString: %v\n", err)
		return
	}

	fmt.Printf(colorString.ColorString())
	fmt.Println(colorString.String())

	/* Or the printf syntax can be printed directly to the terminal */
	/* Keep in mind that when using the printf syntax, color and style arguments MUST come first */
	err = fytecolor.ColoredPrintf(
		"$$$$$$%s$$ $$%s!\n",
		fytecolor.BkgBrightWhite,
		fytecolor.Underlined,
		fytecolor.BrightBlue,
		fytecolor.DefColor,
		fytecolor.BrightRed,
		"Colored",
		"Printf",
	)
	if err != nil {
		fmt.Printf("Error printing ColoredString: %v\n", err)
		return
	}

	fmt.Println("")

	/* So this would cause an error */
	err = fytecolor.ColoredPrintf("%d$$,%d\n", 10, fytecolor.BrightCyan, 10)
	if err != nil {
		fmt.Printf("Error printing ColoredString: %v\n", err)
	}

	/* Create a new colored TemplateString, allowing for sections of the string to be reused */
	helloTemp := fytecolor.NewTemplateString([]fytecolor.FormattedPart{
		&fytecolor.ColoredPart{
			PartColor:  fytecolor.BrightRed,
			PartStyle:  fytecolor.DefStyle,
			PartString: "This is a template, where parts can be ",
		},
		&fytecolor.TemplatePart{
			PartColor: fytecolor.BrightBlue,
			PartStyle: fytecolor.DefStyle,
		},
	})

	fmt.Println("")

	/* Replace the TemplatePart with 'changed!' */
	tempColorStr, err := helloTemp.ToColorString("changed!")
	if err != nil {
		fmt.Printf("Error converting helloTemp to string: %v\n", err)
		return
	}
	fmt.Println(tempColorStr.ColorString())

	/* Replace the TemplatePart with 'replaced!' */
	tempColorStr, err = helloTemp.ToColorString("replaced!")
	if err != nil {
		fmt.Printf("Error converting helloTemp to string: %v\n", err)
		return
	}

	fmt.Println(tempColorStr.ColorString())
}
