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
				out := replace(inputString, numChars)
				fmt.Printf("ans: %s\n\n", out)
			}
		}
	}

}

func replace(in string, numChars int) string {
	found, numFound, indexStart := findFirstNumber(in)
	if !found {
		return in
	}
	paddedNum := pad(numFound, numChars)
	pre := in[0:indexStart] + paddedNum
	indexPost := indexStart + len(numFound)
	post := replace(in[indexPost:len(in)], numChars)
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

func pad(valueToPad string, numChars int) string {
	padString := ""
	if len(valueToPad) < numChars {
		padString = makePadString(numChars - len(valueToPad))
	}
	return (padString + valueToPad)
}

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
