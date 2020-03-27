package color

import "fmt"

//ColoredPart represents a part of a ColoredString
type ColoredPart struct {
	PartColor  color
	PartStyle  style
	PartString string
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

//Translates the raw []ColoredPart to a single formatted string
func (cs *ColoredString) translateColorData() {
	var colorString string = ""
	for _, cp := range cs.coloredParts {
		clrStr, ok := ColorMap[cp.PartColor]
		if ok {
			colorString += clrStr
		} else {
			colorString += ColorMap[ColorDef]
		}

		styleStr, ok := StyleMap[cp.PartStyle]
		if ok {
			colorString += styleStr
		} else {
			colorString += StyleMap[StyleDef]
		}

		colorString += cp.PartString
	}

	colorString += ColorMap[ColorDef]
	cs.formattedString = colorString
}
