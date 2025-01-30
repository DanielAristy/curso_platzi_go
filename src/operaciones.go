package main

import "fmt"

func main() {
	var num1 int = 20
	var num2 int = 10

	suma := suma(num1, num2)
	fmt.Printf("La suma entre %d y %d es: %d\n", num1, num2, suma)

	resta := resta(num1, num2)
	fmt.Printf("La resta entre %d y %d es: %d\n", num1, num2, resta)

	multilicacion := multilicacion(num1, num2)
	fmt.Printf("La multilicacion entre %d y %d es: %d\n", num1, num2, multilicacion)

	division := division(num1, num2)
	fmt.Printf("La division entre %d y %d es: %d\n", num1, num2, division)

	modulo := modulo(num1, num2)
	fmt.Printf("EL modulo entre %d y %d es: %d\n", num1, num2, modulo)
}

func suma(num1, num2 int) int {
	return num1 + num2
}

func resta(num1, num2 int) int {
	return num1 - num2
}

func multilicacion(num1, num2 int) int {
	return num1 * num2
}

func division(num1, num2 int) int {
	return num1 / num2
}

func modulo(num1, num2 int) int {
	return num1 % num2
}
