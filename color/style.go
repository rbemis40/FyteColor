package color

//Style represents a style option, such as Bold, Italics, Underlined, or DefStyle to set the default style
type Style int

const (
	/* Defines the various styles that can be used */
	DefStyle Style = iota + 1
	Bold
	Italics
	Underlined
)

//StyleMap is used to find the ANSI code for each provided style
var StyleMap map[Style]string = map[Style]string{
	DefStyle:   "\033[22;23;24m",
	Bold:       "\033[1m",
	Italics:    "\033[3m",
	Underlined: "\033[4m",
}
