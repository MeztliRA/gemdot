package note

import (
	"fmt"
	"log"

	c "github.com/MeztliRA/gemdot/constants"
	homedir "github.com/mitchellh/go-homedir"
)

var (
	home      = getHomedir()
	Directory = fmt.Sprintf("%s/%s/", home, c.DataDir)
	File      = fmt.Sprintf("%s%s", Directory, c.FileName)
)

func getHomedir() string {
	home, err := homedir.Dir()
	if err != nil {
		log.Fatal(err)
	}
	return home
}
