package color

import "fmt"

//ColoredPart represents a part of a ColoredString
type ColoredPart struct {
	PartColor  color
	PartString string
}

//ColoredString represents a string with colored elements
type ColoredString struct {
	rawString    string
	coloredParts []ColoredPart
}

//String returns the raw string of a ColoredString
func (cs ColoredString) String() string {
	return fmt.Sprint(cs.rawString)
}

//ColorString returns the fully formatted colored string
func (cs ColoredString) ColorString() string {
	var colorString string = ""
	for _, cp := range cs.coloredParts {
		ansiStr, ok := ColorMap[cp.PartColor]
		if ok {
			colorString += ansiStr + cp.PartString
		} else {
			colorString += cp.PartString
		}
	}

	colorString += ColorMap[ColorDef]

	return colorString
}

//NewColoredString returns a fully formed ColoredString
func NewColoredString(colorParts []ColoredPart) ColoredString {
	var rawString string = ""
	for _, cp := range colorParts {
		rawString += cp.PartString
	}

	return ColoredString{
		rawString,
		colorParts,
	}
}
