package chalk

// Escape is the terminal escape literal
const Escape = "\x1b"

// SGR parameters as specified here: http://en.wikipedia.org/wiki/ANSI_escape_code
const (
	Reset Attribute = iota
	Bold
	Faint         // Not widely supported
	Italic        // Not widely supported
	Underline     //
	BlinkSlow     // Less than 150 per min
	BlinkRapid    // 150+ per minute
	Negative      // Swap foreground and background
	Conceal       // Not widely supported
	StrikeThrough // Not widely supported
)

// Foreground colors
const (
	FGBlack Attribute = iota + 30
	FGRed
	FGGreen
	FGYellow
	FGBlue
	FGMagenta
	FGCyan
	FGWhite
)

// Background colors
const (
	BGBlack Attribute = iota + 30
	BGRed
	BGGreen
	BGYellow
	BGBlue
	BGMagenta
	BGCyan
	BGWhite
)
