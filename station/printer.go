package station

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

// Estrutura para a criação de uma impressora
type Printer struct {
	worker   *Worker
	todoList []string
}

// Função para criar uma impressora
func CreatePrinter() *Printer {
	return &Printer{
		worker:   nil,
		todoList: nil,
	}
}

// Função para pegar a lista de tarefas de um funcionário.
func (printer *Printer) GetTasks(worker *Worker, todoList []string) {
	printer.worker = worker
	printer.todoList = todoList
}

// Criando o arquivo da impressora principal
func (printer Printer) CreateFile(path string) {
	file, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()
}

// Escrevendo os dados de um funcionário na impressora principal
func (printer *Printer) WriteContent(path string) {
	file, err := os.OpenFile(
		path,
		os.O_APPEND|os.O_WRONLY,
		0644,
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintln(file, "Worker: "+printer.worker.name)
	fmt.Fprintln(file, "Id: "+strconv.Itoa(printer.worker.id))
	fmt.Fprintln(file, "Todo list: ")

	for _, todo := range printer.todoList {
		fmt.Fprintln(file, todo)
	}
	file.Close()
}

// Exportando esses dados que passaram na impressora principal (Metafora para o papel)
func (p *Printer) ExportHashText(source, desnity, hashCode string) {
	header := sha1.New()
	header.Write([]byte(hashCode))
	hash := hex.EncodeToString(header.Sum(nil))
	newFile, err := os.Create(desnity + hash + ".txt")

	if err != nil {
		log.Fatal(err)
	}

	sourceFile, err := os.Open(source)
	if err != nil {
		log.Fatal(err)
	}

	io.Copy(newFile, sourceFile)
}
