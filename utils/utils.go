package utils

import (
	"fmt"
	"log"
	"os/user"

	"github.com/MeztliRA/gemdot/color"
	c "github.com/MeztliRA/gemdot/constants"
	homedir "github.com/mitchellh/go-homedir"
)

func PrintVersion() {
	fmt.Printf("gemdot %s\n", c.Version)
}

func PrintHelp() {
	PrintVersion()
	color.Magenta(c.HelpMessageHeader)
	fmt.Println(c.HelpMessage)
}

func PrintGreeting() {
	user, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	username := user.Username
	color.Greenf("hello, %s!\n", username)
}

func GetHomedir() string {
	home, err := homedir.Dir()
	if err != nil {
		log.Fatal(err)
	}
	return home
}

func RemoveIndex(notes []string, index int) []string {
	return append(notes[:index], notes[index+1:]...)
}
