package go_goroutine

import (
	"fmt"
	"testing"
	"time"
)

func RunHelloWorld() {
	fmt.Println("Goroutine Hello World")
}

func TestCreateGoroutine(t *testing.T) {
	go RunHelloWorld() // using goroutine with "go" in front of function
	fmt.Println("Uuups")
}

func DisplayNumber(number int) {
	fmt.Println("Display to", number)
}

func TestManyGoroutine(t *testing.T) {
	for i := 0; i < 10000; i++ {
		go DisplayNumber(i)
	}

	time.Sleep(5 * time.Second)
}
