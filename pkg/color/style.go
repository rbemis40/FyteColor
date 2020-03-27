package color

type style int

var (
	StyleDef   style = 1 //Begins at 1 to avoid 0 == nil during translation in ColorString()
	Bold       style = StyleDef + 1
	Italics    style = Bold + 1
	Underlined style = Italics + 1

	StyleMap map[style]string = map[style]string{
		StyleDef:   "\033[22;23;24m",
		Bold:       "\033[1m",
		Italics:    "\033[3m",
		Underlined: "\033[4m",
	}
)
