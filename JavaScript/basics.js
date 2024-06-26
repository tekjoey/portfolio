// basic javascript knowledge.
// taken from lessons @ https://thevalleyofcode.com/#javascript

// basic variable assignment
const num1 = 1
const num2 = 2
const string = "this is a string"
const rune = `this is a rune. You can substitue num1 for ${num1}`
// console.log(rune)
// alert()

const ary = [1,2,3,4,5, "this is an array", ["array", "inception"], 6]
// console.log(ary[6][1])
// length is a property of arays, not a function for arrays
// console.log(ary.length)
// console.log("this is a string".length)

// introducing the spread operator
const fruits = ["apple", 'pear', 'orange']
const more_fruits = ["banana", ...fruits]
// console.log(more_fruits)

const names = [["john", "doe"], ['jane', 'doe'], ['jimmy', 'neutron'], ['captain', 'kirk']]
// without the spread operator, the items of 'names[1]' stay within their own list within the 'name1' list
const name1 = [names[1], "julliet"]
// with the spread operator, the items from 'names[1] get extracted as 'first-degree' elements of the 'name2' list
const name2 = [...names[1], "julliet"]
// console.log(name2)

// adding items to arrays
const fruits2 = [...fruits]
// console.log(fruits2)
// .push and .unshift acept multiple items
fruits2.push('mango', 'pineapple')
fruits2.unshift('lime', 'lemon')
// they even work with spreads!
fruits2.push(...fruits)
fruits2.unshift(...names)

// arrays can also contain empty items. And you can arbitrarilly assign 'future' values
fruits2[fruits2.length] = 'hello'
fruits2[fruits2.length+3] = 'hello'
fruits2[fruits2.length+5] = 'hello'
// console.log(fruits2)

// removing items from arrays. firstly, in ways that affect the original array
fruits2.pop()

let i = 0
while (i < 10) {
    fruits2.pop()
    i++
}
i = 0
// console.log(fruits2)

while (i<10) {
    fruits2.shift()
    i++
}
// console.log(fruits2)

// .splice() will not fail if deleteCount is greater than array.length
fruits2.splice(2, 3)
// console.log(fruits2)

// now for ways that do not mutate the original array.
const fruits3 = ["lime","lemon","apple","pear","orange","mango","pineapple","apple","pear","orange"]
const other_fruits = [...fruits3.slice(1, 5), ...fruits3.slice(6, fruits3.length)]
// console.log(other_fruits)
// or, say we knew which index number we wanted to exclude (1-indexed)
i = 3 //exclude third item; in this case, 'apple'
const more_other_fruits = [...fruits3.slice(0, i-1), ...fruits3.slice(i, fruits3.length)]
// console.log(more_other_fruits)

// different ways to create arrays
const a = [].fill(5) // .fill(5) does nothing in this case, bcause the array is empty
const b = Array(14) //creates an array with 14 empty spaces
b.fill(0, 2, b.length)
// console.log(a, "\t", b)

// array destructuring
const nums = [1,2,3,4,5,6,7,8,9,0]
const [first, second] = nums
// and our friend the ol' spread operator!
const [ffirst, ssecond, ...others] = nums

//I'm assuming this doesnt work...
//const [...oothers, penultimate, ultimate] = nums
// I am correct: "Rest element must be last element"

// console.log(first, second)
// console.log(ffirst, ssecond, others)

// const [ , , third, ...ooothers] = nums
// console.log(third, ooothers)

// // check if item in array
// console.log(ooothers.includes(1)) //returns false
// console.log(nums.includes(1)) // returns true
// // can also use .includes() to search a sub set of an array
// console.log(nums.includes(1, 3))

// conditionals

// if (true) {
//     doSomething()
// } else if (true) {
//     doSomethingElse()
// } else {
//     doNothing()
// }

// if (!true) {
//     console.log("this will never print")
// }

// // can also omit the braces if one line

// if (true) console.log("This is true")

let javascript = "is cool"
switch (javascript) {
    case "is cool":
        console.log("You're right!")
        // switch does not stop after one evaluation, so remember to add breaks
        break
    case "sucks":
        console.log("You're wrong!")
        break
    default:
        console.log("You're indecisive!")
}

javascript = 'cool';

// the almighty ternery operator!
(javascript === "is cool") ? console.log("Ternery: You're right!") : console.log("Ternery: You're wrong")

//loops
for (let index = 0; index < fruits3.length; index++) {
    const element = fruits3[index];
    //console.log(`An item from the list: ${element}`)    
}

i = 0
do {
    i++
    if (i === 42) break
} while (i<100)

//console.log(i)

while (i<200){
    i++
    if (i===150) {
        //console.log("I equals 150")
        continue //continue skips the following code and 'continues' to the next loop element
    }
    //console.log(i)
}


const list1 = [1,2,3,4,5,6,7,8,9,0]
const obj1 = {name: 'john', sex: 'male'}

//for...of is for interable types
for (const letter of "hello world") {
    //console.log(letter)
}
const stri = "world"
for (const letter of `hi ${stri}`) {
    //console.log(letter)
}

// for...in is similar, but also for objects.
for (let item in obj1) {
    //console.log(obj1[item])
}

//functions
function firstFunction(age=16) {
    const yearsTill21 = 21 - age
    if (yearsTill21 > 0) {
        return `You have ${yearsTill21} years left`
    } else {
        return "You are twenty one or older, go party!"
    }
}

// console.log(firstFunction())


function returnEarly(condition) {
    if (condition) return

    const returnary = []
    const unimportant = 1
    const unimport = 2
    let booring = unimport + unimport
    while (booring < 10) {
        returnary.push(booring)
        booring++
    }
    return [booring, returnary, 1]
    //return booring, returnary, 1 // this would result in only returning the last value.
}

const [rtrn1, rtrn2, rtrn3] = returnEarly()

console.log(rtrn1, rtrn2, rtrn3)

//arrow functions
/*
Arrow functions transform

function name() {
    doStuff()
}

to

() => {
    doStuff()
}

they are unnamed (anonymous) functions.

*/

function regular(param) {
    return param
}

const arrowFunc = () => {
    return
}

const arrow2func = param => param+1
console.log(arrow2func(2))
// This just prints '[Function (anonymous)]', so it looks like the function needs to be invoked
console.log (((parm2) => parm2+5))

// It does need to be invoked! This does work!
console.log (((parm2) => parm2+5)(5))

const add = (num1, num2) => num1+num2

const num3 = add(6,5)
console.log(num3)

// Immediately invoked functions, NEED the semi-colon
;(function(parm3) {parm3+157})(3)
; const state = ((parm4) => {
    // just adding complexity to prove that it can be a 'real' functions, not just a one liner
    const states = {'PA':"pennsylvania", 'SC': "south carolinas"}
    for (const state in states) {
        if (parm4 == state) {
            return states[state]
        }
    }
})('SC')

// This will print 'south carolinas
console.log(state)

// what will this print?
//console.log(state('PA'))
// 'state' is not a function, makese sense. The anonymous function is gone, and it's output (a string) was saved to the variable 'state'

// functions can be recursive
function factorial(n) {
    return n >= 1 ? n * factorial(n-1) : 1
}

console.log(factorial(10))




