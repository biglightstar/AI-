// G1.2 Go 环境搭建与工具链
//
// 本文件讲解 Go 开发环境的核心概念：
//   GOROOT / GOPATH / go modules / go run / go build / go install / go fmt
//
// 你的环境：
//   GOROOT      = F:\GO          （Go 安装目录，编译器和标准库在这）
//   GOPATH      = C:\Users\admin\go （旧式工作区，go modules 时代已弱化）
//   GOMODCACHE  = C:\Users\admin\go\pkg\mod （下载的第三方模块缓存在这）
//
// 核心理解：
//   GOROOT = "Go 自己住的地方"（不要动它）
//   GOPATH（旧时代）= "你的 Go 代码放哪"（需要严格按 src/bin/pkg 目录结构）
//   go modules（新时代）= 代码任意放，go.mod 声明模块，不需要 GOPATH 约束

package main

import (
	"fmt"
	"os"
	"runtime"

	"g1.2-go-toolchain/helper"
)

// ============================================================
// 读取环境变量：Go 程序里可以读 GOROOT、GOPATH 等
// ============================================================
func showGoEnv() {
	fmt.Println("=== Go 运行时环境 ===")
	fmt.Printf("操作系统：%s\n", runtime.GOOS)       // windows
	fmt.Printf("CPU 架构：%s\n", runtime.GOARCH)   // amd64
	fmt.Printf("Go 版本：%s\n", runtime.Version()) // go1.26.5
	fmt.Printf("CPU 核数：%d\n", runtime.NumCPU()) // 影响 GOMAXPROCS

	// 也可以用 os.Getenv 读取环境变量
	goroot := os.Getenv("GOROOT")
	gopath := os.Getenv("GOPATH")
	fmt.Printf("\nGOROOT（安装目录）= %s\n", goroot)
	fmt.Printf("GOPATH（工作区）= %s\n", gopath)
}

// ============================================================
// go.mod 的作用：声明模块，管理依赖
// ============================================================
// 当前目录下的 go.mod 声明了本模块：
//
//   module g1.2-go-toolchain
//   go 1.26.5
//
// module 行 = 本模块的"身份证号"，别的模块通过这个名字导入你
// go 行 = 指定 Go 版本

// ============================================================
// 演示：你可以创建多个包，通过 import 组织代码
// ============================================================
// 不用头文件，不用前向声明。import 自动帮你找到依赖。
// 这里演示一个简单的多文件项目：main.go 调用 helper 包的函数

func main() {
	fmt.Println("╔══════════════════════════════════════════════╗")
	fmt.Println("║   G1.2 Go 环境搭建与工具链                    ║")
	fmt.Println("╚══════════════════════════════════════════════╝")

	// 1. 查看环境
	showGoEnv()

	// 2. 演示导包——调用 helper 包
	//    （对比 C++：你不需要 #include "helper.h"，
	//     Go 编译器自动解析 go.mod + import）
	fmt.Println("\n=== 调用 helper 包 ===")
	fmt.Println(helper.Greet("Go 学习者"))

	// 3. 演示 go fmt 的效果
	//    在命令行运行：go fmt ./...
	//    Go 会自动格式化所有代码——统一风格，没有争论空间
	demoGoFmt()

	fmt.Println("\n=== 常用命令速查 ===")
	fmt.Println("  go run main.go        → 编译+运行（适合开发调试）")
	fmt.Println("  go build              → 编译成 .exe 文件（适合发布）")
	fmt.Println("  go build -o myapp.exe → 指定输出文件名")
	fmt.Println("  go install            → 编译并放到 GOBIN 目录")
	fmt.Println("  go fmt ./...          → 格式化所有 Go 代码")
	fmt.Println("  go mod tidy           → 清理 go.mod，添加缺失/删除多余依赖")
	fmt.Println("  go mod init <name>    → 初始化新模块")
	fmt.Println("  go env                → 查看所有 Go 环境变量")
	fmt.Println("  go version            → 查看 Go 版本")
	fmt.Println("  go help               → 查看所有命令")
	fmt.Println("  go help <command>     → 查看某个命令的详细帮助")
}

// go fmt 会把这个函数格式化成这样——缩进统一，空格统一
func demoGoFmt() {
	// 试试故意把这段代码写乱，然后跑 go fmt，看它怎么自动修
	x := 1 + 2*3
	fmt.Printf("没有括号歧义：1 + 2*3 = %d（go fmt 会保留运算符优先级）\n", x)
	fmt.Println("试着：go fmt main.go → 格式化当前文件")
	fmt.Println("试着：go fmt ./...   → 格式化整个项目")
}
