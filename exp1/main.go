package main

import "fmt"

func main() {
	fmt.Println("Whats this experiment?")
	name, surname := "Anubhav", "Rakshit"
	fmt.Println(name, surname)
	someFloat := 3.14
	fmt.Printf("%T %T\n", name, surname)
	fmt.Printf("%T\n", someFloat)
}
