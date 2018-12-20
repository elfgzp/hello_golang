package p1

import (
	"fmt"
	"hello_golang/import/p11"
)

func init() {
	fmt.Println("import run p1 init()")
	fmt.Println(p11.P11)
}

const P1 string = "p1"
