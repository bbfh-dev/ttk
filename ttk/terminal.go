package ttk

const MIN_WIDTH = 20
const MIN_HEIGHT = 8

// The dimensions of the terminal window
//
// NOTE: You must manually set this global variable to match the window size!
var (
	Width  = 0
	Height = 0
)

func IsTerminalTooSmall() bool {
	return Width < MIN_WIDTH || Height < MIN_HEIGHT
}
