package ttk

type Element interface {
	// Initializes internal element fields
	Init()
}

type Orientation bool

const (
	VERTICAL   Orientation = false
	HORIZONTAL Orientation = true
)
