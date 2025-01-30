package main

import "fmt"

func main() {

	//Declaracion de constantes con tipo de dato
	const coordenada_x int = 0
	const coordenada_y float64 = 15

	//Declaracion de constantes sin tipo de dato

	const x = 10
	const y = 10

	fmt.Println("Hola mundo!")
	fmt.Println("Coordenada en x:", coordenada_x)
	fmt.Println("Coordenada en y:", coordenada_y)
	fmt.Println("Coordenada en y:", coordenada_y)
	fmt.Println("Valor de x:", x)
	fmt.Println("Valor de y:", y)

	//Declaracion de variables enteras

	base := 12
	var altura int = 12
	var area int

	fmt.Println(base, altura, area)

	// Valores por defecto

	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c, d)
}
