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
	worker        *Worker
	processesList []string
}

// func createFile(path string) {
// 	file, err := os.Create(path)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	file.Close()
// }

func writeContent(w *Worker, processes []string, path string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintln(file, "Worker: "+w.name)
	fmt.Fprintln(file, "Id: "+strconv.Itoa(w.id))
	fmt.Fprintln(file, "Todo list: ")

	for _, process := range processes {
		fmt.Fprintln(file, process)
	}
}

func (p *Printer) exportHashText(seed string) {
	header := sha1.New()
	header.Write([]byte(seed))
	hash := hex.EncodeToString(header.Sum(nil))
	file, err := os.Create("./output/" + hash)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintln(file, "Worker: "+p.worker.name)
	fmt.Fprintln(file, "Id: "+strconv.Itoa(p.worker.id))
	fmt.Fprintln(file, "Todo list: ")

	for _, process := range p.processesList {
		fmt.Fprintln(file, process)
	}
}

// func deleteFile(path string) {
// 	err := os.Remove(path)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

func CreatePrinter(worker *Worker, processesList []string) *Printer {
	return &Printer{
		worker,
		processesList,
	}
}

func (p *Printer) Impress(path string, seed string) {
	if p == nil {
		return
	}
	// createFile(path)
	writeContent(p.worker, p.processesList, path)
	p.exportHashText(seed)
	// deleteFile(path)
}
