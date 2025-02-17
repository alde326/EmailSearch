package helpers

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/alde326/EmailSearch/Backend/constants"
	indexer "github.com/alde326/EmailSearch/Backend/indexer"
	"github.com/alde326/EmailSearch/Backend/models"
)

var (
	allEmails []models.Email
	mu        sync.Mutex
)

const maxWorkers = constants.MAX_WORKERS

func ProcessFiles(fileNames string) {
	var wg sync.WaitGroup
	semaphore := make(chan struct{}, maxWorkers)
	start := time.Now()

	wg.Add(1)
	go ExploreFolder(fileNames, &wg, semaphore) //se pone el go

	wg.Wait()

	// Guardar emails restantes en el JSON final
	if len(allEmails) > 0 {
		writeBatchToFile(allEmails)
	}

	fmt.Println("Tiempo de inicio:", start)
	fmt.Println("Tiempo de finalización:", time.Since(start))
}

// 1. se procesan los archivos de la base de datos
func ExploreFolder(folderPath string, wg *sync.WaitGroup, semaphore chan struct{}) {
	defer wg.Done()

	files, err := os.ReadDir(folderPath)
	if err != nil {
		fmt.Printf("Error leyendo carpeta %s: %s\n", folderPath, err)
		return
	}

	for _, file := range files {
		filePath := filepath.Join(folderPath, file.Name())
		if file.IsDir() {
			wg.Add(1) //se descomenta
			//Se llama recursivamente a la función ExploreFolder
			ExploreFolder(filePath, wg, semaphore)
		} else {
			wg.Add(1)
			//Se llama a la función readFileLineByLine
			go readFileLineByLine(filePath, wg, semaphore)
		}
	}
}

// 2. se procesa el archivo encontrado de la base de datos
func readFileLineByLine(filePath string, wg *sync.WaitGroup, semaphore chan struct{}) {
	defer wg.Done()
	semaphore <- struct{}{}
	defer func() { <-semaphore }()

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error abriendo archivo %s: %s\n", filePath, err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	buf := make([]byte, 0, 64*1024)
	scanner.Buffer(buf, 1024*1024)
	email := models.Email{}
	hasSubEmails := false

	for scanner.Scan() {
		line := scanner.Text()
		ParseLineMessage(line, &email, &hasSubEmails)
	}
	// Se crean los subemails
	createSubEmails(&email)
	// Se añade el email a la lista de emails
	appendEmail(&email)

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error leyendo archivo %s: %s\n", filePath, err)
	}
}

// Optimización: escribir en lotes en lugar de escribir un email a la vez
func writeBatchToFile(emails []models.Email) {
	file, err := os.Create("emails.json") // Sobrescribe el archivo en cada ejecución
	if err != nil {
		fmt.Printf("Error abriendo el archivo: %s\n", err)
		return
	}
	defer file.Close()

	// Estructura final del JSON
	finalData := map[string]interface{}{
		"index":   "index", // Reemplázalo con el nombre de tu índice
		"records": emails,
	}

	jsonData, err := json.MarshalIndent(finalData, "", "  ") // Formato bonito
	if err != nil {
		fmt.Printf("Error serializando a JSON: %s\n", err)
		return
	}

	_, err = file.Write(jsonData)
	if err != nil {
		fmt.Printf("Error escribiendo en el archivo: %s\n", err)
		return
	}
}

func appendEmail(email *models.Email) {
	mu.Lock()
	defer mu.Unlock()

	allEmails = append(allEmails, *email)

	if len(allEmails) >= constants.EMAILS_AMOUNT {
		writeBatchToFile(allEmails) // Escribe en lote cuando se alcance el límite
		url := "./emails.json"
		indexer.BulkIndexToZingsearch(url)
		//fmt.Println("Emails indexados exitosamente!!")
		allEmails = nil // Limpia la lista después de escribir
	}
}

func RemoveExtraSpaces(s string) string {
	return strings.Join(strings.Fields(s), " ")
}

func parseDateLines(line string, email *models.Email) {
	newLine := RemoveExtraSpaces(strings.TrimPrefix(line, "Date:"))
	firstParsedLine, firstErr := time.Parse(constants.REFERENCE_FIRST_DATE, newLine)
	if firstErr != nil {
		secondParsedLine, secondErr := time.Parse(constants.REFERENCE_SECOND_DATE, newLine)
		if secondErr != nil {
			email.DateSubEmail = newLine
			return
		}
		email.Date = secondParsedLine
		return
	}
	email.Date = firstParsedLine
}

// 3. Se parsean los archivos
func ParseLineMessage(line string, email *models.Email, hasSubEmails *bool) {
	switch {
	case strings.Contains(line, "======================================="):
		email.Body += ""
	case strings.Contains(line, "Message-ID:"):
		email.MessageID = RemoveExtraSpaces(strings.TrimPrefix(line, "Message-ID:"))
	case strings.Contains(line, "Sent:") && !*hasSubEmails:
		email.Sent = RemoveExtraSpaces(strings.TrimPrefix(line, "Sent:"))
	case strings.Contains(line, "From:") && !*hasSubEmails:
		email.From = RemoveExtraSpaces(strings.TrimPrefix(line, "From:"))
	case strings.Contains(line, "To:") && !*hasSubEmails:
		email.To += RemoveExtraSpaces(strings.TrimPrefix(line, "To:")) + " "
	case strings.Contains(line, "Subject:") && !*hasSubEmails:
		email.Subject = RemoveExtraSpaces(strings.TrimPrefix(line, "Subject:"))
	case strings.Contains(line, "Date:") && !*hasSubEmails:
		parseDateLines(line, email)
	case strings.Contains(line, "-----Original Message-----"):
		*hasSubEmails = true
		email.Body += line + "\n"
	default:
		email.Body += line + "\n"
	}
}

func generateSubEmailId(subEmailIndex int, emailId string) string {
	newId := fmt.Sprintf(".JavaSubEmail.%d>", subEmailIndex)
	return strings.Replace(emailId, ">", newId, 1)
}

// 4. Se crean los subemails
func createSubEmails(email *models.Email) {
	if !strings.Contains(email.Body, "-----Original Message-----") {
		return
	}
	subEmails := strings.Split(email.Body, "-----Original Message-----")

	for subEmailIndex := 1; subEmailIndex < len(subEmails); subEmailIndex++ {
		newEmail := models.Email{}
		hasSubEmails := false

		for _, line := range strings.Split(subEmails[subEmailIndex], "\n") {
			ParseLineMessage(line, &newEmail, &hasSubEmails)
		}
		newEmail.MessageID = generateSubEmailId(subEmailIndex, email.MessageID)
		appendEmail(&newEmail)
	}
	email.Body = subEmails[0]
}
