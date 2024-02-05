package main

import (
	"fmt"
)

const aConst int = 64

var outSide string = "this was crteated outside a function"

func main() {
	vars()
}

func vars() {
	// Explicent variable typing, telling the compialer exactly what kind of value to expect
	var aString string = "Go Var"
	// Print variable
	fmt.Println(aString)
	// Print formated text
	fmt.Printf("The variables type is %T\n", aString)

	var anInteger int = 42
	fmt.Println(anInteger)
	fmt.Printf("The variables type is %T\n", anInteger)

	var defaultInt int
	fmt.Println(defaultInt)

	// Implicent typing; now the compialer gueses what the type is
	var implicetString = "new string"
	fmt.Println(implicetString)
	fmt.Printf("This is an implicently typed (%T)\n", implicetString)

	var implicitInt = 24
	fmt.Println(implicitInt)
	fmt.Printf("This is an implicitly typed (%T)\n", implicitInt)

	// implicet typing sans 'var' with := (walrus opperator)
	// this can only be done within a function.
	newImpStr := "*almost* pythonic syntax 🐍"
	fmt.Println(newImpStr)
	fmt.Printf("This is an implicently typed (%T)\n", newImpStr)

	// printing a constant and a var declared ouside this function
	fmt.Println(outSide)
	fmt.Println(aConst)
}
