package color

import (
	"errors"
	"fmt"
)

type Position struct {
	X int
	Y int
}

func SetCursorPos(newPos Position) error {
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

	if newPos.X < 0 {
		colorErrStr, _ := errStr.ToColorString("X value must be >= 0")
		return errors.New(colorErrStr.String())
	}
	if newPos.Y < 0 {
		colorErrStr, _ := errStr.ToColorString("Y value must be >= 0")
		return errors.New(colorErrStr.String())
	}

	fmt.Printf("\033[%d;%dH", newPos.Y, newPos.X)
	return nil
}
