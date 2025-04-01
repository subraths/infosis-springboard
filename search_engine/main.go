package searchengine

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type Worker struct {
	users []User
	ch    chan *User
}

func NewWorker(users []User, ch chan *User) *Worker {
	return &Worker{users: users, ch: ch}
}

func (w *Worker) Find(email string) {
	for i := range w.users {
		user := &w.users[i]
		if strings.Contains(user.Email, email) {
			w.ch <- user
		}
	}
}

func Main() {
	email := os.Args[1]

	users, err := getData()
	if err != nil {
		fmt.Println(err)
		return
	}

	ch := make(chan *User)

	w := NewWorker(users, ch)

	log.Printf("Looking for %s", email)

	go w.Find(email)
	for {
		select {
		case user := <-ch:
			log.Printf("The email %s is owned by %s", user.Email, user.Name)
		case <-time.After(10 * time.Millisecond):
			return
		}
	}
}
