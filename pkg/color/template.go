package color

//TemplatePart holds information about template parts
type TemplatePart struct {
	PartColor  color  //Stores the parts color
	PartStyle  style  //Stores the parts style
	PartString string //Holds the string for the template
}

//Color returns the TemplatePart color
func (tp *TemplatePart) Color() color {
	return tp.PartColor
}

//Style returns the TemplatePart style
func (tp *TemplatePart) Style() style {
	return tp.PartStyle
}

//RawString returns the TemplatePart string
func (tp *TemplatePart) RawString() string {
	return tp.PartString
}

//TemplateString holds crucial data for templated strings
type TemplateString struct {
	formatParts []FormattedPart //Holds the formatted parts
}

//NewTemplateString takes FormattedParts to return a fully formed TemplateString
func NewTemplateString(formatParts []FormattedPart) TemplateString {
	//TODO: Implement NewTemplate String
	return TemplateString{}
}
