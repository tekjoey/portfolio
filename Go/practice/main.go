package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"math"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

const aConst int = 64

var outSide string = "this was crteated outside a function"

func originalFunction() int {
	fmt.Println("hello world for a reassigned function")
	return 0
}

var reassignedFunction = originalFunction

// structs and methods

type Cat struct {
	Lives int
	Age   int
}

// Death is what happens when "Cat" looses a life
func (c *Cat) Death() int {
	c.Lives--
	return c.Lives
}

func (c *Cat) Birthday() int {
	c.Age++
	return c.Age
}

func main() {
	sectionDelimiter("VARS")
	vars()
	sectionDelimiter("INPUTS")
	//input()
	sectionDelimiter("MATH OPERATORS")
	mathOperators()
	sectionDelimiter("TIME")
	playWithTime()
	sectionDelimiter("POINTERS & Arrays")
	pointersArraysMaps()
	sectionDelimiter("STRUCTS")
	structures()
	sectionDelimiter("PROGRAM FLOW")
	programFlow()
	sectionDelimiter("FUNCTIONS")
	reassignedFunction()
	total, length := addAllValues(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Printf("The total sum is %v\nThe number of values was %v\n", total, length)
	cat1 := Cat{9, 9}
	fmt.Println("Lives before death", cat1.Lives)
	catLives := cat1.Death()
	fmt.Println("Lives after death", cat1.Lives, catLives)
	fmt.Println("Age BEfore Birthday", cat1.Age)
	catAge := cat1.Birthday()
	fmt.Println("Age After Birthday", cat1.Age, catAge)
	sectionDelimiter("Files")
	writeToFile("Text to be written from Go", "./text.txt")
	//defer fmt.Println(readFromFile("./text.txt"))
	sectionDelimiter("Network requests")
	postsJSON := networkRequests()
	posts := decodeJSON(postsJSON)
	for _, post := range posts {
		fmt.Printf("User %v posted an article titled [%v]\n", post.UserID, post.Title)
	}
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

func pointersArraysMaps() {
	var colors = [5]string{"red", "green", "blue"}
	for i := 0; i < len(colors); i++ {
		fmt.Println(colors[i])
	}
	fmt.Println(len(colors))
	// important note about arrays is that they have a static size.
	// they can not be increased, or decreased.
	// their values can be changed
	// a better solution is to use slices, which is an abstraction of arrays

	colorSlice := []string{"red", "green", "blue"}
	fmt.Println(colorSlice)

	colorSlice = append(colorSlice, "purple")
	fmt.Println(colorSlice)

	// can also use make() to "make" a new slice, and give it a maximum value

	numbers := make([]int, 0, 10)
	numbers = append(numbers, 1)
	numbers[0] = 6
	fmt.Println(numbers)
	fmt.Println(len(numbers))

	//Maps
	states := make(map[string]string)
	states["SC"] = "'South Carolina'"
	states["PA"] = "Pennsylvania"
	fmt.Println(states)

	penna := states["PA"]
	fmt.Println(penna)

	fmt.Println(states)
	states["NY"] = "New York"
	fmt.Println(states)
	states["CA"] = "California"
	states["ID"] = "Idaho"
	states["HI"] = "Hawaii"
	// maps are unordered, so if i want to print them alphabetically...
	//first i create a slice with the keys
	fmt.Println("--START ALPHABETICAL SORT--")
	stateKeys := []string{}
	for k := range states {
		stateKeys = append(stateKeys, k)
	}

	//then we sort the keys
	sort.Strings(stateKeys)

	// then we can iterate through those keys and print the map in alphabetical order
	for i := range stateKeys {
		fmt.Printf("%v -> %v\n", stateKeys[i], states[stateKeys[i]])
	}

	// and now for some recursive mapping!
	aMap := make(map[string]map[string]int)
	aSecondMap := make(map[string]int)
	aThirdMap := make(map[string]int)

	aSecondMap["first2"] = 1
	aSecondMap["second2"] = 2
	aSecondMap["third2"] = 3
	aSecondMap["fourth2"] = 4

	aThirdMap["first3"] = 1
	aThirdMap["second3"] = 2
	aThirdMap["third3"] = 3

	aMap["first"] = aSecondMap
	aMap["second"] = aThirdMap

	fmt.Println("\n\n---Values in 'aMap'--")
	for k0, v0 := range aMap {
		fmt.Println("Key:", k0)
		fmt.Println("Value:")
		for k1, v1 := range v0 {
			fmt.Printf("\tKey:%v, Value:%v\n\n", k1, v1)
		}
	}

	// can i create a map with initial values?

	aMapWithVals := map[string]string{"String1": "value1", "string2": "value2"}
	fmt.Println(aMapWithVals)

	//indeed i can!
}

// Dog is a struct representation of Mans Best Friend.
// Only some values have been implemented. Dog is not ment to be cainine-complete.
// I wonder Why it does not preserve the new-lines i have been typing.
// I probobly need to use a multi-line comment for that, eh?
type Dog struct {
	Breed        string
	Weight       int
	_age         int
	FavoriteFood string
	Age          int
}

func structures() {
	// structs are similar to classes, but without inheritance

	poodle := Dog{"Poodle", 100, 10, "Steak", 10}

	fmt.Println(poodle.Breed, poodle.Weight, poodle._age)

}

func programFlow() {
	x := 42

	if y := 42; x == y {
		fmt.Println("They are equal")
	} else {
		fmt.Println("they are not equal")
	}

	// now for switch statements!

	// commented out so it doesnt run everytime
	// reader := bufio.NewReader(os.Stdin)
	// fmt.Print("Do you want to learn Go?(y/n) ")
	// input, _ := reader.ReadString('\n')
	// fmt.Println("input equals:", input)

	input := "y"
	switch input {
	case "y\n":
		fmt.Println("Glad to hear it!")
		//fallthrough
	case "n\n":
		fmt.Println("Then what are you doing here?")
		//fallthrough
	default:
		fmt.Println("Umm, you didn't type the right letter..Try again!")

	}

	// and now for 'for' statements.
	// we have already seen some 'for' statements, but GO has a few tricks
	// Go has no 'while' loop, but 'for' can be used in its place

	colors := []string{"Red", "Green", "Blue", "Yellow", "Green", "Purple"}

	//typical C-style for loop
	fmt.Println("---C-Style FOR LOOP---")
	for i := 0; i < len(colors); i++ {
		fmt.Println(colors[i])
	}

	// can skip the index assigment/comparison/incrementation steps by just using 'range' keyword
	// very similar to Pythons "for i in colors"
	// also note that 'range' is a keyword, not a function (as it is in python)
	fmt.Println("---C-Style FOR LOOP Using 'range'---")
	for i := range colors {
		fmt.Println(colors[i])
	}

	// not sure if it's actually called 'direct assignment', but that phrase makes sense for me.
	fmt.Println("---direct assignment using 'range'---")
	for index, color := range colors {
		fmt.Printf("Index number %v in the color slice is: %v\n", index, color)
	}

	//additionally, go has goto statements, which is a pretty cool thing
	// i feel like i'm writing assembly 👨‍💻

	var out int
	j := 0
reStart:
	out = j*j + out
	if out > 12 {
		goto theEnd
	} else {
		j++
		goto reStart
	}
theEnd:
	{
		fmt.Println(out)
		fmt.Println("i assume this is still part of theEnd")
		fmt.Println("i wonder if this is still part of theEnd#1")
		fmt.Println("i wonder if this is still part of theEnd#2")
		fmt.Println("i wonder this is still part of theEnd #3")
	}
}

// can have arbitrary number of input values.
// 'values' is an array
func addAllValues(values ...int) (int, int) {
	var total int
	for _, v := range values {
		total += v
	}

	return total, len(values)
}

func fileError(err error) {
	if err != nil {
		panic(err)
	}
}

func writeToFile(text string, location string) int {
	// first create a file reference
	file, err := os.Create(location)
	fileError(err)
	// Then write text to file
	length, err := io.WriteString(file, text)
	fileError(err)
	fmt.Println("File written to ", location)
	// make sure you close the file at the end!
	defer file.Close()
	return length
}

func readFromFile(fileName string) string {
	data, err := ioutil.ReadFile(fileName)
	fileError(err)
	fmt.Println("Read text from", fileName)
	return string(data)
}

func networkRequests() string {
	url := "https://jsonplaceholder.typicode.com/posts"
	resp, err := http.Get(url)
	fileError(err)
	fmt.Printf("Response type: %T\n", resp)
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	fileError(err)

	content := string(bytes)
	// fmt.Print(content)
	return content
}

type Post struct {
	UserID, ID  int
	Title, Body string
}

func decodeJSON(content string) []Post {
	posts := make([]Post, 0, 20)
	decoder := json.NewDecoder(strings.NewReader(content))
	//fmt.Printf("%T\n", decoder)
	_, err := decoder.Token()
	fileError(err)
	var post Post
	for decoder.More() {
		err := decoder.Decode(&post)
		fileError(err)
		posts = append(posts, post)
	}
	return posts
}
