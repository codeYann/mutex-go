package main

import (
	"fmt"
	"sync"
	"time"

	s "github.com/codeYann/mutex-go/station"
)

var mutex sync.Mutex

func printerDaemon1(worker *s.Worker, list []string, path string) {
	fmt.Println("Printing data from thread 1")
	mutex.Lock()
	ptPrinter := s.CreatePrinter(worker, list)
	time.Sleep(250 * time.Millisecond)
	ptPrinter.Impress(path, "thread1")
	mutex.Unlock()
}

func printerDaemon2(worker *s.Worker, list []string, path string) {
	fmt.Println("Printing data from thread 2")
	mutex.Lock()
	ptPrinter := s.CreatePrinter(worker, list)
	time.Sleep(250 * time.Millisecond)
	ptPrinter.Impress(path, "thread2")
	mutex.Unlock()
}

func main() {
	go printerDaemon1(
		s.CreateWorker("John Doe", "Backend developer", 120),
		[]string{"Create database", "Fixing rest api", "Deploy new application\n"},
		"./output/printer.txt",
	)

	go printerDaemon2(
		s.CreateWorker("Foo bar", "Community manager", 88),
		[]string{"Create new reddit post", "Update trello tasks", "Create new discord server\n"},
		"./output/printer.txt",
	)

	time.Sleep(800 * time.Millisecond)
	fmt.Printf("Finishing")
}
