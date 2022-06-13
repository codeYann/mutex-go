package printers

import (
	"crypto/sha1"
	"encoding/base64"
	"os"
	"sync"

	"github.com/codeYann/mutex-go/scheduler"
	"github.com/codeYann/mutex-go/users"
)

// Define printer struct to simulate a real world printer
// printer has a scheduler, isBusy variable and a mutex to sync
type Printer struct {
	scheduler *scheduler.Fcfs
	isBusy    bool
	mtx       sync.Mutex
}

// Create a printer with a scheduler
func CreatePrinter(scheduler *scheduler.Fcfs) *Printer {
	return &Printer{
		scheduler: scheduler,
		isBusy:    false,
	}
}

// Create a method to print a file
func (p *Printer) Print() {
	// If the printer is not busy
	if !p.isBusy {
		// Lock the mutex
		p.mtx.Lock()
		// Get the front user from the scheduler
		user := p.scheduler.GetFront()
		p.isBusy = true

		if user.GetContent() == "" {
			p.isBusy = false
			p.mtx.Unlock()
			panic("User " + user.GetName() + " has no content to print")
		}

		// Check if user content isn't empty
		// Create a file name
		path := createOutPutFile(user.GetName())
		// Write header in file
		writeHeader(path, user)
		// Write content in the file
		writeContent(path, user.GetContent())

		// Remove the front user from the scheduler
		p.scheduler.Dequeue()

	}

}

// Todo: Create minors methods to print a file
// It should export a txt file with the content of the printer in output folder
// this exported file has to have user name, user id and user role

// Create a func that return hash of user name
func hashNameFile(userName string) string {
	hasher := sha1.New()
	hasher.Write([]byte(userName))
	sha := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	return sha
}

// Create file name
func createOutPutFile(userName string) string {
	// Create a file with hash of user name in output folder
	path := "output/" + hashNameFile(userName) + ".txt"
	_, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	return path
}

// Create a func that write content in a file
func writeContent(path string, content string) {
	file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	_, err = file.WriteString(content)
	if err != nil {
		panic(err)
	}
}

func writeHeader(path string, user *users.User) {
	file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	_, err = file.WriteString("User: " + user.GetName() + "\n")
	if err != nil {
		panic(err)
	}
}
