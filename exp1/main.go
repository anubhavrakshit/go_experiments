package main

import (
	"fmt"
	"math"
)

func addNums(num1 int, num2 int) int {
	return num1 + num2
}
func main() {
	fmt.Println("Whats this experiment?")
	name, surname := "Anubhav", "Rakshit"
	age := 37.3
	// var bookArr [2]string
	// bookArr[0] = "Book1"
	// bookArr[1] = "Book2"
	bookArr := [4]string{"Book 1", "Book 2"}
	fmt.Println(bookArr[0], bookArr[1])
	fmt.Println(name, surname, math.Floor(age))
	someFloat := 3.14
	fmt.Printf("%T %T\n", name, surname)
	fmt.Printf("%T\n", someFloat)
	fmt.Println(addNums(10, 5))
}
