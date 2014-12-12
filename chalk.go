// +build !windows

package chalk

// Format formats the string with the given params
func Format(param Attribute, format string, a ...interface{}) string {
	return NewFormatter(param).format(format, a...)
}

// Black returns the text with black foreground color
func Black(format string, a ...interface{}) string {
	return NewFormatter(FOREGROUND+BLACK).format(format, a...)
}

// Red returns the text with red foreground color
func Red(format string, a ...interface{}) string {
	return NewFormatter(FOREGROUND+RED).format(format, a...)
}

// Green returns the text with green foreground color
func Green(format string, a ...interface{}) string {
	return NewFormatter(FOREGROUND+GREEN).format(format, a...)
}

// Yellow returns the text with yellow foreground color
func Yellow(format string, a ...interface{}) string {
	return NewFormatter(FOREGROUND+YELLOW).format(format, a...)
}

// Blue returns the text with blue foreground color
func Blue(format string, a ...interface{}) string {
	return NewFormatter(FOREGROUND+BLUE).format(format, a...)
}

// Magenta returns the text with magenta foreground color
func Magenta(format string, a ...interface{}) string {
	return NewFormatter(FOREGROUND+MAGENTA).format(format, a...)
}

// Cyan returns the text with cyan foreground color
func Cyan(format string, a ...interface{}) string {
	return NewFormatter(FOREGROUND+CYAN).format(format, a...)
}

// White returns the text with white foreground color
func White(format string, a ...interface{}) string {
	return NewFormatter(FOREGROUND+WHITE).format(format, a...)
}
