package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

const aConst int = 64

var outSide string = "this was crteated outside a function"

func main() {
	sectionDelimiter("VARS")
	vars()
	sectionDelimiter("INPUTS")
	//input()
	sectionDelimiter("MATH OPERATORS")
	mathOperators()
	sectionDelimiter("TIME")
	playWithTime()
	sectionDelimiter("POINTERS % Arrays")
	pointersAndArrays()
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

func mathOperators() {
	var add int = 1 + 5
	var sub int = add - 1
	var mult int = sub * 4
	mult64 := float64(mult)
	var div float64 = (mult64 / 6)
	fmt.Println(div)

	fmt.Println(math.SqrtPi)

	var aFloat = math.Pi
	var roundedFloat = math.Round(aFloat)
	fmt.Printf("Original: %#v, Rounded: %v\n", aFloat, roundedFloat)

	i1, i2, i3 := 12, 45, 68
	intSum := i1 + i2 + i3
	fmt.Println("Sum of integers:", intSum)

	f1, f2, f3 := 23.5, 65.1, 76.3
	floatSum := f1 + f2 + f3
	fmt.Println("Float Sum (before rounding):", floatSum)
	floatSum = math.Round(floatSum*10) / 10
	fmt.Println("Float Sum (after rounding):", floatSum)

	circleRadius := 15.5
	circumfrance := 2 * math.Pi * circleRadius
	fmt.Printf("Circumference: %.2f\n", circumfrance)

}

func playWithTime() {
	now := time.Now()
	fmt.Println(now)
	//EST, _ := time.LoadLocation("America/New_York")
	t := time.Date(1998, time.January, 10, 5, 0, 0, 0, time.UTC)
	fmt.Println(t.Local().Format(time.RFC822))

	t2 := time.Now()
	var timeFormatString string = "Jan 2 Mon 2006 - This is some filler text - 04:05\n"
	fmt.Print(t2.Format(timeFormatString))
}

func pointersAndArrays() {
	var colors = [5]string{"red", "green", "blue"}
	for i := 0; i < len(colors); i++ {
		fmt.Println(colors[i])
	}
	fmt.Println(len(colors))
}
