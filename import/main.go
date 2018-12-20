package _import

import (
	"fmt"
	pack1 "hello_golang/import/p1"
	_ "hello_golang/import/p2"
)

// import 语句类似python中的import
// 这里不同的是import不能导入源码中没有用到的包，否则Go语言编译器会直接报错
/* import 第一种方式：
import "a"
import "b"
import "c"
 */
/* import 第二种方式：
import (
	"a"
    "b"
	"c"
)
 */

// import 原理
// 如果一个main导入其他包，那么这些包将被顺序引入
// 如果导入依赖包B，会先导入包B，然后初始化B中的常量变量，最后如果B中有init()函数，将会执行
// 所有包导入完成后，才会执行main中的初始化，然后执行main中的init(),最后在执行main()
// 如果一个包被导入多次，实际只会导入一次

// import 别名
// 将import包可以起名为别的名字
// 点(.)标示导入包后，调用该包的函数可以不用包名前缀
// 下划线(_) 导入后，导入的包只执行init()

func init() {
	fmt.Println("import run main init()")
}

func ImportPrint() {
	fmt.Println("import run main main()")
	fmt.Println(pack1.P1)
}