package main

import (
	"fmt"
)

func main() {
	for i := 1; i < 10; i++ {
		fmt.Println("loop", i)
	}

	a := []string{"a", "b", "c"}
	for key, value := range a {
		fmt.Println(key, value)
	}

	goto One
	fmt.Println("skip")
One:
	fmt.Println("This is One")

	for {
		fmt.Println("loop and break")
		break
	}

	for i := 1; i < 3; i++ {
		if i > 1 {
			continue
		}
		fmt.Println("loop continue")
	}

}
