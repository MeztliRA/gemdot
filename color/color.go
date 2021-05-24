package color

import "github.com/fatih/color"

var (
	Set         = color.Set
	Unset       = color.Unset
	FgHiMagenta = color.FgHiMagenta
	Magenta     = color.New(color.FgHiMagenta).PrintFunc()
	Magentaln   = color.New(color.FgHiMagenta).PrintlnFunc()
	Green       = color.New(color.FgGreen).PrintFunc()
	Greenf      = color.New(color.FgGreen).PrintfFunc()
	HiGreen     = color.New(color.FgHiGreen).PrintlnFunc()
	Red         = color.New(color.FgHiRed).PrintlnFunc()
	Blue        = color.New(color.FgBlue).PrintfFunc()
)
