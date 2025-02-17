package indexer

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/alde326/EmailSearch/Backend/constants"
)

// Función para preparar la solicitud _bulk para Zingsearch (Elasticsearch)
func BulkIndexToZingsearch(filepath string) error {
	// Username y password para el servidor ZincSearch
	username := constants.USER_NAME
	password := constants.PASSWORD

	// Leer el archivo JSON directamente
	jsonData, err := os.ReadFile(filepath)
	if err != nil {
		return fmt.Errorf("Error leyendo el archivo JSON: %v", err)
	}

	// Crear la solicitud HTTP POST
	req, err := http.NewRequest("POST", "http://localhost:4080/api/_bulkv2", bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("Error creando la solicitud HTTP: %v", err)
	}

	// Configurar encabezados y autenticación
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(username, password)

	// Enviar la solicitud
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("Error enviando la solicitud HTTP: %v", err)
	}
	defer resp.Body.Close()

	// Verificar la respuesta del servidor
	if resp.StatusCode != 200 {
		body, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("Error en la respuesta del servidor: %s", string(body))
	}

	fmt.Println("Emails indexados exitosamente.")
	return nil
}
