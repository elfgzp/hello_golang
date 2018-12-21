package main

import "fmt"

const a string = "a"
const b = "b"

const (
	c string = "c"
	d        = "d"
)

const e, f = "e", "f" //常量定义了可以不一定要使用，但是变量一定要使用

func fc(s string) string {
	return ""
}

// iota在const关键字出现将会被重置为0
// const中每新增一行常量声明将使得iota计数一次

const g  = iota
const h  = iota
const (
	i = iota
	j = iota
	k = iota
	_ = iota
	_ = iota
	l = iota
	m = iota * 2
	n = iota * 2
	o
	p
)
const _  = iota
const (
	q, r = iota, iota + 3
	s, t
	u = iota
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(h)
	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(k)
	fmt.Println(l)
	fmt.Println(m)
	fmt.Println(n)
	fmt.Println(o)
	fmt.Println(p)
	fmt.Println(q, r, s, t, u)
}
