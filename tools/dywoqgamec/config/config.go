package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	Title string `json:"title"`
}

// Fill fills the config instance with the information from the Json config, which located at path.
func (c *Config) Fill(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()
	return json.NewDecoder(f).Decode(c)
}
