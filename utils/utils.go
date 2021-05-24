package utils

import (
	"fmt"
	"strings"

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

func TrimString(str string) string {
	str = strings.Trim(str, "\n")
	str = strings.Trim(str, " ")
	str = strings.Trim(str, "\t")

	return str
}
