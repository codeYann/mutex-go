package main

import (
	"fmt"
	"sync"
	"time"

	s "github.com/codeYann/mutex-go/station"
)

var mutex sync.Mutex

// Função para executar na primeira thread
func printerService1() {
	john := s.CreateWorker("John Doe", "Backend developer", 190)
	johnPrinter := s.CreatePrinter()
	johnPrinter.GetTasks(john, []string{
		"Fixing database issues",
		"Create new api route",
		"Create new docker container",
		"Fixing cors problem\n",
	})
	mutex.Lock()
	time.Sleep(250 * time.Millisecond)
	johnPrinter.CreateFile("./printer.txt")
	johnPrinter.WriteContent("./printer.txt")
	johnPrinter.ExportHashText("./printer.txt", "output/", "T1")
	mutex.Unlock()
}

// Função para executar na segunda thread
func printerService2() {
	foo := s.CreateWorker("Foo Bar", "Social Midia Manager", 182)
	fooPrinter := s.CreatePrinter()
	fooPrinter.GetTasks(foo, []string{
		"New instagram post",
		"New TikTok post",
		"Create reddit community",
		"Create discord server\n",
	})
	mutex.Lock()
	time.Sleep(250 * time.Millisecond)
	fooPrinter.CreateFile("./printer.txt")
	fooPrinter.WriteContent("./printer.txt")
	fooPrinter.ExportHashText("./printer.txt", "output/", "T2")
	mutex.Unlock()
}

func main() {
	go printerService1() // Executando a primeira thread
	go printerService2() // Executando a segunda thread
	time.Sleep(700 * time.Millisecond)
	fmt.Println("Finishing!")
}
