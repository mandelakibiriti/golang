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
```