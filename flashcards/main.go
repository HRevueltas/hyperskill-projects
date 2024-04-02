package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Card struct {
	Term, Definition string
}

func main() {
	// slice para guardar las cartas
	var card []Card

	// un lector para leer la entrada del usuario
	reader := bufio.NewReader(os.Stdin)

	// solicitar el numero de cartas agregar
	fmt.Println("Ingrese el numero de cartas: ")
	var numCardStr string
	_, err := fmt.Scanln(&numCardStr)
	if err != nil {
		fmt.Println("Error al leer el numero de cartas", err)
		return
	}

	// convertir a entero el valor ingresado
	numCard, err := strconv.Atoi(strings.TrimSpace(numCardStr))
	if err != nil {
		fmt.Println("Error al convertir el número de cartas a entero:", err)
		return
	}

	for i := 0; i < numCard; i++ {

		// Leer término
		fmt.Printf("Termino para la carta #%d: \n", i+1)
		term, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error al leer la entrada del usuario:", err)
			return
		}
		term = strings.TrimSpace(term)

		// Leer definición
		fmt.Printf("Definición para la carta #%d: \n", i+1)
		definition, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error al leer la entrada del usuario:", err)
			return
		}
		definition = strings.TrimSpace(definition)

		card = append(card, Card{term, definition})

	}

	for i := 0; i < len(card); i++ {
		fmt.Printf("Escribe la definicion para el termino '%s': \n", card[i].Term)
		var resp string
		_, err := fmt.Scanln(&resp)
		if err != nil {
			fmt.Println("Error al obtener la respuesta", err)
		}
		if card[i].Definition == resp {
			fmt.Println("Correct!")
		} else {

			fmt.Printf("Incorrecto. La respuesta correcta es: '%s' \n", card[i].Definition)
		}

	}

}
