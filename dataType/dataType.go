package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

// 字符串编码统一为 UTF-8

// 默认类型零值 int->0 string->"" bool->false
// 类型也可以设置别名

func main() {
	var i8 uint8 = 1
	var i32 uint32 = 1
	var i int = 1
	var ui uint = 1
	// int 根据操作系统的位数
	fmt.Println(unsafe.Sizeof(i8))
	fmt.Println(unsafe.Sizeof(i32))
	fmt.Println(unsafe.Sizeof(i))
	fmt.Println(unsafe.Sizeof(ui))

	var f32 float32 = 1.000
	var f64 float64 = 1.000
	fmt.Println(unsafe.Sizeof(f32))
	fmt.Println(unsafe.Sizeof(f64))

	var bl bool = true
	fmt.Println(unsafe.Sizeof(bl))

	var bs byte = 1
	fmt.Println(unsafe.Sizeof(bs))

	var r rune = 1
	fmt.Println(unsafe.Sizeof(r))

	var di int32
	var df float32
	var dbl bool
	var dc complex64 // 复数
	var str string
	fmt.Println(di)
	fmt.Println(df)
	fmt.Println(dbl)
	fmt.Println(dc)
	fmt.Println(str)

	// 类型别名
	type 整型 int

	var it 整型
	fmt.Println(it)
	fmt.Println(reflect.TypeOf(it))
}
