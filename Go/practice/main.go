package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const aConst int = 64

var outSide string = "this was crteated outside a function"

func main() {
	sectionDelimiter("VARS")
	vars()
	sectionDelimiter("INPUTS")
	input()
}

func sectionDelimiter(name string) {
	fmt.Println("")
	fmt.Printf("---%s---", name)
	fmt.Println("")

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

func input() {
	reader := bufio.NewReader(os.Stdin)
	// fmt.Print does not place a line feed at the end
	fmt.Print("Enter text: ")
	// using _ as a variable name creates a 'null' variable
	input, _ := reader.ReadString('\n')
	fmt.Println("You entered:", input)

	fmt.Print("Enter a number: ")
	numInput, _ := reader.ReadString('\n')
	aFloat, err := strconv.ParseFloat(strings.TrimSpace(numInput), 64)
	// deal with error object
	if err != nil {
		fmt.Println("\n---ERROR---")
		fmt.Printf("Error type was: %T\n", err)
		fmt.Println("Error msg: ", err)
	} else {
		fmt.Println("Your number was:", aFloat)
	}
}
