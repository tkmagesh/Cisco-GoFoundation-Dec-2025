# Go Foundation

## Magesh Kuppan
- tkmagesh77@gmail.com

## Schedule
| What | When |
|-----|----|
| Commence | 9:00 AM |
| Tea Break | 10:30 AM (20 mins) |
| Lunch Break | 12:30 PM (1 hour) |
| Tea Break | 3:00 PM (20 mins) |
| Wind up | 4:30 PM |

## Software Prerequisites
- Go Tools (https://go.dev/dl)
- Visual Studio Code (http://code.visualstudio.com)
    - Go Extension (https://marketplace.visualstudio.com/items?itemName=golang.Go)

## Repository
- https://github.com/tkmagesh/cisco-gofoundation-dec-2025

## Methodology
- No powerpoint
- Code & Discussion

## Why Go?
- Simplicity
    - ONLY 25 keywords
    - No access modifiers (public, private, protected)
    - No reference types (everything is a value)
    - No class (only structs)
    - No inheritance (only composition)
    - No pointer arithmatic
    - No exceptions (only errors)
    - No try..catch..finally construct
    - No implicit type conversion
- Performance
    - Equivalent to C++
    - Close to hardware
        - Compile to machine code
        - Tooling support for cross-compilation
- Support for Concurrency
    - Concurrency problems are typically solved using OS Threads
        - OS Threads are costly (~2MB of memory)
    - Go offers a cheaper & efficient alternatives called "Goroutines"
        - Goroutines (2KB of memory)
        - Built in scheduler
    - Language support for concurrency
        - "go" `keyword`
        - "chan" `data type`
        - "<-" `operator`
        - "range" `construct`
        - "select-case" `construct`
    - Standard Library support
        - `sync` package
        - `sync/atomic` package

## Compilation
```shell
go build <filename.go>
# ex : go build 01-hello-world.go

go build -o <binary_name> <filename.go>
# ex : go build -o hw 01-hello-world.go
```

### Compilation & Execution
```shell
go run <filename.go>
# ex : go run 01-hello-world.go
```

### Cross Compilation

#### Environment variables for cross compilatio
- GOOS
- GOARCH

#### To get the list of support platforms
```shell
go tool dist list
```

#### To get the environment variables used by the go tool
```shell
# list all 
go env

# get specific
go env <var_1> <var_2> ...
```

#### To change the environment variable values
```shell
go env -w <var_1>=<new_value_1> <var_2>=<new_value_2> ...
```

#### To cross compile (in mac & linux)
```shell
GOOS=windows GOARCH=amd64 go build 01-hello-world.go
```

#### In windows (powershell)
```powershell
$env:GOOS="windows"; $env:GOARCH="amd64"; go build <filename.go>
```

## Standard Library documentation
- https://pkg.go.dev/std

## Data Types
- string
- bool
- integers family
    - int8
    - int16
    - int32
    - int64
    - int
- unsigned integers family
    - uint8
    - uint16
    - uint32
    - uint64
    - uint
- floating points family
    - float32
    - float64
- complex family
    - complex64 ( real[float32] + imaginary[float32] )
    - complex128 ( real[float64] + imaginary[float64] )
- aliases
    - byte (alias for unsigned int)
    - rune (alias for unicode code point)

### Zero values
| Data Type | Zero value |
------------ | ------------- |
|int family     | 0 |
|uint family    | 0 |
|complex family | (0+0i) |
|string         | "" (empty string) |
|bool           | false |
|byte           | 0 |
|interface      | nil |
|pointer        | nil |
|function       | nil |
|struct         | struct instance |

## Scope
### Block Scope
### Function Scope
- CAN use ":="
- CANNOT have unused variables

### Package Scope
- CANNOT use ":="
- CAN have unused variables

## iota
Auto generated values for a `group` of constants

## constructs
- if else
- for
- switch case
- function

## Functions
- Functions can return more than one result
    - Named results
- Variadic functions
- Anonymous functions
    - function defined in another function
    - Cannot have a name
    - Has to be immediately invoked (if not assigned to a variable)
- Higher Order Functions (functions can be treated like data)
    - Functions can be assigned as values to variables
    - Functions can be passed as arguments to other functions
    - Functions can be returned as return values from other functions

## Collections
### Arrays
- Fixed sized typed collection
### Slices
### Maps

## Error Handling
## Panic & Recovery
## Modules & Packages
## OOP
### Structs
### Struct Composition
### Methods
### Type Assertion
### Interfaces