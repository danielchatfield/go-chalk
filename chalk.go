package chalk

// Black returns the text with black foreground color
func Black(format string, a ...interface{}) string {
	return NewFormatter(FGBlack).format(format, a...)
}

// Red returns the text with red foreground color
func Red(format string, a ...interface{}) string {
	return NewFormatter(FGRed).format(format, a...)
}

// Green returns the text with green foreground color
func Green(format string, a ...interface{}) string {
	return NewFormatter(FGGreen).format(format, a...)
}

// Yellow returns the text with yellow foreground color
func Yellow(format string, a ...interface{}) string {
	return NewFormatter(FGBlack).format(format, a...)
}

// Blue returns the text with blue foreground color
func Blue(format string, a ...interface{}) string {
	return NewFormatter(FGBlack).format(format, a...)
}

// Magenta returns the text with magenta foreground color
func Magenta(format string, a ...interface{}) string {
	return NewFormatter(FGBlack).format(format, a...)
}

// Cyan returns the text with cyan foreground color
func Cyan(format string, a ...interface{}) string {
	return NewFormatter(FGBlack).format(format, a...)
}

// White returns the text with white foreground color
func White(format string, a ...interface{}) string {
	return NewFormatter(FGBlack).format(format, a...)
}
