package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestReplace(t *testing.T) {
	type Input struct {
		in       string
		numChars int
	}
	inputs := []Input{
		0: Input{"area 59", 4},
		1: Input{"123 Foo st", 5},
		2: Input{"123 Foo st", 2},
		3: Input{"Area59asdf234", 4},
		4: Input{"ABCDEFGH", 4},
	}
	expected := []string{
		0: "area 0059",
		1: "00123 Foo st",
		2: "123 Foo st",
		3: "Area0059asdf0234",
		4: "ABCDEFGH",
	}
	for i, in := range inputs {
		out := replace(in.in, in.numChars)
		require.Equal(t, expected[i], out)
	}
}

func TestFindFirstNumber(t *testing.T) {
	inputs := []string{
		0: "area 59",
		1: "123 Foo st",
		2: "Area59asdf234",
		3: "ABCDEFGH",
	}
	type Expected struct {
		found      bool
		out        string
		indexStart int
	}
	expected := []Expected{
		0: Expected{true, "59", 5},
		1: Expected{true, "123", 0},
		2: Expected{true, "59", 4},
		3: Expected{false, "", 0},
	}
	for i, in := range inputs {
		found, out, indexStart := findFirstNumber(in)
		//tprint("  found(%v), out(%v), indexStart(%d), expectedindexStart(%d)", found, out, indexStart, expected[i].indexStart)
		assert.Equal(t, expected[i].found, found)
		assert.Equal(t, expected[i].out, out)
		assert.Equal(t, expected[i].indexStart, indexStart)
	}
}

func TestPad(t *testing.T) {
	type Input struct {
		in       string
		numChars int
	}
	inputs := []Input{
		0: Input{in: "123", numChars: 4},
		1: Input{in: "123", numChars: 2},
		2: Input{in: "5", numChars: 0},
		3: Input{in: "", numChars: 3},
		4: Input{in: "ab", numChars: 4},
		5: Input{in: "25", numChars: -3},
	}
	expected := []string{
		0: "0123",
		1: "123",
		2: "5",
		3: "000",
		4: "00ab",
		5: "25",
	}
	for i, input := range inputs {
		actual := pad(input.in, input.numChars)
		assert.Equal(t, expected[i], actual)
	}
}

func TestMakePadString(t *testing.T) {
	inputs := []int{
		0: 5,
		1: 20,
		2: 0,
		3: -2,
	}
	expected := []string{
		0: "00000",
		1: "00000000000000000000",
		2: "",
		3: "",
	}
	for i, in := range inputs {
		actual := makePadString(in)
		assert.Equal(t, expected[i], actual)
	}
}

func TestIsNumber(t *testing.T) {
	inputs := []string{
		0: "123",
		1: "abc",
		2: "1a2b",
		3: "0",
		4: " ",
	}
	expected := []bool{
		0: true,
		1: false,
		2: false,
		3: true,
		4: false,
	}

	for i, in := range inputs {
		actual := isNumber(in)
		assert.Equal(t, expected[i], actual, "input(%d) %s failed", i, in)
	}

}

func tprint(in string, args ...interface{}) {
	if in == "\n" {
		fmt.Println()
	} else {
		fmt.Printf("[test] "+in+"\n", args...)
	}
}
