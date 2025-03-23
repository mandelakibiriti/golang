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