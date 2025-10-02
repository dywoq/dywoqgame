package entry

// Window containing the information about a game, like title and dimension size.
type Window struct {
	Title    string `json:"title"`
	Width    int    `json:"width"`
	Height   int    `json:"height"`
	Resizing bool   `json:"resizing"`
}
