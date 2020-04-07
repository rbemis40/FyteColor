# FyteColor
A simple package capable of adding color and styling to strings for output within seconds.

# Usage
## ColoredString
This struct stores all of the information for printing either the raw string, or the formatted form. Currently, colored strings can not be altered after creation, instead a new one must be generated.
### Methods
  * String() : Simply returns the unformatted, raw string of a `ColoredString`
  * ColorString() : Returns the fully formatted `ColoredString`

### Creating A ColoredString
A `ColoredString` can be created several ways, incuding:
  * NewColoredString(colorParts []ColoredPart) ColeredString : This takes an array of `ColoredPart` to manually construct a `ColoredString`
  * NewFormattedString(fStr string, formatArgs ...interface{}) ColoredString : This takes a printf-style set of arguments, however the formatArgs has several differences a user __needs__ to be aware of:
    1. The format specifier used to denote a __color__ is `$$`, rather than `%$`, important to avoid errors during negation
    2. Any color format args __must__ be first in the formatArgs
  * ColoredPrintf(fStr string, formatArgs ...interface{}) error : This directly prints the formatted `ColoredString`, rather than returning a `ColoredString`

## ColoredPart
A simple structure representing a part of a `ColoredString`, implements `FormattedPart`
### Creating A ColoredPart
Although useful when using `NewColoredString`, this structure should otherwise be avoided in favor of `ColoredString` methods when possible

## TemplateString
This struct is one of the most powerful features of FyteColor, allowing for the reuse of formatting, while still allowing for customization.
### Creating A TemplateString
Creating a `TemplateString` is currently only possible by assembling a series of `FormattedPart`. This is done using the NewTemplateString function.
  * NewTemplateString(formatParts []FormattedPart) TemplateString : Takes an array of `FormattedPart` and returns a new TemplateString

### Methods
  * ToColorString(insertions ...string) ColoredString : Converts the `FormattedPart` array provided in `NewTemplateString` into a fully formatted `ColoredString` , replacing the `TemplatePart` with the provided _insertions_

## TemplatePart
Represents a replaceable section of a `TemplateString` , implements `FormattedPart`. This should be used in `NewTemplateString` to add parts to the `TemplateString` that should be later replaced by _insertions_.
### Creating A TemplatePart
Currently, the only way to create a `TemplatePart` is by manually filling in the `PartColor` and `PartStyle` members of the struct. `PartString` serves no purpose, as it will later be filled by an _insertion_ when converting to a `ColoredString`.

## FormattedPart
An interface representing important attributes for constructing `ColoredString` . Types implementing this interface should contain a `PartString` , `PartColor` , and `PartStyle`.
### Methods
  * Color() Color : Returns the `Color` of the current `FormattedPart`
  * Style() Style : Returns the `Style` of the current `FormattedPart`
  * RawString() string : Returns a raw, unformatted string for the `FormattedPart`

# Why did I create this?
I created this as just a side project for fun, and to help me learn the Go language (which is very nice). Because of this, _feel free_ to both use and edit it.

