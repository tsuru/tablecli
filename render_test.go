// Copyright 2018 tsuru authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tablecli

import (
	"fmt"
	"sort"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddOneRow(t *testing.T) {
	table := NewTable()
	table.AddRow(Row{"Three", "foo"})
	assert.Equal(t, "+-------+-----+\n| Three | foo |\n+-------+-----+\n", table.String())
}

func TestAddRows(t *testing.T) {
	table := NewTable()
	table.AddRow(Row{"One", "1"})
	table.AddRow(Row{"Two", "2"})
	table.AddRow(Row{"Three", "3"})
	expected := `+-------+---+
| One   | 1 |
| Two   | 2 |
| Three | 3 |
+-------+---+
`
	assert.Equal(t, expected, table.String())
}

func TestRows(t *testing.T) {
	table := NewTable()
	assert.Equal(t, 0, table.Rows())
	table.AddRow(Row{"One", "1"})
	assert.Equal(t, 1, table.Rows())
	table.AddRow(Row{"One", "1"})
	assert.Equal(t, 2, table.Rows())
	table.AddRow(Row{"One", "1"})
	table.AddRow(Row{"One", "1"})
	table.AddRow(Row{"One", "1"})
	assert.Equal(t, 5, table.Rows())
}

func TestSort(t *testing.T) {
	table := NewTable()
	table.AddRow(Row{"Three", "3"})
	table.AddRow(Row{"Zero", "0"})
	table.AddRow(Row{"One", "1"})
	table.AddRow(Row{"Two", "2"})
	expected := `+-------+---+
| One   | 1 |
| Three | 3 |
| Two   | 2 |
| Zero  | 0 |
+-------+---+
`
	table.Sort()
	assert.Equal(t, expected, table.String())
}

func TestColumnsSize(t *testing.T) {
	table := NewTable()
	table.AddRow(Row{"One", "1"})
	table.AddRow(Row{"Two", "2"})
	table.AddRow(Row{"Three", "3"})
	assert.Equal(t, []int{5, 1}, table.columnsSize())
}

func TestSeparator(t *testing.T) {
	table := NewTable()
	expected := "+-------+---+\n"
	buf := &strings.Builder{}
	table.separator(buf, []int{5, 1})
	assert.Equal(t, expected, buf.String())
}

func TestHeadings(t *testing.T) {
	table := NewTable()
	table.Headers = Row{"Word", "Number"}
	table.AddRow(Row{"One", "1"})
	table.AddRow(Row{"Two", "2"})
	table.AddRow(Row{"Three", "3"})
	expected := `+-------+--------+
| Word  | Number |
+-------+--------+
| One   | 1      |
| Two   | 2      |
| Three | 3      |
+-------+--------+
`
	assert.Equal(t, expected, table.String())
}

func TestString(t *testing.T) {
	table := NewTable()
	table.AddRow(Row{"One", "1"})
	table.AddRow(Row{"Two", "2"})
	table.AddRow(Row{"Three", "3"})
	expected := `+-------+---+
| One   | 1 |
| Two   | 2 |
| Three | 3 |
+-------+---+
`
	assert.Equal(t, expected, table.String())
}

func TestStringWithSeparator(t *testing.T) {
	table := NewTable()
	table.LineSeparator = true
	table.AddRow(Row{"One", "1"})
	table.AddRow(Row{"Two", "2"})
	table.AddRow(Row{"Three", "3"})
	expected := `+-------+---+
| One   | 1 |
+-------+---+
| Two   | 2 |
+-------+---+
| Three | 3 |
+-------+---+
`
	assert.Equal(t, expected, table.String())
}

func TestStringWithNewLineMultipleColumns(t *testing.T) {
	table := NewTable()
	table.AddRow(Row{"One", "1", ""})
	table.AddRow(Row{"Two", "xxx\nyyy", "aa\nbb\ncc\ndd"})
	table.AddRow(Row{"Three", "3", ""})
	expected := `+-------+-----+----+
| One   | 1   |    |
| Two   | xxx | aa |
|       | yyy | bb |
|       |     | cc |
|       |     | dd |
| Three | 3   |    |
+-------+-----+----+
`
	assert.Equal(t, expected, table.String())
}

func TestStringWithNewLine(t *testing.T) {
	table := NewTable()
	table.AddRow(Row{"One", "xxx\nyyy"})
	table.AddRow(Row{"Two", "2"})
	table.AddRow(Row{"Three", "3"})
	expected := `+-------+-----+
| One   | xxx |
|       | yyy |
| Two   | 2   |
| Three | 3   |
+-------+-----+
`
	assert.Equal(t, expected, table.String())
}

func TestStringWithNewLineWithSeparator(t *testing.T) {
	table := NewTable()
	table.LineSeparator = true
	table.AddRow(Row{"One", "xxx\nyyy\nzzzz"})
	table.AddRow(Row{"Two", "2"})
	table.AddRow(Row{"Three", "3"})
	expected := `+-------+------+
| One   | xxx  |
|       | yyy  |
|       | zzzz |
+-------+------+
| Two   | 2    |
+-------+------+
| Three | 3    |
+-------+------+
`
	assert.Equal(t, expected, table.String())
}

func TestRenderNoRows(t *testing.T) {
	table := NewTable()
	table.Headers = Row{"Word", "Number"}
	expected := `+------+--------+
| Word | Number |
+------+--------+
+------+--------+
`
	assert.Equal(t, expected, table.String())
}

func TestRenderEmpty(t *testing.T) {
	table := NewTable()
	assert.Equal(t, "", table.String())
}

func TestBytes(t *testing.T) {
	table := NewTable()
	table.AddRow(Row{"One", "1"})
	table.AddRow(Row{"Two", "2"})
	table.AddRow(Row{"Three", "3"})
	assert.Equal(t, []byte(table.String()), table.Bytes())
}

func TestRowListAdd(t *testing.T) {
	l := rowSlice([]Row{{"one", "1"}})
	l.add(Row{"two", "2"})
	assert.Len(t, l, 2)
}

func TestRowListLen(t *testing.T) {
	l := rowSlice([]Row{{"one", "1"}})
	assert.Equal(t, 1, l.Len())
	l.add(Row{"two", "2"})
	assert.Equal(t, 2, l.Len())
}

func TestRowListLess(t *testing.T) {
	l := rowSlice([]Row{{"zero", "0"}, {"one", "1"}, {"two", "2"}})
	assert.Equal(t, false, l.Less(0, 1))
	assert.Equal(t, false, l.Less(0, 2))
	assert.Equal(t, true, l.Less(1, 2))
	assert.Equal(t, true, l.Less(1, 0))
}

func TestRowListLessDifferentCase(t *testing.T) {
	l := rowSlice([]Row{{"Zero", "0"}, {"one", "1"}, {"two", "2"}})
	assert.Equal(t, false, l.Less(0, 1))
	assert.Equal(t, false, l.Less(0, 2))
	assert.Equal(t, true, l.Less(1, 2))
	assert.Equal(t, true, l.Less(1, 0))
}

func TestRowListSwap(t *testing.T) {
	l := rowSlice([]Row{{"zero", "0"}, {"one", "1"}, {"two", "2"}})
	l.Swap(0, 2)
	assert.Equal(t, true, l.Less(0, 2))
}

func TestRowListIsSortable(t *testing.T) {
	var _ sort.Interface = rowSlice{}
}

func TestResizeLargestColumn(t *testing.T) {
	tb := NewTable()
	tb.AddRow(Row{"1", "abcdefghijk"})
	tb.AddRow(Row{"2", "1234567890"})
	sizes := tb.resizeLargestColumn(11)
	assert.Equal(t, []int{1, 3}, sizes)
	assert.Equal(t, Row{"1", `ab↵
cd↵
ef↵
gh↵
ij↵
k`}, tb.rows[0])
	assert.Equal(t, Row{"2", `12↵
34↵
56↵
78↵
90`}, tb.rows[1])
}

func TestResizeLargestColumnOnMiddle(t *testing.T) {
	tb := NewTable()
	tb.AddRow(Row{"1", "abcdefghijk", "x"})
	tb.AddRow(Row{"2", "1234567890", "y"})
	sizes := tb.resizeLargestColumn(15)
	assert.Equal(t, []int{1, 3, 1}, sizes)
	assert.Equal(t, Row{"1", `ab↵
cd↵
ef↵
gh↵
ij↵
k`, "x"}, tb.rows[0])
	assert.Equal(t, Row{"2", `12↵
34↵
56↵
78↵
90`, "y"}, tb.rows[1])
}

func TestResizeLargestColumnNoTTYSize(t *testing.T) {
	tb := NewTable()
	tb.AddRow(Row{"1", "abcdefghijk"})
	tb.AddRow(Row{"2", "1234567890"})
	sizes := tb.resizeLargestColumn(0)
	assert.Equal(t, []int{1, 11}, sizes)
	assert.Equal(t, Row{"1", "abcdefghijk"}, tb.rows[0])
	assert.Equal(t, Row{"2", "1234567890"}, tb.rows[1])
}

func TestResizeLargestColumnNotEnoughSpace(t *testing.T) {
	tb := NewTable()
	tb.AddRow(Row{"1", "abcdefghijk"})
	tb.AddRow(Row{"2", "1234567890"})
	sizes := tb.resizeLargestColumn(9)
	assert.Equal(t, []int{1, 11}, sizes)
	assert.Equal(t, Row{"1", "abcdefghijk"}, tb.rows[0])
	assert.Equal(t, Row{"2", "1234567890"}, tb.rows[1])
}

func TestResizeLargestColumnWithLineBreaks(t *testing.T) {
	tb := NewTable()
	tb.AddRow(Row{"1", "abcde\nfgh\ni\njklm"})
	sizes := tb.resizeLargestColumn(12)
	assert.Equal(t, []int{1, 4}, sizes)
	assert.Equal(t, Row{"1", `abc↵
de
fgh
i
jkl↵
m`}, tb.rows[0])
}

func withColor(s string) string {
	return fmt.Sprintf("\033[0;31;10m%s\033[0m", s)
}

func TestResizeLargestColumnWithColors(t *testing.T) {
	tb := NewTable()
	color1 := withColor("abcdefghijk")
	color2 := withColor("1234567890")
	color3 := "123" + withColor("456789") + "012"
	tb.AddRow(Row{"1", color1})
	tb.AddRow(Row{"2", color2})
	tb.AddRow(Row{"3", color3})
	sizes := tb.resizeLargestColumn(11)
	assert.Equal(t, []int{1, 3}, sizes)
	redInit := "\033[0;31;10m"
	colorReset := "\033[0m"
	colorResetBreak := "\033[0m\n"
	assert.Equal(t, Row{"1", redInit + "ab↵" + colorResetBreak +
		redInit + "cd↵" + colorResetBreak +
		redInit + "ef↵" + colorResetBreak +
		redInit + "gh↵" + colorResetBreak +
		redInit + "ij↵" + colorResetBreak +
		redInit + "k" + colorReset}, tb.rows[0])
	assert.Equal(t, Row{"2", redInit + "12↵" + colorResetBreak +
		redInit + "34↵" + colorResetBreak +
		redInit + "56↵" + colorResetBreak +
		redInit + "78↵" + colorResetBreak +
		redInit + "90" + colorReset}, tb.rows[1])
	assert.Equal(t, Row{"3", "12↵\n" +
		"3" + redInit + "4↵" + colorResetBreak +
		redInit + "56↵" + colorResetBreak +
		redInit + "78↵" + colorResetBreak +
		redInit + "9" + colorReset + "0↵\n" +
		"12"}, tb.rows[2])
}

func TestResizeLargestColumnUnicode(t *testing.T) {
	tb := NewTable()
	tb.AddRow(Row{"1", "åß∂¬ƒ˚©“œ¡™"})
	tb.AddRow(Row{"2", "åß∂¬ƒ˚©“œ¡"})
	sizes := tb.resizeLargestColumn(11)
	assert.Equal(t, []int{1, 3}, sizes)
	assert.Equal(t, Row{"1", `åß↵
∂¬↵
ƒ˚↵
©“↵
œ¡↵
™`}, tb.rows[0])
	assert.Equal(t, Row{"2", `åß↵
∂¬↵
ƒ˚↵
©“↵
œ¡`}, tb.rows[1])
}

func TestColoredString(t *testing.T) {
	table := NewTable()
	two := withColor("str")
	two = two + " - " + two
	table.AddRow(Row{"Some large string", "1"})
	table.AddRow(Row{two, "2"})
	table.AddRow(Row{"Three", "3"})
	expected := `+-------------------+---+
| Some large string | 1 |
| ` + two + `         | 2 |
| Three             | 3 |
+-------------------+---+
`
	assert.Equal(t, expected, table.String())
}

// TestRuneLenWithFatihColorFormats tests that runeLen correctly handles
// all ANSI escape sequences produced by github.com/fatih/color library.
// This prevents regression when handling colored output.
func TestRuneLenWithFatihColorFormats(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		// Basic colors (single number): \x1b[31m (red), \x1b[32m (green), etc.
		{
			name:     "fatih/color basic red",
			input:    "\x1b[31mhello\x1b[0m",
			expected: 5,
		},
		{
			name:     "fatih/color basic green",
			input:    "\x1b[32mworld\x1b[0m",
			expected: 5,
		},
		// Bold attribute (single number): \x1b[1m
		{
			name:     "fatih/color bold",
			input:    "\x1b[1mbold text\x1b[0m",
			expected: 9,
		},
		// Hi-intensity colors: \x1b[90m to \x1b[97m
		{
			name:     "fatih/color hi-intensity",
			input:    "\x1b[91mhi-red\x1b[0m",
			expected: 6,
		},
		// Combined attributes (two numbers): \x1b[1;31m (bold red)
		{
			name:     "fatih/color bold+red",
			input:    "\x1b[1;31mbold red\x1b[0m",
			expected: 8,
		},
		// Background colors: \x1b[41m (red bg)
		{
			name:     "fatih/color background",
			input:    "\x1b[41mred bg\x1b[0m",
			expected: 6,
		},
		// Foreground + Background: \x1b[31;47m
		{
			name:     "fatih/color fg+bg",
			input:    "\x1b[31;47mred on white\x1b[0m",
			expected: 12,
		},
		// 256-color mode: \x1b[38;5;196m (foreground) or \x1b[48;5;196m (background)
		{
			name:     "fatih/color 256-color foreground",
			input:    "\x1b[38;5;196m256 color\x1b[0m",
			expected: 9,
		},
		{
			name:     "fatih/color 256-color background",
			input:    "\x1b[48;5;21mblue bg\x1b[0m",
			expected: 7,
		},
		// 24-bit RGB: \x1b[38;2;255;128;0m (foreground orange)
		{
			name:     "fatih/color RGB foreground",
			input:    "\x1b[38;2;255;128;0morange\x1b[0m",
			expected: 6,
		},
		{
			name:     "fatih/color RGB background",
			input:    "\x1b[48;2;0;0;255mblue bg\x1b[0m",
			expected: 7,
		},
		// Mixed RGB foreground + background
		{
			name:     "fatih/color RGB fg+bg",
			input:    "\x1b[38;2;255;255;255m\x1b[48;2;0;0;0mwhite on black\x1b[0m",
			expected: 14,
		},
		// Underline and other attributes: \x1b[4m
		{
			name:     "fatih/color underline",
			input:    "\x1b[4munderlined\x1b[0m",
			expected: 10,
		},
		// Nested/chained attributes
		{
			name:     "fatih/color nested styles",
			input:    "\x1b[1m\x1b[4m\x1b[31mbold underline red\x1b[0m",
			expected: 18,
		},
		// Specific reset codes: \x1b[22m (reset bold), \x1b[24m (reset underline)
		{
			name:     "fatih/color specific resets",
			input:    "\x1b[1mbold\x1b[22m normal\x1b[0m",
			expected: 11,
		},
		// Multiple resets in sequence
		{
			name:     "fatih/color multiple resets",
			input:    "\x1b[1m\x1b[4mtext\x1b[22m\x1b[24m\x1b[0m",
			expected: 4,
		},
		// Empty string with just escape codes
		{
			name:     "fatih/color only escapes",
			input:    "\x1b[31m\x1b[0m",
			expected: 0,
		},
		// Real-world example: colored status output
		{
			name:     "fatih/color real status",
			input:    "\x1b[32m✓\x1b[0m \x1b[1mSuccess\x1b[0m",
			expected: 9, // "✓ Success" = 1 + 1 + 7
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := runeLen(tt.input)
			assert.Equal(t, tt.expected, got, "runeLen(%q) = %d, want %d", tt.input, got, tt.expected)
		})
	}
}

// TestTableColumnWidthWithFatihColor ensures table column width calculation
// works correctly with various fatih/color escape sequences.
func TestTableColumnWidthWithFatihColor(t *testing.T) {
	tests := []struct {
		name          string
		rows          []Row
		expectedSizes []int
	}{
		{
			name: "basic colors",
			rows: []Row{
				{"\x1b[31mred\x1b[0m", "normal"},
				{"plain", "\x1b[32mgreen\x1b[0m"},
			},
			expectedSizes: []int{5, 6},
		},
		{
			name: "256 colors",
			rows: []Row{
				{"\x1b[38;5;196mcolor256\x1b[0m", "test"},
			},
			expectedSizes: []int{8, 4},
		},
		{
			name: "RGB colors",
			rows: []Row{
				{"\x1b[38;2;255;128;0mRGB orange\x1b[0m", "data"},
			},
			expectedSizes: []int{10, 4},
		},
		{
			name: "mixed styles",
			rows: []Row{
				{"\x1b[1;31mbold red\x1b[0m", "\x1b[4munderline\x1b[0m"},
				{"\x1b[38;2;0;255;0mRGB green\x1b[0m", "plain"},
			},
			expectedSizes: []int{9, 9},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			table := NewTable()
			for _, row := range tt.rows {
				table.AddRow(row)
			}
			sizes := table.columnsSize()
			assert.Equal(t, tt.expectedSizes, sizes)
		})
	}
}

func TestResizeLargestColumnOnWhitespace(t *testing.T) {
	tb := NewTable()
	tb.AddRow(Row{"1", "abc def ghi jk"})
	tb.AddRow(Row{"2", "12 3 456 7890"})
	tb.AddRow(Row{"3", "1 2 3 4"})
	sizes := tb.resizeLargestColumn(12)
	assert.Equal(t, []int{1, 4}, sizes)
	assert.Equal(t, Row{"1", `abc↵
def↵
ghi↵
jk`}, tb.rows[0])
	assert.Equal(t, Row{"2", `12 ↵
3  ↵
456↵
789↵
0`}, tb.rows[1])
	assert.Equal(t, Row{"3", `1 2↵
3 4`}, tb.rows[2])
}

func TestResizeLargestColumnOnAnyWithBreakAny(t *testing.T) {
	TableConfig.BreakOnAny = true
	defer func() { TableConfig.BreakOnAny = false }()
	tb := NewTable()
	tb.AddRow(Row{"1", "abc def ghi jk"})
	tb.AddRow(Row{"2", "12 3 456 7890"})
	tb.AddRow(Row{"3", "1 2 3 4"})
	sizes := tb.resizeLargestColumn(12)
	assert.Equal(t, []int{1, 4}, sizes)
	assert.Equal(t, Row{"1", `abc↵
 de↵
f g↵
hi ↵
jk`}, tb.rows[0])
	assert.Equal(t, Row{"2", `12 ↵
3 4↵
56 ↵
789↵
0`}, tb.rows[1])
	assert.Equal(t, Row{"3", `1 2↵
 3 ↵
4`}, tb.rows[2])
}

func TestResizeLargestColumnOnBreakableChars(t *testing.T) {
	tb := NewTable()
	tb.AddRow(Row{"1", "abc:def ghi jk"})
	tb.AddRow(Row{"2", "12:3 456 7890"})
	tb.AddRow(Row{"3", "1 2 3: 4"})
	sizes := tb.resizeLargestColumn(12)
	assert.Equal(t, []int{1, 4}, sizes)
	assert.Equal(t, Row{"1", `abc↵
:de↵
f  ↵
ghi↵
jk`}, tb.rows[0])
	assert.Equal(t, Row{"2", `12:↵
3  ↵
456↵
789↵
0`}, tb.rows[1])
	assert.Equal(t, Row{"3", `1 2↵
3: ↵
4`}, tb.rows[2])
}

func TestStringTabWriter(t *testing.T) {
	TableConfig.UseTabWriter = true
	defer func() {
		TableConfig.UseTabWriter = false
	}()
	table := NewTable()
	table.Headers = Row{"Word", "Number"}
	table.AddRow(Row{"One", "1"})
	table.AddRow(Row{"Two", "2"})
	table.AddRow(Row{"Three", "3"})
	expected := `WORD    NUMBER
One     1
Two     2
Three   3
`
	assert.Equal(t, expected, table.String())
}

func TestStringTabWriterTruncate(t *testing.T) {
	TableConfig.UseTabWriter = true
	defer func() {
		TableConfig.UseTabWriter = false
	}()
	table := NewTable()
	table.TableWriterTruncate = true
	table.AddRow(Row{"One", "1", ""})
	table.AddRow(Row{"Two", "xxx\nyyy", "aa|bb|cc|dd"})
	table.AddRow(Row{"Three", "3", ""})
	expected := "One     1         \nTwo     xxx ...   aa|bb|cc|dd\nThree   3         \n"
	assert.Equal(t, expected, table.String())
}

func TestStringTabWriterTruncateEnabled(t *testing.T) {
	TableConfig.UseTabWriter = true
	TableConfig.TabWriterTruncate = true
	defer func() {
		TableConfig.UseTabWriter = false
		TableConfig.TabWriterTruncate = true
	}()
	table := NewTable()
	table.TableWriterTruncate = true
	table.Headers = Row{"Name", "Description"}
	table.AddRow(Row{"Item1", "Single line"})
	table.AddRow(Row{"Item2", "First line\nSecond line\nThird line"})
	table.AddRow(Row{"Item3", "Another single"})
	expected := `NAME    DESCRIPTION
Item1   Single line
Item2   First line ...
Item3   Another single
`
	assert.Equal(t, expected, table.String())
}

func TestStringTabWriterTruncateDisabled(t *testing.T) {
	TableConfig.UseTabWriter = true
	TableConfig.TabWriterTruncate = false
	defer func() {
		TableConfig.UseTabWriter = false
		TableConfig.TabWriterTruncate = true
	}()
	table := NewTable()
	table.Headers = Row{"Name", "Description"}
	table.AddRow(Row{"Item1", "Single line"})
	table.AddRow(Row{"Item2", "First line\nSecond line"})
	table.AddRow(Row{"Item3", "Another single"})
	expected := `NAME    DESCRIPTION
Item1   Single line
Item2   First line Second line
Item3   Another single
`
	assert.Equal(t, expected, table.String())
}

func TestStringTabWriterDisableTableTruncate(t *testing.T) {
	TableConfig.UseTabWriter = true
	TableConfig.TabWriterTruncate = false
	defer func() {
		TableConfig.UseTabWriter = false
		TableConfig.TabWriterTruncate = false
	}()
	table := NewTable()
	table.TableWriterTruncate = false
	table.Headers = Row{"Name", "Description"}
	table.AddRow(Row{"Item1", "Single line"})
	table.AddRow(Row{"Item2", "First line\nSecond line"})
	expected := `NAME    DESCRIPTION
Item1   Single line
Item2   First line Second line
`
	assert.Equal(t, expected, table.String())
}

func TestStringTabWriterPadding(t *testing.T) {
	TableConfig.UseTabWriter = true
	defer func() {
		TableConfig.UseTabWriter = false
	}()
	table := NewTable()
	table.TableWriterPadding = 2
	table.Headers = Row{"Word", "Number"}
	table.AddRow(Row{"One", "1"})
	table.AddRow(Row{"Two", "2"})
	expected := `  WORD   NUMBER
  One    1
  Two    2
`
	assert.Equal(t, expected, table.String())
}

func TestStringTabWriterFormFeedAndCarriageReturn(t *testing.T) {
	TableConfig.UseTabWriter = true
	TableConfig.TabWriterTruncate = true
	defer func() {
		TableConfig.UseTabWriter = false
		TableConfig.TabWriterTruncate = true
	}()
	table := NewTable()
	table.TableWriterTruncate = true
	table.Headers = Row{"Name", "Value"}
	table.AddRow(Row{"FormFeed", "before\fafter"})
	table.AddRow(Row{"CarriageReturn", "before\rafter"})
	table.AddRow(Row{"Mixed", "a\fb\nc\rd"})
	expected := `NAME             VALUE
FormFeed         before ...
CarriageReturn   before ...
Mixed            a ...
`
	assert.Equal(t, expected, table.String())
}

func TestStringTabWriterFormFeedAndCarriageReturnNoTruncate(t *testing.T) {
	TableConfig.UseTabWriter = true
	TableConfig.TabWriterTruncate = false
	defer func() {
		TableConfig.UseTabWriter = false
		TableConfig.TabWriterTruncate = true
	}()
	table := NewTable()
	table.Headers = Row{"Name", "Value"}
	table.AddRow(Row{"FormFeed", "before\fafter"})
	table.AddRow(Row{"CarriageReturn", "before\rafter"})
	table.AddRow(Row{"Mixed", "a\fb\nc\rd"})
	expected := `NAME             VALUE
FormFeed         before after
CarriageReturn   before after
Mixed            a b c d
`
	assert.Equal(t, expected, table.String())
}

func TestStringTabWriterWithANSIColors(t *testing.T) {
	TableConfig.UseTabWriter = true
	defer func() {
		TableConfig.UseTabWriter = false
	}()
	table := NewTable()
	table.Headers = Row{"ID", "Status", "Name"}
	// ANSI color codes: \033[33;1;1m = yellow, \033[0m = reset
	yellow := func(s string) string { return "\033[33;1;1m" + s + "\033[0m" }
	red := func(s string) string { return "\033[31;1;1m" + s + "\033[0m" }
	table.AddRow(Row{"abc123", "ok", "app1"})
	table.AddRow(Row{yellow("def456"), yellow("running"), yellow("…")})
	table.AddRow(Row{red("ghi789"), red("error"), red("app3")})
	output := table.String()

	// Verify alignment is correct (columns should align despite ANSI codes)
	lines := strings.Split(output, "\n")
	assert.Len(t, lines, 5) // header + 3 rows + trailing newline

	// Check that ANSI codes are preserved in output
	assert.Contains(t, output, "\033[33;1;1m")
	assert.Contains(t, output, "\033[31;1;1m")
	assert.Contains(t, output, "\033[0m")

	// Verify the colored content is present
	assert.Contains(t, output, "def456")
	assert.Contains(t, output, "running")
	assert.Contains(t, output, "error")
	assert.Contains(t, output, "…")

	assert.Equal(t, 23, runeLen(lines[1])) // First row
	assert.Equal(t, 20, runeLen(lines[2])) // Second row
	assert.Equal(t, 23, runeLen(lines[3])) // Third row
	assert.Equal(t, 0, runeLen(lines[4]))  // Trailing newline
}

// Tests for expandRow function and TableWriterExpandRows flag

func TestExpandRowSimple(t *testing.T) {
	result := expandRow(Row{"A", "X\nY", "1"})
	expected := [][]string{
		{"A", "X", "1"},
		{"", "Y", ""},
	}
	assert.Equal(t, expected, result)
}

func TestExpandRowNoNewlines(t *testing.T) {
	result := expandRow(Row{"A", "B", "C"})
	expected := [][]string{
		{"A", "B", "C"},
	}
	assert.Equal(t, expected, result)
}

func TestExpandRowEmptyRow(t *testing.T) {
	result := expandRow(Row{})
	assert.Len(t, result, 0)
}

func TestExpandRowSingleCell(t *testing.T) {
	result := expandRow(Row{"A\nB\nC"})
	expected := [][]string{
		{"A"},
		{"B"},
		{"C"},
	}
	assert.Equal(t, expected, result)
}

func TestExpandRowEmptyCells(t *testing.T) {
	result := expandRow(Row{"", "X\nY", ""})
	expected := [][]string{
		{"", "X", ""},
		{"", "Y", ""},
	}
	assert.Equal(t, expected, result)
}

func TestExpandRowAllEmpty(t *testing.T) {
	result := expandRow(Row{"", "", ""})
	expected := [][]string{
		{"", "", ""},
	}
	assert.Equal(t, expected, result)
}

func TestExpandRowWhitespaceOnly(t *testing.T) {
	result := expandRow(Row{"  ", "X\nY", "   "})
	expected := [][]string{
		{"", "X", ""},
		{"", "Y", ""},
	}
	assert.Equal(t, expected, result)
}

func TestExpandRowMultipleNewlines(t *testing.T) {
	result := expandRow(Row{"A", "1\n2\n3\n4\n5", "X"})
	expected := [][]string{
		{"A", "1", "X"},
		{"", "2", ""},
		{"", "3", ""},
		{"", "4", ""},
		{"", "5", ""},
	}
	assert.Equal(t, expected, result)
}

func TestExpandRowDifferentLineCounts(t *testing.T) {
	result := expandRow(Row{"A\nB", "1\n2\n3", "X"})
	expected := [][]string{
		{"A", "1", "X"},
		{"B", "2", ""},
		{"", "3", ""},
	}
	assert.Equal(t, expected, result)
}

func TestExpandRowLeadingTrailingNewlines(t *testing.T) {
	// TrimSpace removes leading/trailing whitespace including newlines
	result := expandRow(Row{"\nA\n", "X", "1"})
	expected := [][]string{
		{"A", "X", "1"},
	}
	assert.Equal(t, expected, result)
}

func TestExpandRowUnicode(t *testing.T) {
	result := expandRow(Row{"日本語", "こんにちは\nさようなら", "🎉"})
	expected := [][]string{
		{"日本語", "こんにちは", "🎉"},
		{"", "さようなら", ""},
	}
	assert.Equal(t, expected, result)
}

func TestExpandRowWithColors(t *testing.T) {
	colored := withColor("red") + "\n" + withColor("blue")
	result := expandRow(Row{"A", colored, "1"})
	assert.Len(t, result, 2)
	assert.Equal(t, "A", result[0][0])
	assert.Contains(t, result[0][1], "red")
	assert.Equal(t, "1", result[0][2])
	assert.Equal(t, "", result[1][0])
	assert.Contains(t, result[1][1], "blue")
	assert.Equal(t, "", result[1][2])
}

func TestTableWriterExpandRowsBasic(t *testing.T) {
	TableConfig.UseTabWriter = true
	defer func() {
		TableConfig.UseTabWriter = false
	}()
	table := NewTable()
	table.TableWriterExpandRows = true
	table.Headers = Row{"Name", "Values", "Status"}
	table.AddRow(Row{"Item1", "A\nB\nC", "OK"})
	table.AddRow(Row{"Item2", "X", "Done"})
	output := table.String()
	lines := strings.Split(output, "\n")
	assert.Len(t, lines, 6) // header + 4 data lines + trailing newline
	assert.Contains(t, lines[0], "NAME")
	assert.Contains(t, lines[1], "Item1")
	assert.Contains(t, lines[1], "A")
	assert.Contains(t, lines[1], "OK")
	assert.Contains(t, lines[2], "B")
	assert.Contains(t, lines[3], "C")
	assert.Contains(t, lines[4], "Item2")
}

func TestTableWriterExpandRowsNoHeaders(t *testing.T) {
	TableConfig.UseTabWriter = true
	defer func() {
		TableConfig.UseTabWriter = false
	}()
	table := NewTable()
	table.TableWriterExpandRows = true
	table.AddRow(Row{"A", "1\n2", "X"})
	table.AddRow(Row{"B", "3", "Y"})
	output := table.String()
	lines := strings.Split(output, "\n")
	assert.Len(t, lines, 4) // 3 data lines + trailing newline
	assert.Contains(t, lines[0], "A")
	assert.Contains(t, lines[0], "1")
	assert.Contains(t, lines[0], "X")
	assert.Contains(t, lines[1], "2")
	assert.Contains(t, lines[2], "B")
	assert.Contains(t, lines[2], "3")
	assert.Contains(t, lines[2], "Y")
}

func TestTableWriterExpandRowsWithPadding(t *testing.T) {
	TableConfig.UseTabWriter = true
	defer func() {
		TableConfig.UseTabWriter = false
	}()
	table := NewTable()
	table.TableWriterExpandRows = true
	table.TableWriterPadding = 4
	table.Headers = Row{"Col1", "Col2"}
	table.AddRow(Row{"A", "X\nY"})
	expected := `    COL1   COL2
    A      X
           Y
`
	assert.Equal(t, expected, table.String())
}

func TestTableWriterExpandRowsDisabled(t *testing.T) {
	TableConfig.UseTabWriter = true
	TableConfig.TabWriterTruncate = false
	defer func() {
		TableConfig.UseTabWriter = false
		TableConfig.TabWriterTruncate = false
	}()
	table := NewTable()
	table.TableWriterExpandRows = false
	table.AddRow(Row{"A", "X\nY", "1"})
	// When disabled, newlines should be replaced with spaces (default behavior)
	output := table.String()
	assert.Contains(t, output, "X Y")
	assert.NotContains(t, output, "X\nY")
}

func TestTableWriterExpandRowsVsTruncate(t *testing.T) {
	// ExpandRows should take precedence - when enabled, truncate is ignored for newlines
	TableConfig.UseTabWriter = true
	defer func() {
		TableConfig.UseTabWriter = false
	}()
	table := NewTable()
	table.TableWriterExpandRows = true
	table.TableWriterTruncate = true // This should be ignored when ExpandRows is true
	table.AddRow(Row{"A", "X\nY\nZ", "1"})
	output := table.String()
	lines := strings.Split(output, "\n")
	assert.Len(t, lines, 4) // 3 rows + trailing newline
}

func TestTableWriterExpandRowsAllCellsMultiline(t *testing.T) {
	TableConfig.UseTabWriter = true
	defer func() {
		TableConfig.UseTabWriter = false
	}()
	table := NewTable()
	table.TableWriterExpandRows = true
	table.AddRow(Row{"A\nB", "1\n2", "X\nY"})
	expected := `A   1   X
B   2   Y
`
	assert.Equal(t, expected, table.String())
}

func TestTableWriterExpandRowsUnevenMultiline(t *testing.T) {
	TableConfig.UseTabWriter = true
	defer func() {
		TableConfig.UseTabWriter = false
	}()
	table := NewTable()
	table.TableWriterExpandRows = true
	table.AddRow(Row{"A", "1\n2\n3\n4", "X\nY"})
	output := table.String()
	lines := strings.Split(output, "\n")
	assert.Len(t, lines, 5) // 4 data lines + trailing newline
	assert.Contains(t, lines[0], "A")
	assert.Contains(t, lines[0], "1")
	assert.Contains(t, lines[0], "X")
	assert.Contains(t, lines[1], "2")
	assert.Contains(t, lines[1], "Y")
	assert.Contains(t, lines[2], "3")
	assert.Contains(t, lines[3], "4")
}

func TestTableWriterExpandRowsUnicode(t *testing.T) {
	TableConfig.UseTabWriter = true
	defer func() {
		TableConfig.UseTabWriter = false
	}()
	table := NewTable()
	table.TableWriterExpandRows = true
	table.Headers = Row{"名前", "値"}
	table.AddRow(Row{"テスト", "あ\nい\nう"})
	output := table.String()
	assert.Contains(t, output, "テスト")
	assert.Contains(t, output, "あ")
	assert.Contains(t, output, "い")
	assert.Contains(t, output, "う")
	lines := strings.Split(output, "\n")
	assert.Len(t, lines, 5) // header + 3 data lines + trailing newline
}

func TestTableWriterExpandRowsWithColors(t *testing.T) {
	TableConfig.UseTabWriter = true
	defer func() {
		TableConfig.UseTabWriter = false
	}()
	table := NewTable()
	table.TableWriterExpandRows = true
	colored := withColor("line1") + "\n" + withColor("line2")
	table.AddRow(Row{"A", colored, "X"})
	output := table.String()
	assert.Contains(t, output, "\033[0;31;10m")
	assert.Contains(t, output, "line1")
	assert.Contains(t, output, "line2")
	lines := strings.Split(output, "\n")
	assert.Len(t, lines, 3) // 2 rows + trailing newline
}

func TestTableWriterExpandRowsEmptyTable(t *testing.T) {
	TableConfig.UseTabWriter = true
	defer func() {
		TableConfig.UseTabWriter = false
	}()
	table := NewTable()
	table.TableWriterExpandRows = true
	table.Headers = Row{"A", "B"}
	expected := `A   B
`
	assert.Equal(t, expected, table.String())
}

func TestTableWriterExpandRowsMixedRows(t *testing.T) {
	TableConfig.UseTabWriter = true
	defer func() {
		TableConfig.UseTabWriter = false
	}()
	table := NewTable()
	table.TableWriterExpandRows = true
	table.Headers = Row{"ID", "Items", "Count"}
	table.AddRow(Row{"1", "apple", "1"})
	table.AddRow(Row{"2", "banana\norange\ngrape", "3"})
	table.AddRow(Row{"3", "melon", "1"})
	output := table.String()
	lines := strings.Split(output, "\n")
	assert.Len(t, lines, 7) // header + 5 data lines + trailing newline
	assert.Contains(t, lines[0], "ID")
	assert.Contains(t, lines[0], "ITEMS")
	assert.Contains(t, lines[1], "1")
	assert.Contains(t, lines[1], "apple")
	assert.Contains(t, lines[2], "2")
	assert.Contains(t, lines[2], "banana")
	assert.Contains(t, lines[2], "3")
	assert.Contains(t, lines[3], "orange")
	assert.Contains(t, lines[4], "grape")
	assert.Contains(t, lines[5], "3")
	assert.Contains(t, lines[5], "melon")
}

func TestTableWriterExpandRowsOnlyNewlines(t *testing.T) {
	TableConfig.UseTabWriter = true
	defer func() {
		TableConfig.UseTabWriter = false
	}()
	table := NewTable()
	table.TableWriterExpandRows = true
	// Cell with only newlines should result in empty strings after TrimSpace
	table.AddRow(Row{"A", "\n\n\n", "X"})
	output := table.String()
	// After TrimSpace, "\n\n\n" becomes "" which splits to [""]
	lines := strings.Split(output, "\n")
	assert.Len(t, lines, 2) // 1 row + trailing newline
}

func TestTableWriterExpandRowsCarriageReturn(t *testing.T) {
	TableConfig.UseTabWriter = true
	defer func() {
		TableConfig.UseTabWriter = false
	}()
	table := NewTable()
	table.TableWriterExpandRows = true
	// Only \n triggers expansion, \r should remain in the cell
	table.AddRow(Row{"A", "X\rY", "1"})
	output := table.String()
	// \r is not a split character in expandRow, but TrimSpace might affect it
	lines := strings.Split(output, "\n")
	assert.Len(t, lines, 2) // 1 row + trailing newline
}

func TestTableWriterExpandRowsFormFeed(t *testing.T) {
	TableConfig.UseTabWriter = true
	defer func() {
		TableConfig.UseTabWriter = false
	}()
	table := NewTable()
	table.TableWriterExpandRows = true
	// Only \n triggers expansion, \f should remain in the cell
	table.AddRow(Row{"A", "X\fY", "1"})
	output := table.String()
	lines := strings.Split(output, "\n")
	assert.Len(t, lines, 2) // 1 row + trailing newline
}

func TestTableWriterExpandRowsLongContent(t *testing.T) {
	TableConfig.UseTabWriter = true
	defer func() {
		TableConfig.UseTabWriter = false
	}()
	table := NewTable()
	table.TableWriterExpandRows = true
	longLine1 := strings.Repeat("a", 50)
	longLine2 := strings.Repeat("b", 50)
	table.AddRow(Row{"ID", longLine1 + "\n" + longLine2, "OK"})
	output := table.String()
	assert.Contains(t, output, longLine1)
	assert.Contains(t, output, longLine2)
	lines := strings.Split(output, "\n")
	assert.Len(t, lines, 3) // 2 rows + trailing newline
}

func BenchmarkString(b *testing.B) {
	b.StopTimer()
	table := NewTable()
	table.Headers = Row{"row 1", "row 2", "row 3", "row 4"}
	for i := 0; i < 100; i++ {
		table.AddRow(Row{"my big string", "other string", "small", `largest string in the whole table
continuing string
another line
yet another big line`})
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = table.String()
	}
}

func BenchmarkStringWithColor(b *testing.B) {
	b.StopTimer()
	table := NewTable()
	table.Headers = Row{"row 1", "row 2", "row 3", "row 4"}
	for i := 0; i < 100; i++ {
		table.AddRow(Row{"my " + withColor("big") + " string", withColor("other string"), "small", `largest string in the whole table
continuing string
another line
yet another big line`})
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = table.String()
	}
}

func BenchmarkStringWithResize(b *testing.B) {
	TableConfig.MaxTTYWidth = 52
	defer func() {
		TableConfig.MaxTTYWidth = 0
	}()
	b.StopTimer()
	table := NewTable()
	table.Headers = Row{"row 1", "row 2", "row 3", "row 4"}
	for i := 0; i < 100; i++ {
		table.AddRow(Row{"my big string", "other string", "small", `largest string in the whole table
continuing string
another line
yet another big line`})
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = table.String()
	}
}
