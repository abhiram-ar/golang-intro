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
- we have two types of switch statements,
  - one without input variable for `switch`
  - other with a input variable/condition for `switch`




#### Arrays
- fixed length - size of the array cannot change once it is initialized
- same type 
- indexable
- contiguous memory location
- in go we cannot have a composite literal without type (go is very strict - `intArr := [3]int32{1,2,3}`)

#### Slices
- "slices wrap arrays to give a more general, powerful and convenient interface to sequence of data"
- the syntax of slice is similar to arrays, but we will not mention size (only `[]`)
- we can add more elements to slice using the in build `append` function - appending returns the new slice therefore its better to save it to the initial variable

**`cap` vs `len`**
- "capacity" is the total size of the underlying array slice uses
- "length" is the currently occupied elements
- if we try to access an element e, `e > len` and `e < cap` we will get an index out of range error

`make` 
- given a type like slice/map as argument,it allocate that type returns the that type for explicitly memory management
- for slices the second numeric argument corresponds to length of underlying array 
- and optional third argument specifies the capacity of the array - default value being length of the array


#### map
- map is a set of key value pairs
- if we try to access a key that does not exist, go will return the default value of for that `value` type
- this can be cause may logical errors.To handle this, when we access a key go returns a second value indicating weather the key exists or not

#### loops
- for loops - with `range` and without `range`
- no while loop in go, but can be archived with for loop
- we can also have normal 3 part loops in go
