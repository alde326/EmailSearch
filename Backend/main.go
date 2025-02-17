package main

import (
	"fmt"
	"time"

	"github.com/alde326/EmailSearch/Backend/constants"
	helpers "github.com/alde326/EmailSearch/Backend/helpers"
	//indexer "github.com/alde326/EmailSearch/Backend/indexer"
)

func main() {
	// Marcar el tiempo de inicio
	start := time.Now()

	// Ejecutar las funciones
	//indexer.FetchCreateZincIndex()
	helpers.ProcessFiles(constants.FILE_NAME)

	// Calcular el tiempo transcurrido
	elapsed := time.Since(start)

	// Imprimir el tiempo de ejecución
	fmt.Printf("El tiempo de ejecución fue: %s\n", elapsed)
}
