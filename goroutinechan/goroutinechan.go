package main
import (
	"fmt"
	"time"
)
func say(s string, done chan string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
	done <- "Terminei 1"
}

func say2(s string, done chan string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
	done <- "Terminei 2"
}

func main() {
	done := make(chan string)
	go say("Hello", done)
	go say2("world", done)
	fmt.Println(<-done)
}