package conf

import (
	"io"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Settings struct {
	OrderPort   int `yaml:"orderPort"`
	AccountPort int `yaml:"accountPort"`
	StockPort   int `yaml:"stockPort"`
}

func New() *Settings {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	file, err := os.Open(dir + "/conf.yml")
	if err != nil {
		log.Fatal(err)
	}
	b, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	var (
		settings Settings
	)
	if err = yaml.Unmarshal(b, &settings); err != nil {
		log.Fatal(err)
	}
	return &settings
}
