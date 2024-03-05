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
console.log(fruits2)

// removing items from arrays. firstly, in ways that affect the original array
fruits2.pop()


let i = 0
while (i < 10) {
    fruits2.pop()
    i++
}
i = 0
console.log(fruits2)

while (i<10) {
    fruits2.shift()
    i++
}
console.log(fruits2)