# evalexpr

EvalExpr lets you evaluate a string based condition, in the same behaviour as an if statement

## Install
Make sure you have Go installed and have set your GOPATH.
`go get -u github.com/Quen98/evalexpr`

## Example

```
package main

func main() {
    expression := "22 & (31 | (44 & 32))"
    parameters := []string{
        "22",
        "31"
    }
    evalexpr.IsFulfillingCondition(expression, parameters) // returns true
}
```
