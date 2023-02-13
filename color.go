package userio

type Color string

// escape codes for premixed theme
const (
	ColorInfo        Color = "\x1b[33m"
	ColorInstruction Color = "\x1b[38;5;140m"
	ColorMuted       Color = "\x1b[38;5;242m"
	ColorStandard    Color = "\x1b[38;5;69m"
)

// escape codes for highlighters (background + foreground)
const (
	HighlightGreen  Color = "\x1b[42;39m"
	HighlightRed    Color = "\x1b[41;97m"
	HighlightWhite  Color = "\x1b[107;30m"
	HighlightYellow Color = "\x1b[43;30m"
)

// escape codes for additional hues
const (
	ColorCyan    Color = "\x1b[36m"
	ColorGreen   Color = "\x1b[32m"
	ColorMagenta Color = "\x1b[35m"
	ColorRed     Color = "\x1b[31m"
)

const TextReset Color = "\x1b[0m"
