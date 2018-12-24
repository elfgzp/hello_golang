package main

import "fmt"

func main() {
	a := 3
	if a > 0 {
		fmt.Println("a > 0")
	} else {
		fmt.Println("a > 0")
	}

	if a > 1 {
		fmt.Println("a > 1")
		if a < 4 {
			fmt.Println("a < 4")
		}
	}

	b := 3
	switch b {
	case 1:
		fmt.Println("b = 1")
	case 2:
		fmt.Println("b = 2")
	default:
		fmt.Println("b unknown")
	}

	var c interface{}
	c = "c"
	switch c. (type) {
	case int:
		fmt.Println("c type is int")
	case string:
		fmt.Println("c type is string")
	default:
		fmt.Println("c type unknown")
	}
}
