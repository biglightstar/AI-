// Package helper 演示多文件项目中的另一个包
//
// 这和 main.go 在同一个模块下，但属于不同的 package。
// Go 通过目录名（不是文件名）来识别包：
//
//	main.go       → package main（程序入口）
//	helper/*.go   → package helper（被 main 导入）
//
// 对比 C++：
//
//	C++：main.cpp 需要 #include "helper.h"，编译时链接 helper.o → 生成一个可执行文件
//	Go：  main.go import "模块名/helper"，go build 自动编译所有依赖 → 生成一个可执行文件
//	Go 不需要 .h 文件、不需要前向声明、不需要手动链接！
package helper

import "fmt"

// Greet 是导出的函数（大写开头），别的包可以调用
func Greet(name string) string {
	return fmt.Sprintf("你好，%s！这是来自 helper 包的问候。", name)
}

// internalNote 是私有的（小写开头），只有本包内部能用
func internalNote() string {
	return "这个函数只能在 helper 包内部调用"
}
