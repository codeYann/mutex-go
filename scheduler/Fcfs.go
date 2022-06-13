package scheduler

import (
	"sync"

	"github.com/codeYann/mutex-go/users"
)

// Define Fcfs struct to schedule users and its files
type Fcfs struct {
	// Define a list of users
	users []*users.User
	// Define a mutex to sync users
	mutex sync.Mutex
}

// Create a new scheduler
func CreateScheduler() *Fcfs {
	return &Fcfs{
		users: make([]*users.User, 0),
	}
}

// Enqueue a user to the list of users
func (f *Fcfs) Enqueue(user *users.User) {
	// Lock the mutex
	f.mutex.Lock()
	f.users = append(f.users, user)
	f.mutex.Unlock()
}

// Dequeue a user from the list of users
func (f *Fcfs) Dequeue() *users.User {
	// Lock the mutex
	f.mutex.Lock()
	user := f.users[0]
	// Remove the first user from the list
	f.users = f.users[1:]
	f.mutex.Unlock()
	return user
}

// Get the number of users in the list of users
func (f *Fcfs) GetSize() int {
	return len(f.users)
}

// Get front user from the list of users
func (f *Fcfs) GetFront() *users.User {
	return f.users[0]
}

// Get back user from the list of users
func (f *Fcfs) GetBack() *users.User {
	return f.users[len(f.users)-1]
}
