// Copyright 2026 tsuru authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tablecli

import (
	"fmt"
	stdColor "image/color"
	"strings"

	"github.com/fatih/color"
)

func SetBorderColorByString(tableColor string) {
	var tblColor *color.Color

	if fgColor := colorMap[tableColor]; fgColor != 0 {
		tblColor = color.New(fgColor)
	} else if strings.HasPrefix(tableColor, "#") {
		c, err := parseHexColor(tableColor)

		if err == nil {
			tblColor = color.RGB(int(c.R), int(c.G), int(c.B))
		}
	}

	if tblColor != nil {
		TableConfig.BorderColorFunc = func(s string) string {
			return tblColor.Sprint(s)
		}
	}
}

func parseHexColor(s string) (c stdColor.RGBA, err error) {
	c.A = 0xff
	switch len(s) {
	case 7:
		_, err = fmt.Sscanf(s, "#%02x%02x%02x", &c.R, &c.G, &c.B)
	case 4:
		_, err = fmt.Sscanf(s, "#%1x%1x%1x", &c.R, &c.G, &c.B)
		// Double the hex digits:
		c.R *= 17
		c.G *= 17
		c.B *= 17
	default:
		err = fmt.Errorf("invalid length, must be 7 or 4")

	}
	return
}

var colorMap = map[string]color.Attribute{
	"black":      color.FgBlack,
	"red":        color.FgRed,
	"green":      color.FgGreen,
	"yellow":     color.FgYellow,
	"blue":       color.FgBlue,
	"magenta":    color.FgMagenta,
	"cyan":       color.FgCyan,
	"white":      color.FgWhite,
	"hi-black":   color.FgHiBlack,
	"hi-red":     color.FgHiRed,
	"hi-green":   color.FgHiGreen,
	"hi-yellow":  color.FgHiYellow,
	"hi-blue":    color.FgHiBlue,
	"hi-magenta": color.FgHiMagenta,
	"hi-cyan":    color.FgHiCyan,
	"hi-white":   color.FgHiWhite,
}
