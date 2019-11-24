package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Whats this experiment?")
	name, surname := "Anubhav", "Rakshit"
	age := 37.3
	fmt.Println(name, surname, math.Floor(age))
	someFloat := 3.14
	fmt.Printf("%T %T\n", name, surname)
	fmt.Printf("%T\n", someFloat)
}
