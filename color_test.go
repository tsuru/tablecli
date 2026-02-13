// Copyright 2026 tsuru authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tablecli

import (
	"image/color"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseHexColor(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    color.RGBA
		wantErr string
	}{
		{"full hex", "#FF8800", color.RGBA{R: 0xFF, G: 0x88, B: 0x00, A: 0xFF}, ""},
		{"short hex", "#F80", color.RGBA{R: 0xFF, G: 0x88, B: 0x00, A: 0xFF}, ""},
		{"black", "#000000", color.RGBA{R: 0, G: 0, B: 0, A: 0xFF}, ""},
		{"white", "#FFFFFF", color.RGBA{R: 0xFF, G: 0xFF, B: 0xFF, A: 0xFF}, ""},
		{"short white", "#FFF", color.RGBA{R: 0xFF, G: 0xFF, B: 0xFF, A: 0xFF}, ""},
		{"short black", "#000", color.RGBA{R: 0, G: 0, B: 0, A: 0xFF}, ""},
		{"lowercase", "#ff8800", color.RGBA{R: 0xFF, G: 0x88, B: 0x00, A: 0xFF}, ""},
		{"mixed case", "#Ff8800", color.RGBA{R: 0xFF, G: 0x88, B: 0x00, A: 0xFF}, ""},
		{"invalid length", "#FF88", color.RGBA{}, "invalid length, must be 7 or 4"},
		{"too long", "#FF880000", color.RGBA{}, "invalid length, must be 7 or 4"},
		{"too short", "#FF", color.RGBA{}, "invalid length, must be 7 or 4"},
		{"empty", "", color.RGBA{}, "invalid length"},
		{"no hash", "FF8800", color.RGBA{}, "invalid length"},
		{"invalid chars", "#ZZZZZZ", color.RGBA{}, "expected integer"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseHexColor(tt.input)
			if tt.wantErr != "" {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.wantErr)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestSetBorderColorByString(t *testing.T) {
	t.Run("named colors", func(t *testing.T) {
		namedColors := []string{
			"black", "red", "green", "yellow",
			"blue", "magenta", "cyan", "white",
			"hi-black", "hi-red", "hi-green", "hi-yellow",
			"hi-blue", "hi-magenta", "hi-cyan", "hi-white",
		}
		for _, name := range namedColors {
			t.Run(name, func(t *testing.T) {
				TableConfig.BorderColorFunc = nil
				SetBorderColorByString(name)
				assert.NotNil(t, TableConfig.BorderColorFunc, "BorderColorFunc should be set for color %q", name)
			})
		}
		TableConfig.BorderColorFunc = nil
	})

	tests := []struct {
		name    string
		input   string
		wantNil bool
	}{
		{"hex color", "#FF0000", false},
		{"short hex", "#F00", false},
		{"invalid name", "invalid-color", true},
		{"empty string", "", true},
		{"invalid hex", "#ZZZZZZ", true},
		{"bad hex length", "#FF", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			TableConfig.BorderColorFunc = nil
			SetBorderColorByString(tt.input)
			if tt.wantNil {
				assert.Nil(t, TableConfig.BorderColorFunc)
				return
			} else {
				assert.NotNil(t, TableConfig.BorderColorFunc)
			}

			result := TableConfig.BorderColorFunc("hello")
			assert.NotEmpty(t, result)
			assert.Contains(t, result, "hello")
			TableConfig.BorderColorFunc = nil
		})
	}
}
