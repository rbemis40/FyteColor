package color

//TemplatePart holds information about template parts
type TemplatePart struct {
	PartString string //Holds the string for the template
}

//TemplateString holds crucial data for templated strings
type TemplateString struct {
	rawString string         //Holds the raw string data
	tempParts []TemplatePart //Holds the templated parts
}

//NewTemplateString takes TemplateParts to return a fully formed TemplateString
func NewTemplateString(tempParts []TemplatePart) TemplateString {
	//TODO: Implement NewTemplate String
	return TemplateString{}
}
