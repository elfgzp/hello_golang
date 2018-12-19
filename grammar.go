// 程序所属包, 必须在首行
package hello_golang

// 依赖包
import "fmt"

// 单行注释

/*
多行注释
 */

const HELLOGO string = "hello go" // 常量定义
var helloGo string = "hello go"   // 变量定义

type goInt int // 一般类型声明

type GoStruct struct { // 声明结构体
}

type IGo interface {  // 声明接口
}


func goFunc()  {  // 函数定义
	fmt.Println("goFunc")
}

func main() {
	fmt.Println(HELLOGO)
	fmt.Println(helloGo)
	goFunc()
}
