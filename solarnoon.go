package main

// From the go tour package.go file for the discussion in file and program structure

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("My favorite number is", rand.Intn(10))
	
	calculate()
}

func calculate() {
	fmt.Println("Calculating noon")
}
