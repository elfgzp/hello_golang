package main

import "fmt"

func main() {
	a := 2
	b := 3
	fmt.Println(a)
	fmt.Println(b)
	a ++
	b --
	fmt.Println(a)
	fmt.Println(b)

	fmt.Println(a >= b)
	fmt.Println(a == b)
	fmt.Println(a != b)

	fmt.Println(a != b && a >= b)
	fmt.Println(a == b || a >= b)

	c := byte(0)
	d := byte(1)

	fmt.Println(c & d)
	fmt.Println(c | d)
	fmt.Println(c ^ d)
	fmt.Println(d>>1)
	fmt.Println(d<<1)

	a += b
	fmt.Println(a)
}
