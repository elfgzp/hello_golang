package main

import (
	"fmt"
	"time"
)

func main() {
	go fmt.Println("Go!")
	time.Sleep(100 * time.Millisecond)
}
