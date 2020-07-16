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
At the prompt, follow the instructions. Make sure you put quotes around the string. Type *exit* to exit.

#### Time Complexity
Let *T(n)* denote the num of elementary operations that *replace* runs. *T(1)* performs a fixed number of operations. If n > 1, the function will make a recursive call which will perform at most T(n-1) operations, and so on. This is *O(n)* complexity.

#### Space Complexity
Each function call of *replace* takes *O(n)* space. If the maximum depth of the recursion tree is *m* then the space complexity of the recursive function is *O(mn)*.
