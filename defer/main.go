package main

import "fmt"

func a() {
	i := 0
	defer fmt.Println(i)
	i++
	return
}

func b() {
	i := 0
	defer fmt.Println(i) //输出0，因为i此时就是0
	i++
	defer fmt.Println(i) //输出1，因为i此时就是1
	return
}

func c() (i int) {
	defer func() { i++ }()
	return 1
}

func main() {
	a()
	b()
	c()
}
