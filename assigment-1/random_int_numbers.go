
// to get Intn Type Random Number 
package main 

import ( 
	"fmt"
	"math/rand"
) 

// Function 
func intnrandom(value_1, value_2 int) int { 
	return value_1 + value_2 + rand.Intn(3) 

} 

// Main function 
func main() { 

	// Getting result from Intnrandom() function 
	res1 := intnrandom(10, 3) 
	res2 := intnrandom(44, 59) 
	res3 := intnrandom(130, 50) 

	// Displaying results 
	fmt.Println("Result 1: ", res1) 
	fmt.Println("Result 2: ", res2) 
	fmt.Println("Result 3: ", res3) 
} 
