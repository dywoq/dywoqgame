package core

type Window struct {
	Title         string
	Width, Height int
}

func NewWindow(title string, width, height int) *Window {
	return &Window{title, width, height}
}
