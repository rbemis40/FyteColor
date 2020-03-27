package color

import "fmt"

//ColoredPart represents a part of a ColoredString
type ColoredPart struct {
	partColor  color
	partString string
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
