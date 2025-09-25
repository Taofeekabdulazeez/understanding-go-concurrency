package main

import (
	"fmt"
	"time"
)

func speak(arg string, ch chan<- string) {
	fmt.Println("Trying to speak....")
	time.Sleep(3 * time.Second)
	ch <- arg
}

func main() {
	ch := make(chan string, 2)
	go speak("Hello", ch)
	data := <-ch
	close(ch)

	fmt.Println("Data received from channel ==> ", data)

}
