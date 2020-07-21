package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
		5: Input{"abx4x5x6", 4},
	}
	expected := []string{
		0: "area 0059",
		1: "00123 Foo st",
		2: "123 Foo st",
		3: "Area0059asdf0234",
		4: "ABCDEFGH",
		5: "abx0004x0005x0006",
	}
	for i, in := range inputs {
		out := replace([]rune(in.in), in.numChars)
		assert.Equal(t, expected[i], string(out))

	}
}

// func tprint(in string, args ...interface{}) {
// 	if in == "\n" {
// 		fmt.Println()
// 	} else {
// 		fmt.Printf("[test] "+in+"\n", args...)
// 	}
// }
