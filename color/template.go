package color

import "fmt"

//TemplatePart holds information about template parts
type TemplatePart struct {
	PartColor  Color  //Stores the parts Color
	PartStyle  Style  //Stores the parts Style
	PartString string //Holds the string for the template
}

//Color returns the TemplatePart Color
func (tp *TemplatePart) Color() Color {
	return tp.PartColor
}

//Style returns the TemplatePart Style
func (tp *TemplatePart) Style() Style {
	return tp.PartStyle
}

//RawString returns the TemplatePart string
func (tp *TemplatePart) RawString() string {
	return tp.PartString
}

//TemplateString holds crucial data for templated strings
type TemplateString struct {
	formatParts []FormattedPart //Holds the formatted parts
	stringSegs  []string        //Holds the constant formatted segments of the string
}

//NewTemplateString takes FormattedParts to return a fully formed TemplateString
func NewTemplateString(formatParts []FormattedPart) TemplateString {
	var templateString TemplateString = TemplateString{ //Create a new TemplateString
		formatParts: formatParts,
		stringSegs:  make([]string, 0, 2),
	}

	for _, curPart := range formatParts {
		coloredPart, ok := curPart.(*ColoredPart) //Check if the current formatted part is a coloredPart

		if ok { //If it is, add it to the stringSegs array
			segArray := &templateString.stringSegs                 //Get the pointer to the stringSegs array to avoid tediousness
			*segArray = append(*segArray, coloredPart.translate()) //Translate the ColoredPart into a string and store it
		}
	}
	return templateString
}

//ToColorString converts a template string into a usable color string using the insertions to replace TemplateParts
func (ts TemplateString) ToColorString(insertions ...string) (ColoredString, error) {
	insertLen := len(insertions)
	templateLen := len(ts.formatParts) - len(ts.stringSegs)

	if insertLen != templateLen { //Check if the number of insertions is equal to the number of TemplateParts
		return ColoredString{}, fmt.Errorf("failed to convert template to colored string, received %d insertions, expected %d", insertLen, templateLen)
	}

	var newColorStr ColoredString = ColoredString{ //Represents the ColoredString that will be returned
		rawString:       "",
		formattedString: "",
		coloredParts:    make([]ColoredPart, 0, 2),
	}

	var insertIndex int = 0  //Stores the current insertion, incremented once TemplatePart is found
	var segmentIndex int = 0 //Stores the current segment index for when a colored string is found, uses the constant variant

	for _, curPart := range ts.formatParts { //Build the newColorStr
		var newColorPart ColoredPart //Stores the current ColoredPart to be added onto the newColorStr.coloredParts

		switch v := curPart.(type) { //Check if the curPart is a TemplatePart or a ColoredPart
		case *TemplatePart:
			newColorPart = ColoredPart{ //Create the ColoredPart manually
				PartColor:  v.PartColor,
				PartStyle:  v.PartStyle,
				PartString: insertions[insertIndex],
			}

			newColorStr.formattedString += newColorPart.translate() //Add the newly translated ColorPart onto the ColoredString

			insertIndex++ //Increment the insertIndex to the next insertion
		case *ColoredPart:
			newColorStr.formattedString += ts.stringSegs[segmentIndex] //Add the previously constructed string to the newColorStr
			newColorPart = *v

			segmentIndex++ //Increment the segmentIndex for when the next ColoredPart is found
		}

		colorParts := &newColorStr.coloredParts
		*colorParts = append(*colorParts, newColorPart)
	}

	newColorStr.formattedString += ColorMap[DefColor] //Reset the color at the end of the string
	newColorStr.formattedString += StyleMap[DefStyle] //Reset the style at the end of the string

	return newColorStr, nil
}
