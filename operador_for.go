package main

import "fmt"

func main() {
	fmt.Println("Números divisíveis por 3 entre 1 e 100:")
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Println(i)
		}
	}
}

// Resolvendo Problemas Numéricos com Go - parte 01
/*
Criar um código que exiba todos os números entre 1 e 100 que são divisíveis por 3.
*/
