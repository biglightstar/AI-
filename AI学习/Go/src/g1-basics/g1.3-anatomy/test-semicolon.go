// 小实验：return 后面的换行会发生什么？
package main

import "fmt"

func test() int {
	return
	42
}

func main() {
	fmt.Println(test())
}
