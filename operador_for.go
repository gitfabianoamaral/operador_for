package main

import "fmt"

func main() {
	fmt.Println("Números divisíveis por 3 entre 1 e 100:")
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Println(i)
		}
	}

	fmt.Println("Número múltiplo de 3 = 'Pin'")
	fmt.Println("Número múltiplo de 5 = 'Pan'")
	fmt.Println("Número múltiplo de 3 e 5 = 'PinPan'")
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("PinPan")
		} else if i%3 == 0 {
			fmt.Println("Pin")
		} else if i%5 == 0 {
			fmt.Println("Pan")
		} else {
			fmt.Println(i)
		}
	}
}

// Resolvendo Problemas Numéricos com Go - parte 01
/*Criar um código que exiba todos os números entre 1 e 100 que são divisíveis por 3.*/

// Resolvendo Problemas Numéricos com Go - parte 02
/*Criar um código que deve imprimir números de 1 a 100.
No lugar de número multiplo de 3 imprimir "Pin",
no lugar de múltiplo de 5 imprimir "Pan",
e caso seja divisivel por ambos, imprimir "PinPan".*/
