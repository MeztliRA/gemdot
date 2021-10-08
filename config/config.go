package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	c "github.com/MeztliRA/gemdot/constants"
	u "github.com/MeztliRA/gemdot/utils"
	"github.com/olebedev/config"
)

type Config struct {
	Greeting bool `json:"greeting"`
	NoColor  bool `json:"no-color"`
}

// default config
var Default = Config{
	Greeting: true,
	NoColor:  false,
}

var (
	home      = u.GetHomedir()
	Directory = fmt.Sprintf("%s/%s/", home, c.DataDir)
	File      = fmt.Sprintf("%s%s", Directory, c.ConfigFileName)
)

// read config file
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

// check for config file existance
func Check() {
	if _, err := os.Stat(File); os.IsNotExist(err) {
		if _, err := os.Stat(Directory); os.IsNotExist(err) {
			dirErr := os.Mkdir(Directory, 0755)
			if dirErr != nil {
				log.Fatal(dirErr)
			}
		}

		jsonData, err := json.MarshalIndent(Default, "", "	")
		if err != nil {
			log.Fatal(err)
		}

		writeErr := os.WriteFile(File, jsonData, 0644)
		if writeErr != nil {
			log.Fatal(writeErr)
		}
	}
}
