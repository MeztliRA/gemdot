package color

import "github.com/fatih/color"

var (
	Magenta   = color.New(color.FgHiMagenta).PrintFunc()
	Magentaln = color.New(color.FgHiMagenta).PrintlnFunc()
	Green     = color.New(color.FgGreen).PrintFunc()
	HiGreen   = color.New(color.FgHiGreen).PrintlnFunc()
	Red       = color.New(color.FgHiRed).PrintlnFunc()
	Blue      = color.New(color.FgBlue).PrintfFunc()
)
