# Go 语言学习进度追踪

**学习启动日期**：2026-07-13
**目标完成时间**：2026-08-31（7 周）
**学习方向**：Go 语言系统深入学习
**学习状态**：进行中
**学习节奏**：系统推进，每 2-3 天一个话题，代码练习驱动

本文档追踪所有 Go 语言学习进度，包括：
- 各领域话题掌握情况
- 知识盲区识别
- 优先学习计划

---

## 快速统计

- **总体进度**：2/52 话题掌握 = **4%**
- **上次学习**：2026-07-14（G1.2 环境搭建与工具链，掌握程度：高）
- **学习方向**：Go 基础 → 控制流与函数 → 方法与接口 → 错误处理 → 包与模块 → **并发编程（核心）** → 标准库 → 测试 → Web

---

## 领域进度汇总

| 领域 | 权重 | 话题数 | 掌握 | 状态 | 优先级 |
|------|------|--------|------|------|--------|
| **G1. Go 基础入门** | 10% | 5 | 2 | 学习中 | 高 |
| **G2. 数据类型与复合类型** | 15% | 8 | 0 | 未开始 | 高 |
| **G3. 控制流与函数** | 12% | 5 | 0 | 未开始 | 高 |
| **G4. 方法与接口** | 15% | 5 | 0 | 未开始 | 中 |
| **G5. 错误处理** | 8% | 3 | 0 | 未开始 | 中 |
| **G6. 包与模块系统** | 8% | 5 | 0 | 未开始 | 中 |
| **G7. 并发编程** ⭐ | 16% | 6 | 0 | 未开始 | 高 |
| **G8. 标准库核心** | 8% | 6 | 0 | 未开始 | 中 |
| **G9. 测试与工具链** | 5% | 5 | 0 | 未开始 | 低 |
| **G10. Web 后端实战** | 3% | 4 | 0 | 未开始 | 低 |

---

## G1. Go 基础入门（10%，5 话题）

> 🎯 **目标**：理解 Go 的设计哲学，能独立搭建和运行 Go 项目

### 未学习（3/5）

- [x] **G1.1** Go 语言概述与设计哲学（诞生背景/核心设计理念/Go vs Python vs JS/适用场景）✅
- [x] **G1.2** 环境搭建与工具链（安装/GOROOT/GOPATH/go mod/go run/build/install/go fmt）✅
- [ ] **G1.3** 第一个 Go 程序深度解剖（package main/import/func main/分号自动插入/编译流程）
- [ ] **G1.4** Go 基本语法速览（注释/命名规范/导出规则—大写public小写private/语句分隔）
- [ ] **G1.5** Go REPL 与交互式学习工具（gore/yaegi/go play ground/go 命令交互模式/与 Python REPL 对比）

---

## G2. 数据类型与复合类型（15%，8 话题）

> 🎯 **目标**：掌握 Go 的类型体系，理解值类型/引用类型的本质区别

### 未学习（0/8）

- [ ] **G2.1** 基本数据类型（int 系列/float/bool/string/零值哲学/byte=uint8/rune=int32）
- [ ] **G2.2** 变量声明与类型推断（var/:=/多变量声明/未使用变量报错/空白标识符 `_`）
- [ ] **G2.3** 常量与 iota（const 关键字/无类型常量/iota 枚举模式/新块重置/`_` 跳过）
- [ ] **G2.4** 操作符（算术/比较/逻辑/位运算符/赋值运算/地址操作符 & */`++` 是语句非表达式）
- [ ] **G2.5** 类型转换（Type(val) 语法/strconv 包/截断不四舍五入/类型断言 vs 类型转换）
- [ ] **G2.6** 复合类型 — 数组与切片（数组值类型 vs 切片引用类型/len vs cap/append 扩容/截取共享底层/copy 独立拷贝）
- [ ] **G2.7** 复合类型 — Map 与 Struct（map 创建/逗号ok模式/delete/遍历随机序/struct 定义与零值/Map+Struct 组合）
- [ ] **G2.8** 指针（*T/&/值传递 vs 指针传递/无指针运算/理解为"传地址而非传拷贝"）

---

## G3. 控制流与函数（12%，5 话题）

> 🎯 **目标**：掌握 Go 的控制结构和函数特性，特别是 Go 独有的设计

### 未学习（0/5）

- [ ] **G3.1** 条件与循环（if—无括号必须花括号/可带初始化语句/for—Go唯一循环关键字/range 遍历/break/continue/嵌套循环/没有while）
- [ ] **G3.2** Switch（自动break/fallthrough/无表达式switch）
- [ ] **G3.3** 函数 — 基础/多返回值/命名返回值/可变参数/一等公民/闭包/递归
- [ ] **G3.4** 作用域（包级 vs 函数级 vs 块级/Shadowing 变量遮蔽/:= 块级陷阱—同块复用新块遮蔽/与 Python LEGB JS var/let 对比）
- [ ] **G3.5** Defer（LIFO 后进先出/参数注册时求值/defer+命名返回值改返回值/trace 工厂/循环陷阱）

---

## G4. 方法与接口（15%，5 话题）

> 🎯 **目标**：理解 Go 的"面向对象"哲学——组合优于继承，接口隐式实现

### 未学习（0/5）

- [ ] **G4.1** Struct 进阶（嵌套与提升字段/有名嵌套 vs 匿名嵌套/struct tag—json omitempty "-"/空 struct{} 零内存做 Set 和信号通道）
- [ ] **G4.2** 方法（值接收者副本改不动 vs 指针接收者改原值/Go 自动值↔指针转换/字面量不可寻址不能调指针方法）
- [ ] **G4.3** 接口 Interface（行为抽象/隐式实现无需 implements/多态/空接口 any/类型断言 switch v.(type)/接口由消费者定义非生产者）
- [ ] **G4.4** 组合 vs 继承（struct 嵌入实现复用/接口组合/FileLogger ≠ Logger 类型—组合不产生类型等同/与 Python 多继承 JS 原型链对比）
- [ ] **G4.5** 泛型 Generics（类型参数 [T Constraint]/any comparable 内置约束/自定义 Ordered 约束/泛型 Stack[T]/类型推断/与 TS <T> 对比）

---

## G5. 错误处理（8%，3 话题）

> 🎯 **目标**：掌握 Go 显式错误处理模式，理解 panic/recover 的边界

### 未学习（0/3）

- [ ] **G5.1** Error 接口（error 内置接口/errors.New/fmt.Errorf/错误哨兵/错误包装与解包—%w/errors.Is/As）
- [ ] **G5.2** Panic 与 Recover（panic 严重错误/recover 只能在 defer 中/return 三步走—赋值→defer→返回/Go 哲学—能用 error 就不要 panic/与 try/except 对比）
- [ ] **G5.3** 错误处理最佳实践（不要忽略错误/错误包装保留上下文/自定义错误类型/fail fast/哨兵+errors.Is 分类处理）

---

## G6. 包与模块系统（8%，5 话题）

> 🎯 **目标**：理解 Go 的包管理和导入机制，掌握模块化开发

### 未学习（0/5）

- [ ] **G6.1** Package（大写导出/小写私有/internal 包限制/package main 特殊性/跨包导入/函数 vs 方法设计决策）
- [ ] **G6.2** Go Modules（go mod init/go.mod 结构/go get/go mod tidy/语义化版本/go.sum/replace 替换/本地模块引用/私有仓库配置—GOPRIVATE）
- [ ] **G6.3** 导入机制（标准库/第三方/别名/匿名导入副作用/点导入/循环导入禁止）
- [ ] **G6.4** 资源嵌入 Embed（//go:embed 指令/embed.FS/编译进二进制/单文件部署）
- [ ] **G6.5** 工作区模式 go.work（多模块协同/monorepo 场景/use 指令）

---

## G7. 并发编程 ⭐（16%，6 话题）

> 🎯 **目标**：掌握 Go 最核心的竞争力——goroutine 和 channel

### 未学习（0/6）

- [ ] **G7.1** Goroutine（go 关键字/协程 vs OS 线程/M:N 调度/GOMAXPROCS/与 Python asyncio JS Promise 对比）
- [ ] **G7.2** Channel（无缓冲—同步通信/有缓冲—异步队列/发送接收/关闭 channel/通信共享内存哲学）
- [ ] **G7.3** Select（多路复用/default 非阻塞/time.After 超时/与 asyncio.wait 对比）
- [ ] **G7.4** 同步原语（sync.Mutex/RWMutex/sync.WaitGroup/sync.Once/sync/atomic 原子操作）
- [ ] **G7.5** 并发模式（Pipeline 管道/Fan-out Fan-in/Worker Pool 工作池/errgroup）
- [ ] **G7.6** Context 深入（4个方法—Deadline/Done/Err/Value/4种创建方式/取消传播/ctx.Value 边界/传入 struct 禁忌）

---

## G8. 标准库核心（8%，6 话题）

> 🎯 **目标**：熟悉 Go 最常用的标准库

### 未学习（0/6）

- [ ] **G8.1** 输入输出 fmt/io/os/bufio（fmt 格式化动词/io.Reader io.Writer 接口核心/os 文件操作/bufio 缓冲读写/流式处理）
- [ ] **G8.2** 编码 encoding/json（Marshal/Unmarshal/struct tag 控制/自定义序列化/流式 Encoder Decoder）
- [ ] **G8.3** 网络编程 net（net/http HTTP 服务—HandleFunc/Handler/客户端 / TCP 通信 / UDP 通信 / 建立连接/数据读写/关闭连接）
- [ ] **G8.4** 时间与字符串（time.Time/Duration/strings 包/strconv 类型转换/regexp 正则）
- [ ] **G8.5** 反射 reflect（TypeOf/ValueOf/Kind vs Type/通过反射修改值 CanSet/struct tag 解析）
- [ ] **G8.6** 结构化日志 log/slog（Go 1.21+/JSON Handler/结构化字段/级别/与 Python structlog 对比）

---

## G9. 测试与工具链（5%，5 话题）

> 🎯 **目标**：写出有测试的 Go 代码，会用 Go 工具链

### 未学习（0/5）

- [ ] **G9.1** 单元测试（_test.go 命名约定/TestXxx 函数签名/表格驱动测试/t.Run 子测试/Mock 技巧—接口mock/gomock/monkey patch）
- [ ] **G9.2** 基准测试与示例测试（BenchmarkXxx/ExampleXxx 自动验证输出/go test -bench/go test -cover）
- [ ] **G9.3** 核心工具链（go fmt/go vet/go doc/go mod tidy/gopls 语言服务器）
- [ ] **G9.4** 性能分析 pprof（CPU profile/Memory heap profile/Goroutine profile/go tool pprof/火焰图）
- [ ] **G9.5** 代码生成 go generate（//go:generate 指令/stringer 枚举/mockgen/工具链哲学）

---

## G10. Web 后端实战（3%，4 话题）

> 🎯 **目标**：用 Go 构建一个完整的 RESTful API 服务

### 未学习（0/4）

- [ ] **G10.1** 标准库 HTTP 服务（net/http REST API/中间件模式链式包装/自定义 Handler）
- [ ] **G10.2** Gin 框架入门（Gin vs 标准库 vs FastAPI/路由路径参数查询参数/中间件 Logger Recovery CORS/请求绑定验证）
- [ ] **G10.3** 模板引擎 html/template（Parse/Execute/自动 HTML 转义防 XSS/嵌套模板自定义函数/与 Jinja2 对比）
- [ ] **G10.4** 完整 API 项目（cmd/internal/pkg/api 项目结构/Viper 配置管理/集成 GORM or sqlx/构建 Python 项目的 Go 版本）

---

## 已掌握话题

| 话题 | 掌握日期 | 掌握程度 | 关键要点 |
|------|----------|----------|----------|
| G1.1 Go 概述与设计哲学 | 2026-07-14 | 高 | Go 从 C++ 编译慢痛点诞生；9 大设计哲学（编译快/25关键字/大小写导出/错误是值/struct+组合/goroutine/channel/GC/零值）；深刻理解"通过通信共享内存"vs"共享内存加锁" |
| G1.2 环境搭建与工具链 | 2026-07-14 | 高 | GOROOT/GOPATH/GOMODCACHE 三环境变量；go run vs go build 区别（调试 vs 发布）；go fmt 自动格式化；go mod init/tidy；多包项目结构（main + helper）；无头文件自动依赖解析 |

---

## 知识盲区

### 高优先级

| 盲区 | 严重程度 | 说明 |
|------|----------|------|
| G1-G10 全部（除G1.1/G1.2外） | 高 | 50 个话题待掌握 |

### 中优先级

| 盲区 | 严重程度 | 说明 |
|------|----------|------|
| — | — | 暂无 |

### 低优先级

| 盲区 | 严重程度 | 说明 |
|------|----------|------|
| — | — | 暂无 |

### 最近解决

| 话题 | 解决日期 | 说明 |
|------|----------|------|
| — | — | 暂无 |

---

## 学习计划

### 当前学习状态

**学生背景**：
- 已掌握：HTML 基础
- Python：之前学过但现在已经忘完了，学习 Go 的过程中如需对比会尽量用简单的 Python 例子
- 薄弱环节：后端开发、数据库、版本控制、AI 集成
- 定位：编程初学者，从 Go 开始系统学习第一门后端语言

### 七周学习路径（2026-07-13 → 2026-08-31）

**第一阶段：Go 语言基础（第 1 周，7/13 - 7/19）**
1. G1.1-G1.5 Go 基础入门含 REPL（5 话题）
2. G2.1-G2.5 数据类型/变量/常量/操作符/类型转换（5 话题）
3. G3.1-G3.2 控制流（2 话题）

**第二阶段：Go 核心机制（第 2 周，7/20 - 7/26）**
4. G2.6-G2.8 数组切片/Map/指针（3 话题）
5. G3.3-G3.5 函数含递归/作用域/Defer（3 话题）
6. G4.1-G4.3 Struct进阶/方法/接口（3 话题）

**第三阶段：接口/错误/模块（第 3 周，7/27 - 8/2）**
7. G4.4-G4.5 组合vs继承/泛型（2 话题）
8. G5.1-G5.3 错误处理（3 话题）
9. G6.1-G6.3 Package/Modules/导入机制（3 话题）

**第四阶段：并发基础与深度（第 4 周，8/3 - 8/9）**
10. G6.4-G6.5 Embed/go.work（2 话题）
11. G7.1-G7.3 Goroutine/Channel/Select（3 话题）

**第五阶段：并发深度与标准库（第 5 周，8/10 - 8/16）**
12. G7.4-G7.6 同步原语/并发模式/Context（3 话题）
13. G8.1-G8.4 fmt+io+os/json/net/http/time+strings（4 话题）

**第六阶段：标准库进阶与工程化（第 6 周，8/17 - 8/23）**
14. G8.5-G8.6 反射/slog（2 话题）
15. G9.1-G9.5 测试/benchmark/pprof/go generate（5 话题）

**第七阶段：Web 实战（第 7 周，8/24 - 8/31）**
16. G10.1-G10.4 Web 后端实战（4 话题）

### 下次会话建议

**当前话题**：✅ G1.2 环境搭建与工具链 — 已完成
**下次话题**：**G1.3 第一个 Go 程序深度解剖** — package main、import、func main、分号自动插入、编译流程

**教学要求**：
- 从头开始，扎实打好基础
- 一个小知识点一个小知识点的讲
- 一文件夹一知识点的渐进式推进
- 每个文件夹 = 完整 main.go + go.mod（非 TODO 练习）

---

## 学习原则

- **不只学语法，要学设计哲学。** Go 的简洁是有代价的——理解它为什么这样取舍。
- **不只写 Demo，要做对比。** 每个 Go 概念都要和 Python/JS 对应概念对比，加深理解。
- **不只背规则，要写代码验证。** Go 编译器的严格性是最好老师——让代码跑起来证明你对了。
- **不只学标准库 API，要理解接口设计。** io.Reader/Writer 是 Go 设计的精髓。

---

## 仓库结构

```
/Go/                           # Go 代码根目录
  /src/
    /g1-basics/                # G1 基础入门练习
    /g2-types/                 # G2 数据类型练习
    /g3-control-flow/          # G3 控制流练习
    /g4-methods-interfaces/    # G4 方法与接口练习
    /g5-errors/                # G5 错误处理练习
    /g6-packages/              # G6 包与模块练习
    /g7-concurrency/           # G7 并发编程练习
    /g8-stdlib/                # G8 标准库练习
    /g9-testing/               # G9 测试练习
    /g10-web/                  # G10 Web 实战
/sessions/                     # 学习会话记录（跨 Python/Go 共用）
/progress/
  go-study-tracker.md          # ← 本文件：Go 学习唯一追踪文件
  study-tracker.md             # Python 全栈学习追踪文件
```
