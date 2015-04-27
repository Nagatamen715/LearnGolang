
// Single line comment
/* Multi-
   line comment*/

// A package clause starts every source file.
// Main is a special name declaring an executable rather than a library.
package main

// Import declaration declares library packages referenced in this file.
import (
	"fmt"		// A package in the Go standard library.
	"io/ioutil"	// Implements some I/O utility functions.
	m "math"	// Math library with local alias m.
	"net/http"	// Yes, a web server!
	"strconv"	// String conversions.
)

// A function definition. Main is special. It is the entry point for the 
// executable program. Love it or hate it, Go uses brace brackets.
func main() {
	// Println outputs a line to stdout.
	// Qualify it with the package name, fmt.
	fmt.Println("Hello world!")
	
	// Call another function within this package.
	beyondHello()
}

// Functions have parameters in parentheses.
// If there are no parameters, empty parentheses are still required.
func beyondHello() {
	var x int 	// Variable declaration. Variables must be declared before use.
	x = 3		// Variable assignment.
	// "Short" declarations use := to infer the type, declare and assign.
	y := 4
	sum, prod := learnMultiple(x, y)			// Function returns two values.
	fmt.Println("sum:", sum, "prod:", prod)	// Simple output.
	learnTypes()
}

// Functions can have parameters and (multiple!) return values.
func learnMultiple(x, y int) (sum, prod int) {
	return x + y, x * y	// Return two values.
}

// Some built-in types and literals.
func learnTypes() {
	// Short declaration usually gives you want you want.
	str := "Learn Go!"	// string type.
	
	s2 := `A "raw" string 
literal can include line breaks.`	// Same string type.
	
	// Non-ASCII literal. Go source is UTF-8.
	g := 'Σ'	// rune type, an alias for int32, holds a unicode code point.
	
	f := 3.14195	// float64, an IEEE-754 64-bit floating point number.
	c := 3 + 4i		// complex128, represented internally with two float64's.
	
	// Var syntax with an initializers.
	var u uint = 7	// Unsigned, but implementation dependent size as with int.
	var pi float32 = 22. / 7
	
	// Conversion syntax with a short declaration.
	n := byte('\n')	// byte is an alias for uint8.
	
	// Arrays have size fixed at compile time.
	var a4 [4]int			// An array of 4 ints, initialized to all 0.
	a3 := [...]int{3, 1, 5}	// An array initialized with a fixed size of three
	// elements, with values 3, 1, and 5.
	
	// Slices have dynamic size. Arrays and slices each have advantages
	// but use cases for slices are much more common. 
	s3 := []int{4, 5, 9}	// Compare to a3. No ellipsis here.
	s4 := make([]int, 4)	// Allocates slice of 4 ints, initialized to all 0.
	var d2 [][]float64		// Declarations only, nothing allocated here.
	bs := []byte("a slice")	// Type conversion syntax.
	
	// Because they are dynamic, slices can be appended to on-demand.
	// To append elements to a slice, built-in append() function is used.
	// First argument is a slice to which we are appending. Commonly,
	// the array variable is updated in place, as in example below.
	s := []int{1, 2, 3}		// Result is a slice of length 3.
	s = append(s, 4, 5, 6)	// Added 3 elements. Slice now has length of 6.
	fmt.Println(s)	// Updated slice is now [1 2 3 4 5 6]
	// To append another slice, instead of list of atomic elements we can
	// pass a reference to a slice or a slice literal like this, with a 
	// trailing elipsis, meaning take a slice and unpack its elements,
	//appending them to slice s.
	s = append(s, []int{7, 8, 9}...)	// Second argument is a slice literal.
	fmt.Println(s)		// Updated slice is now [1 2 3 4 5 6 7 8 9]
	
	p, q := learnMemory()	// Declares p, q to be type pointed to int.
	fmt.Println(*p, *q)		// * follows a pointer. This prints two ints.
	
	// Maps are a dynamically growable associative array type, like the
	// hash or dictionary types of some other languages.
	m := map[string]int{"three":3, "four":4}
	m["one"] = 14195
	
	// Unused variables are an error in Go.
	// The underbar lets you "use" a variable but discard its value.
	_, _, _, _, _, _, _, _, _, _ = str, s2, g, f, u, pi, n, a3, s4, bs
	
	// Output of course counts as using a variable.
	fmt.Println(s, c, a4, s3, d2, m)
	
	learnFlowControl() // Back in the flow
}
