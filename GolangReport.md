# Golang Report

#### Author: Jacky

## 1. Primitive Types 

#### 1.1 Boolean types
| Types | Kinds of Values | Range of Values | Memories |
| ----- | --------------- | --------------- | -------- |
| bool  | true, false     | true, false     | 8-bit    |

#### 1.2 Numeric types
| Types      | Kinds of Values     | Range of Values                               | Memories |
| ---------- | ------------------- | --------------------------------------------- | -------- |
| uint8      | 0, 1, 2, …          | (0 to 255)                                    | 8-bit    |
| uint16     | 0, 1, 2, …          | (0 to 65535)                                  | 16-bit   |
| uint32     | 0, 1, 2, …          | (0 to 4294967295)                             | 32-bit   |
| uint64     | 0, 1, 2, …          | (0 to 18446744073709551615)                   | 64-bit   |
| int8       | -1, 0, 1, …         | (-128 to 127)                                 | 8-bit    |
| int16      | -1, 0, 1, …         | (-32768 to 32767)                             | 16-bit   |
| int32      | -1, 0, 1, …         | (-2147483648 to 2147483647)                   | 32-bit   |
| int64      | -1, 0, 1, …         | (-9223372036854775808 to 9223372036854775807) | 64-bit   |
| float32    | 0.1, -0.2, …        | IEEE-754                                      | 32-bit   |
| float64    | 0.1, -0.2, …        | IEEE-754                                      | 64-bit   |
| complex64  | 1+4i, 0+1i, 1+0i, … | float32 real and imaginary parts              | 64-bit   |
| complex128 | 1+4i, 0+1i, 1+0i, … | float64 real and imaginary parts              | 128-bit  |
| byte       | 0, 1, 2, …          | alias for uint8                               | 8-bit    |
| rune       | -1, 0, 1, …         | alias for int32                               | 32-bit   |

#### 1.3 String types
| Types  | Kinds of Values | Range of Values     | Memories |
| ------ | --------------- | ------------------- | -------- |
| string | “Hello”, “”, …  | Array of characters | 8-bit    |

#### 1.4 Reference
1.	[Primitive Data Types in Golang](https://medium.com/golang-jedi-knight/primitive-data-types-in-golang-35a291df3bbe#:~:text=In%20golang%2C%20we%20have%20int,coming%20from%20other%20programming%20languages)
2.	[Golang.org/Types/](https://golang.org/ref/spec#Types)

## 2. Pointers/references
The declartion of pointer in golang almost identical in C/C++. A pointer type denotes the set of all pointers to variables of a given type, called the base type of the pointer. The value of an uninitialized pointer is `nil`. We can access the variable by pointers and manipulate them by pointers directly in golang. For example, in the following code, we use a pointer to access the value of `i` and change the value of `i` through the pointer as well.

```go
package main
import "fmt"

func foo(p *int){
    *p = 1
}

func main(){
    var i int= 0;
    var p1 *int = &i
    fmt.Println("The value pointed by p1 is", *p1)
    *p1 = 2
    fmt.Println("After change by pointer, i is", i)
}
```

#### Reference
1. [Pointer types](https://golang.org/ref/spec#Pointer_types)

## 3. Background

Go was designed at Google in 2007 to improve programming productivity in an era of multicore, networked machines and large codebases. The designers wanted to address criticism of other languages in use at Google, but keep their useful characteristics:
•	static typing and run-time efficiency (like C),
•	readability and usability (like Python or JavaScript)
•	high-performance networking and multiprocessing.

Go compiles quickly to machine code yet has the convenience of garbage collection and the power of run-time reflection. It's a fast, statically typed, compiled language that feels like a dynamically typed, interpreted language.

Technology companies would hire golang programmers as developers.

#### Reference
1.	[Wikipedia-Go](https://en.wikipedia.org/wiki/Go_(programming_language))
2.	[Golang.org/doc](https://golang.org/doc/)



## 4. User defined types

Go support user defined types. The program below runs (`go run main.go`) correctly and prints 1. Here we use `type` to declare an user-defined alias to type `int`, which is similar to `typedef` in C

```go
package main

import "fmt"

type Int int

func main(){
    var s Int
	s = 1
    fmt.Printf("%d\n", s) // Compiled correctly Print 1
}
```

## 5. Static vs Dynamic typing

The golang is static typed. If we run the program below, it will return error `cannot use "X" (type untyped string) as type int in assignment` Thus, in golang, the variable holds the type.

```go
package main

import "fmt"

func main(){
    s := 1
    s := "X" // Error
}
```

## 6. Implicitly or explicitly typed

Golang supports both implicitly and explicitly typed. See the program below. Here, we use explicitly typed to declare variable `x` and use implicitly typed to declare variable `y`. 

```go
package main

import "fmt"

func main(){
    var x int = 10 // Explicitly typed
    y := 10 // Implicitly typed
}
```

## 7. Nominal vs Structural typing

In golang, we use nominal typing. To test this, we declare two type with the same structure (same attributes) but difffernt names in the following program. This program does not run with error message `cannot use Child{...} (type Child) as type Person in assignment`. If golang supports structural typing, then this program should run correctly, where we can assign `Child` type to `Person` type variable. Thus, we can conclude that golang uses nominal typing, since assignment across different name is not allowed even if they have the same structure. 

```go
package main
	
type Person struct {
    name string
    age  int
}

type Child struct {
    name string
    age  int
}

func main(){
    var x Person
	x = Child{"Name",11}
}
```

## 8. Typing Strength
Golang is a strongly typed language. As we can see from the following program:
* In part 1, you cannot implicitly coerse a int to a float, which would print error `cannot use i (type int) as type float32 in assignment`
* In part 2, you can implicitly coerse a char to a float or a int
* In part 3, you cannot neither implicitly convert a bool to int/ float nor a int to bool.  

```go

// go run implicit-type.go
package main

import "fmt"

func main(){
    // Part 1
    var i int = 1
    var f float32 = i
    // cannot use i (type int) as type float32 in assignment
    // Part 2
    var c int = 'a'
    // constant 0.3 truncated to integer
    // Part 3
    var b int = true
    var b float32 = true
    //cannot use true (type untyped bool) as type int in assignment
    var b int = 1
    //cannot use 1 (type untyped int) as type bool in assignment
}
```

#### Reference
1. [Golang/Introduction](https://golang.org/ref/spec#Introduction)

## 9. Reflection
By using the reflect package implementing run-time reflection, we are allowed a program to manipulate objects with arbitrary types. In the following program, we use `fmt.Scanln()` to accept input from user and store it to variable `first`, which is a type of structural reflection. We also use `reflect.TypeOf()` to check the type of variable `first`, which is a type computational reflection. Thus, golang supports both reflections. 

```go
package main
  
import "fmt"
import "reflect"

func main() {
    fmt.Println("Enter Your First Name: ")
    var first string
    // Taking input from user
	// Structural Reflection
    fmt.Scanln(&first)
    fmt.Print(first)
	// Computational Reflection
	fmt.Print(reflect.TypeOf(first))
}
```

#### Reference
1. [Package reflect](https://golang.org/pkg/reflect/)
2. [Package fmt](https://golang.org/pkg/fmt/)

## 10. Memory Hazards

#### 10.1 Alias
We create a pointer `p` to `int` variable `i`, which is an example of alias in golang

```go
var i int = 1;
var p *int;
p = &i;
fmt.Println(*p) // Print 1
```

#### 10.2 Dangling references 
When clearing a pointer by setting it to `nil`, the call pointed object again will be taken care of by the Garbage Collection

```go
var i int = 1
var p *int
p = &i
p = nil 
fmt.Println(*p) // Cause Error

```

#### 10.3 Memory leak 

Subslices may also cause kind-of memory leaking. In the following code, after the g function is called, most memory occupied by the memory block hosting the elements of s1 will be lost (if no more values reference the memory block).

```go
var s0 []int 
func g(s1 []int) { 	
    // Assume the length of s1 is much larger than 30. 	
    s0 = s1[len(s1)-30:] 
}
```

#### Reference
1. [Memory Leaking Scenarios](https://go101.org/article/memory-leaking.html)

## 11. Static vs Dynamic Scope
According to official documents, go is lexically scoped using blocks:
1. The scope of a predeclared identifier is the universe block.
2. The scope of an identifier denoting a constant, type, variable, or function (but not method) declared at top level (outside any function) is the package block.
3. The scope of the package name of an imported package is the file block of the file containing the import declaration.
4. The scope of an identifier denoting a method receiver, function parameter, or result variable is the function body.
5. The scope of a constant or variable identifier declared inside a function begins at the end of the ConstSpec or VarSpec (ShortVarDecl for short variable declarations) and ends at the end of the innermost containing block.
6. The scope of a type identifier declared inside a function begins at the identifier in the TypeSpec and ends at the end of the innermost containing block.

```go
package main

import "fmt"

var i = 1

func foo(){
    i += 2
    fmt.Println("[foo] i = ", i)
}

func main(){
    var i = 0
    fmt.Println("[main] i = ", i)
    foo()
}
```

We can verify the scoping using the above program, if golang is dynamic scoping, then `i += 2` in function `foo()` would refers to the variable declared in `main()`, which would prints
```
[main] i =  0
[foo] i =  2
```
On the contrary, if the golang is static scoping, then `i += 2` in function `foo()` would refers to the variable declared globally, which would prints
```
[main] i =  0
[foo] i =  3
```
The above program prints the following which matches that goalng is static scoping.
```
[main] i =  0
[foo] i =  3
```

#### Reference
1. [Declarations and scope](https://golang.org/ref/spec#Declarations_and_scope)

## 12. Polymorphism

Golang is a light-Object Oriented language and supports polymorphism through interfaces only. Such a polymorphism only contains ad hoc polymophism. For example, program below declare a `Area()` function for two different types of structure. This is an example of ad hoc polymorphism.

```go
package main
  
import "fmt"

type Figure interface{
    Area() int
}

type Rectangle struct{
    length int
    width int
}

type Square struct{
    side int
}
  
// function to calculate area of a rectangle
func (rect Rectangle) Area() int{
    area := rect.length * rect.width
    return area
}
  
// function to calculate area of a square
func (sq Square) Area() int{
    area := sq.side * sq.side
    return area
}
  
func main() {
    // declaring a rectangle instance
    rectangle := Rectangle{length: 2,width: 1}
    // declaring a square instance
    square := Square{side: 2}
    // printing the calculated result
    fmt.Printf("Area of rectangle: %.3f.\n", rectangle.Area())
    fmt.Printf("Area of square: %.3f.\n", square.Area())
}
```

#### Reference
1. [Polymorphism in GoLang](https://www.geeksforgeeks.org/polymorphism-in-golang/)


## 13. Strict vs Non-strict Evaluation

Golang is strict evaluation. In the following program, the strict evaluation would print `bar` first before prints `foo`. On the other hand, non-strict evaluation would would print `foo` first before prints `bar`.
```go
package main

import "fmt"

func foo(i int) int{
    fmt.Println("foo")
	return i + 2
}

func bar() int{
    fmt.Println("bar")
	return 1
}

func main(){
    foo(bar())
}
```
Golang prints the following. Thus, golang is strict evaluation.
```
[bar]
[foo]
```

## 14. Parameter Passing
From official FAQ, as in all languages in the C family, everything in Go is passed by value. That is, a function always gets a copy of the thing being passed, as if there were an assignment statement assigning the value to the parameter.

```go
package main

import "fmt"

func foo(i int){
    i = 1
}

func bar(i *int){
    *i = 1
}

func main(){
    var i = 0
    fmt.Println("[main] i = ", i)
    foo(i)
    fmt.Println("[foo] i = ", i)
    bar(&i)
    fmt.Println("[bar] i = ", i)
}
```
The above program prints
```
[main] i =  0
[foo] i =  0
[bar] i =  1
```
Thus, we could conclude that we pass value of `i` by value into function `foo`, since `i` doesn't change after we modify the value of `i` in the `foo` function. This is the default scenario. We can pass by reference in golang by using `&` which return the memory address of variable. Thus, after we pass `&i` into bar function, and modify the value in the memory address of `i`, the value of `i` changes. Conclusively, goalng supports both pass by value and pass by reference.

#### Reference
1. [When are function parameters passed by value?](https://golang.org/doc/faq#pass_by_value)

## 15. The string type
String is a primitive type in golang. String is the set of all strings of 8-bit bytes, conventionally but not necessarily representing UTF-8-encoded text. A string may be empty, but not `nil`. Values of string type are immutable. Unlike C/C++, string in Go is not null-terminated.

Here some example of string operation, more can be found on [Package strings](https://golang.org/pkg/strings/)
```go
package main

import (
	"fmt"
	s "strings"
)

func main() {
    var s1 = "hello"
    var s2 = " world"
    var s3 = s1 + s2
	fmt.Println("Contains:  ", s.Contains("test", "es"))
	fmt.Println("Count:     ", s.Count("test", "t"))
	fmt.Println("HasPrefix: ", s.HasPrefix("test", "te"))
	fmt.Println("HasSuffix: ", s.HasSuffix("test", "st"))
	fmt.Println("Index:     ", s.Index("test", "e"))
	fmt.Println("Join:      ", s.Join([]string{"a", "b"}, "-"))
	fmt.Println("Repeat:    ", s.Repeat("a", 5))
	fmt.Println("Replace:   ", s.Replace("foo", "o", "0", -1))
	fmt.Println("Replace:   ", s.Replace("foo", "o", "0", 1))
	fmt.Println("Split:     ", s.Split("a-b-c-d-e", "-"))
	fmt.Println("ToLower:   ", s.ToLower("TEST"))
	fmt.Println("ToUpper:   ", s.ToUpper("test"))
}
```

#### Reference
1. [Strings and string operations in Go](https://kuree.gitbooks.io/the-go-programming-language-report/content/29/text.html)
2. [Go by Example: String Functions](https://gobyexample.com/string-functions)

## 16. Math operations on numbers
The following are a list of math operations in golang, which includes some non-stardard math operations, such as `<<`, `>>`.
| Operator | Name                | Types                                     |
| -------- | ------------------- | ----------------------------------------- |
| `+`      | sum                 | integers, floats, complex values, strings |
| `-`      | difference          | integers, floats, complex values          |
| `*`      | product             | integers, floats, complex values          |
| `/`      | quotient            | integers, floats, complex values          |
| `%`      | remainder           | integers                                  |
| `&`      | bitwise AND         | integers                                  |
| `|`      | bitwise OR          | integers                                  |
| `^`      | bitwise XOR         | integers                                  |
| `&^`     | bit clear (AND NOT) | integers                                  |
| `<<`     | left shift          | integer << unsigned integer               |
| `>>`     | right shift         | integer >> unsigned integer               |

## 17. Multi-dimentional arrays
Go programming language allows multidimensional arrays with no limits. Here is the general form of a multidimensional array declaration
```go
var variable_name [SIZE1][SIZE2]...[SIZEN] variable_type
```
For example, we could declare a 4 dimension array
```go
var ndarray [1][2][3][4] int
```
We can access the fields by using index. For example, the following program would print each number in the array `a` by call the index in the array.
```go
package main

import "fmt"

func main() {
   /* an array with 5 rows and 2 columns*/
   var a = [5][2]int{ {0,0}, {1,2}, {2,4}, {3,6},{4,8}}
   var i, j int

   /* output each array element's value */
   for  i = 0; i < 5; i++ {
      for j = 0; j < 2; j++ {
         fmt.Printf("a[%d][%d] = %d\n", i,j, a[i][j] )
      }
   }
}
```

#### Reference
1. [Go - Multidimensional Arrays in Go](https://www.tutorialspoint.com/go/go_multi_dimensional_arrays.htm)

## 18. Variables

The types of variables golang supports are: locals, globals, struct member variables, static variables, and constants. See the code below.
```go
package main

import "fmt"

// Global variable
var a int = 1
 
type V struct {
	// Member variables
    d  int
}

func main() {
	// Local variable
   	var b int = 2
	// Constant variable
   	const c = 1 
	var k = V{2}
   	fmt.Printf("a = %d, b = %d, c = %d, d = %d\n",a,b,c,k.d);
}
```
## 19. Dangling Else

The dangling else is a problem in computer programming in which an optional else clause in an if then statement results in nested conditionals being ambiguous. 

By design, Go doesn't have dangling else problem as it requires block (wrapped in curly brackets). Each "if", "for", and "switch" statement is considered to be in its own implicit block

#### Reference 
1. [Dangling else](https://en.wikipedia.org/wiki/Dangling_else)
2. [Dangling else in Go](https://kuree.gitbooks.io/the-go-programming-language-report/content/17/text.html)

## 20. Garbage Collecting
Golang has garbage collector. Go takes a more traditional approach by addressing object lifetimes through garbage collection, and garbage collection alone. The current implementation is a mark-and-sweep collector. If the machine is a multiprocessor, the collector runs on a separate CPU core in parallel with the main program. 

If we allocate a very large amount of memory say a long list. It will triger garbage collection. For example:
```go
package main

import (
   "fmt"
)

func printStats(mem runtime.MemStats) {
   for i := 0; i < 10; i++ {
         s := make([]byte, 50000000)
         if s == nil {
                fmt.Println("Operation failed!")
         }
   }
}
```

#### Reference
1. [Why do garbage collection? Won't it be too expensive? ](https://golang.org/doc/faq#garbage_collection)

## 21. Short Circuit Evaluation

Golang uses the short circuit evaluation. As we can see from the program below. If go uses short circuit evaluation, then it will only evaluates `val1()` in line1 and `val2()` in line 2.

```go
package main

import "fmt"

func val1() bool{
    fmt.Println("val1 gets called")
    return true
}

func val2() bool{
    fmt.Println("val2 gets called")
    return false
}

func main(){
    if val1() || val2(){ // line 1
        fmt.Println("Or Test")
    }
    if val2() && val1() { // line 2
        fmt.Println("And Test")
    }
}
```

After run the program, we obtain the following prints, which matches our expectation to short circuit evaluation. Otherwise, the program should print `val1 gets called` and `val2 gets called` twice for each sentence.

```
val1 gets called
Or Test
val2 gets called
```

## 22. Error Handling
 Go code uses `error` values to indicate an abnormal state. The `error` type is an interface type. An error variable represents any value that can describe itself as a string. A simple example of error handling when opening a file in golang is 
```go
package main

import (
   "fmt"
)

func printStats(mem runtime.MemStats) {
    f, err := os.Open("filename.ext")
    if err != nil {
        log.Fatal(err)
    }
}
```
We can use to `errorString` create a user defined error in golang:
```go
    type errorString struct {
        s string
    }

    func (e *errorString) Error() string {
        return e.s
    }
```

#### Reference
1. [Error handling and Go](https://blog.golang.org/error-handling-and-go)

## 23. Deep/Shallow Copy
In golang, assignment is a shallow copy. We can use fucntion `func copy(dst, src []Type) int` to do a deep copy. As demonstrate by the program below:
```go
package main

import "fmt"

func main() {
	//Deep copy
    src := []int{1, 2, 3, 4, 5}
    dst := make([]int, 5)
    copy(dst, src)
    
    fmt.Println("[Deep copy] After copy")
    fmt.Printf("dst: %v\n", dst)
    fmt.Printf("src: %v\n", src)

    // changing dst
    dst[0] = 10
    fmt.Println("[Deep copy] After changing dst")
    fmt.Printf("dst: %v\n", dst)
    fmt.Printf("src: %v\n", src)
    
    //Shallow Copy
    src2 := []int{1, 2, 3, 4, 5}
    dst2 := src2
    
    fmt.Println("[Shallow Copy] After copy")
    fmt.Printf("dst: %v\n", dst2)
    fmt.Printf("src: %v\n", src2)

    // changing dst2
    dst2[0] = 10
    fmt.Println("[Shallow Copy] After changing dst")
    fmt.Printf("dst: %v\n", dst2)
    fmt.Printf("src: %v\n", src2)

}
```
This program prints the following:
```
[Deep copy] After copy
dst: [1 2 3 4 5]
src: [1 2 3 4 5]
[Deep copy] After changing dst
dst: [10 2 3 4 5]
src: [1 2 3 4 5]
[Shallow Copy] After copy
dst: [1 2 3 4 5]
src: [1 2 3 4 5]
[Shallow Copy] After changing dst
dst: [10 2 3 4 5]
src: [10 2 3 4 5]
```
As we can observe from the printed result, when we use `copy()` function, after we change the value in `dst`, the `src` doesn't change at all. However, when we use assignment, after we change the value in `dst`, the `src` changes at the same place as well. Thus, we could conclude that `copy()` function in golang execute a deep copy, while assignment in golang execute a shallow copy.
#### Reference
1. [Copy function in Go (Golang)](https://golangbyexample.com/copy-function-in-golang/)