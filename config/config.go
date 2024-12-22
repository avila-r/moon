package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/pelletier/go-toml/v2"
)

type Properties struct {
	Info struct {
		Command          string `toml:"command"`
		ShortDescription string `toml:"short"`
		LongDescription  string `toml:"long"`

		License    string `toml:"license"`
		Author     string `toml:"author"`
		Version    string `toml:"version"`
		Repository string `toml:"repository"`
	} `toml:"info"`

	Advanced struct {
		Debug bool `toml:"debug"`
	} `toml:"advanced"`

	Genesis struct {
		Port int `toml:"port"`
	} `toml:"genesis"`
}

var (
	config = func() Properties {
		cfg := Properties{}

		dir, err := os.Getwd()
		if err != nil {
			log.Fatalf("failed to get current working directory: %v", err)
		}

		path := filepath.Join(dir, "moon.toml")

		if _, err := os.Stat(path); err != nil {
			log.Printf("config file not found at %s. using default values.", path)
			return cfg
		}

		file, err := os.ReadFile(path)
		if err != nil {
			log.Fatalf("failed to read moon.toml: %v", err)
		}

		if err := toml.Unmarshal(file, &cfg); err != nil {
			log.Fatalf("failed to parse moon.toml: %v", err)
		}

		return cfg
	}()
)

func Get() *Properties {
	return &config
}

func (p *Properties) ToString() string {
	return fmt.Sprintf(`
  Info:
    Command: %s
    Short Description: %s
    Long Description: %s
    License: %s
    Author: %s
    Version: %s
    Repository: %s

  Advanced:
    Debug: %v

  Genesis:
    Port: %d`,
		p.Info.Command,
		p.Info.ShortDescription,
		p.Info.LongDescription,
		p.Info.License,
		p.Info.Author,
		p.Info.Version,
		p.Info.Repository,
		p.Advanced.Debug,
		p.Genesis.Port)
}

func (p *Properties) Log() {
	log.Printf("Loaded configuration:\n%s", p.ToString())
}
