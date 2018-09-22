# Early stage notes for learning Go

 - `go build` puts the code into binary. Doing so turns it into an executable command `./[file] rather than a file to be referenced and run by `go run [file]` 

 - Go doesn't support "#" comments within an executable file. Use "//" before comment line
 
 - `fmt.Println("1+1 =", 1+1)` Prints: "1+1 = 2" The calc function is done w/o quotes seperated by comma

 - Variables: `var [variable name] type` declares 1 or more varibles. Ex `var a = "blue"` will print "blue" when `fmt.Println(a)` is run
    - Declare multi variables at a time separated w/ comma
    - `:=` syntax is shorthand for declaring and initializing a variable, but only usable for a var inside the function. Outside the func, every satement begings withe a keyword. 
	- EX: Use `a := "apple"` instead of `var a = "apple"    

 - Types:
    - bool aka  boolean
    - string
    - int  int8  int16  int32  int64
    - uint uint8 uint16 uint32 uint64 uintptr
    - byte // alias for uint8
    - rune // alias for int32
      - represents a Unicode code point
    - float32 float64
    - complex64 complex128
    - When you need an integer value you should use int unless you have a specific reason to use a sized or unsigned integer type.
    
 - Truncate float to specific decimal ex: `%.2f` 2 dec places












