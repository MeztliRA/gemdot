package config

import (
	"fmt"
	"log"
	"os"

	c "github.com/MeztliRA/gemdot/constants"
	u "github.com/MeztliRA/gemdot/utils"
	"github.com/olebedev/config"
)

type Config struct {
	Greeting bool `json:"greeting"`
}

var Default = Config{
	Greeting: true,
}

var (
	home      = u.GetHomedir()
	Directory = fmt.Sprintf("%s/%s/", home, c.DataDir)
	File      = fmt.Sprintf("%s%s", Directory, c.ConfigFileName)
)

func Read() *config.Config {
	file, err := os.ReadFile(File)
	if err != nil {
		log.Fatal(err)
	}
	JSONString := string(file)

	config, err := config.ParseJson(JSONString)
	if err != nil {
		log.Fatal(err)
	}

	return config
}
