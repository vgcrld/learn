# Java Script Notes

[Javascript](javascript.md) | [Other](other.md) | [Author](author.md)

## VS code

Use the VS plugin of `live-server`. Be sure to turn off autosave. It will start a session in chrome. From there you can use the debugger and such. You can use autosave and it will update continually but that's kind of overkill.

## var, let and const

`const` is makeing a value so it cannot be reset or changed. Use uppercase: `const DATA='richard'`. Even with `const` can still mutate and array. So use `Object.freeze(OBJ)` to make it unmutable.

```javascript
"use strict";
const ARR = [];
ARR.push('bob');
Object.freeze(ARR);
ARR.push('sue');     // Error
```

`var` is local scoped but not block. To force block you should use `let`.

`let` is for loop and if scoped. It will also not `let` you delare a variable twice in the same scope. `let` makes your programing safer. Many people will only use `let` and `const`.

`"use strict"` will enforce more checks. Notice it is just a hanging string in your file. So something like this will let you know that `val` was not declared.

```javascript
"use strict"
val = 10
```

If you don't use `var` or `let` then it will be global if called.

```javascript
for (let i=1; i<10; i++) {
   console.log(i);
}
console.log("Stuff")
console.log(i)
```

using let here makes sure that the i says in the for. It is undefined outside. 
Using var in the same will make it available outside of the for.

```javascript
function inAFunc() {
   for (let i = 1; i < 10; i++) {
       console.log(i);
   }
   itsGlobal = 'damn'
}
```

## Functions

Functions must return something or undefined is returned.


Use `JSON.stringify(value)` to make json a string.

Compare is `==`
Strict compare is `===`
so 

```javascript
3 == 3 // true 
3 === '3' // false

99 != '99'  // false
99 !==  '99'  // true 
```

## Conditions

There are a few examples:

```javascript
if (val > 99 && val < 105) {
  console.log('yes')
}
```

The switch is simple but needs a break:

```javascript
   switch (val) {
       case 1:
           break;
       case 2:
       case 3:
           answer = "Other";
           break;
       case 1:
           answer = "Alpha";
           break;
       default:
           answer = "<undefined>"
           break;
   }

```

## Objects 

JSON and Javascript Objects are one in the same:

```javascript
var ourDog = {
   "name": "Paul",
   "legs": 6,
   "tails": 99,
   "friends": ["bob", "sue"]
}
ourDog.name
ourDog['legs']

```

You can delete an object with the `delete` keyword.

`delete ourObject[0].name`

This will allow you to create a new object if it does not already exist:

```javascript
objlist[id][prop] = obj || []
```

This is similar to the ruby `||=` but it's seems to only be for objects.

## While and For loop

```javascript
var myArray = [];

var i = 0;

while(i < 5) {
    myArray.push(i);
    i++   
}

```

This will push the 12 times tables onto an array of mytable.

```javascript
function makeTable(val) {
    let myArray = [];
    for (var i = 1; i <= 12; i++) {
        myArray.push(val * i)
    }
    return myArray
}

var mytable = []
for (let i = 0; i <= 12; i++ ) {
    mytable.push(makeTable(i))
}
```

There is also a `do...while` loop. This will always execute once.

```javascript
var i = 0 
var myArray = []
do {
    myArray.push(i);
    i++;
} while (i<5)
console.log(myArray)
```

## Random

Use `Math.random()` to create a random value. Must use other fucntions to get Integer.

```javascript

function randNumber(val) {
    addIt = Math.floor(Math.random() * 2)
    return Math.floor(Math.random() * val) + addIt
}

arr = new Set
for (let i = 0; i <= 1000; i++) {
  arr.add(randNumber(12))
}



arr.forEach(element  => {
    console.log(element)
});
```

## Fat Arrow

The fat arrow is a shortcut to creating `annonymous` fuctions. Rather than doing the long way:

```javascript

const x = function (val1, val2) {
    console.log(val1);
    console.log(val2);
}

x('bob', 'sue')
```

You can use the new way

```javascript
const y = (val1, val2) => {
    console.log(val1);
    console.log(val2);
}

y('mark', 'rich')
```

It says that `var y` is poiting to a function with `val1` and `val2` as inputs. I like it. I like it a lot.

By default the arrow function implies the return value. 

```javascript
// The old way
const reg = function() {
    1 + 1;
}

// The new way
const fat = () => 1+1;

xx = reg()  // undefined
yy = fat()  // 2
```


## Parse to Integer

Sometimes you need a string to an it.

```javascript
var val = '2';
var num = parseInt(val);
var nan = parseInt('richard');

var log = val => { 
    console.log(val)
    console.log(typeof val)
}
log(val) // 2
log(num) // 2
log(nan) // NaN
```

The conversion of `richard` using `parseInt` is `NaN` but the type is `number`.

You can provide the radix and convert from a different base to decimal. Octal, Hex, etc...

```javascript
// Quick loger function 
var log = val => { console.log(val) }

// Convert from binary, radix 2
var fromBinary = val => {
    return parseInt(val, 2)
}

// Convert from Hex, radix 16 
var fromHex = val => {
    return parseInt(val, 16)
}

log(fromBinary('100'))      // 4
log(fromBinary('1001'))     // 9

log(fromHex('FF'))          // 255
log(fromHex('a8'))          // 168
log(fromHex('10'))          // 16
```

The `radix` can go up to 36. Over that `nan` is returned.

## Ternary Operator

Just like in ruby you can test with a single line:

```javascript
val = false
name1 = val ? 'Rich' : "bob";  // bob
val = true
name2 = val ? 'Rich' : "bob";  // rich
```

## More on Set

Set is pretty awesome. It has a great way of adding items that are not duplicate. 

```javascript
// Quick console
const log = msg => console.log(msg);

// sqrt method
const sr = o => Math.sqrt(o);

// Create a set add all unique floor values
let myset = new Set;
for (let i = 0; i <= 100; i++) {
    //   myset.push(math.floor(sr(i)))
    let val = Math.sqrt(i)
    let fl = Math.floor(val)
    myset.add(fl)
}

// Iterate over the set
for (let item of myset){
    log(item)
}
```