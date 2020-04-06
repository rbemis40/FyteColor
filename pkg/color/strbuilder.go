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

	newString.translateColorData()

	return newString
}

//NewFormattedStr creates a colored string from a printf Style string
//
//- **Note:** Color based arguments must come before printf arguments
func NewFormattedStr(fStr string, formatArgs ...interface{}) (ColoredString, error) {
	var isNegated bool = false
	var curString string = ""

	const FormatSpecifier = "$$"

	partStr := make([]string, 0, 3)

	for i := 0; i < len(fStr); { //Loop through the format string to detect FormatSpecifier

		if fStr[i] == '\\' { //Check if negation is needed
			if isNegated { //If already negated, print the missed backslash, but keep negation
				curString += "\\"
			} else { //If not already negated, set the negation to true, but don't print the backslash in case it is negating the Color code
				isNegated = true
			}
		} else { //The current character is not negative
			if len(fStr)-i >= 2 && fStr[i:i+2] == FormatSpecifier { //Before proceeding, check if the Color code specifier is present
				if isNegated { //If it is and negated, remove the negation and print the characters
					isNegated = false
				} else { //If not negated, store the previous characters, as they will be differently colored
					partStr = append(partStr, curString) //Add the current str to the slice, this is the part that will be uniquely colored
					curString = ""                       //Reset the current string

					i += 2   //Since already found, skip pass the next two characters
					continue //Dont add the characters to the new curString
				}
			} else { //The format specifier was not found
				if isNegated { //If it is negated but not a format specifier, than set negation back to false and add the backslashe for later formatting
					isNegated = false
					curString += "\\"
				}
			}

			curString += string(fStr[i]) //Add the character to the current string
		}

		i++ //Change the current string index
	}

	partStr = append(partStr, curString)

	if argLen := len(formatArgs); argLen < len(partStr)-1 { //Subtract 1, as there is always at least one string
		return ColoredString{}, fmt.Errorf("not enough arguments passed to NewFormattedStr, only received %d", argLen)
	}

	/* Generate the ColoredParts needed to translate into fully formatted string */
	colorPartSlice := make([]ColoredPart, 1, 3)
	colorPartSlice[0] = ColoredPart{
		PartString: partStr[0],
	}

	/* Loop through the string parts found above, and check if the corresponding argument is a Color or a Style. Then, construct a ColoredPart */
	for i, curPartStr := range partStr[1:] {
		switch v := formatArgs[i].(type) {
		case Color:
			colorPartSlice = append(colorPartSlice, ColoredPart{
				PartColor:  v,
				PartString: curPartStr,
			})
		case Style:
			colorPartSlice = append(colorPartSlice, ColoredPart{
				PartStyle:  v,
				PartString: curPartStr,
			})
		default:
			return ColoredString{}, fmt.Errorf("argument %d of formatArgs was of the wrong type '%T' (hint: Color arguments must be first)", i, formatArgs[i])
		}
	}

	/* Create a new ColoredString and then format it with Sprintf */
	partialColoredStr := NewColoredString(colorPartSlice)
	partialColoredStr.formattedString = fmt.Sprintf(partialColoredStr.formattedString, formatArgs[len(partStr)-1:]...)

	return partialColoredStr, nil
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
func (cs *ColoredString) translateColorData() {
	var colorString string = ""
	for _, cp := range cs.coloredParts {

		/* If the Color code exists, find the string for it and use it, else continue using the last Color */
		ansiStr, ok := ColorMap[cp.PartColor]
		if ok {
			colorString += ansiStr
		}

		/* If the Style code exists, find the string for it and use it, else continue using the last Style */
		ansiStr, ok = StyleMap[cp.PartStyle]
		if ok {
			colorString += ansiStr
		}

		colorString += cp.PartString
	}

	/* Ensure the colored string always resets the colors at the end */
	colorString += ColorMap[DefColor]
	colorString += StyleMap[DefStyle]

	cs.formattedString = colorString
}
