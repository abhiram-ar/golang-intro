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


#### more on string
- go uses utf-8 encoding for string
- utf-8 uses 7-bits (`uint8`) to represent  a character  - (128 characters)
- to handle extended character we could have used utf-32, but this will have a lot of memory wasted 
- go handle this using unicode code points - which has varialbe bytes sizing depending on the character
- when we are indexing a string in go, we are indexing the underlying byte array
- however, if we are using `range` to iterate over the string, the `range` does some heavy lifting under the hood to encoding the variable length byte array to its appropritate character
- key takeaway: when dealing with string in go under the hood we have an array of bytes
- when taking the `len` of the string, we are taking the length of bytes not the number of character in the string

an easier way of iterating string:   
  - cast the string to array of runes (`[]rune`)
  - runes are unicode point numbers which represent the character
  - runes are alias for `int32`
  - we can easier represent `rune` using single quote `''`
   
- strings are immutable in go - we cannot modify them once created 


#### structs 
- structs allows to define our own typed
- structs allows us to define combine/mixed types
- default value of a field is the  default value if its type
- while setting the value of a struct,
  - we can mention  the field and initialize the value
  - or, use short hand and declare the values in order (in this case all fields values should to be initialized)
  - or, directly assign the value using the `.` operator

- when composing another struct we can exclude the name of the field, and only mention its struct type. This causes the fields in the nested struct to be same level as the nested struct


**anoymous struct** 
- struct without a name
- but we have to define and initialize the fields in  the same location
- note: this type of struct is not reusable

- struct also have the concept of methods
- these are functions that are directly tied to the struct, and have access to struct itself

#### interface
- allows us to be more general about structs,
- interface allows processing of multiple struct which has the same signature (common fields) as defined interface
- this make our code more reusable

#### pointers 
- pointers are special data type in go
- we use * just before the data type to indicate that the variable is a pointer to that specific data type in memory eg:`var a *int32` 
- the default value of a pointer is `nil`
- we can allocate memory for a type using `new(type)` which returns the allocated memory address 
- note: `make()` returns the new type after allocating the type in memory

- to dereference a pointer(getting/setting the value at that memory location) we can use the `*` operator
- if we try to set a value to a pointer which does not points to any memory location we will get a *nill pointers dereference error*

- to assign the memory location of another variable to a pointer we can use the `&` operator
```go
var i 10
var *p = &i
```
```


**slices and pointers**
- when we assing a alice it a variable, we are assinging the pointer to the slice to the variable
- if we duplicate this slice variable and modify the slice of duplicate variable, we will mutate the original slice

**pointers and functions**
- when we pass an array to a functions, we are passing the arguments as values (deep copy of array is created)
- to pass my reference, make the function arguments as pointers and send the address of the array
