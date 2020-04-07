package color

//Color represents a color option such as Black, Red, Green, or DefColor for the default color
type Color int

const (
	/* Defines the various foreground colors, used in ColorMap */
	_              = iota //Offset the values to begin at 1, prevents DefColor == nil in strbuilder
	DefColor Color = iota
	Black    Color = iota
	Red      Color = iota
	Green    Color = iota
	Yellow   Color = iota
	Blue     Color = iota
	Magenta  Color = iota
	Cyan     Color = iota
	White    Color = iota

	BrightBlack   Color = iota
	BrightRed     Color = iota
	BrightGreen   Color = iota
	BrightYellow  Color = iota
	BrightBlue    Color = iota
	BrightMagenta Color = iota
	BrightCyan    Color = iota
	BrightWhite   Color = iota

	/* Defines the various background colors, used in ColorMap */
	BkgBlack   Color = iota
	BkgRed     Color = iota
	BkgGreen   Color = iota
	BkgYellow  Color = iota
	BkgBlue    Color = iota
	BkgMagenta Color = iota
	BkgCyan    Color = iota
	BkgWhite   Color = iota

	BkgBrightBlack   Color = iota
	BkgBrightRed     Color = iota
	BkgBrightGreen   Color = iota
	BkgBrightYellow  Color = iota
	BkgBrightBlue    Color = iota
	BkgBrightMagenta Color = iota
	BkgBrightCyan    Color = iota
	BkgBrightWhite   Color = iota
)

//ColorMap provides a map between a Color and its ANSI string value
var ColorMap map[Color]string = map[Color]string{
	/* Foreground colors */
	DefColor:      "\033[39;49m",
	Black:         "\033[30m",
	Red:           "\033[31m",
	Green:         "\033[32m",
	Yellow:        "\033[33m",
	Blue:          "\033[34m",
	Magenta:       "\033[35m",
	Cyan:          "\033[36m",
	White:         "\033[37m",
	BrightBlack:   "\033[90m",
	BrightRed:     "\033[91m",
	BrightGreen:   "\033[92m",
	BrightYellow:  "\033[93m",
	BrightBlue:    "\033[94m",
	BrightMagenta: "\033[95m",
	BrightCyan:    "\033[96m",
	BrightWhite:   "\033[97m",

	/* Background colors */
	BkgBlack:         "\033[40m",
	BkgRed:           "\033[41m",
	BkgGreen:         "\033[42m",
	BkgYellow:        "\033[43m",
	BkgBlue:          "\033[44m",
	BkgMagenta:       "\033[45m",
	BkgCyan:          "\033[46m",
	BkgWhite:         "\033[47m",
	BkgBrightBlack:   "\033[100m",
	BkgBrightRed:     "\033[101m",
	BkgBrightGreen:   "\033[102m",
	BkgBrightYellow:  "\033[103m",
	BkgBrightBlue:    "\033[104m",
	BkgBrightMagenta: "\033[105m",
	BkgBrightCyan:    "\033[106m",
	BkgBrightWhite:   "\033[107m",
}
