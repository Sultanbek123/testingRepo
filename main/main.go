package main

import (
	"fmt"
	"testingRepo/internal/entities"
	"time"
)

func firstAction(ch chan string) {
	ch <- "message"
	fmt.Println("Message has been sent")
}
func secondAction(ch chan string) {
	time.Sleep(5 * time.Second)
	message := <-ch
	fmt.Println("received message :", message)
}

func main() {
	var message string = "Hi my name is Sultanbek"
	fmt.Println(message)
	ch := make(chan string)
	go firstAction(ch)
	go secondAction(ch)
	user := entities.NewUser("Sultanbek", 24, 5345.789)
	fmt.Println(user)
	time.Sleep(10 * time.Second)
}
