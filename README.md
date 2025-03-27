### Go Development
Running Go code is as simple as:
> `go run .`
## Concept 1 - Local Module Implementation
Importing local modules in a conversion script where the main function imports the conversion module as a package from the conversion folder.
In setting up the development environment the `go.mod` file defines the packages needed for the application to run.
- The `convo.go` file defines the conversion function. 
- The `convo_test.go` file has the test cases for the conversion function.
- The `main.go` file calls the `conv` function and runs the application.

## Concept 2 - Third party Package Implementation
The `main.go` file imports the **mux** third party package and by running the command below creates the `go.sum` file which highlights the imported packages and third party packages are installed and are accessible: 
> `go mod init dev_folder_name` 
- The main function assigns an r object to create a route to serve text in the route in a html page.

## Concept 3 - Variable Type Assignment
Show cases how to assign variables as types, values or values via function. Of note you can assign mulitiple values to multiple variables.

## Concept 4 - Operators
Using operators with numbers. Note variable assignment with the`:=` notation which allows Go's compiler to simultaneuously assign a value and implicitly define an appropriate type to that variable.
- Note: the use of shorthand operators ++ (increament by 1), --(decrease value by 1), += increase and assign by x value which is the opposite of -= which assigns and decreases a variable by that x assigned value.

## Concept 5 - Zero Values
The zero value of a variable is the empty or default value for that variable’s type. Go has a set of rules stating that the zero values are for all the core types. Using `fmt.Printf` we can get more detail about a value’s type which uses a template language that allows us to transform passed values.

|Type      | Default Value |
| -------- | ------------- |
|bool      |false          | 
|Numbers(floats and integers)|0|
|String|""(empty string)|
|Pointers, Structs, Maps, Interfaces, Slices, Channels|nil|
- Of importance using `%+v` prints out the variable name and default value. Others such as `%T` print out the type of the variable, `%d` prints out the decimal(base10) and `%s` prints out a string.

## Concept 6 - Pointers
Passing values by copying tends to result in code that has fewer bugs. With this method of passing
values, Go can use its simple memory management system, called the stack. The downside is that
copying uses up more and more memory as values get passed from function to function.

You can use the built-in new function for this which is intended to be used to get some memory for a type
and return a pointer to that address. 

- There is an alternative to copying that uses less memory. Instead of passing a value, we create something
called a pointer and then pass that to functions. A pointer is not a value itself, and you can’t do anything
useful with a pointer other than getting a value using it

- When creating a pointer to a value, Go can’t manage the value’s memory using the stack. This is because
the stack relies on simple scope logic to know when it can reclaim the memory that’s used by a value,
and having a pointer to a variable means these rules don’t work.

- When a value gets copied, Go needs CPU cycles to get that memory and then release it later. Having a value on the heap means that it then needs to be managed by the complex garbage collection process. This process can become a CPU bottleneck in certain situations – for example, if there are lots of values on the heap. When this happens, the garbage
collector has to do lots of checking, which uses up CPU cycles.
### Function pointers
- If you have a pointer variable or have passed a pointer of a variable to a function, any changes that are made to the value of the variable in the function also affect the value of the variable outside of the function
- Passing values by a pointer has traditionally been shown to be ***more error-prone, so use
this design sparingly***. It’s also common to use pointers in functions to create more efficient code, which Go’s standard library does a lot

## Concept 7 - Const and Enums
- The const keyword is used to create a constant value. Constants are like variables, but you can’t change their initial values. These are useful for situations where the value of a constant doesn’t need to or shouldn’t change when your code is running.

## Concept 8 - Scopes
All the variables in Go live in a scope. The top-level scope is the package scope. A scope can have child scopes within it. The parent-child relationship is defined when the code compiles, not when the code runs. When accessing a variable, Go looks at the scope the code was defined in. If it can’t find a variable with that name, it looks in the parent scope, then the grandparent scope, all the way until it gets to the package scope. It stops looking once it finds a variable with a matching name or raises an error if it can’t find a match.

- To put it another way, when your code uses a variable, Go needs to work out where that variable was defined. It starts its search in the scope of the code using the variable it’s currently running in. If a variable definition using that name is in that scope, then it stops looking and uses the variable definition to complete its work. If it can’t find a variable definition, then it starts walking up the stack of scopes, stopping as soon as it finds a variable with that name. This searching is all done based on a variable name. If a variable with that name is found but is of the wrong type, Go raises an error.

## Concept 9 - Control
It’s common to need to call a function but not care too much about the returned value. Often, you’ll want to check that it executed correctly and then discard the returned value; for example, sending an email, writing to a file, or inserting data into a database: most of the time, if these types of operations execute successfully, you don’t need to worry about the variables they return. 
- Unfortunately, the variables don’t go anywhere as they are still in scope. To stop these unwanted variables from hanging around, we can use what we know about scope rules to get rid of them. The best way to check for errors is to use initial statements on if statements. The notation looks like this: 
```golang
    // If 
    if <initial statement>; <boolean expression> { 
        <code block> 
    } else if <condition> {
        <code block>
    }

    // Switch
    switch <initial statement>; <expression> {
        case <expression>:
            <statements>
        case <expression>, <expression>:
            <statements>
        default:
            <statements>
    }

    // Endless loop with break and continue
    for {
        if <condition> {
            <statment>
            continue;
        } else if <condition> {
            <statement>
            break;
        }
    }
    
    // For loop with condition
    for <initial statement>; <condition>; <post statement> {
        <statements>    
    }

    // Defer function - code recovery
    defer func () {
        <code block to run recovery logic>
    }()

    <code block with error>

    // Defer to handle operations post error
    <code block to open file>

    defer function to close file
```

## Concept 10 - Basic Data types
Types are needed to make data easier for humans to work with. Computers only think about data in
binary. Binary is hard for people to work with. By adding a layer of abstraction to binary data and
labelling it as a number or some text, humans have an easier time reasoning about it. Reducing the
cognitive load allows people to build more complex software because they’re not overwhelmed by
managing the details of the binary data.

The way data is stored is also a large part of what defines a type. To allow for the building of efficient
software, Go places limits on how large some of its types can be. For example, the largest amount
of storage for a number in Go’s core types is 64 bits of memory. 

### Number Types
- uint and int are either 32 or 64 bits, depending on whether you compile your code for a 32-bit
system or a 64-bit system. It’s rare nowadays to run applications on a 32-bit system, as most systems
nowadays are 64-bit. ***Only think about using the other types when using an int type is causing a problem.***

| Type      | Size (bits)| Size (bytes) | Value Range (int Signed)       | Value Range (uint Unsigned)    |
|-----------|------------|--------------|--------------------------------|--------------------------------|
| `int8`    | 8          | 1            | -128 to 127                   | 0 to 255                      |
| `int16`   | 16         | 2            | -32,768 to 32,767             | 0 to 65,535                   |
| `int32`   | 32         | 4            | -2,147,483,648 to 2,147,483,647 | 0 to 4,294,967,295          |
| `int64`   | 64         | 8            | -9,223,372,036,854,775,808 to 9,223,372,036,854,775,807 | 0 to 18,446,744,073,709,551,615 |
| `int`     | 32 or 64   | 4 or 8       | Depends on system architecture| N/A                       |
| `uint`    | 32 or 64   | 4 or 8       | N/A                           | Depends on system architecture |
| `uintptr` | 32 or 64   | 4 or 8       | N/A                           | Used for storing pointer addresses |
- Go has two floating-point number types, float32 and float64. The bigger float64 allows for
more precision in numbers. ***float64’s bigger space for storage, it can store more whole numbers and/or more
decimal numbers than float32 can.***
- The byte type in Go is just an alias for uint8, which is a number that has eight bits of storage. Grouping bits into groups of eight was a common standard in early computing and became a near-universal way to encode data. 8-bits have 256 possible combinations of “off ” and “on,” so uint8 has 256 possible integer values from 0 to 255
- When you try to initialize a number with a value that’s too big for the type you are using, you get an
overflow error. The highest number you can have in an int8 type is 127. Wraparound means the number
goes from its highest possible value to its lowest possible value. In the example below the `b` will move from 255 to 0 once it gets to the highest uint8 value.
```go
    var a int8 = 125
	var b uint8 = 253
	for i := 0; i < 5; i++ {
		a++
		b++
		fmt.Println(i, ")", "int8", a, "uint8", b)
	}
```
### Text Types
- When you write text to a string variable it’s called a string literal. There are two kinds of string
literals in Go:
    - Raw – defined by wrapping text in a pair of `
    - Interpreted – defined by surrounding the text in a pair of "
- Raw literals, what ends up in your variable is precisely the text that you see on the screen. With
interpreted literals, Go scans what you’ve written and then applies transformations based on its own
set of rules.
- **`Rune`** is a type with enough storage to store a single UTF-8 multi-byte character. String literals are
encoded using UTF-8. UTF-8 is a massively popular and common multi-byte text encoding standard.
The string type itself is not limited to UTF-8, as Go also needs to support other text encoding
types.
```go
	username := "Sir_King_Über"
	runes := []rune(username)

	for i := 0; i < len(runes); i++ {
		fmt.Print(string(runes[i]), " ")
	}
```

## Concept 11 - Complex Data Types
### Array
When you define an array, you must specify what type of data it may contain and how big the array is in the following form: [<size>]<type>. For example,[10]int is an array of size 10 that contains integers, while [5]string is an array of size 5 that contains strings. 
> The key to making this an array is specifying the size. If your definition didn’t have the size, it would seem like it works, but it would not be an array – it’d be a slice.
```go
[<size>]<type>{<value1>,<value2>,…<valueN>}.
```
Keys refer to the indices or positions used to initialize specific values in an array. By setting values at specific indices using keys, you can initialize an array with desired values at specific positions while leaving other elements at their default values.
```go
[<size>]<type>{<key1>:<value1>,…<keyN>:<valueN>}.
```