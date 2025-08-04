package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go sayHello(&wg)
	wg.Wait()
	sayGoodBye()
}

func sayHello(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Hello, world")
}

func sayGoodBye() {
	fmt.Println("Goodbye!")
}
