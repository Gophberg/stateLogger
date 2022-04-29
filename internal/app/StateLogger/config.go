package StateLogger

import (
	yaml "gopkg.in/yaml.v2"
	"image"
	"io/ioutil"
	"log"
	"os"
)

// Config ...
type Config struct {
	Obj Objects

	//Token string `yaml:"token"`
	//URL   string `yaml:"url"`
}

type Objects struct {
	plots       PlotColor
	coordinates Coordinates
}

type PlotColor struct {
	onlineColor  *image.RGBA
	offlineColor *image.RGBA
}

type Coordinates struct {
	X, Y int
}

// NewConfig ...
func NewConfig() *[]Config {
	var config []Config
	f, err := os.Open("configs/config.yaml")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	cfg, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}
	err = yaml.Unmarshal(cfg, &config)
	if err != nil {
		log.Fatal(err)
	}
	return &config
}
