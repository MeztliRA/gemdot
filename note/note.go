package note

import (
	"fmt"

	c "github.com/MeztliRA/gemdot/constants"
	u "github.com/MeztliRA/gemdot/utils"
)

var (
	home      = u.GetHomedir()
	Directory = fmt.Sprintf("%s/%s/", home, c.DataDir)
	File      = fmt.Sprintf("%s%s", Directory, c.FileName)
)
