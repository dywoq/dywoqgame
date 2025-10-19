package entry

// Metadata contains methods to get information about the game.
type Metadata interface {
	Name() string
	Email() string
	Authors() []string
	Year() string
}
