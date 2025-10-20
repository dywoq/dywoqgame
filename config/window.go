package config

// Window provides information about the game window.
type Window interface {
	Title() string
	Width() int
	Height() int
}
