package help

import (
	"fmt"

	"github.com/MeztliRA/gemdot/color"
	c "github.com/MeztliRA/gemdot/constants"
)

func Print() {
	fmt.Printf("gemdot %s\n", c.Version)
	color.Magenta(c.HelpMessageHeader)
	fmt.Println(c.HelpMessage)
}
