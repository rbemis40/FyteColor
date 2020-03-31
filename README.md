# FyteColor
A simple package capable of adding color to strings and output within minutes.

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
A simple structure representing a part of a _ColoredString_
### Creating A ColoredPart
Although useful when using _NewColoredString_, this structure should otherwise be avoided in favor of _ColoredString_ methods when possible

# Why did I create this?
I created this as just a side project for fun, and to help me learn the Go language (which is very nice). Because of this, _feel free_ to both use and edit it.

