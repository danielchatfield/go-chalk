package chalk

import (
	"fmt"
	"strconv"
)

// Formatter represents a formatter with a list of attributes to apply
type Formatter struct {
	params []Attribute
}

// Attribute defines the SGR codes
type Attribute int

// NewFormatter returns a new Formatter with the given params
func NewFormatter(value ...Attribute) *Formatter {
	f := &Formatter{
		params: make([]Attribute, 0),
	}

	f.Add(value...)
	return f
}

// Add adds an attribute to the formatter
func (f *Formatter) Add(value ...Attribute) *Formatter {
	f.params = append(f.params, value...)
	return f
}

func (f *Formatter) format(format string, a ...interface{}) string {
	return fmt.Sprintf(
		"%s[%sm%s%s[%dm",
		ESC,
		f.sequence(),
		fmt.Sprintf(format, a...),
		ESC,
		RESET,
	)
}

// sequence returns an SGR sequence
func (f *Formatter) sequence() string {
	seq := ""
	for i, v := range f.params {
		if i > 0 {
			seq += ";"
		}
		seq += strconv.Itoa(int(v))
	}
	return seq
}
