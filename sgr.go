package chalk

// Escape is the terminal escape literal
const ESC = "\x1b"

// SGR parameters as specified here: http://en.wikipedia.org/wiki/ANSI_escape_code
const (
	RESET Attribute = iota
	BOLD
	FAINT         // Not widely supported
	ITALIC        // Not widely supported
	UNDERLINE     //
	BLINKSLOW     // Less than 150 per min
	BLINKRAPID    // 150+ per minute
	NEGATIVE      // Swap foreground and background
	CONCEAL       // Not widely supported
	STRIKETHROUGH // Not widely supported
)

// Foreground, background, and intense offset
const (
	FOREGROUND Attribute = 30
	BACKGROUND Attribute = 40
	INTENSE    Attribute = 60
)

// Colors
const (
	BLACK Attribute = iota
	RED
	GREEN
	YELLOW
	BLUE
	MAGENTA
	CYAN
	WHITE
)
