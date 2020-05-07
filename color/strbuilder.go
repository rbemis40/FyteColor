package color

import (
	"fmt"
)

//FormattedPart represents a general part of a string
type FormattedPart interface {
	Color() Color      //Color returns the PartColor
	Style() Style      //Style returns the PartStyle
	RawString() string //String returns the PartString
}

//ColoredPart represents a part of a ColoredString
type ColoredPart struct {
	PartColor  Color
	PartStyle  Style
	PartString string
}

//Color returns the ColoredPart Color
func (cp *ColoredPart) Color() Color {
	return cp.PartColor
}

//Style returns the ColoredPart Style
func (cp *ColoredPart) Style() Style {
	return cp.PartStyle
}

//RawString returns the ColoredPart string
func (cp *ColoredPart) RawString() string {
	return cp.PartString
}

//ColoredString represents a string with colored elements
type ColoredString struct {
	rawString       string
	formattedString string
	coloredParts    []ColoredPart
}

//String returns the raw string of a ColoredString
func (cs *ColoredString) String() string {
	return fmt.Sprint(cs.rawString)
}

//ColorString returns the fully formatted colored string
func (cs *ColoredString) ColorString() string {
	return cs.formattedString
}

//NewColoredString returns a fully formed ColoredString
func NewColoredString(colorParts []ColoredPart) ColoredString {
	var rawString string = ""
	for _, cp := range colorParts {
		rawString += cp.PartString
	}

	newString := ColoredString{
		rawString:    rawString,
		coloredParts: colorParts,
	}

	newString.translate()

	return newString
}

func findNegation(searchStr string) bool {
	if searchStr[0] == '\\' {
		return true
	}

	return false
}

func findSpecifier(searchStr, specStr string) bool {
	if len(searchStr) >= len(specStr) && searchStr[:len(specStr)] == specStr {
		return true;
	}

	return false;
}

//NewFormattedStr creates a colored string from a printf Style string
//
//- **Note:** Color based arguments must come before printf arguments
func NewFormattedStr(fStr string, formatArgs ...interface{}) (ColoredString, error) {
	var isNegated bool = false
	var isSpecifier bool = false

	const ColorSpecifier = "$c"
	var curPart ColoredPart = ColoredPart {
		PartColor: DefColor,
		PartStyle: DefStyle,
		PartString: "",
	}
	var fullParts []ColoredPart = make([]ColoredPart, 0, 1)

	for i := 0; i < len(fStr); i++ {
		fmt.Printf("[%d] : isNegated: %t\n", i, isNegated)

		isSpecifier = findSpecifier(fStr[i:], ColorSpecifier)
		if isSpecifier {
			if isNegated {
				
			}
		} else {
			
		}


		/*
		if findNegation(fStr[i:]) {
			isNegated = true

			isSpecifier = findSpecifier(fStr[i+1:], ColorSpecifier)
			if isSpecifier {
				continue
			}
		} else {
			isSpecifier = findSpecifier(fStr[i:], ColorSpecifier)
			if isSpecifier && !isNegated{
				fmt.Println(curPart.PartString)
				fullParts = append(fullParts, curPart)
				curPart = ColoredPart {
					PartColor: DefColor,
					PartStyle: DefStyle,
					PartString: "",
				}

				//Manually increment i to skip the rest of the characters in the colorspecifier
				i += len(ColorSpecifier)-1
				continue
			}
		}
		isNegated = false
		curPart.PartString += string(fStr[i])

		*/
	}

	fullParts = append(fullParts, curPart)
	fmt.Printf("FullPartLen: %d\n", len(fullParts))


	return ColoredString{}, nil
}

//ColoredPrintf creates than prints a new coloredstring using printf-like arguments
//
//- **Note:** Color based arguments must come before the printf arguments
func ColoredPrintf(fStr string, formatArgs ...interface{}) error {

	coloredString, err := NewFormattedStr(fStr, formatArgs...)
	if err != nil {
		return fmt.Errorf("coloredprintf: error creating coloredString: %v", err)
	}

	fmt.Print(coloredString.ColorString())

	return nil
}

//Translates the raw []ColoredPart to a single formatted string
func (cs *ColoredString) translate() {
	var colorString string = ""
	for _, cp := range cs.coloredParts {
		colorString += cp.translate() //Use the current color part and translate it into a usable string
	}

	/* Ensure the colored string alwyas resets the colors at the end of the string */
	colorString += ColorMap[DefColor]
	colorString += StyleMap[DefStyle]

	cs.formattedString = colorString //Set the formattedString to the newly created string
}

func (cp *ColoredPart) translate() string {
	var colorString string = ""

	/* If the Color code exists, find the string for it and use it, else continue using the last color */
	ansiStr, ok := ColorMap[cp.PartColor]
	if ok {
		colorString += ansiStr
	}

	/* If the Style code exists, find the string for it and use it, else continue using the last style */
	ansiStr, ok = StyleMap[cp.PartStyle]
	if ok {
		colorString += ansiStr
	}

	return colorString + cp.PartString
}
