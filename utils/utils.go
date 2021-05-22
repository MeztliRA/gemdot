package utils

import (
	"fmt"

	"github.com/MeztliRA/gemdot/color"
	c "github.com/MeztliRA/gemdot/constants"
)

func PrintVersion() {
	fmt.Printf("gemdot %s\n", c.Version)
}

func PrintHelp() {
	PrintVersion()
	color.Magenta(c.HelpMessageHeader)
	fmt.Println(c.HelpMessage)
}
