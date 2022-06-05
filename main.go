// Nome: Yan Rodrigues da Silva
// Matrícula: 495532

package main

import (
	"fmt"
	"sync"
	"time"

	s "github.com/codeYann/mutex-go/station"
)

var mutex sync.Mutex

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
	johnPrinter.CreateFile("output/printer.txt")
	johnPrinter.WriteContent("output/printer.txt")
	johnPrinter.ExportHashText("output/", "T1")
	mutex.Unlock()
}

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
	fooPrinter.CreateFile("output/printer.txt")
	fooPrinter.WriteContent("output/printer.txt")
	fooPrinter.ExportHashText("output/", "T2")
	mutex.Unlock()
}

func main() {
	go printerService1()
	go printerService2()
	time.Sleep(700 * time.Millisecond)
	fmt.Println("Finishing!")
}
