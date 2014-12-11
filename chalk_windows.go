package chalk

import "fmt"

// Black on windows is simply a wrapper around fmt.Sprintf
func Black(format string, a ...interface{}) string {
	return fmt.Sprintf(format, a...)
}

// Red on windows is simply a wrapper around fmt.Sprintf
func Red(format string, a ...interface{}) string {
	returfmt.Sprintf(format, a...)
}

// Green on windows is simply a wrapper around fmt.Sprintf
func Green(format string, a ...interface{}) string {
	return fmt.Sprintf(format, a...)
}

// Yellow on windows is simply a wrapper around fmt.Sprintf
func Yellow(format string, a ...interface{}) string {
	return fmt.Sprintf(format, a...)
}

// Blue on windows is simply a wrapper around fmt.Sprintf
func Blue(format string, a ...interface{}) string {
	return fmt.Sprintf(format, a...)
}

// Magenta on windows is simply a wrapper around fmt.Sprintf
func Magenta(format string, a ...interface{}) string {
	return fmt.Sprintf(format, a...)
}

// Cyan on windows is simply a wrapper around fmt.Sprintf
func Cyan(format string, a ...interface{}) string {
	return fmt.Sprintf(format, a...)
}

// White on windows is simply a wrapper around fmt.Sprintf
func White(format string, a ...interface{}) string {
	return fmt.Sprintf(format, a...)
}
