package color

//Color represents a color option such as Black, Red, Green, or DefColor for the default color
type Color int

var (
	DefColor Color = 1 //Begins at 1 to avoid 0 == nil during ColorString() translation
	Black    Color = DefColor + 1
	Red      Color = Black + 1
	Green    Color = Red + 1
	Yellow   Color = Green + 1
	Blue     Color = Yellow + 1
	Magenta  Color = Blue + 1
	Cyan     Color = Magenta + 1
	White    Color = Cyan + 1

	BrightBlack   Color = White + 1
	BrightRed     Color = BrightBlack + 1
	BrightGreen   Color = BrightRed + 1
	BrightYellow  Color = BrightGreen + 1
	BrightBlue    Color = BrightYellow + 1
	BrightMagenta Color = BrightBlue + 1
	BrightCyan    Color = BrightMagenta + 1
	BrightWhite   Color = BrightCyan + 1

	ColorMap map[Color]string = map[Color]string{
		DefColor:      "\033[39m",
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
	}
)
