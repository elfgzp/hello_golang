package main

import (
	"fmt"
	"reflect"
)

func main() {
	var (
		i int
		j string
	)
	i = 1
	j = "j"

	var a, b, c int = 1, 2, 3
	d, e, f := 1, 2, 3 // 只能用在函数体内

	var _ = 4
	g := float64(a)
	//h := string("12")
	//i := int(h) 无法转换

	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(a, b, c)
	fmt.Println(d, e, f)
	//fmt.Println(_) // 下划线为垃圾桶
	fmt.Println(reflect.TypeOf(g))
}
