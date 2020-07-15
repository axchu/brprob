package main

import (
	"fmt"
	"strconv"
)

func main() {

}

func replace(in string, numChars int) string {
	found, numFound, indexStart := findFirstNumber(in)
	if !found {
		//dprint("not found. returning (%v)", in)
		return in
	}
	//dprint("(%v) found in (%v)", numFound, in)
	paddedNum := pad(numFound, numChars)
	//dprint("padded num(%v)", paddedNum)
	pre := in[0:indexStart] + paddedNum
	//dprint("pre(%v)", pre)
	indexPost := indexStart + len(numFound)
	//dprint("indexPost(%d)", indexPost)
	//dprint("calling replace on (%v)", in[indexPost:len(in)])
	post := replace(in[indexPost:len(in)], numChars)
	//dprint("post came back with (%v)", post)
	//dprint("returning(%v)", pre+post)
	return (pre + post)
}

func findFirstNumber(in string) (found bool, out string, indexStart int) {
	out = ""
	indexStart = 0
	found = false
	for i, cLetter := range in {
		letter := string(cLetter)
		if isNumber(letter) {
			out += letter
			if found == false {
				indexStart = i
			}
			found = true
			if i+1 >= len(in) || !isNumber(string(in[i+1])) {
				break
			}
		}
	}
	return found, out, indexStart
}

// assumes numChars is valid
func pad(valueToPad string, numChars int) string {
	padString := ""
	if len(valueToPad) < numChars {
		padString = makePadString(numChars - len(valueToPad))
	}
	return (padString + valueToPad)
}

// assumes numChars is valid
func makePadString(numChars int) (out string) {
	count := 0
	out = ""
	for count < numChars {
		out += "0"
		count++
	}
	return out
}

func isNumber(in string) bool {
	if _, err := strconv.Atoi(in); err == nil {
		return true
	}
	return false
}

func dprint(in string, args ...interface{}) {
	if in == "\n" {
		fmt.Println()
	} else {
		fmt.Printf("[debug] "+in+"\n", args...)
	}
}
