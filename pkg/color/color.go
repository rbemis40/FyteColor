package color

type color int

var (
	ColorDef color = 1 //Begins at 1 to avoid 0 == nil during ColorString() translation
	Black    color = ColorDef + 1
	Red      color = Black + 1
	Green    color = Red + 1
	Yellow   color = Green + 1
	Blue     color = Yellow + 1
	Magenta  color = Blue + 1
	Cyan     color = Magenta + 1
	White    color = Cyan + 1

	BrightBlack   color = White + 1
	BrightRed     color = BrightBlack + 1
	BrightGreen   color = BrightRed + 1
	BrightYellow  color = BrightGreen + 1
	BrightBlue    color = BrightYellow + 1
	BrightMagenta color = BrightBlue + 1
	BrightCyan    color = BrightMagenta + 1
	BrightWhite   color = BrightCyan + 1

	ColorMap map[color]string = map[color]string{
		ColorDef:      "\033[39m",
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
