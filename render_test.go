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
	expected := `Word      Number
One       1
Two       2
Three     3
`
	assert.Equal(t, expected, table.String())
}

func TestStringTabWriterMultiline(t *testing.T) {
	TableConfig.UseTabWriter = true
	defer func() {
		TableConfig.UseTabWriter = false
	}()
	table := NewTable()
	table.AddRow(Row{"One", "1", ""})
	table.AddRow(Row{"Two", "xxx|yyy", "aa|bb|cc|dd"})
	table.AddRow(Row{"Three", "3", ""})
	expected := `One       1         
Two       xxx|yyy   aa|bb|cc|dd
Three     3         
`
	assert.Equal(t, expected, table.String())
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
