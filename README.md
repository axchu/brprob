# brprob

#### The problem
Given a string, zero pad any whole number to X char.

Examples:
```
"area 59", 4
return "area 0059"
"123 Foo st", 5
return "00123 Foo st"
"123 Foo st", 2
return "123 Foo st"
"Area59asdf234", 4
return "Area0059asdf0234"
```

#### How to run on your local box
```
> go get github.com/axchu/brprob
> go run main.go
```
At the prompt, follow the instructions. Make sure you put double quotes around the string. Type *exit* to exit.

#### Time Complexity
The first pass goes once through the input, the 2nd pass goes through the output once, building that output. There are only 2 passes, so the complexity is *O(n)*.

#### Space Complexity
The input has *n* slots, padPoints to record the padding has at most *len(input)* slots, and the output buffer has *x + n* slots where *x* is the number of additional 0's added to pad. At most *x* can be *(numChars - 1)(n / 2)* so the complexity is still linear *O(n)*. (Note: this does not include the string parsing done for i/o in the main function.)
