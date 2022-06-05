// Nome: Yan Rodrigues da Silva
// Matrícula: 495532package station

package station

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Printer struct {
	worker   *Worker
	todoList []string
}

func CreatePrinter() *Printer {
	return &Printer{
		worker:   nil,
		todoList: nil,
	}
}

func (printer *Printer) GetTasks(worker *Worker, todoList []string) {
	printer.worker = worker
	printer.todoList = todoList
}

func (printer Printer) CreateFile(path string) {
	file, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()
}

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

func (p *Printer) ExportHashText(path, hashCode string) {
	header := sha1.New()
	header.Write([]byte(hashCode))
	hash := hex.EncodeToString(header.Sum(nil))
	file, err := os.Create(path + hash + ".txt")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintln(file, "Worker: "+p.worker.name)
	fmt.Fprintln(file, "Id: "+strconv.Itoa(p.worker.id))
	fmt.Fprintln(file, "Todo list: ")

	for _, todo := range p.todoList {
		fmt.Fprintln(file, todo)
	}
}
