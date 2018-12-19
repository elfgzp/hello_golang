package _package

// package是最基本的分发单位和工程管理中依赖关系的体现
// 每个Go语言源代码开头都必须要有package的声明，表示源码文件所属的代码包
// 要生成Go语言可执行程序，必须要有main的package包，而且必须该包下有main
// 同一个路径下只能存在一个package，一个package可以拆成多个源文件组成