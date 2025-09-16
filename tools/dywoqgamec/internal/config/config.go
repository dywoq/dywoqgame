package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	Title   string   `json:"title"`
	Version string   `json:"version"`
	Owner   string   `json:"owner"`
	Link    string   `json:"link"`
	Year    int      `json:"year"`
	Authors []string `json:"authors"`
	Scripts string   `json:"scripts"`
	Assets  string   `json:"assets"`
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

// Save saves the config back to JSON with the specified path.
func (c *Config) Save(path string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	enc := json.NewEncoder(f)
	enc.SetIndent("", "  ")
	return enc.Encode(c)
}
