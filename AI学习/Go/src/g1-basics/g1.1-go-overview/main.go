// G1.1 Go 语言概述与设计哲学
//
// 本文件展示了 Go 的核心设计哲学，每个函数对应一个关键概念。
// 对照你的 C++ 经验，体会 Go 的"少即是多"。

package main

import (
	"fmt"
	"time"
)

// ============================================================
// 哲学 1：编译必须快
// ============================================================
// 你不需要写头文件 (.h)，不用写前向声明。
// Go 编译器自己解析依赖关系。
// 试试：go run main.go  — 秒级编译

// ============================================================
// 哲学 2：25 个关键字，语法极简
// ============================================================
// Go 总共只有 25 个关键字。对比 C++ 有 90+ 个。
// 下面是 Go 全部关键字，一眼就能看完：
//
// break    default     func   interface  select
// case     defer       go     map        struct
// chan     else        goto   package    switch
// const    fallthrough if     range      type
// continue for         import return     var
//
// 没有 class、没有 try/catch、没有 public/private/protected
// 没有 virtual、没有 explicit、没有 constexpr...

// ============================================================
// 哲学 3：大写 = public，小写 = private（导出规则）
// ============================================================
// 你不需要写 public: / private:，Go 用首字母大小写决定可见性：
//
//   HelloWorld()  → 别的包可以调用（导出）
//   helloWorld()  → 只能在当前包内调用（私有）
//
// 就是这么简单，没有 protected，没有 friend。

// 导出的函数：别的包可用
func ShowCase() {
	fmt.Println("我是导出的函数，首字母大写！")
}

// 私有的函数：只有本包可用
func internalHelper() {
	fmt.Println("我是私有函数，首字母小写！")
}

// ============================================================
// 哲学 4：错误是值，不是异常（error handling）
// ============================================================
// C++ 有 try/catch/throw，Go 说：错误就是一个普通的返回值。
// 这让你必须显式处理每一个错误，不能"忘了"。

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("除数不能为零：%d/%d", a, b)
	}
	return a / b, nil // nil 表示没有错误，类似 nullptr 但更安全
}

func demoErrorHandling() {
	fmt.Println("\n=== 哲学 4：错误是值 ===")

	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("出错：", err)
	} else {
		fmt.Printf("10 / 2 = %d\n", result)
	}

	result, err = divide(10, 0)
	if err != nil {
		fmt.Println("出错：", err) // 不会 crash，你显式处理了错误
	} else {
		fmt.Printf("10 / 0 = %d\n", result)
	}
}

// ============================================================
// 哲学 5：没有类，用 struct + 方法（组合优于继承）
// ============================================================
// Go 没有 class 关键字。数据用 struct，行为用"方法"挂在 struct 上。
// 没有继承，想复用代码？把 struct 嵌套进去（组合）。

type Animal struct {
	Name string
}

func (a Animal) Speak() string {
	return a.Name + " 发出了声音"
}

// Go 用组合代替继承：Dog 可以嵌入 Animal
type Dog struct {
	Animal             // 嵌入（组合）：Dog 自动拥有 Animal 的所有方法和字段
	Breed  string
}

// Dog 可以覆盖（override）父 struct 的方法
func (d Dog) Speak() string {
	return d.Name + "（" + d.Breed + "）汪汪叫！"
}

func demoStructAndMethod() {
	fmt.Println("\n=== 哲学 5：struct + 方法，不用 class ===")

	a := Animal{Name: "某动物"}
	fmt.Println(a.Speak())

	d := Dog{
		Animal: Animal{Name: "旺财"},
		Breed:  "金毛",
	}
	fmt.Println(d.Speak()) // 调用 Dog 自己的 Speak
	// d.Name 可以直接访问，因为 Dog 嵌入了 Animal
	fmt.Printf("狗的名字：%s，品种：%s\n", d.Name, d.Breed)
}

// ============================================================
// 哲学 6：用 goroutine 做并发（"go" 一个关键字就是线程）
// ============================================================
// C++ 里你写 std::thread，Go 里你写 go。
// goroutine 不是 OS 线程——它更轻（~2KB 栈），一个程序可以跑百万个。

func demoGoroutine() {
	fmt.Println("\n=== 哲学 6：goroutine 并发 ===")

	// go 关键字启动一个 goroutine
	go func() {
		for i := 1; i <= 3; i++ {
			fmt.Printf("  goroutine 里：%d\n", i)
			time.Sleep(50 * time.Millisecond)
		}
	}()

	// 主 goroutine 也做点事
	for i := 1; i <= 3; i++ {
		fmt.Printf("  main 里：%d\n", i)
		time.Sleep(50 * time.Millisecond)
	}

	// 等一下让 goroutine 跑完（后面会学用 sync.WaitGroup 更优雅）
	time.Sleep(100 * time.Millisecond)
}

// ============================================================
// 哲学 7：通过通信来共享内存（channel）
// ============================================================
// C++ 用 mutex + 共享变量，Go 用 channel 传递数据。
// 数据不是"放在那里抢"，而是"发给你"。

func sum(s []int, ch chan int) {
	total := 0
	for _, v := range s {
		total += v
	}
	ch <- total // 把结果发送到 channel，而不是写入共享变量
}

func demoChannel() {
	fmt.Println("\n=== 哲学 7：通过通信共享内存 ===")

	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// 创建 channel：数据从这里流动
	ch := make(chan int)

	// 两个 goroutine 各自计算一半，把结果发送到 channel
	go sum(nums[:len(nums)/2], ch) // 前一半：1+2+3+4+5 = 15
	go sum(nums[len(nums)/2:], ch) // 后一半：6+7+8+9+10 = 40

	// 从 channel 接收：不用锁，没有 data race
	x := <-ch
	y := <-ch

	fmt.Printf("前一半和：%d，后一半和：%d，总和：%d\n", x, y, x+y)
	// 整个过程没有 mutex.lock，没有 mutex.unlock！
}

// ============================================================
// 哲学 8：自动垃圾回收（GC）—— 不用 delete/free
// ============================================================
// C++ 你要管理内存：new/delete, malloc/free, RAII, smart pointers
// Go 有垃圾回收器，你只管创建，不用管释放。

func demoGC() {
	fmt.Println("\n=== 哲学 8：自动 GC ===")

	// 你可以 new，但更常用字面量
	p := new(int)
	*p = 42
	fmt.Printf("new 出来的指针指向的值：%d\n", *p)

	// 用完了不用 delete——GC 会在没人用的时候自动回收
	// 没有悬空指针（dangling pointer），没有 double free

	// 更常见的写法：用 & 取地址
	x := 100
	ptr := &x
	fmt.Printf("&x 取地址：%d\n", *ptr)
	// ptr 指向的 x 在栈上还是堆上？Go 编译器自动逃逸分析，你不用管
}

// ============================================================
// 哲学 9：零值初始化 —— 每个类型都有默认值
// ============================================================
// C++ 里未初始化的变量是"未定义行为"的经典来源。
// Go 里每个变量都有"零值"：int=0, string="", bool=false, 指针=nil

func demoZeroValue() {
	fmt.Println("\n=== 哲学 9：零值初始化 ===")

	var i int        // 0（不是未定义！）
	var s string      // ""（不是乱码！）
	var b bool        // false
	var p *int        // nil（不是野指针！）

	fmt.Printf("int 零值=%d, string 零值=%q, bool 零值=%v, 指针零值=%v\n",
		i, s, b, p)
}

// ============================================================
// main 入口 —— 运行所有演示
// ============================================================
func main() {
	fmt.Println("╔══════════════════════════════════════════════╗")
	fmt.Println("║   G1.1 Go 语言概述与设计哲学                 ║")
	fmt.Println("║   主题：C++ 开发者的 Go 初体验               ║")
	fmt.Println("╚══════════════════════════════════════════════╝")

	// 演示导出规则
	ShowCase()
	internalHelper()

	// 演示各设计哲学
	demoErrorHandling()
	demoStructAndMethod()
	demoGoroutine()
	demoChannel()
	demoGC()
	demoZeroValue()

	fmt.Println("\n=== 总结 ===")
	fmt.Println("Go 从 C++ 的痛点中诞生，设计哲学是：")
	fmt.Println("  1. 编译要快（没有头文件）")
	fmt.Println("  2. 语法要少（25 个关键字）")
	fmt.Println("  3. 导出靠大小写（不用 public/private）")
	fmt.Println("  4. 错误是值（不用 try/catch）")
	fmt.Println("  5. struct + 组合（不用 class + 继承）")
	fmt.Println("  6. goroutine 轻量并发（go 关键字）")
	fmt.Println("  7. channel 传递数据（不抢锁）")
	fmt.Println("  8. 自动 GC（不用 delete/free）")
	fmt.Println("  9. 零值安全（没有未定义行为）")
}