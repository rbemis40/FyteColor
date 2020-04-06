# FyteColor
A simple package capable of adding color to strings and output within seconds.

# Usage
## ColoredString
This struct stores all of the information for printing either the raw string, or the formatted form. Currently, colored strings can not be altered after creation, instead a new one must be generated.
### Methods
  * String() : Simply returns the unformatted, raw string of a _ColoredString_
  * ColorString() : Returns the fully formatted _ColoredString_

### Creating A ColoredString
A _ColoredString_ can be created several ways, incuding:
  * NewColoredString(colorParts []ColoredPart) ColeredString : This takes an array of _ColoredParts_ to manually construct a _ColoredString_
  * NewFormattedString(fStr string, formatArgs ...interface{}) ColoredString : This takes a printf-style set of arguments, however the formatArgs has several differences a user __needs__ to be aware of:
    1. The format specifier used to denote a __color__ is `$$`, rather than `%$`, important to avoid errors during negation
    2. Any color format args __must__ be first in the formatArgs
  * ColoredPrintf(fStr string, formatArgs ...interface{}) error : This directly prints the formatted _ColoredString_, rather than returning a _ColoredString_

## ColoredPart
A simple structure representing a part of a _ColoredString_, implements _FormattedPart_
### Creating A ColoredPart
Although useful when using _NewColoredString_, this structure should otherwise be avoided in favor of _ColoredString_ methods when possible

## TemplateString
This struct is one of the most powerful features of FyteColor, allowing for the reuse of formatting, while still allowing for customization.
### Creating A TemplateString
Creating a _TemplateString_ is currently only possible by assembling a series of _FormattedPart_s. This is done using the NewTemplateString function.
  * NewTemplateString(formatParts []FormattedPart) TemplateString : Takes an array of _FormattedPart_s and returns a new TemplateString

### Methods
  * ToColorString(insertions ...string) ColoredString : Converts the _FormattedPart_s provided in _NewTemplateString_ into a fully formatted ColoredString, replacing the _TemplatePart_s with the provided _insertions_

## TemplatePart
Represents a replaceable section of a _TemplateString_, implements _FormattedPart_. This should be used in _NewTemplateString_ to add parts to the _TemplateString_ that should be later replaced by _insertions_.
### Creating A TemplatePart
Currently, the only way to create a _TemplatePart_ is by manually filling in the _PartColor_ and _PartStyle_ members of the struct. _PartString_ serves no purpose, as it will later be filled by an _insertion_ when converting to a _ColoredString_.

## FormattedPart
An interface representing important attributes for constructing _ColoredString_s. Types implementing this interface should contain a _PartString_, _PartColor_, and _PartStyle_.
### Methods
  * Color() Color : Returns the _Color_ of the current _FormattedPart_
  * Style() Style : Returns the _Style_ of the current _FormattedPart_
  * RawString() string : Returns a raw, unformatted string for the _FormattedPart_

# Why did I create this?
I created this as just a side project for fun, and to help me learn the Go language (which is very nice). Because of this, _feel free_ to both use and edit it.

