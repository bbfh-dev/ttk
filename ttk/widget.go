package ttk

type Widget interface {
	Element
	Render(width, height int) string
	ExpandWidth() bool
	ExpandHeight() bool
}
