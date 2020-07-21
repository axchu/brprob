package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter <\"stringToPad\", num0s>")
	fmt.Println("Example: <\"area 59\", 4>")

	for {

		fmt.Print("> ")
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)
		if text == "exit" {
			fmt.Println("goodbye!")
			break
		}
		input := strings.Split(text, "\",")
		if len(input) != 2 {
			fmt.Println("check formatting and try again")
		} else {
			inputString := strings.Replace(input[0], "\"", "", -1)
			numChars, err := strconv.Atoi(strings.TrimSpace(input[1]))
			if err != nil || numChars < 0 {
				fmt.Println(err.Error())
				fmt.Println("check number and try again")
			} else {
				out := replace([]rune(inputString), numChars)
				fmt.Printf("ans: %v\n\n", string(out))
			}
		}
	}

}

type IndexPointer struct {
	index int
	found bool
}

func newIndexPointer(index int, found bool) (indexPointer *IndexPointer) {
	indexPointer = new(IndexPointer)
	indexPointer.index = index
	indexPointer.found = found
	return indexPointer
}

func (ip *IndexPointer) reset() {
	ip.index = 0
	ip.found = false
}

type PadPointer struct {
	index int
	pad   int
}

func getInfo(input []rune, numChars int) (size int, padPoints []*PadPointer) {
	iStart := newIndexPointer(0, false)
	iEnd := newIndexPointer(0, false)
	padPoints = make([]*PadPointer, 0, len(input))
	padPointsIndex := 0
	numAdditionalSlots := 0

	for i, in := range input {

		if isNumber(in) && !iStart.found { // new num seq
			iStart.index = i
			iStart.found = true
			iEnd.index = i
		}

		if iStart.found {
			iEnd.index = i
			if i+1 >= len(input) || !isNumber(input[i+1]) {
				iEnd.found = true
			}
		}

		if iEnd.found {
			diff := iEnd.index - iStart.index + 1
			if diff < numChars {
				pad := numChars - diff

				// record where to pad and how many
				pp := new(PadPointer)
				pp.index = iStart.index + numAdditionalSlots
				pp.pad = pad
				padPoints = append(padPoints, pp)
				padPointsIndex++
				numAdditionalSlots += pad
			}
			iStart.reset()
			iEnd.reset()
		}

	}
	return len(input) + numAdditionalSlots, padPoints
}

func replace(input []rune, numChars int) []rune {

	// 1st pass, figure out how big the new slice needs to be
	newSize, padPoints := getInfo(input, numChars)
	if newSize == len(input) {
		return input
	}

	// make the new output slice
	output := make([]rune, newSize)

	// 2nd pass, insert the pads
	inIndex := 0
	padPointsIndex := 0
	padCounter := 0

	for i := range output {
		if padPointsIndex < len(padPoints) { // check if there are still points to pad
			if padCounter > 0 { // already padding
				output[i] = '0'
				padCounter--
				if padCounter == 0 {
					padPointsIndex++
				}
			} else { // not already padding
				if padPoints[padPointsIndex].index == i { // found an index to pad
					padCounter = padPoints[padPointsIndex].pad
					output[i] = '0'
					padCounter--
				} else { // did not find an index to pad
					output[i] = input[inIndex] // do a straight copy
					inIndex++
				}
			}
		} else { // no more points to pad, just keep copying
			output[i] = input[inIndex]
			inIndex++
		}
	}
	return output
}

func isNumber(in rune) bool {
	if _, err := strconv.Atoi(string(in)); err == nil {
		return true
	}
	return false
}

// func dprint(in string, args ...interface{}) {
// 	if in == "\n" {
// 		fmt.Println()
// 	} else {
// 		fmt.Printf("[debug] "+in+"\n", args...)
// 	}
// }
