package color

import "fmt"

type color int

var (
	Def     color = 0
	Black   color = Def + 1
	Red     color = Black + 1
	Green   color = Red + 1
	Yellow  color = Green + 1
	Blue    color = Yellow + 1
	Magenta color = Blue + 1
	Cyan    color = Magenta + 1
	White   color = Cyan + 1

	colorStr map[color]string = map[color]string{
		Black:   "\033[30m",
		Red:     "\033[31m",
		Green:   "\033[32m",
		Yellow:  "\033[33m",
		Blue:    "\033[34m",
		Magenta: "\033[35m",
		Cyan:    "\033[36m",
		White:   "\033[37m",
	}

	curColor color = Def
)

//ColoredPart provides information for ColoredString
type ColoredPart struct {
	partColor  color
	partString string
}

//ColoredString represents a colored string
type ColoredString struct {
	rawString   string
	stringParts []ColoredPart
}

//SetColor sets the current color for string building
func SetColor(newColor color) {
	curColor = newColor
}

//BuildStr builds the input string, adding color as needed
func BuildStr(str string) (string, error) {
	colorANSI, ok := colorStr[curColor]
	if !ok {
		return str, fmt.Errorf("BuildStr: error building str '%s', color '%v' is not defined", str, curColor)
	}

	return colorANSI + str, nil
}
