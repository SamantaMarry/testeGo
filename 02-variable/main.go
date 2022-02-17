package main

import "fmt"

func main() {

	var name string

	name = "Samanta"

	fmt.Println(name)


	const name2 = "Samanta"

	fmt.Println(name2)

	a := "Samanta"

	fmt.Println(a)

	var firstName, lastName = "Samanta Marry", "da Silva Britto"

	fmt.Println(firstName, lastName)


	fmt.Println("Olá, meu nome é ", name)
	fmt.Printf("meu nome completo é: %s %s \n ", firstName, lastName) // %d (double) %f (float)
	
}