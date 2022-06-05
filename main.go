// Nome: Yan Rodrigues da Silva
// Matrícula: 495532

package main

import (
	"fmt"

	"github.com/codeYann/mutex-go/printers"
	"github.com/codeYann/mutex-go/scheduler"
	"github.com/codeYann/mutex-go/users"
)

func main() {
	// Create a scheduler
	scheduler := scheduler.CreateScheduler()
	// Create a printer
	printer := printers.CreatePrinter(scheduler)
	// Create a user
	user := users.CreateUser("John Doe", "admin")
	user.SetContent("Hello World, esse é o texto que estou criando para testar o metodo Print()")
	// Enqueue the user to the scheduler
	scheduler.Enqueue(user)
	// Print the scheduler
	fmt.Println(scheduler)
	// Print the printer
	fmt.Println(printer)
	// Print the user
	fmt.Println(user)

	printer.Print()

}
