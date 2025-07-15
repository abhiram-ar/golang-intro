# golang-init

- go is a statically typed language - once a type is inferred for a variable it cannot be changed
- go is a strictly typed language - we can only do strict rules on which operation can be done on different types

#### packages

- **packages** is  a folder of go files
- every go file belong to a packages
- within a packages all go files should have the same package name
- `main` packages tell the compiler that this packages is the executions starts

#### modules

- **modules** is the collection of these packages
- so when we are initializing a project we are initializing a module

```bash
go mod init githubPath/name
```



#### constants variables and basic data types

##### integer

- `int` will automatically cast the integer value based on system architecture
- `int8`,`int32`, `int64` allows explicit memory allocations
- when initializing we will given an warning that the particular integer type cannot hold larger value
- but we wont get any waning when declaring large value causing buffer over flow

##### float

- we don't have a `float` type we need to be explicit about weather it is `float32` or `float64`

##### string

- when we use the `len` function on a string we get back the number of bytes it use not the actual character length
- go uses `utf-8` for other other characters out of this scope we will have a extra bytes

##### constants

- should be assigned the value during variable initializing

##### runes
- go uses utf-8 for stirng encoding, any non-ascii character will require an extra byte to encode to binary
- for single character non-ascii, we can use runes - which is `int` under the hood but can handle unicode characters


#### control structures
- `error` is a datatype in go with default value being `nil`
- `elseif` and `else` should be in the same line as the previous closing `}` bracker
- for `switch` unlike other languages, we don't want to explicitly break a `case` as if one case is executed it control will break out of switch block
