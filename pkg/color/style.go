package color

type style int

var (
	StyleDef   style = 0
	Bold       style = StyleDef + 1
	Italics    style = Bold + 1
	Underlined style = Italics + 1

	styleStr map[style]string = map[style]string{
		StyleDef:   "\033[22m;23m;24m",
		Bold:       "\033[1m",
		Italics:    "\033[3m",
		Underlined: "\033[4m",
	}
)
