package main

import (
	"fmt"
)


/* I pull the vars p and q from the user as ints and then pass their memory addresses to this fuction which takes pointer of type int and
   do a modulus on the memory addresses to var z.
*/

func remainder (pPtr *int, qPtr *int){
	var z int
	z = *pPtr % *qPtr
	fmt.Println("The remainder is", z)
}


func main() {

x := "Paul"
var name string
var p int
var q int

fmt.Println("Hello, my name is", x)

fmt.Println("Please enter your name: ")
fmt.Scan(&name)
fmt.Println("Hello,", name)

	fmt.Println("Please enter a small number")
	fmt.Scan(&p)
	fmt.Println("Please enter a large number")
	fmt.Scan(&q)

// Passing the memory values to the function
remainder(&p, &q)


}