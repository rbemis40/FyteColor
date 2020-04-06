package color

//Style represents a style option, such as Bold, Italics, Underlined, or DefStyle to set the default style
type Style int

var (
	DefStyle   Style = 1 //Begins at 1 to avoid 0 == nil during translation in ColorString()
	Bold       Style = DefStyle + 1
	Italics    Style = Bold + 1
	Underlined Style = Italics + 1

	StyleMap map[Style]string = map[Style]string{
		DefStyle:   "\033[22;23;24m",
		Bold:       "\033[1m",
		Italics:    "\033[3m",
		Underlined: "\033[4m",
	}
)
