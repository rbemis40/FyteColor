package color

import (
	"errors"
	"fmt"
)

func SetCursorPos(x, y int) error {
	errStr := NewTemplateString([]FormattedPart{
		&ColoredPart{
			PartColor:  DefColor,
			PartStyle:  DefStyle,
			PartString: "Error setting cursor position: ",
		},
		&TemplatePart{
			PartColor: DefColor,
			PartStyle: DefStyle,
		},
	})

	if x < 0 {
		colorErrStr, _ := errStr.ToColorString("X value must be >= 0")

		return errors.New(colorErrStr.String())
	}

	fmt.Printf("\033[%d;%dH", y, x)
	return nil
}
