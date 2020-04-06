package color

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
	for _, curPart := range formatParts {
		switch v := curPart.(type) {
		case *TemplatePart:
			continue
		case *ColoredPart:
			continue
		}
	}
	return TemplateString{}
}
