# Golang 一本通

整理自：[小满zs](https://go-docs.pages.dev)

[toc] 

# 1 Golang 基础入门

Go（又称 Golang）是**Google**在 2007 年推出、2009 年正式开源的静态编译型编程语言，由

- 肯·汤普森 (Ken Thompson)：图灵奖得主，是C语言的前身B语言的作者、UNIX操作系统的发明人之一。他还参与了Plan 9操作系统和UTF-8编码的设计
- 罗伯·派克 (Rob Pike)：曾是贝尔实验室UNIX开发团队成员，也是Plan 9操作系统和Inferno操作系统开发的主要领导人。他与肯·汤普森共同设计了UTF-8编码
- 罗伯特·格瑞史莫 (Robert Griesemer)：曾为谷歌的V8 JavaScript引擎、Chubby以及HotSpot JVM（Java虚拟机）做出重要贡献

我们带入神话故事来理解Go语言：

1. Unix —— 鸿钧老祖（定系统天道规则）
2. C 语言 —— 造化玉碟（鸿钧核心基础功法，万物底层根基）
3. BSD / System V / Linux —— 三清三大圣人（鸿钧三大弟子，各自演化道统）
4. C++ —— 太上老君（在基础功法上炼化灵宝，功能最全，复杂大型基建专用）
5. Go —— 鸿钧后人新创简化功法（Ken Thompson 是 Unix/C 元老，简化 C，适配云洪荒）
6. Java —— 西方接引道人（套一层虚拟机结界，跨天地通用，但厚重笨重）
7. Python —— 散修懒人功法（上手简单，封装底层，不用苦修基础内功）
8. Windows —— 域外魔神，自成一方小世界，早年排斥鸿钧道统

![Go 语言特点](https://go-docs.pages.dev/post/1.png)

### Go 语言特点

Go 的设计哲学是「少即是多」—— 砍掉冗余特性，专注工程效率与可维护性。以下是它最核心的几大特点：

1. **原生并发**：基于 goroutine 与 channel 实现 CSP 并发模型，协程初始栈仅数 KB，远比传统 OS 线程轻量，由运行时统一调度，天然适合高并发服务
2. **自动内存管理**：内置 GC 垃圾回收器，无需像 C/C++ 那样手动申请与释放内存，在开发效率与运行安全之间取得平衡
3. **静态编译**：编译产物为单一可执行文件，无外部运行时依赖，拷贝到目标机器即可运行，部署极其简单
4. **跨平台编译**：内置交叉编译能力，在 macOS 上一条 `go build` 命令即可产出 Linux、Windows 等平台的二进制文件
5. **语法简洁**：仅 25 个关键字，用组合替代继承，拒绝过度设计，代码易读易写，学习曲线平缓
6. **编译速度极快**：大型项目完整编译通常在秒级完成，修改代码后几乎即编即跑，开发体验流畅
7. **标准库强大**：`net/http`、`encoding/json`、`context` 等核心库久经生产验证，网络、序列化、并发控制等常见需求开箱即用
8. **工具链一体化**：`go mod` 管理依赖、`go test` 运行测试、`go fmt` 统一格式，从编码到发布全流程有官方工具支撑

### 如何安装 Go 语言

#### 1 下载安装包

Go 语言的安装包可以在 [Go 语言官网](https://golang.org/dl/) 下载。

![Go 语言特点](https://go-docs.pages.dev/post/2.png)

选择你所对应的操作系统，点击下载安装包。

#### 2 安装 Go 语言

一路下一步，直到安装完成。

![Go 语言特点](https://go-docs.pages.dev/post/3.png)

#### 3 验证安装

在终端输入 `go version` 命令，如果显示 Go 语言的版本信息，则说明安装成功。(目前主包用的是`1.26.4`版本)

![Go 语言特点](https://go-docs.pages.dev/post/4.png)

#### 4 编辑器下载

推荐使用 [VS Code](https://code.visualstudio.com/) 编辑器，因为Go语言官方推荐使用VS Code作为开发工具。下载完成一路下一步，直到安装完成。

#### 5 配置go语言插件

在VS Code中安装Go语言插件，安装完成后，VS Code会自动配置Go语言环境。

![Go 语言特点](https://go-docs.pages.dev/post/5.png)

### 为什么要出这个教程以及为什么要学习 Go 语言

#### 为什么要写这个教程

Go 的中文资料不少，但质量参差不齐——有的太浅，有的直接翻译官方文档，缺少一条清晰的学习路径。写这个系列，是希望给零基础或从其他语言转过来的同学一份**系统化、能跟着动手**的入门指南，少走弯路。

#### 为什么要学 Go

1. **就业市场需求大**：字节、腾讯、美团、百度等大厂的后端、基础架构、中间件岗位大量招聘 Go 开发；云原生、微服务方向几乎绕不开它
2. **薪资有竞争力**：具备 Go + 云原生（K8s、Docker）技能的工程师，在市场上通常比纯 CRUD 后端更有议价空间
3. **入门门槛相对友好**：语法简洁、关键字少，有 Python、Java 或前端 JS 基础的同学，通常一两周就能写出可运行的服务
4. **前端转全栈的务实选择**：纯前端岗位竞争加剧，掌握 Go 后可以独立完成后端 API、中间层逻辑，向全栈或平台工程方向发展
5. **云原生时代的「官方语言」**：Docker、Kubernetes、Prometheus、Etcd 等核心基础设施均用 Go 编写，学 Go 有助于真正理解这些工具的底层原理
6. **区块链 / Web3 的主流技术栈**：以太坊 Go 客户端 Geth、Hyperledger Fabric 等均基于 Go，相关岗位普遍将其列为硬性要求
7. **适用场景广**：微服务 API、命令行工具、爬虫、消息队列、网关、DevOps 脚本——从业务到基建都能覆盖
8. **性能与开发效率的平衡**：编译快、部署简单、并发模型清晰，在需要高并发但又不想陷入 C++ 复杂度的场景里，Go 是性价比很高的选择

如果你属于以下任一情况，都很适合从这里开始：

- 想转行后端或全栈，不知道从哪下手
- 做运维 / 云计算，想读懂 K8s 生态的代码
- 在校或转行，希望掌握一门「好找工作」的实用语言
- 已经会一点编程，想系统补全 Go 的知识体系

# 2 Hello World

我们可以新建一个`hello.go`文件，然后输入以下代码：

```go
package main
import "fmt"
func main() {
	fmt.Println("Hello, World!")
}
```

然后我们打开终端，输入`go run hello.go`命令，运行程序。

```bash
go run hello.go
```

输出> Hello, World!

## package main

```go
package main
```

1. package main = 告诉 Go 编译器，这是程序的入口包（可执行程序）。
2. Go 官方规定，程序入口必须叫 main，名字是固定的，不能修改。

例如我改成 `aaa` 会报错：

```go
package aaa

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}
package command-line-arguments is not a main package # 错误提示
```

## import “fmt”

导入（Import）Go 标准库中的 fmt 包，让当前文件可以使用它提供的功能。

1. fmt：Go 官方提供的格式化输入输出包。
2. Println()：fmt 包中的一个函数，用来打印内容并换行。

## func main()

1. func main()：定义一个名为 main 的函数，这是 Go 程序的入口函数。
2. {}：main 函数的代码块，必须有。
3. fmt.Println(“Hello, World!")：调用 fmt 包中的 Println 函数，打印 “Hello, World!” 字符串。

## 作业

1. 修改 `hello.go` 文件，将 “Hello, World!” 改成 “Hello, Go!"。
2. 运行程序，观察输出结果。

# 3 变量声明

Go 变量标准完整声明格式： `var 变量名 数据类型 = 初始值`

```go
package main

import "fmt"

func main() {
	//关键字 变量名 类型 = 值
	var name string = "小满"
	fmt.Println(name)
}
```

以上便是变量声明的规则，我们声明了一个字符串类型的变量name，并赋值为"小满"，那什么时候该用字符串呢？你在网页中所看到大部分文本和文字都是字符串，比如：“你好，世界！"、“Hello, World!"、“こんにちは、世界！“等。所以，当你需要存储文本、文字或任何字符序列时，就可以使用字符串类型。

### 类型推断

类型推断是Go语言的一种特性，它可以根据变量的值自动推断变量的类型。比如：

```go
package main

import "fmt"

func main() {
	//关键字 变量名 类型 = 值
	var name = "小满"
	var age = 18 //int 表示整数类型
	fmt.Println(name, age) //输出：小满 18
}
```

当你鼠标悬浮在变量名上时，会自动显示变量的类型和值。如上所示，name 的类型为 string，age 的类型为 int。

![Go 语言特点](https://go-docs.pages.dev/post/6.png)

### 自动零值

```go
package main

import "fmt"

func main() {
	var name string
	var age int
	fmt.Println(name, age) //输出："" 0
}
```

我们可以声明变量不赋值，那么变量就会有一个默认值。 输出结果为：”” 0，因为name和age没有赋值，所以默认值为空字符串和0。

### 变量简写

```go
package main

import "fmt"

func main() {
	//变量名 := 值
	name := "小满"
	fmt.Println(name)
}
```

我们可以用 `:=`来简写变量声明，可以省略类型和关键字，但是不能省略变量名和值。

> 注意点：变量简写只能用于函数内部，不能用于全局变量。

![Go 语言特点](https://go-docs.pages.dev/post/7.png)

### 多变量声明

```go
package main

import "fmt"

func main() {
	//多个变量用逗号隔开
	age, name := 18, "小满"
	fmt.Println(age, name) //输出：18 小满
}
```

多个变量声明需要用`,`隔开，并且赋值的时候也是根据顺序赋值。例如 age = 18, name = “小满”。

## 作业

1. 声明一个变量，类型为字符串，值为"小满”。
2. 声明一个变量，类型为整数，值为18。
3. 并且使用多变量声明以及简写的方式，以及输出到终端。

# 4 数据类型

## 字符串

```go
package main

import "fmt"

func main() {
	text := "这是一段文本，苍天已死，黄天当立，岁在甲子，天下大吉。"
    text += "\n这是第二段文本，天地玄黄，宇宙洪荒，日月盈昃，辰宿列张。"
	fmt.Println(text)
}
```

Go 语言中字符串需要通过双引号包裹文本内容，如果有多段文本需要拼接，可以使用加号`+`进行拼接。如果需要换行，可以使用反斜杠`\n`进行换行。

输出内容如下:

```
这是一段文本，苍天已死，黄天当立，岁在甲子，天下大吉。
这是第二段文本，天地玄黄，宇宙洪荒，日月盈昃，辰宿列张。
```

## 数字

```go
package main

import "fmt"

func main() {
	var age int = 18
	fmt.Println(age)
}
```

Go 语言中数字类型有多种，包括整数、浮点数、复数等。

整数包括有符号和无符号两种，有符号包括int8、int16、int32、int64，无符号包括uint8、uint16、uint32、uint64。

> 有符号和无符号的区别是：有符号可以表示正数和负数，无符号只能表示正数。

| 类型  | 占用字节数 | 取值范围                                   |
| ----- | ---------- | ------------------------------------------ |
| int8  | 1 字节     | -128 ~ 127                                 |
| int16 | 2 字节     | -32768 ~ 32767                             |
| int32 | 4 字节     | -2147483648 ~ 2147483647                   |
| int64 | 8 字节     | -9223372036854775808 ~ 9223372036854775807 |
| int   | 自适应     | 32 位系统同 int32，64 位系统同 int64       |

| 类型    | 占用字节数 | 取值范围                               |
| ------- | ---------- | -------------------------------------- |
| uint8   | 1 字节     | 0 ~ 255                                |
| uint16  | 2 字节     | 0 ~ 65535                              |
| uint32  | 4 字节     | 0 ~ 4294967295                         |
| uint64  | 8 字节     | 0 ~ 18446744073709551615               |
| uint    | 自适应     | 32 位系统同 uint32，64 位系统同 uint64 |
| uintptr | 自适应     | 与 uint 范围一致，用于存储内存地址     |

浮点数包括单精度浮点数和双精度浮点数，分别对应float32和float64。

| 类型    | 名称   | 字节 | 有效精度 |
| ------- | ------ | ---- | -------- |
| float32 | 单精度 | 4    | 5~7 位   |
| float64 | 双精度 | 8    | 15~16 位 |

## 布尔

```go
package main

import "fmt"

func main() {
	var isShow bool = true
	fmt.Println(isShow)
}
```

布尔值只有两个，true和false。一般用于条件判断，和状态表示。

例如：开关 如果开启则是(true)，如果关闭则是(false)。

## 数组 和 切片

```go
package main

import "fmt"

func main() {
	//数组
	var arr [5]int // 定义一个长度为5的整型数组 但是未赋值 所以输出为[0 0 0 0 0]
	fmt.Println(arr)
	//----------------------------------------------------------------------
	arr2 := [5]int{1, 2, 3, 4, 5} // 定义一个长度为5的整型数组 并赋值为[1 2 3 4 5]
	fmt.Println(arr2)
	//----------------------------------------------------------------------
	//切片
	arr3 := []int{}                          // 定义一个空切片
	arr3 = append(arr3, 6, 6, 6, 6, 6, 6, 6) // 追加元素
	fmt.Println(arr3)                        // 输出为[6 6 6 6 6 6 6]
	fmt.Println(len(arr3))                   // 获取切片的长度 输出为7
}
```

**数组**是长度固定的连续内存块，定义后长度不可变，未赋值的元素会使用对应类型的零值（整型为 `0`）。

**切片**是对数组的抽象，长度可变，底层仍依赖数组。常用 `append` 追加元素，用 `len` 获取当前长度，用 `cap` 获取底层数组容量。

| 特性     | 数组                   | 切片                   |
| -------- | ---------------------- | ---------------------- |
| 长度     | 固定，定义时确定       | 可变                   |
| 声明     | `[5]int`               | `[]int`                |
| 传参     | 值传递，会复制整个数组 | 引用传递，共享底层数据 |
| 适用场景 | 固定大小、性能敏感     | 日常开发中最常用       |

## 映射

```go
package main

import "fmt"

func main() {
	// 定义一个 map，键为 string，值为 int
	scores := map[string]int{
		"张三": 90,
		"李四": 85,
		"王五": 92,
	}
	fmt.Println(scores["张三"]) // 输出：90

	// 添加或修改键值对
	scores["赵六"] = 88
	delete(scores, "李四") // 删除键值对

	// 判断键是否存在
	score, ok := scores["张三"]
	if ok {
		fmt.Println("张三的分数是", score)
	}
}
```

映射（`map`）是一种键值对集合，通过键快速查找对应的值。键的类型必须是可比较类型（如 `string`、`int`），值可以是任意类型。

常用操作：

- `map[key]` — 读取值，键不存在时返回零值
- `map[key] = value` — 添加或修改
- `delete(map, key)` — 删除键值对
- `value, ok := map[key]` — 判断键是否存在，`ok` 为 `true` 表示存在

## 结构体

```go
package main

import "fmt"

// 定义结构体类型
type Person struct {
	Name string
	Age  int
}

func main() {
	// 创建结构体实例
	p := Person{
		Name: "小满",
		Age:  18,
	}
	fmt.Println(p.Name, p.Age) // 输出：小满 18

	// 访问和修改字段
	p.Age = 20
	fmt.Println(p.Age) // 输出：20
}
```

结构体（`struct`）用于将多个不同类型的字段组合成一个自定义类型，是面向对象编程的基础。通过 `type 名称 struct { ... }` 定义，用点号 `.` 访问字段。

结构体适合描述具有多个属性的实体，例如用户信息、商品详情、坐标点等。

## 接口

```go
package main

import "fmt"

// 定义接口：包含一个 Speak 方法
type Animal interface {
	Speak() string
}

type Dog struct{}
type Cat struct{}

func (d Dog) Speak() string { return "汪汪" }
func (c Cat) Speak() string  { return "喵喵" }

func main() {
	var a Animal

	a = Dog{}
	fmt.Println(a.Speak()) // 输出：汪汪

	a = Cat{}
	fmt.Println(a.Speak()) // 输出：喵喵
}
```

接口（`interface`）定义了一组方法签名，任何实现了这些方法的类型都自动满足该接口，无需显式声明。这是 Go 实现多态的核心机制。

接口的特点：

- **隐式实现**：只要类型实现了接口的所有方法，就满足该接口
- **空接口** `interface{}`（或 `any`）可以接收任意类型的值
- **组合接口**：多个接口可以嵌入组合成更大的接口

> 接口让代码更灵活——函数可以接收接口类型参数，从而兼容所有实现了该接口的类型，便于扩展和测试。

## 作业

1. **字符串**：声明两个字符串变量，分别存储你的姓名和学校名称，用 `+` 拼接后打印，中间用换行符 `\n` 分隔。
2. **数字**：声明一个 `int` 类型的年龄变量和一个 `float64` 类型的身高变量（如 `175.5`），并打印输出。
3. **布尔**：声明一个 `bool` 变量表示「是否在线」，初始值为 `true`，打印后将其改为 `false` 再打印一次。
4. **数组与切片**：
   - 定义一个长度为 3 的整型数组，赋值为 `[10, 20, 30]` 并打印；
   - 定义一个空切片，用 `append` 依次追加 `1, 2, 3`，打印切片内容和长度 `len`。
5. **映射**：创建一个 `map[string]int`，存储 3 位同学的姓名和分数，打印其中一位的分数，再添加一位新同学并删除其中一位，最后遍历打印所有键值对。
6. **结构体**：定义一个 `Book` 结构体，包含 `Title`（书名）和 `Price`（价格）两个字段，创建一本书的实例并打印书名和价格。
7. **接口**：定义一个 `Shape` 接口，包含 `Area() float64` 方法；再定义 `Rectangle`（矩形）结构体并实现该接口，创建实例后打印面积。
8. **综合练习**：将上述 `Book` 结构体放入切片中，存储至少 3 本书，遍历切片打印每本书的信息。

# 5 类型转换

实际开发里，不同类型之间经常需要互相转换：字符串转数字、数字转字符串、浮点数转整数、整数转浮点数，等等。

### 数字的转换

比如有一个浮点数 `3.1415926535`，想把它变成整数，可以用 `int()` 来转换：

```go
package main
import "fmt"

func main() {
	num := 3.1415926535
	intNum := int(num)
	fmt.Println(intNum) // 3
}
```

可以看到，`int()` 会把浮点数转成整数，并直接舍去小数部分（不是四舍五入）。

其他数值类型的转换写法也类似，记起来很简单——跟定义类型时的名字一模一样：

- `int()`：转为 `int`
- `int8()`：转为 `int8`
- `int16()`：转为 `int16`
- `int32()`：转为 `int32`
- `int64()`：转为 `int64`
- `uint()`：转为 `uint`
- `uint8()`：转为 `uint8`
- `uint16()`：转为 `uint16`
- `uint32()`：转为 `uint32`
- `uint64()`：转为 `uint64`
- `float32()`：转为 `float32`
- `float64()`：转为 `float64`

### 数字转字符串

数字和字符串之间的转换要用 `strconv`（string convert，字符串转换）包。比如把整数 `3` 转成字符串：

```go
package main

import (
	"fmt"
	"strconv"
)

func main() {
	num := 3
	strNum := strconv.Itoa(num)
	fmt.Printf("类型: %T,值: %v", strNum, strNum)
	// 类型: string,值: 3
}
```

几点说明：

1. 导入多个包时，用 `()` 分组，每个包独占一行。
2. `strconv.Itoa()` 就是 **Int to ASCII**：把整数转成字符串（字符串本质上是字符序列）。
3. `fmt.Printf()` 里，`%T` 打印类型，`%v` 打印值。

### 字符串转数字

反过来，字符串转数字也很常见。Go 里字符串和数字不能直接运算，必须先转成同一类型。比如这样会报错：

```go
var a = "10"
var b = 20
var c = a + b // 错误：字符串和数字不能直接运算
```

这时可以用 `strconv.Atoi()` 把字符串转成整数：

```go
package main

import (
	"fmt"
	"strconv"
)

func main() {
	str := "10"
	num, _ := strconv.Atoi(str)
	fmt.Printf("类型: %T,值: %v", num, num)
	// 类型: int,值: 10
}
```

`Atoi` 是 **ASCII to Int** 的缩写。它会返回两个值：转换结果和错误。转换失败时第二个返回值不为 `nil`；这里用 `_` 忽略错误，表示暂时不处理失败情况。

### 字符串转布尔值

字符串转布尔用 `strconv.ParseBool()`。它同样返回「结果 + 错误」两个值：

```go
package main

import (
	"fmt"
	"strconv"
)

func main() {
	str := "true"
	b, _ := strconv.ParseBool(str)
	fmt.Printf("类型: %T,值: %v", b, b)
	// 类型: bool,值: true
}
```

`ParseBool` 能识别的真值有：`1`、`t`、`T`、`TRUE`、`true`、`True`；假值有：`0`、`f`、`F`、`FALSE`、`false`、`False`。其他字符串都会转换失败。

### 布尔值转字符串

反过来，用 `strconv.FormatBool()` 把布尔值转成字符串，结果只会是 `"true"` 或 `"false"`：

```go
package main

import (
	"fmt"
	"strconv"
)

func main() {
	flag := true
	str := strconv.FormatBool(flag)
	fmt.Printf("类型: %T,值: %v", str, str)
	// 类型: string,值: true
}
```

## 作业

1. **数值转换**：声明一个 `float64` 变量，值为 `9.99`，用 `int()` 转成整数并打印；再把结果用 `float64()` 转回去，观察小数部分是否还在。
2. **数字转字符串**：声明一个 `int` 变量表示年龄，用 `strconv.Itoa()` 转成字符串，再用 `fmt.Printf` 打印类型和值（确认类型是 `string`）。
3. **字符串转数字**：声明两个字符串 `"15"` 和 `"27"`，分别用 `strconv.Atoi()` 转成整数后相加，打印和。
4. **布尔转换**：把字符串 `"1"`、`"false"`、`"yes"` 分别传给 `strconv.ParseBool()`，打印每次转换的结果和错误（这次不要用 `_` 忽略错误，看看哪些能成功）。
5. **综合练习**：声明字符串价格 `"199.5"` 和折扣标记 `"true"`。先把价格用 `strconv.ParseFloat(str, 64)` 转成 `float64`，折扣标记用 `ParseBool` 转成布尔；若享受折扣则打八折，最后用 `strconv.FormatFloat` 把应付金额转成字符串并打印。

# 6 数组和切片

数组和切片都是用来存放一组同类型数据的结构，写法也很像，所以放在一起讲。先记住一个关键区别：**数组长度固定，切片长度可变**。

## 数组

数组的长度在声明时就定死了，之后不能改。元素类型随意，但同一数组里必须是同一种类型。

语法：

```go
var 数组名 = [长度]类型{值1, 值2, 值3, ...}
数组名 := [长度]类型{值1, 值2, 值3, ...}
```

例子：

```go
package main

import "fmt"

func main() {
	var arr = [5]int{1, 2, 3, 4, 5}
	arr2 := [5]int{6, 7, 8, 9, 10}
	fmt.Println(arr, arr2)
}
```

两种写法都能声明数组：`var` 完整声明，或 `:=` 短声明。`[5]int` 表示「长度为 5、元素类型为 `int`」——**长度是类型的一部分**，`[5]int` 和 `[4]int` 是两种不同类型。

### 读写元素

通过下标访问和修改元素。下标从 `0` 开始：

```go
package main

import "fmt"

func main() {
	arr := [4]string{"唱", "跳", "rap", "篮球"}
	// 			下标：     0     1     2      3

	fmt.Println(arr[0]) // 唱

	arr[0] = "篮球"
	fmt.Println(arr[0]) // 篮球
}
```

### 获取长度

用内置函数 `len` 获取数组长度：

```go
package main

import "fmt"

func main() {
	arr := [4]string{"唱", "跳", "rap", "篮球"}
	fmt.Println(len(arr)) // 4
}
```

## 切片

切片可以看作「可变长度的数组」。日常开发里用得比数组多得多。

语法上，把数组的长度去掉就行：

```go
var 切片名 = []类型{值1, 值2, 值3, ...}
切片名 := []类型{值1, 值2, 值3, ...}
```

例子：

```go
package main

import "fmt"

func main() {
	var s1 = []int{1, 2, 3}
	s2 := []string{"唱", "跳", "rap"}
	fmt.Println(s1, s2)
}
```

`[]int` 没有写死长度，所以这是切片，不是数组。

### 读写元素

读写方式和数组一样，也是按下标操作：

```go
package main

import "fmt"

func main() {
	s := []string{"唱", "跳", "rap", "篮球"}

	fmt.Println(s[1]) // 跳

	s[1] = "打球"
	fmt.Println(s) // [唱 打球 rap 篮球]
}
```

### 长度与容量

切片有两个常用属性：

- `len`：当前元素个数
- `cap`：底层数组的容量（还能装多少，不用立刻扩容）

```go
package main

import "fmt"

func main() {
	s := []int{1, 2, 3}
	fmt.Println(len(s)) // 3
	fmt.Println(cap(s)) // 3
}
```

### 切片截取

用 `切片[低:高]` 可以从已有切片（或数组）上「切」出一段，得到一个新切片。区间是**左闭右开**：包含下标 `低`，不包含下标 `高`。

```go
package main

import "fmt"

func main() {
	s := []int{10, 20, 30, 40, 50}
	// 下标：  0   1   2   3   4

	fmt.Println(s[1:4]) // [20 30 40]  取下标 1、2、3
	fmt.Println(s[:3])  // [10 20 30]  从头切到下标 3（不含 3）
	fmt.Println(s[2:])  // [30 40 50]  从下标 2 切到末尾
	fmt.Println(s[:])   // [10 20 30 40 50]  整段复制一份切片头
}
```

几种常见写法：

| 写法       | 含义                         |
| ---------- | ---------------------------- |
| `s[低:高]` | 从 `低` 到 `高`（不含 `高`） |
| `s[:高]`   | 从开头到 `高`（不含 `高`）   |
| `s[低:]`   | 从 `低` 到末尾               |
| `s[:]`     | 整段截取                     |

也可以对数组做同样的事，结果是一个切片：

```go
package main

import "fmt"

func main() {
	arr := [5]int{10, 20, 30, 40, 50}
	s := arr[1:4]
	fmt.Println(s) // [20 30 40]
}
```

截取出来的切片和原数据共享底层数组：改一边，另一边也会变。后面做前置追加、中间插入时，会经常用到 `s[:idx]`、`s[idx:]` 这种写法。

### 追加元素

切片长度可变，最常见的操作就是往里面加元素。Go 只内置了往**尾部**追加的 `append`；往前面插、往中间插，都要靠 `append` 和切片截取自己拼。

#### 尾部追加

`append` 接收原切片和要追加的元素，返回一个新切片。容量不够时会自动扩容：

```go
package main

import "fmt"

func main() {
	s := []int{1, 2, 3}
	s = append(s, 4, 5)
	fmt.Println(s) // [1 2 3 4 5]
}
```

一次可以追加多个值。记得把返回值接回来（`s = append(...)`），否则原切片看不到新增内容。

#### 前置追加

Go 没有「往头部插入」的内置函数，思路是：先建一个装着新元素的切片，再用 `append` 把原切片接在后面。

这里会用到 `...`（展开操作符）：`s...` 表示把切片 `s` 里的元素一个个拆开，当作多个参数传给 `append`。

```go
package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5}
	s = append([]int{-1, 0}, s...)
	// 等同于：append([]int{-1, 0}, 1, 2, 3, 4, 5)
	fmt.Println(s) // [-1 0 1 2 3 4 5]
}
```

拆开看就两步：

1. `[]int{-1, 0}`：要插到前面的新元素
2. `s...`：把原来的 `1, 2, 3, 4, 5` 展开，接到后面

#### 中间插入

中间插入同样没有内置方法，本质是把切片拆成「前半段 + 新元素 + 后半段」，再拼回去。

假设切片是 `[-1 0 1 2 3 4 5]`，想在 `2` 后面（下标 `4` 的位置）插入 `100, 200`：

```go
package main

import "fmt"

func main() {
	s := []int{-1, 0, 1, 2, 3, 4, 5}
	// 下标：  0  1  2  3  4  5  6
	idx := 4 // 从下标 4 开始插入，也就是 2 的后面

	s = append(s[:idx], append([]int{100, 200}, s[idx:]...)...)
	fmt.Println(s) // [-1 0 1 2 100 200 3 4 5]
}
```

这行可以拆成三块理解：

| 部分              | 含义         | 结果         |
| ----------------- | ------------ | ------------ |
| `s[:idx]`         | 插入点之前   | `[-1 0 1 2]` |
| `[]int{100, 200}` | 要插入的内容 | `[100 200]`  |
| `s[idx:]`         | 插入点及之后 | `[3 4 5]`    |

先用 `append([]int{100, 200}, s[idx:]...)` 得到 `[100 200 3 4 5]`，再接到 `s[:idx]` 后面，就拼成了最终结果。

写法有点绕，但记口诀就行：**前半段 + 新元素 + 后半段**。

## 数组 vs 切片

|          | 数组                   | 切片                     |
| -------- | ---------------------- | ------------------------ |
| 长度     | 固定，声明时确定       | 可变，可用 `append` 增减 |
| 类型写法 | `[5]int`               | `[]int`                  |
| 传递方式 | 值拷贝（复制整个数组） | 引用底层数组（更轻量）   |
| 使用场景 | 长度已知且不变         | 日常开发更常用           |

入门阶段可以先记住：**不确定长度、或需要增删，就用切片。**

## 作业

1. **数组基础**：声明一个长度为 `5` 的 `int` 数组，依次赋值为 `1`～`5`，打印第 `3` 个元素（注意下标），再用 `len` 打印数组长度。
2. **切片声明与追加**：声明切片 `[]string{"唱", "跳", "rap"}`，用 `append` 在尾部追加 `"篮球"`，打印切片内容、`len` 和 `cap`。
3. **切片截取**：给定 `s := []int{10, 20, 30, 40, 50}`，分别打印 `s[1:4]`、`s[:3]`、`s[2:]`、`s[:]`，对照文中表格核对结果。
4. **前置与中间插入**：以 `[]int{1, 2, 3, 4, 5}` 为起点，先把 `-1, 0` 插到最前面；再在元素 `3` 后面插入 `100, 200`，最终打印完整切片。
5. **综合练习**：用切片存储若干门课程分数（如 `[]int{95, 58, 76, 42, 88}`）。遍历切片：低于 60 的改为 `0`（表示重考），然后截取前 `3` 门课的分数打印，并统计及格科目数量。

# 7 结构体

结构体是 Go 中的一种复合数据类型，用于将多个不同类型的数据组合在一起，形成一个新的数据类型。日常开发里，用结构体描述「一个人」「一辆车」「一条订单」这类对象非常常见。

## 定义结构体

用 `type` + `struct` 定义一种新类型：

```go
type 结构体名 struct {
	字段名 类型
	字段名 类型
	字段名 类型
}
```

实操：定义一个 `Person`，包含姓名、年龄和爱好：

```go
package main

import "fmt"

type Person struct {
	Name  string
	Age   int
	Hobby []string
}

func main() {
	person := Person{
		Name:  "John",
		Age:   20,
		Hobby: []string{"reading", "swimming"},
	}
	fmt.Println(person)
}
```

`Person{...}` 是字面量初始化：按字段名赋值，顺序可以随便写，没写的字段会是该类型的零值。

### 字段读写

通过 `.` 访问和修改字段：

```go
package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	p := Person{Name: "张三", Age: 18}
	fmt.Println(p.Name) // 张三

	p.Age = 20
	fmt.Println(p.Age) // 20
}
```

### 几种初始化方式

```go
package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	// 1. 按字段名初始化（推荐，可读性好）
	p1 := Person{Name: "张三", Age: 18}

	// 2. 按字段顺序初始化（字段多时容易写错，不推荐）
	p2 := Person{"李四", 20}

	// 3. 只给部分字段赋值，其余为零值
	p3 := Person{Name: "王五"} // Age 为 0

	// 4. 先声明再赋值（全部为零值）
	var p4 Person
	p4.Name = "赵六"
	p4.Age = 22

	// 5. 取地址，得到 *Person
	p5 := &Person{Name: "钱七", Age: 25}

	fmt.Println(p1, p2, p3, p4, p5)
}
```

指针形式 `p5` 访问字段时，Go 会自动解引用，写 `p5.Name` 即可，不必写 `(*p5).Name`。

## 结构体嵌套

一个结构体里可以包含另一个结构体类型的字段，用来组合更复杂的数据：

```go
package main

import "fmt"

type Car struct {
	Brand string
	Model string
	Year  int
}

type Person struct {
	Name  string
	Age   int
	Hobby []string
	Car   Car
}

func main() {
	person := Person{
		Name:  "John",
		Age:   20,
		Hobby: []string{"reading", "swimming"},
		Car: Car{
			Brand: "Toyota",
			Model: "Camry",
			Year:  2020,
		},
	}
	fmt.Println(person)
	fmt.Println(person.Car.Brand) // Toyota
}
```

这里 `Car` 是命名字段，访问时要写 `person.Car.Brand`。

## 结构体嵌入

如果字段类型前**不写字段名**，就是嵌入（匿名字段）。嵌入后，被嵌入类型的字段和方法会「提升」到外层，可以直接用：

```go
package main

import "fmt"

type Car struct {
	Brand string
	Model string
	Year  int
}

type Person struct {
	Name string
	Age  int
	Car        // 嵌入：没有字段名，只有类型
}

func main() {
	person := Person{
		Name: "John",
		Age:  20,
		Car: Car{
			Brand: "Toyota",
			Model: "Camry",
			Year:  2020,
		},
	}

	// 提升后可以直接访问
	fmt.Println(person.Brand) // Toyota
	// 也可以走完整路径
	fmt.Println(person.Car.Model) // Camry
}
```

嵌套是「有名字的组合」；嵌入是「把另一个类型的能力提升过来」。两者都能组合数据，嵌入更适合表达「is-a / 带有某种能力」的关系。

## 方法

方法是「绑在某个类型上的函数」。接收者写在 `func` 和函数名之间：

```go
func (接收者 类型) 方法名(参数列表) 返回值 {
	// ...
}
```

### 值接收者

```go
package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) Introduce() {
	fmt.Printf("我叫 %s，今年 %d 岁\n", p.Name, p.Age)
}

func main() {
	p := Person{Name: "张三", Age: 18}
	p.Introduce() // 我叫 张三，今年 18 岁
}
```

值接收者拿到的是副本，方法里改字段**不会**影响原变量。

### 指针接收者

需要修改结构体本身时，用指针接收者：

```go
package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p *Person) GrowUp() {
	p.Age++
}

func (p Person) Introduce() {
	fmt.Printf("我叫 %s，今年 %d 岁\n", p.Name, p.Age)
}

func main() {
	p := Person{Name: "张三", Age: 18}
	p.GrowUp()
	p.Introduce() // 我叫 张三，今年 19 岁
}
```

调用时写 `p.GrowUp()` 即可，Go 会自动取地址。习惯上：要改状态用指针接收者；只读、小结构体可以用值接收者。同一类型的方法最好统一风格，不要有的用值、有的用指针混着来。

## 匿名结构体

不提前 `type` 定义，直接在用的地方写 `struct { ... }`，适合只临时用一次的场景：

```go
package main

import "fmt"

func main() {
	person := struct {
		Name string
		Age  int
	}{
		Name: "John",
		Age:  20,
	}
	fmt.Println(person.Name, person.Age)
}
```

配置、单次返回、测试数据等「用完就丢」的结构，用匿名结构体很方便。若多处复用，还是起个名字更清晰。

## 泛型结构体

Go 1.18 起，结构体也可以带类型参数：

```go
package main

import "fmt"

type Person[T string | int] struct {
	Name  string
	Age   int
	Phone T
}

func main() {
	person := Person[string]{
		Name:  "John",
		Age:   20,
		Phone: "010-10086",
	}
	person2 := Person[int]{
		Name:  "John",
		Age:   20,
		Phone: 10086,
	}
	fmt.Println(person)
	fmt.Println(person2)
}
```

`Person[T string | int]` 表示 `Phone` 只能是 `string` 或 `int`。创建时要写明具体类型，如 `Person[string]`、`Person[int]`。

## 作业

1. **定义与初始化**：定义 `Book` 结构体（书名、作者、页数），用至少两种方式创建实例并打印。
2. **嵌套**：定义 `Address`（城市、街道）和 `User`（姓名、地址），创建用户并打印「姓名住在城市的街道」。
3. **嵌入**：把上一题的 `Address` 改成嵌入到 `User`，用提升后的字段直接打印城市。
4. **方法**：给 `Book` 写值接收者方法 `Info()` 打印信息；再写指针接收者方法 `AddPages(n int)` 增加页数，调用后验证页数有变。
5. **综合**：定义 `Rectangle`（宽、高），写方法 `Area()` 和 `Perimeter()` 返回面积和周长，并在 `main` 里验证。

# 8 Map

Map 是一种**无序**的键值对集合，适合按「键」快速查找「值」。日常开发里常用它存配置、统计次数、缓存结果等。

Map 是引用类型，零值是 `nil`。**键必须是可比较的类型**（如 `string`、`int`、数组等），不能是切片、map、函数；值可以是任意类型。

## 语法

```go
var 变量名 map[键类型]值类型
变量名 := map[键类型]值类型{键: 值, ...}
变量名 := make(map[键类型]值类型)
```

## 定义与初始化

有三种常见写法：字面量、`make`、以及先声明再赋值。

```go
package main

import "fmt"

func main() {
	// 1. 字面量初始化（推荐，一眼能看出内容）
	scores := map[string]int{
		"语文": 90,
		"数学": 85,
		"英语": 92,
	}

	// 2. make 创建空 map，再往里塞
	person := make(map[string]string)
	person["name"] = "张三"
	person["city"] = "杭州"

	// 3. 先声明（此时是 nil），再赋值一个已有 map
	var ages map[string]int
	ages = map[string]int{"张三": 18, "李四": 20}

	fmt.Println(scores)
	fmt.Println(person)
	fmt.Println(ages)
}
```

注意：只写 `var m map[string]int` 而不初始化时，`m` 是 `nil`。**向 nil map 写入会 panic**，读却可以（得到零值）。需要写入时，务必先用 `make` 或字面量初始化。

```go
package main

func main() {
	var m map[string]int
	// m["a"] = 1  // panic: assignment to entry in nil map
	_ = m["a"]     // 读可以，得到 0
}
```

也可以给 `make` 传第二个参数，预估容量，减少扩容次数（不是长度上限）：

```go
m := make(map[string]int, 10)
```

## 读写元素

用 `map[键]` 读写，和切片按下标访问很像：

```go
package main

import "fmt"

func main() {
	scores := map[string]int{
		"语文": 90,
		"数学": 85,
	}

	fmt.Println(scores["语文"]) // 90

	scores["数学"] = 95        // 修改
	scores["英语"] = 88        // 新增
	fmt.Println(scores)       // map[语文:90 数学:95 英语:88]
}
```

键不存在时，读出来的是**值类型的零值**，不会报错：

```go
package main

import "fmt"

func main() {
	scores := map[string]int{"语文": 90}
	fmt.Println(scores["体育"]) // 0（int 的零值）
}
```

所以单靠读到的值，分不清「键不存在」和「值刚好是零值」。这时要用「逗号 ok」写法。

## 判断键是否存在

`value, ok := map[键]`：`ok` 为 `true` 表示键存在。

```go
package main

import "fmt"

func main() {
	scores := map[string]int{
		"语文": 90,
		"体育": 0, // 值就是 0
	}

	if score, ok := scores["语文"]; ok {
		fmt.Println("语文成绩：", score) // 语文成绩： 90
	}

	if score, ok := scores["体育"]; ok {
		fmt.Println("体育成绩：", score) // 体育成绩： 0（键存在，值就是 0）
	} else {
		fmt.Println("没有体育这门课")
	}

	if _, ok := scores["音乐"]; !ok {
		fmt.Println("没有音乐这门课")
	}
}
```

只要判断存在与否、不关心值时，用 `_` 丢掉第一个返回值即可：`_, ok := scores["音乐"]`。

## 删除元素

用内置函数 `delete`。键不存在时调用也是安全的，不会报错：

```go
package main

import "fmt"

func main() {
	scores := map[string]int{
		"语文": 90,
		"数学": 85,
		"英语": 92,
	}

	delete(scores, "数学")
	fmt.Println(scores) // map[语文:90 英语:92]

	delete(scores, "体育") // 键不存在，什么都不发生
	fmt.Println(scores)
}
```

`delete` 没有返回值，也不会返回「是否删成功」。若要确认，删除后再用逗号 ok 查一次即可。

## 获取长度

用 `len` 获取键值对个数：

```go
package main

import "fmt"

func main() {
	scores := map[string]int{
		"语文": 90,
		"数学": 85,
	}
	fmt.Println(len(scores)) // 2

	delete(scores, "数学")
	fmt.Println(len(scores)) // 1
}
```

Map 没有「容量」概念，不像切片那样有 `cap`。

## 遍历 map

用 `range` 遍历。每次迭代返回**键**和**值**：

```go
package main

import "fmt"

func main() {
	person := map[string]string{
		"name": "张三",
		"city": "杭州",
		"job":  "程序员",
	}

	for key, value := range person {
		fmt.Println(key, value)
	}
}
```

只要键或只要值时：

```go
for key := range person {
	fmt.Println(key)
}

for _, value := range person {
	fmt.Println(value)
}
```

注意：map 的遍历**顺序不固定**，每次运行打印顺序可能不同。若需要稳定顺序，可以把键收集到切片里再排序后遍历。

## 作为函数参数

Map 是引用类型，传给函数后，函数里对它的增删改会反映到外面：

```go
package main

import "fmt"

func bump(scores map[string]int, subject string, delta int) {
	scores[subject] += delta
}

func main() {
	scores := map[string]int{"语文": 90}
	bump(scores, "语文", 5)
	fmt.Println(scores["语文"]) // 95
}
```

但若在函数里把参数重新指向另一个 map（`scores = make(...)`），外面的变量不会跟着变——改的是局部变量本身，不是底层数据。

## 常见注意点

| 点       | 说明                                                |
| -------- | --------------------------------------------------- |
| 零值     | `nil`，写入会 panic，读取得到值类型零值             |
| 键类型   | 必须可比较，不能是 slice / map / function           |
| 无序     | `range` 顺序不固定                                  |
| 键不存在 | 读到零值；用 `value, ok := m[k]` 区分               |
| 并发     | 默认**不是**并发安全的，多 goroutine 同时写会出问题 |

## 作业

1. **声明与读写**：创建一个 `map[string]string`，存自己的姓名、城市、爱好；修改城市，再新增一个「职业」键，打印整个 map。
2. **逗号 ok**：给定 `scores := map[string]int{"语文": 90, "数学": 0}`，分别判断「语文」「数学」「英语」是否存在，并正确打印成绩或「不存在」。
3. **删除与长度**：在上一题的 map 上删除「数学」，用 `len` 打印剩余键值对数量。
4. **遍历统计**：用 `map[string]int` 统计字符串切片 `[]string{"苹果", "香蕉", "苹果", "橘子", "香蕉", "苹果"}` 中每种水果出现的次数，并遍历打印结果。
5. **综合练习**：写一个函数 `addScore(scores map[string]int, name string, score int)`，把某人的分数累加进去（没有则新建）。在 `main` 里多次调用后，打印总分最高的姓名和分数。

# 9 条件判断

```go
package main

import "fmt"

type Person struct {
	Name string
	Age  int
	Sex  string
}

func main() {
	person := Person{
		Name: "张三",
		Age:  18,
		Sex:  "男",
	}
	if person.Age >= 18 {
		fmt.Println("成年")
	} else {
		fmt.Println("未成年")
	}
}
```

我们定义了一个 `Person` 结构体，包含 `Name`、`Age`、`Sex` 三个字段。然后创建了一个实例，判断其年龄是否满 18 岁：满 18 岁输出「成年」，否则输出「未成年」。这就是最基本的 `if` / `else` 二选一判断。

## 多条件判断

```go
package main

import "fmt"

type Person struct {
	Status string
}

func main() {
	person := Person{
		Status: "在线",
	}
	if person.Status == "在线" {
		fmt.Println("在线")
	} else if person.Status == "离线" {
		fmt.Println("离线")
	} else if person.Status == "忙碌" {
		fmt.Println("忙碌")
	} else {
		fmt.Println("未知状态")
	}
}
```

当需要处理**三种及以上**的分支时，可以在 `if` 后面连续使用 `else if`，从上到下依次判断，**命中第一个为真的分支后就不再往下执行**。最后的 `else` 用于兜底，处理所有未匹配的情况。

上面示例模拟了用户在线状态：`Status` 为「在线」「离线」「忙碌」时分别输出对应文字，其他值则输出「未知状态」。

同样的逻辑，用 `switch` 写会更简洁——把要判断的值写在 `switch` 后面，每个 `case` 对应一种取值：

```go
package main

import "fmt"

type Person struct {
	Status string
}

func main() {
	person := Person{
		Status: "在线",
	}
	switch person.Status {
	case "在线":
		fmt.Println("在线")
	case "离线":
		fmt.Println("离线")
	case "忙碌":
		fmt.Println("忙碌")
	default:
		fmt.Println("未知状态")
	}
}
```

`switch` 会拿 `person.Status` 与每个 `case` 逐一比较，匹配到就执行对应分支。Go 的 `switch` **默认不会穿透**到下一个 `case`（不需要写 `break`）。所有 `case` 都不匹配时，执行 `default` 分支。

| 写法             | 适用场景                                       |
| ---------------- | ---------------------------------------------- |
| `if` / `else if` | 条件各不相同，不好抽成一个表达式               |
| `switch 表达式`  | 同一个值匹配多种固定取值（如状态、等级、星期） |

## 逻辑嵌套

```go
package main

import "fmt"

type Person struct {
	Name string
	Age  int
	Sex  string
}

func main() {
	person := Person{
		Name: "张三",
		Age:  18,
		Sex:  "男",
	}
	if person.Age >= 18 {
		if person.Sex == "男" {
			fmt.Println("成年男性")
		} else {
			fmt.Println("成年女性")
		}
	} else {
		fmt.Println("未成年")
		if person.Sex == "男" {
			fmt.Println("未成年男性")
		} else {
			fmt.Println("未成年女性")
		}
	}
}
```

多个if条件判断可以嵌套使用，我们首先判断年龄是否等于18，如果等于18，则再判断性别是否等于"男"，如果等于"男"，则输出"成年男性"，否则输出"成年女性"。如果不等于18，则输出"未成年"，再判断性别是否等于"男"，如果等于"男"，则输出"未成年男性"，否则输出"未成年女性"。这就是一个基本的逻辑嵌套。当然也不是只能嵌套两层，可以嵌套多层，但是要注意逻辑的正确性。

## 逻辑运算符

判断条件经常需要组合多个表达式，常用的逻辑运算符如下：

| 运算符 | 含义 | 示例     | 为真条件              |
| ------ | ---- | -------- | --------------------- |
| `&&`   | 并且 | `a && b` | `a`、`b` 都为真       |
| `||`   | 或者 | `a || b` | `a`、`b` 至少一个为真 |
| `!`    | 取反 | `!a`     | `a` 为假              |

### 并且（`&&`）

```go
age := 18
sex := "男"

if age >= 18 && sex == "男" {
	fmt.Println("成年男性")
}
```

`&&` 表示**并且**：左右两边条件都为真，整个表达式才为真。上面只有「年满 18 岁」且「性别为男」时，才会打印「成年男性」。

### 或者（`||`）

```go
age := 33

if age == 18 || age == 19 {
	fmt.Println("18 岁或 19 岁")
} else {
	fmt.Println("其他年龄")
}
```

`||` 表示**或者**：左右两边只要有一个为真，整个表达式就为真。上面当年龄是 18 **或** 19 时走 `if` 分支，否则走 `else`。

### 取反（`!`）

```go
isLogin := false

if !isLogin {
	fmt.Println("请先登录")
}
```

`!` 把布尔值**反过来**：`!true` 为 `false`，`!false` 为 `true`。常用于「不满足某条件时」的分支判断。

## 作业

1. **基础判断**：声明一个 `int` 类型的分数变量，用 `if` / `else if` / `else` 判断等级：90 分及以上为「优秀」，60～89 为「及格」，60 分以下为「不及格」，并打印结果。
2. **switch 练习**：定义一个 `string` 变量表示星期几（如 `"周一"`），用 `switch` 匹配并打印对应提示；未匹配时走 `default`，输出「无效的星期」。
3. **状态判断**：参照文中的 `Person` 结构体，将 `Status` 分别设为「在线」「离线」「忙碌」和一个自定义值，分别运行 `if` / `else if` 和 `switch` 两个版本，观察输出是否一致。
4. **逻辑运算符**：声明 `age` 和 `hasTicket` 两个变量，用 `&&` 判断「年满 18 岁且已购票」方可入场；再改用 `||` 判断「会员或年龄大于 60 岁」可享受优惠，分别打印结果。
5. **综合练习**：定义包含 `Name`、`Age`、`Score` 的结构体，结合 `if` 嵌套判断：年满 18 且分数不低于 60 才输出「合格」，否则输出「不合格」，并带上姓名。

# 10 循环

Go 里**只有一种循环关键字：`for`**。没有 `while`、也没有 `do-while`，所有循环形态都靠 `for` 表达。

## 普通 for 循环

我们声明了一个切片，用经典的「初始化；条件；步进」写法遍历：

```go
package main

import "fmt"

func main() {
	var hobbys = []string{"吃饭", "睡觉", "打豆豆"}
	for i := 0; i < len(hobbys); i++ {
		fmt.Println(hobbys[i])
	}
}
```

`for` 后面分成三段，用分号隔开：

| 部分   | 作用                       | 上面示例          |
| ------ | -------------------------- | ----------------- |
| 初始化 | 循环开始前执行一次         | `i := 0`          |
| 条件   | 每次循环前判断，为真才继续 | `i < len(hobbys)` |
| 步进   | 每次循环体结束后执行       | `i++`             |

输出依次为：`吃饭`、`睡觉`、`打豆豆`。

三段都可以省略。只留条件时，就相当于其他语言里的 `while`：

```go
package main

import "fmt"

func main() {
	i := 0
	for i < 3 {
		fmt.Println(i)
		i++
	}
}
```

条件也省略时，就是**无限循环**，需要靠 `break` 自己跳出：

```go
package main

import "fmt"

func main() {
	i := 0
	for {
		fmt.Println(i)
		i++
		if i >= 3 {
			break
		}
	}
}
```

## range 循环

遍历切片、数组、字符串、map、channel 时，更常用 `range`，写法更简洁：

```go
package main

import "fmt"

func main() {
	var hobbys = []string{"吃饭", "睡觉", "打豆豆"}
	for index, hobby := range hobbys {
		fmt.Println(index, hobby)
	}
}
```

`range` 每次迭代会返回两个值：

- 第一个是**下标**（或 map 的 key）
- 第二个是**对应的元素值**

输出为：

```text
0 吃饭
1 睡觉
2 打豆豆
```

如果只要值、不要下标，用 `_` 丢掉第一个返回值：

```go
for _, hobby := range hobbys {
	fmt.Println(hobby)
}
```

如果只要下标、不要值，写一个变量即可：

```go
for index := range hobbys {
	fmt.Println(index)
}
```

### 遍历 map

```go
package main

import "fmt"

func main() {
	person := map[string]string{
		"name": "张三",
		"city": "杭州",
	}
	for key, value := range person {
		fmt.Println(key, value)
	}
}
```

注意：map 的遍历**顺序不固定**，每次运行打印顺序可能不同。

### 遍历字符串

```go
package main

import "fmt"

func main() {
	text := "你好Go"
	for index, char := range text {
		fmt.Printf("下标 %d，字符 %c\n", index, char)
	}
}
```

用 `range` 遍历字符串时，拿到的是 **Unicode 码点（rune）**，不是单个字节。像「你」「好」这种汉字各占 3 个字节，下标会按字节位置跳着走。

| 写法                          | 适用场景                                      |
| ----------------------------- | --------------------------------------------- |
| `for i := 0; i < len(x); i++` | 需要精确控制下标、步长，或从中间开始/倒序遍历 |
| `for i, v := range x`         | 顺序遍历切片、数组、字符串、map               |

## break 与 continue

循环中经常要「提前结束」或「跳过本次」：

| 关键字     | 含义                             |
| ---------- | -------------------------------- |
| `break`    | 立刻跳出**当前**循环             |
| `continue` | 跳过本次剩余代码，进入下一次循环 |

### break：跳出循环

```go
package main

import "fmt"

func main() {
	hobbys := []string{"吃饭", "睡觉", "打豆豆"}
	for _, hobby := range hobbys {
		if hobby == "睡觉" {
			break
		}
		fmt.Println(hobby)
	}
}
```

遇到「睡觉」就 `break`，后面的元素不再遍历。输出只有：`吃饭`。

### continue：跳过本次

```go
package main

import "fmt"

func main() {
	hobbys := []string{"吃饭", "睡觉", "打豆豆"}
	for _, hobby := range hobbys {
		if hobby == "睡觉" {
			continue
		}
		fmt.Println(hobby)
	}
}
```

遇到「睡觉」就 `continue`，跳过打印，但循环继续。输出为：`吃饭`、`打豆豆`。

### 带标签的 break

多层循环嵌套时，普通 `break` 只跳出最内层。若要一次跳出外层，可以给循环加**标签**：

```go
package main

import "fmt"

func main() {
Outer:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == 1 && j == 1 {
				break Outer
			}
			fmt.Println(i, j)
		}
	}
}
```

`break Outer` 会直接跳出名为 `Outer` 的那一层循环。`continue` 也可以配合标签使用，表示跳到指定循环的下一次迭代。

## goto

`goto` 可以跳转到同一函数内的某个标签位置。能用，但可读性往往更差，日常代码里更推荐 `break` / `continue` / `return`。

```go
package main

import "fmt"

func main() {
	i := 0
Loop:
	fmt.Println(i)
	i++
	if i < 3 {
		goto Loop
	}
}
```

上面用 `goto` 模拟了一个简单循环：打印 `0`、`1`、`2`。实际项目中除非处理错误清理这类特殊场景，否则尽量少用。

## 作业

1. **普通 for**：声明一个 `1` 到 `10` 的循环，打印所有偶数。
2. **range 练习**：定义切片 `[]string{"春", "夏", "秋", "冬"}`，分别用「只要下标」「只要值」「下标+值」三种 `range` 写法打印。
3. **break / continue**：遍历 `1` 到 `20`，遇到能被 `3` 整除的数就 `continue` 跳过；一旦遇到大于 `15` 的数就 `break` 结束。观察最终打印了哪些数。
4. **类 while**：用「只写条件」的 `for`（不要初始化段和步进段）计算 `1 + 2 + ... + 100` 的和并打印。
5. **综合练习**：定义 `map[string]int` 表示各科分数（如语文、数学、英语），用 `range` 遍历：低于 60 的科目打印「不及格」，其余打印「及格」，并统计及格科目数量。

# 11 函数

函数是一段可复用的代码块：接收参数、执行逻辑，再通过 `return` 把结果交还给调用方。Go 的入口也是函数——程序从 `main` 开始跑。

## 函数定义

基本语法：

```go
func 函数名(参数名 参数类型) 返回值类型 {
	// 函数体
	return 返回值
}
```

拆开看每一部分：

| 部分       | 含义                                                       |
| ---------- | ---------------------------------------------------------- |
| `func`     | 声明函数的关键字                                           |
| 函数名     | 包内可调用的名字；首字母大写表示对外导出                   |
| 参数列表   | `名字 类型`，多个参数用逗号分隔；同类型可简写为 `a, b int` |
| 返回值类型 | 写在参数列表后面；无返回值时可省略                         |
| `return`   | 结束函数并把结果返回给调用方；返回值类型必须与声明一致     |

最小例子：

```go
package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func main() {
	fmt.Println(add(1, 2)) // 3
}
```

参数同类型时可以合并写法：`func add(a, b int) int`。

### 多返回值

Go 支持一次返回多个值，常见于「结果 + 错误」（详见下一节「错误处理」）：

```go
package main

import "fmt"

func calc(a, b int) (int, int) {
	return a + b, a - b
}

func main() {
	sum, sub := calc(1, 2)
	fmt.Println(sum, sub) // 3 -1
}
```

也可以给返回值命名（命名返回值）。此时 `return` 可以不带表达式，会返回当前命名变量的值：

```go
func calc(a, b int) (sum int, sub int) {
	sum = a + b
	sub = a - b
	return // 等价于 return sum, sub
}
```

不需要某个返回值时，用 `_` 丢弃：

```go
sum, _ := calc(1, 2)
```

### 注意事项

1. **不能在函数内部声明具名函数**。这不只是 `main`，任意函数里都不允许再写 `func foo() {}` 这种带名字的声明。需要嵌套逻辑时，用**匿名函数**（见下文）。
2. 函数必须先声明再调用的说法不成立：同一包内函数可以任意顺序定义。
3. Go 是**按值传递**：传入的是副本。想改调用方的变量，要传指针，或使用切片 / map 这类引用语义类型。
4. 可变参数用 `...类型`，例如 `func sum(nums ...int) int`，调用时写成 `sum(1, 2, 3)`。

错误示例（无法编译）：

```go
func main() {
	// 不允许：函数里再声明具名函数
	func add(a, b int) int {
		return a + b
	}
}
```

正确写法：把内部逻辑写成匿名函数，赋给变量：

```go
func main() {
	add := func(a, b int) int {
		return a + b
	}
	fmt.Println(add(1, 2))
}
```

## 错误处理

Go **没有** `try / catch`。函数出错时，惯例是多返回值：最后一个是 `error`。调用方必须自己检查。

`error` 是内置接口，只有一个方法 `Error() string`。没出错时返回 `nil`，有错时返回非 `nil`。

```go
package main

import (
	"errors"
	"fmt"
)

// 除数为 0 时返回错误
func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("除数不能为 0")
	}
	return a / b, nil
}

func main() {
	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("出错了:", err)
		return
	}
	fmt.Println("结果:", result) // 5

	_, err = divide(10, 0)
	if err != nil {
		fmt.Println("出错了:", err) // 出错了: 除数不能为 0
	}
}
```

常用写法说明：

| 写法                         | 作用                     |
| ---------------------------- | ------------------------ |
| `errors.New("说明")`         | 创建一个简单错误         |
| `fmt.Errorf("a=%d 非法", a)` | 带格式化信息的错误       |
| `if err != nil { ... }`      | 检查错误的标准姿势       |
| `return 0, err`              | 把错误继续往上抛给调用方 |

注意：

1. **不要忽略 `error`**。写成 `result, _ := divide(...)` 等于把错误吞掉，出了问题很难排查。
2. 出错时，第一个返回值通常给零值（如 `0`、`""`、`nil`），真正该看的是 `err`。
3. 很多标准库函数也是这个模式，例如 `os.Open`、`strconv.Atoi`：

```go
package main

import (
	"fmt"
	"strconv"
)

func main() {
	n, err := strconv.Atoi("123")
	if err != nil {
		fmt.Println("转换失败:", err)
		return
	}
	fmt.Println(n) // 123

	_, err = strconv.Atoi("abc")
	if err != nil {
		fmt.Println("转换失败:", err)
	}
}
```

和 `panic` 的区别（先建立直觉即可）：日常可预期的失败（文件不存在、输入不合法）用 `error`；程序无法继续的严重问题才用 `panic`。入门阶段优先掌握 `if err != nil`。

## 匿名函数

没有名字的函数叫匿名函数。可以立刻执行，也可以赋给变量、当作参数传递。

立刻执行（IIFE）：

```go
package main

import "fmt"

func main() {
	func() {
		fmt.Println("匿名函数")
	}() // 末尾 () 表示立刻调用
}
```

赋给变量后再调用：

```go
package main

import "fmt"

func main() {
	greet := func(name string) {
		fmt.Println("你好,", name)
	}
	greet("Go")
}
```

### 注意事项

1. **闭包会捕获外层变量**。匿名函数引用外层变量时，拿到的是变量本身（引用语义），不是当时的拷贝。循环里起 goroutine 时尤其容易踩坑，要额外注意。
2. 匿名函数可以访问外层作用域的变量，适合做回调、延迟逻辑（配合 `defer`）。
3. 只有「函数类型」的值才能赋值、传递；具名函数本身不是变量，但可以写成 `f := add` 把函数当值用。
4. 需要反复复用、导出给其他包用的逻辑，优先写成包级具名函数，可读性和调试都更好。

闭包示例：

```go
package main

import "fmt"

func makeCounter() func() int {
	n := 0
	return func() int {
		n++
		return n
	}
}

func main() {
	next := makeCounter()
	fmt.Println(next()) // 1
	fmt.Println(next()) // 2
}
```

## 泛型函数

没有泛型时，想对 `int`、`float64` 做同样的加法，往往要写两套几乎一样的函数，很繁琐：

```go
func addInt(a, b int) int { return a + b }
func addFloat(a, b float64) float64 { return a + b }
```

Go 1.18 起支持**类型参数**。把「类型」也参数化，一套函数就能服务多种类型——这就是泛型的中心思想：

```go
package main

import "fmt"

// T 被约束为 int、uint 或 float64，这些类型都能做 +
func add[T int | uint | float64](a, b T) T {
	return a + b
}

func main() {
	fmt.Println(add[int](1, 2))       // 显式指定类型参数
	fmt.Println(add[uint](1, 2))
	fmt.Println(add(1.5, 2.5))        // 编译器从实参推断 T 为 float64
}
```

- `[T int | uint | float64]`：类型参数 `T`，以及它允许的类型集合（类型约束）。
- 调用时可以写 `add[int](1, 2)`，多数情况也能靠实参**类型推断**省略。
- 约束里的类型必须都支持函数体里用到的运算；否则编译报错。

也可以用 `comparable`、自定义 interface 约束等更复杂的写法；入门先记住：泛型用来消灭「同逻辑、多类型」的重复函数。

## defer 关键字

`defer` 会把一次函数调用推迟到**当前函数即将返回之前**执行，常用于关闭文件、解锁、收尾打印等「成对出现」的清理工作。

基本例子：

```go
package main

import "fmt"

func main() {
	fmt.Println("main 1")
	defer fmt.Println("defer 1")
	fmt.Println("main 2")
}
```

输出：

```text
main 1
main 2
defer 1
```

### 多个 defer：后进先出（LIFO）

多个 `defer` 按栈顺序执行，**最后注册的最先跑**：

```go
package main

import "fmt"

func main() {
	defer fmt.Println("第一")
	defer fmt.Println("第二")
	defer fmt.Println("第三")
	fmt.Println("函数体")
}
```

输出：

```text
函数体
第三
第二
第一
```

### 实用案例：保证资源释放

打开文件后立刻 `defer Close()`，后面无论正常返回还是提前 `return`，都会关闭：

```go
package main

import (
	"fmt"
	"os"
)

func readFile(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close() // 函数返回前关闭，避免泄漏

	buf := make([]byte, 64)
	n, err := f.Read(buf)
	if err != nil {
		return err
	}
	fmt.Println(string(buf[:n]))
	return nil
}
```

### 再记两点

1. **参数在 defer 语句执行时就已经算好**，不是等到真正调用时才算：

```go
package main

import "fmt"

func main() {
	x := 1
	defer fmt.Println("defer:", x) // 这里已经把 x=1 传进去了
	x = 2
	fmt.Println("main:", x)
}
```

输出是 `main: 2`，然后 `defer: 1`。

1. `defer` 绑定的是**函数**，不是代码块。写在 `for` 里的 `defer` 会等到外层函数返回才执行；循环里反复 `defer` 可能堆积很多调用，一般要包一层函数或换写法。

## 用接口模拟「面向对象」能力

先把话说清楚：

- Go **没有** class、没有基于类的继承，也**不是**传统 Java / C++ 那种面向对象语言。
- Go 提供的是：**结构体存数据**、**方法绑行为**、**接口描述能力**。多态主要靠 interface 实现。
- 所以更准确的说法是：Go 用「组合 + 方法 + 接口」达到类似面向对象里「多态、面向行为编程」的效果，而不是复制一套 class 体系。

接口只声明「要有哪些方法」，不关心具体类型：

```go
package main

import "fmt"

// Speaker 描述「能说话」这种能力
type Speaker interface {
	Speak() string
}

type Dog struct {
	Name string
}

func (d Dog) Speak() string {
	return d.Name + ": 汪汪"
}

type Cat struct {
	Name string
}

func (c Cat) Speak() string {
	return c.Name + ": 喵喵"
}

// 只依赖接口，不依赖具体类型
func say(s Speaker) {
	fmt.Println(s.Speak())
}

func main() {
	say(Dog{Name: "旺财"})
	say(Cat{Name: "咪咪"})
}
```

要点：

1. 类型**不用显式写 implements**。只要方法集满足接口，就自动实现了该接口。
2. 函数参数写成接口类型后，传入任何实现了这些方法的值都可以——这就是接口带来的多态。
3. 需要「复用字段 / 方法」时，优先用结构体**嵌入（组合）**，而不是继承。

和「有没有面向对象」相关的常见误区：看到方法、接口就以为 Go 是 class OOP——其实它刻意避开了继承树，把重点放在**接口**和**组合**上。

## 作业

1. **基本函数**：写一个 `max(a, b int) int`，返回较大的那个数；在 `main` 里调用并打印结果。
2. **多返回值**：写一个 `swap(a, b int) (int, int)`，交换两个整数并返回；用 `x, y := swap(1, 2)` 验证。
3. **错误处理**：写一个 `sqrt(n float64) (float64, error)`：`n < 0` 时返回错误（可用 `errors.New`），否则返回平方根（可用 `math.Sqrt`）。在 `main` 里分别用正数和负数调用，正确处理 `err`。
4. **匿名函数**：在 `main` 里用匿名函数实现「打印 1 到 5」，立刻执行；再把同一个逻辑赋给变量后调用一次。
5. **综合练习**：写一个 `parseAndDouble(s string) (int, error)`，用 `strconv.Atoi` 把字符串转成整数后乘以 2；转换失败则把错误返回。在 `main` 里测试 `"21"` 和 `"abc"` 两种输入。

# 12 错误处理

Go **没有** `try / catch`。日常约定是：

- **能预期的失败**（参数不对、查无此人、除数为 0）→ 返回 `error`，由调用方决定怎么处理。
- **不该发生、程序状态已不可信**（下标越界、断言失败、严重逻辑崩坏）→ 用 `panic`，必要时再在上层用 `defer` + `recover` 兜底。

本文用「九九乘法表」同一个业务场景，对比这两种写法的差别。

## 业务错误 vs 程序错误

| 类型     | 常见手段             | 出错后同函数里后面的代码                  | 典型场景                   |
| -------- | -------------------- | ----------------------------------------- | -------------------------- |
| 业务错误 | 返回 `(结果, error)` | **还能继续走**（只要你没 `return`）       | 用户输入非法、业务校验失败 |
| 程序错误 | `panic`              | **不再执行**（即使 `recover` 了也是如此） | 严重异常、不可恢复状态     |

记住口诀：**业务错误是「报告给调用方」；程序错误是「当前这条执行路径断了」。**

## 业务错误：返回 `error`

`error` 是内置接口，只有一个方法 `Error() string`。没出错返回 `nil`，有错返回非 `nil`。

下面要求参数必须是 `9` 才打印九九表；传 `8` 时返回错误。注意：处理完错误之后，`main` 里的「其余逻辑」**仍然会执行**。

```go
package main

import (
	"errors"
	"fmt"
)

// 业务校验失败：返回 error，不中断整个程序
func nn(a int) (bool, error) {
	if a != 9 {
		return false, errors.New("请输入9")
	}

	for i := 1; i <= a; i++ {
		for k := 1; k <= i; k++ {
			fmt.Printf("%d * %d = %d  ", k, i, i*k)
		}
		fmt.Println()
	}
	return true, nil
}

func main() {
	success, err := nn(8)
	if err != nil {
		fmt.Println(err, "执行失败")
	} else {
		fmt.Println(success, "执行成功")
	}

	// 业务错误只是「这次调用失败」，后面的代码照常走
	fmt.Println("其余逻辑")
}
```

输出类似：

```text
请输入9 执行失败
其余逻辑
```

要点：

1. 用 `errors.New("说明")` 或 `fmt.Errorf("a=%d 非法", a)` 创建错误。
2. 调用方标准姿势：`if err != nil { ... }`。
3. 返回错误后，**函数正常结束**；调用方可以打日志、换参数重试，也可以继续干别的事。

## 程序错误：`panic`

把「请输入 9」改成 `panic`。一旦触发，从 `panic` 那一行开始，**同函数里后面的代码不会再跑**。

先看**没有** `recover` 的情况——进程会直接崩溃：

```go
package main

import "fmt"

func nn(a int) (bool, error) {
	if a != 9 {
		panic("请输入9") // 程序错误：直接中断当前执行路径
	}

	for i := 1; i <= a; i++ {
		for k := 1; k <= i; k++ {
			fmt.Printf("%d * %d = %d  ", k, i, i*k)
		}
		fmt.Println()
	}
	return true, nil
}

func main() {
	success, err := nn(8)
	if err != nil {
		fmt.Println(err, "执行失败")
	} else {
		fmt.Println(success, "执行成功")
	}

	// 上面已经 panic，这行永远走不到
	fmt.Println("其余逻辑")
}
```

运行后会看到 panic 栈，且没有「其余逻辑」。这就是和业务错误最大的差别。

## defer 用法

`defer` 会把一次函数调用推迟到 **当前函数即将返回之前** 执行，常用于关文件、解锁、以及配合 `recover` 捕获 `panic`。

基本顺序：先写正常逻辑，`defer` 的内容最后跑。

```go
package main

import "fmt"

func main() {
	fmt.Println("开始")
	defer fmt.Println("defer：收尾")
	fmt.Println("业务逻辑")
}
```

输出：

```text
开始
业务逻辑
defer：收尾
```

多个 `defer` 按 **后进先出（LIFO）** 执行——最后注册的最先跑：

```go
package main

import "fmt"

func main() {
	defer fmt.Println("第一")
	defer fmt.Println("第二")
	defer fmt.Println("第三")
	fmt.Println("函数体")
}
```

输出：

```text
函数体
第三
第二
第一
```

和错误处理相关的关键点：

1. **`recover` 必须写在 `defer` 里**才有效；普通代码里调用 `recover()` 拿不到 panic。
2. `defer` 常用来保证「出错也要收尾」：关文件、释放锁、打印现场。
3. 参数在写 `defer` 那一刻就算好（不是真正执行时才算）。更多细节见「Go 函数」一文的 defer 小节。

## 用 `defer` + `recover` 兜住 panic

加上 `recover` 后，进程**不会崩溃**，可以打印 panic 信息；但 `panic` 发生点之后的代码**仍然不会执行**。

```go
package main

import "fmt"

func nn(a int) (bool, error) {
	if a != 9 {
		panic("请输入9")
	}

	for i := 1; i <= a; i++ {
		for k := 1; k <= i; k++ {
			fmt.Printf("%d * %d = %d  ", k, i, i*k)
		}
		fmt.Println()
	}
	return true, nil
}

func main() {
	defer func() {
		if info := recover(); info != nil {
			fmt.Println("捕获到程序错误:", info)
		}
	}()

	success, err := nn(8) // 这里 panic
	if err != nil {
		fmt.Println(err, "执行失败")
	} else {
		fmt.Println(success, "执行成功")
	}

	// panic 之后：同函数里这些代码依然不走
	fmt.Println("其余逻辑")
}
```

输出类似：

```text
捕获到程序错误: 请输入9
```

没有「执行失败 / 执行成功」，也没有「其余逻辑」。
`recover` 的作用是：**拦住崩溃**，不是**从 panic 下一行接着跑**。

## 怎么选

| 情况                                 | 建议                                             |
| ------------------------------------ | ------------------------------------------------ |
| 用户输错、查不到数据、权限不够       | 返回 `error`                                     |
| 库函数约定用 `error`（如 `os.Open`） | 跟着返回 / 包装 `error`                          |
| 明显的编程错误、不可继续的状态       | 可以 `panic`                                     |
| 服务入口想避免整个进程挂掉           | 边界处 `defer` + `recover`，打日志后返回错误响应 |

入门阶段优先练好 **`if err != nil`**；`panic` / `recover` 留给真正的异常路径，不要拿来代替普通业务校验。

## 作业

1. **返回 error**：写 `divide(a, b int) (int, error)`，`b == 0` 时返回错误，否则返回商。在 `main` 里分别用 `(10, 2)` 和 `(10, 0)` 调用并处理。
2. **业务错误对比**：改写文中的 `nn`，允许传入 `1`～`9` 任意整数打印对应乘法表；小于 `1` 或大于 `9` 时返回 `error`。调用失败后打印一句「其余逻辑」，确认它仍会执行。
3. **defer 顺序**：在 `main` 里注册三个 `defer`，打印 `A`、`B`、`C`，再打印一句「函数体」，观察输出顺序是否为「函数体 → C → B → A」。
4. **panic 与 recover**：写一个函数 `mustPositive(n int)`，`n <= 0` 时 `panic`。在 `main` 里用 `defer` + `recover` 捕获，打印 panic 信息；并在调用后面再写一句 `fmt.Println("其余逻辑")`，验证它不会执行。
5. **综合练习**：写 `printTable(n int) error`（业务错误版）和 `mustPrintTable(n int)`（`n != 9` 就 `panic`）。在 `main` 里先演示返回 `error` 后后续代码仍运行；再用 `recover` 演示 `panic` 被兜住但后续代码不运行。

# 13 值传递与引用类型

调用函数时，实参怎么传给形参？改形参会不会影响外面的变量？
Go 的答案比「值传递 / 引用传递」二分法更精确：**Go 里永远是值传递**——传参时会复制一份。但有些类型的「那一份」只是描述符（header），底层数据仍然共享，所以看起来就像「引用传递」。

## 先记住一句话

> **Go 只有值传递，没有 C++ 那种引用传递（`&`）。**
> 区别不在于语法，而在于**被复制的到底是什么**：复制整个 `int`，还是复制一个指向底层数据的「头」。

| 分类     | 常见类型                                              | 传参时复制什么        | 函数内改形参，外面会变吗       |
| -------- | ----------------------------------------------------- | --------------------- | ------------------------------ |
| 值类型   | `int`、`bool`、`float`、`array`、`struct`（普通字段） | 完整拷贝一份数据      | **不会**（改的是副本）         |
| 引用类型 | `slice`、`map`、`func`、`interface`                   | 复制 header，共享底层 | **会**（改的是同一份底层数据） |
| 指针     | `*T`                                                  | 复制指针值（地址）    | **会**（通过地址改原数据）     |

下面分开看。

## 值传递

值传递时，函数拿到的是实参的**完整副本**。形参和实参是两个变量，互不影响。

```go
package main

import "fmt"

func change(x int) {
	x = 20 // 改的是副本
}

func main() {
	x := 10
	change(x)
	fmt.Println(x) // 10，外面的 x 没变
}
```

数组也是值类型——传参会复制**整个数组**，开销比切片大，日常更常用切片：

```go
package main

import "fmt"

func change(arr [3]int) {
	arr[0] = 100
}

func main() {
	arr := [3]int{1, 2, 3}
	change(arr)
	fmt.Println(arr) // [1 2 3]，原数组没变
}
```

## 引用类型有哪些

Go 没有官方术语「引用类型」，但社区通常把下面几类归为一组——它们的变量里藏着**指向底层数据的指针**：

| 类型        | 说明                                                         |
| ----------- | ------------------------------------------------------------ |
| `slice`     | 切片，可变长序列                                             |
| `map`       | 映射，键值对                                                 |
| `func`      | 函数值，指向同一段代码                                       |
| `interface` | 接口值，内部含类型与数据指针                                 |
| `chan`      | 通道，用于 goroutine 通信（**本节仅了解名字，详细用法见后续专题**） |

下面用 `slice` 和 `map` 举例（入门阶段重点掌握这两个）；`func`、`interface` 会在函数、接口相关章节深入，这里知道「也属于引用类型」即可。

### 切片 slice

传参时复制的是 slice header（指针、长度、容量），**底层数组共享**。所以在函数里改元素，外面看得见：

```go
package main

import "fmt"

func change(s []string) {
	s[0] = "猫来"
}

func main() {
	s := []string{"牛来", "羊来", "码来"}
	change(s)
	fmt.Println(s) // [猫来 羊来 码来]
}
```

用 `append` 追加元素时，只要**没触发扩容**（未超过 `cap`），同样共享底层数组：

```go
package main

import "fmt"

func appendOne(s []int) {
	s = append(s, 100) // 在共享的底层数组上追加
}

func main() {
	s := make([]int, 0, 4) // len=0, cap=4
	appendOne(s)
	fmt.Println(s, len(s)) // [] 0 —— 注意：形参 s 的 len 变了，但 main 里的 s 仍是原来的 header
  fmt.Println(s[:1])     // [100]  ← 手动把 len 撑到 1，就能看到数据了
}
```

上面这个例子要仔细看：`appendOne` 里 `s` 的 `len` 变了，但 `main` 里的 `s` 还是原来的 header（`len` 仍为 0）。
**改已有下标的元素**会生效；**`append` 可能只改形参自己的 header**，外面不一定变长。若 `append` 触发扩容、分配了新数组，则完全与外面无关。

```go
package main

import "fmt"

func appendMany(s []int) {
	s = append(s, 1, 2, 3) // cap 不够时会新分配数组
}

func main() {
	s := make([]int, 0, 1)
	appendMany(s)
	fmt.Println(s, len(s)) // [] 0 —— 扩容后外面完全不受影响
}
```

要点：**改 `s[i]` 共享；`append` 是否影响外面，要看有没有扩容、以及是否把新 slice 返回给调用方。**

### 映射 map

`map` 传参后，内外指向同一张哈希表。改值、删键、增键，外面都看得见：

```go
package main

import "fmt"

func change(m map[string]int) {
	m["语文"] = 800
}

func main() {
	m := map[string]int{
		"语文": 100,
		"数学": 99,
		"英语": 98,
	}
	change(m)
	fmt.Println(m) // map[语文:800 数学:99 英语:98]
}
```

调用 `delete` 等 API，同样共享底层：

```go
package main

import "fmt"

func change(m map[string]int) {
	m["语文"] = 800
	delete(m, "数学")
}

func main() {
	m := map[string]int{
		"语文": 100,
		"数学": 99,
		"英语": 98,
	}
	change(m)
	fmt.Println(m) // map[语文:800 英语:98]
}
```

### 通道 chan（了解即可）

`chan` 和 `slice`、`map` 一样，传参后内外指向同一个通道。
通道用于 goroutine 之间收发数据，用法较多，**本教程不展开**；后续会有单独一篇。这里只要记住：**它也是引用类型之一**。

## 结构体：值传递 vs 指针传递

结构体本身是**值类型**。传 `Person` 会复制整个结构体，函数里改字段不影响外面：

```go
package main

import "fmt"

type Person struct {
	Name  string
	Age   int
	Email string
}

func change(person Person) {
	person.Name = "张满月"
	person.Age = 91
	person.Email = "zmy@qq.com"
}

func main() {
	person := Person{
		Name:  "小满",
		Age:   25,
		Email: "1195566313@qq.com",
	}
	change(person)
	fmt.Println(person) // {小满 25 1195566313@qq.com}，原结构体未变
}
```

若希望函数内修改能反映到外面，传**指针** `*Person`：

```go
package main

import "fmt"

type Person struct {
	Name  string
	Age   int
	Email string
}

func change(person *Person) {
	person.Name = "张满月" // 语法糖，等价于 (*person).Name
	person.Age = 91
	person.Email = "zmy@qq.com"
}

func main() {
	person := Person{
		Name:  "小满",
		Age:   25,
		Email: "1195566313@qq.com",
	}
	change(&person)
	fmt.Println(person) // {张满月 91 zmy@qq.com}
}
```

补充说明：

- `*Person` 是**指针类型**，不是「引用类型」表格里的那一类；传指针时复制的是**地址**，通过地址仍能改到原结构体。
- 结构体里某字段写成 `*string`，表示该字段是指针，和「把整个人当成 `*Person` 传递」是不同层面的概念。入门阶段优先掌握 **`Person` 与 `\*Person` 的区别** 即可。

## 怎么选

| 场景                             | 建议                                                         |
| -------------------------------- | ------------------------------------------------------------ |
| 基本类型、小结构体，函数只读不改 | 直接传值                                                     |
| 需要在函数里改调用方的结构体     | 传 `*Struct`                                                 |
| 一组可变长数据                   | 用 `slice`，注意 `append` 与共享的规则                       |
| 按键查找、统计                   | 用 `map`                                                     |
| 不确定要不要改外面               | 先想清楚「要不要共享」；要共享就指针或引用类型，不要共享就传值或返回新值 |

## 作业

1. **值传递**：写 `setZero(n int)`，在函数内把 `n` 设为 `0`。在 `main` 里传入 `42`，调用前后各打印一次，确认外面仍是 `42`。
2. **切片改元素**：写 `firstToX(s []string)`，把第一个元素改成 `"X"`。在 `main` 里传入 `[]string{"A", "B", "C"}`，调用后打印切片，确认已改变。
3. **map 增删改**：写 `updateScore(scores map[string]int, subject string, delta int)`，给某科加分；若科目不存在则新建。再写 `removeSubject(scores map[string]int, subject string)` 删除一科。在 `main` 里验证增、改、删后 map 的变化。
4. **结构体值 vs 指针**：定义 `Counter struct { Value int }`。分别写 `addByValue(c Counter)` 和 `addByPointer(c *Counter)`，都让 `Value + 1`。在 `main` 里各调用一次，对比哪种方式能让外面的 `Counter` 真的加 1。
5. **综合练习**：用 `map[string][]int` 表示「学生姓名 → 各科成绩列表」。写函数 `addScore(scores map[string][]int, name string, score int)` 追加一门成绩；写 `fixFirst(scores map[string][]int, name string, newScore int)` 把该生第一门课分数改成 `newScore`（若存在）。在 `main` 里多次调用后打印，体会 map 与 slice 嵌套时的共享行为。

# 14 协程

普通函数按顺序执行：前面的不结束，后面的不会开始。
协程（goroutine）让你可以**同时跑多段逻辑**——比如一边查订单、一边发短信、一边写日志，彼此不互相堵着。

Go 的协程由运行时调度，比操作系统线程更轻。入门阶段先记住三件事：

1. 用 `go` 启动协程。
2. `main` 结束，整个程序就结束，其他协程也会被直接掐掉。
3. 多个协程要「等它们都做完」，常用 `sync.WaitGroup`。

> 协程之间怎么**传数据**，主要靠 `channel`。`channel` 内容较多，**下一篇单独讲**；本文只解决「怎么开」和「怎么等」。

## 创建协程

在函数调用前加 `go`，就会在新的协程里执行，**不会阻塞当前代码**：

```go
package main

import (
	"fmt"
	"time"
)

func say(msg string) {
	fmt.Println(msg)
}

func main() {
	go say("你好，协程")
	fmt.Println("你好，main")

	// 先凑合等一下，否则 main 太快结束，协程可能还没打印
	time.Sleep(time.Millisecond * 100)
}
```

输出顺序不一定固定，常见类似：

```text
你好，main
你好，协程
```

也可以直接 `go` 一个匿名函数：

```go
go func() {
	fmt.Println("匿名协程")
}()
```

## 坑：`main` 结束，协程也没了

下面这段**常常什么都打印不出来**（或偶尔能看到一行）：

```go
package main

import "fmt"

func main() {
	go func() {
		fmt.Println("我可能来不及打印")
	}()
	// main 立刻返回，程序退出
}
```

原因：`go` 只是启动，不保证立刻跑完。`main` 一结束，进程就退出，其他协程直接被终止。

### 为什么不能靠 `Sleep` 硬等

多个协程时，每个任务耗时往往**不一样、也不好猜**。你在 `main` 里写 `time.Sleep(xxx)`，时间是自己拍的：

- 睡短了：有的协程还没做完，程序就退出了，结果丢了。
- 睡长了：任务早做完了，还干等着，浪费时间。

所以正式写法不用「自己估延时」，而是用标准库 **`sync`** 包里的 **`WaitGroup`**：谁做完谁报到，全部报到再往下走。

## 用 WaitGroup 等待

`sync.WaitGroup` 像一个计数器：

| 方法     | 作用                                   |
| -------- | -------------------------------------- |
| `Add(n)` | 计数器加 `n`（准备启动几个任务就加几） |
| `Done()` | 计数器减 `1`（一个任务结束就调一次）   |
| `Wait()` | 卡住，直到计数器变为 `0`               |

约定：每个协程结束时调用 `Done()`，通常写成 `defer wg.Done()`，避免漏调。

### 例子：并发处理多个订单

模拟给 3 个订单分别「发货通知」。每个通知耗时不同——如果用 `Sleep` 等，你根本不知道该睡多久；用 `WaitGroup` 就不需要猜。

先看**错误示范**（靠猜延时）：

```go
package main

import (
	"fmt"
	"time"
)

func notify(orderID string, cost time.Duration) {
	time.Sleep(cost) // 这里的 Sleep 只是「模拟业务耗时」，不是用来等协程
	fmt.Println("已通知订单:", orderID)
}

func main() {
	go notify("A1001", 200*time.Millisecond)
	go notify("A1002", 100*time.Millisecond)
	go notify("A1003", 150*time.Millisecond)

	// 自己拍一个 120ms：A1002 能出来，A1001 / A1003 可能来不及
	time.Sleep(120 * time.Millisecond)
	fmt.Println("全部通知完成？——不一定")
}
```

再看**正确写法**（`sync.WaitGroup`）：

`wg` 写在 `main` 里，外面的 `notify` **本来读不到**这个变量——这和普通局部变量一样。所以要把地址 **`&wg` 当参数传进去**，函数签名写成 `wg *sync.WaitGroup`。

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

// 第三个参数接收 WaitGroup 的指针，否则外面读不到 main 里的 wg
func notify(orderID string, cost time.Duration, wg *sync.WaitGroup) {
	defer wg.Done() // 这个协程结束时，计数器 -1

	time.Sleep(cost) // 仅模拟「发短信 / 调接口」耗时，与「等不等协程」无关
	fmt.Println("已通知订单:", orderID)
}

func main() {
	var wg sync.WaitGroup // 定义在 main 里：只有 main 和「被传进去的函数」能用它

	orders := []struct {
		id   string
		cost time.Duration
	}{
		{"A1001", 200 * time.Millisecond},
		{"A1002", 100 * time.Millisecond},
		{"A1003", 150 * time.Millisecond},
	}

	wg.Add(len(orders)) // 告诉 WaitGroup：一共要等 3 个
	for _, o := range orders {
		go notify(o.id, o.cost, &wg) // 传地址，不是传副本
	}

	wg.Wait() // 不猜时间：三个 Done 都调用完才继续
	fmt.Println("全部通知完成")
}
```

如果嫌传参麻烦，也可以把逻辑写在 `main` 里的匿名函数中，直接闭包用到 `wg`（效果一样）：

```go
wg.Add(1)
go func(id string, cost time.Duration) {
	defer wg.Done()
	time.Sleep(cost)
	fmt.Println("已通知订单:", id)
}(o.id, o.cost)
```

两种写法二选一即可：**独立函数 → 必须传 `\*sync.WaitGroup`；匿名函数 → 可直接用外层的 `wg`。**

输出类似（顺序可能变）：

```text
已通知订单: A1002
已通知订单: A1003
已通知订单: A1001
全部通知完成
```

对比记住：

|            | `time.Sleep` 硬等            | `sync.WaitGroup`           |
| ---------- | ---------------------------- | -------------------------- |
| 依据       | 自己定的毫秒数，不准         | 每个协程调用 `Done()` 报到 |
| 任务变慢时 | 容易提前退出、丢结果         | 自动多等一会儿             |
| 任务变快时 | 白白多睡                     | 做完立刻往下走             |
| 适用       | 演示「协程能跑起来」可以凑合 | **多协程正式等待**         |

注意三点：

1. `wg` 在 `main` 里时，外层命名函数**读不到**，必须加参数 `wg *sync.WaitGroup`，调用时传 `&wg`。
2. 一定要传**指针**。若写成 `wg sync.WaitGroup`（值传递），协程里 `Done` 的是副本，`main` 里的 `Wait` 永远等不到。
3. `Add` 在启动协程**之前**调用；`Done` 在协程内部调用。

### 例子：并发算总分

多个学生各有一科成绩，分别在协程里累加到总分（这里用局部变量返回思路，先避免多协程同时写同一个变量）：

```go
package main

import (
	"fmt"
	"sync"
)

func sumScores(name string, scores []int, wg *sync.WaitGroup) {
	defer wg.Done()

	total := 0
	for _, s := range scores {
		total += s
	}
	fmt.Printf("%s 总分: %d\n", name, total)
}

func main() {
	var wg sync.WaitGroup

	students := map[string][]int{
		"小满": {90, 85, 88},
		"张三": {70, 92, 80},
		"李四": {95, 91, 89},
	}

	wg.Add(len(students))
	for name, scores := range students {
		// 把循环变量拷一份再传进协程，避免闭包踩坑
		n, s := name, scores
		go sumScores(n, s, &wg)
	}

	wg.Wait()
	fmt.Println("统计结束")
}
```

> 循环里 `go` 时，务必把 `name`、`scores` **拷到新变量**再传入。否则多个协程可能都用到「最后一次循环」的值（Go 1.22 之前尤其常见）。习惯上写成 `n, s := name, scores` 最稳妥。

## 闭包参数再强调一次

错误示范（容易踩坑）：

```go
for _, name := range names {
	go func() {
		fmt.Println(name) // 可能全是最后一个 name
	}()
}
```

正确写法：把值作为参数传进去：

```go
for _, name := range names {
	go func(n string) {
		fmt.Println(n)
	}(name)
}
```

## 怎么选（入门版）

| 场景                                       | 建议                                                 |
| ------------------------------------------ | ---------------------------------------------------- |
| 只是顺序执行业务                           | 普通函数调用即可，不必上协程                         |
| 多个互不依赖的任务要同时跑                 | `go` + `sync.WaitGroup`                              |
| 想「等全部完成再继续」                     | 用 `wg.Wait()`，**不要**在 `main` 里猜一个 `Sleep`   |
| `Sleep` 可以用在哪                         | 模拟业务耗时（发短信、调接口）；**不**用来等协程结束 |
| 协程之间要传结果、收发消息                 | 用 `channel`（**下篇**）                             |
| 多个协程同时改同一块数据（如同一个 `map`） | 先别这么做；需要时用锁或 `channel` 串行化            |

入门阶段先练熟：**启动 → 别让 main 提前走 → 用 `sync.WaitGroup` 等齐（别靠猜延时）**。共享数据的安全写法放到后续专题。

## 作业

1. **启动协程**：写函数 `hello(name string)` 打印 `你好, xxx`。在 `main` 里用 `go` 调用它，并用 `time.Sleep` 短暂等待后结束（体会「协程能跑起来」即可；这里 Sleep 只是凑合演示）。
2. **猜延时不准**：启动 3 个协程，内部分别 `Sleep` 50ms / 200ms / 100ms 后打印。在 `main` 里只 `Sleep(80ms)` 再打印「结束」。观察哪些输出丢了，用一两句话说明为什么多协程不该靠自己定延时。
3. **WaitGroup 基础**：把上一题改成 `sync.WaitGroup`：三个任务都 `Done` 后，再打印「结束」。确认每次都能看到全部输出。
4. **并发通知**：给定订单号切片 `[]string{"O1", "O2", "O3", "O4"}`，每个订单在协程里 `Sleep` 50 毫秒后打印「处理完成: 订单号」。用 `WaitGroup` 等全部处理完再打印「收银台下班」。
5. **综合练习**：定义函数 `download(file string, wg *sync.WaitGroup)`，内部 `Sleep` 100 毫秒后打印「已下载: 文件名」。文件列表为 `[]string{"a.txt", "b.txt", "c.txt"}`。用循环 + `go` 启动下载，注意循环变量拷贝（或参数传入），最后 `Wait` 并打印「全部下载完成」。



# 15 通道

上一篇讲了怎么用 `go` 启动协程、用 `WaitGroup` 等待。
协程之间要**传数据**，Go 的推荐方式是 **channel（通道）**：一边往里送，一边从里取，比多协程同时改同一个变量更安全、更好懂。

入门先记住四件事：

1. 用 `make(chan T)` 创建通道，`T` 是收发的数据类型。
2. `ch <- v` 发送，`v := <-ch` 接收。
3. **无缓冲**和**有缓冲**行为不一样，卡不卡在这一步。
4. 多个通道同时等，用 `select`。

## 创建与收发

```go
ch := make(chan int) // 无缓冲 int 通道
ch <- 1              // 发送 1
x := <-ch            // 接收到 x
```

也可以只声明发送或接收方向（函数参数里常见）：

```go
func producer(out chan<- int) { // 只能发送
	out <- 1
}

func consumer(in <-chan int) { // 只能接收
	fmt.Println(<-in)
}
```

## 无缓冲通道

`make(chan T)` 不写容量，就是**无缓冲**：

- 发送方 `ch <- v` 会**卡住**，直到有人来接收。
- 接收方 `<-ch` 也会**卡住**，直到有人来发送。

两边必须「对上」才能继续——像当面递东西：你递、对方接，动作同时完成。

```go
package main

import "fmt"

func main() {
	ch := make(chan string)

	go func() {
		ch <- "订单已支付" // 没人接时，这里会一直等
	}()

	msg := <-ch // main 在这里接住
	fmt.Println(msg)
}
```

输出：

```text
订单已支付
```

注意：如果**只有发送、永远没有接收**（或反过来），程序会**死锁**，运行时报 `fatal error: all goroutines are asleep - deadlock!`。

### 例子：协程算完，通过通道把结果交回 main

```go
package main

import "fmt"

func sum(nums []int, ch chan int) {
	total := 0
	for _, n := range nums {
		total += n
	}
	ch <- total // 把结果送出去
}

func main() {
	ch := make(chan int)
	go sum([]int{1, 2, 3, 4}, ch)

	result := <-ch // 等协程送来结果
	fmt.Println("总和:", result)
}
```

无缓冲通道本身就能「等对方」：**接收会阻塞到有发送**，所以这里不必再写 `WaitGroup`。

## 有缓冲通道

`make(chan T, n)` 的 `n` 是缓冲区容量：

- 缓冲区**没满**时，发送可以立刻成功，不必等接收方。
- 缓冲区**空**时，接收会卡住，等有人发送。
- 缓冲区**满了**再发，才会卡住等接收腾出空位。

像信箱：最多放 `n` 封信；没满就能投，空了再取就得等。

```go
package main

import "fmt"

func main() {
	ch := make(chan int, 2) // 容量 2

	ch <- 10 // 不阻塞
	ch <- 20 // 不阻塞
	// ch <- 30 // 再写就会卡住（缓冲区已满），除非另有协程在收

	fmt.Println(<-ch) // 10
	fmt.Println(<-ch) // 20
}
```

对比记住：

|              | 无缓冲 `make(chan T)` | 有缓冲 `make(chan T, n)` |
| ------------ | --------------------- | ------------------------ |
| 发送何时不卡 | 必须立刻有人接收      | 缓冲区还有空位           |
| 接收何时不卡 | 必须立刻有人发送      | 缓冲区里已有数据         |
| 典型用途     | 同步、交接结果        | 削峰、批量暂存           |

入门选型：

- 要「算完立刻交回 / 两边对齐」→ 无缓冲够用。
- 生产者偶尔比消费者快一点，想先攒几条 → 有缓冲。

## 关闭通道与 range

发送方不再发时，应 `close(ch)`。接收方可以用 `range` 一直取，直到通道关闭：

```go
package main

import "fmt"

func main() {
	ch := make(chan int, 3)

	ch <- 1
	ch <- 2
	ch <- 3
	close(ch) // 关闭后不能再发送；还能继续接收剩余数据

	for v := range ch {
		fmt.Println(v)
	}
}
```

约定：

1. **谁发送，谁关闭**；接收方一般不要 `close`。
2. 对已关闭的通道再发送会 **panic**。
3. 从已关闭且空的通道接收，会立刻拿到类型零值；可用 `v, ok := <-ch`，`ok == false` 表示已关闭且没有数据了。

```go
v, ok := <-ch
if !ok {
	fmt.Println("通道已关闭")
}
```

## select：同时等多个通道

`select` 像通道版的 `switch`：哪个 case 能立刻收发，就走哪个；多个都能走时**随机选一个**。

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(50 * time.Millisecond)
		ch1 <- "来自 ch1"
	}()
	go func() {
		time.Sleep(30 * time.Millisecond)
		ch2 <- "来自 ch2"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg := <-ch1:
			fmt.Println(msg)
		case msg := <-ch2:
			fmt.Println(msg)
		}
	}
}
```

常见输出（`ch2` 先就绪）：

```text
来自 ch2
来自 ch1
```

### default：非阻塞尝试

所有 case 都暂时走不了时，如果写了 `default`，就立刻执行 `default`，**不会卡死**：

```go
select {
case msg := <-ch:
	fmt.Println("收到:", msg)
default:
	fmt.Println("暂时没有数据")
}
```

适合「试一下有没有消息，没有就干别的」。

### 例子：带超时

用 `time.After` 做一个超时分支，避免一直等某个通道：

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	go func() {
		time.Sleep(200 * time.Millisecond)
		ch <- "慢请求的结果"
	}()

	select {
	case msg := <-ch:
		fmt.Println(msg)
	case <-time.After(100 * time.Millisecond):
		fmt.Println("超时了")
	}
}
```

输出：

```text
超时了
```

（把超时改成 `300ms`，就能收到「慢请求的结果」。）

## 怎么选（入门版）

| 场景                               | 建议                               |
| ---------------------------------- | ---------------------------------- |
| 协程算完把结果交回                 | 无缓冲通道，`main` 里 `<-ch`       |
| 多个结果依次送出                   | 发送完 `close`，接收端 `for range` |
| 想先攒几条再处理                   | 有缓冲 `make(chan T, n)`           |
| 同时等好几个通道 / 要超时          | `select`                           |
| 试一下有没有数据，没有就继续       | `select` + `default`               |
| 多协程只是「等全部做完」、不传数据 | 继续用上一篇的 `WaitGroup`         |

入门阶段先练熟：**创建 → 无缓冲/有缓冲区别 → 关闭与 range → select 多路与超时**。

## 作业

1. **无缓冲交接**：写函数 `double(n int, ch chan int)`，把 `n * 2` 发送到通道。在 `main` 里 `go double(21, ch)`，再接收并打印结果（应是 `42`）。
2. **体会阻塞**：无缓冲通道上，先在 `main` 里发送一个值、再启动协程接收——观察是否死锁。改成「先 `go` 接收（或先 `go` 发送），再在另一边操作」，说明无缓冲为什么要两边配合。
3. **有缓冲**：创建容量为 `2` 的 `chan string`，连续发送 `"A"`、`"B"`（不启动其他协程），再依次接收打印。思考：若再发送第三个而不先接收，会发生什么？
4. **关闭与 range**：启动一个协程，向通道依次发送 `1..5` 后 `close`；`main` 里用 `for range` 打印所有数字。
5. **select 二选一**：两个协程分别向 `ch1`、`ch2` 发送字符串（可加不同的 `Sleep`）。用 `select` 接收两次并打印，观察谁先到。
6. **综合：超时查询**：模拟查询：协程 `Sleep` 150ms 后往通道送 `"查到了"`。`main` 用 `select`：100ms 内收到就打印结果，否则打印「查询超时」。再把超时改成 200ms，确认能收到结果。

# 16 互斥锁与读写锁

上一篇用 **channel** 在协程之间**传数据**。
有时协程要**共改同一块内存**——比如几个协程一起往同一个切片里 `append`。这时通道不是唯一答案，标准库 `sync` 里的**锁**也很常用。

入门先记住三件事：

1. **互斥锁 `Mutex`**：同一时刻只允许一个协程进「临界区」改共享数据。
2. 不加锁并发改切片 / map，可能丢数据、甚至崩溃（数据竞争）。
3. **读写锁 `RWMutex`**：读可以多人同时读；写还是独占。适合「读多写少」。

> 口诀对照：channel 偏「把结果递出去」；锁偏「大家轮流改同一份」。

## 为什么要锁：奶茶菜单被同时加料

场景：菜单里已有三杯奶茶，三个协程各自再加一杯。
共享的是同一个切片 `arr`，改它时要排队。

### 加了互斥锁：每次都能累加成功

```go
package main

import (
	"fmt"
	"sync"
)

func worker(arr *[]string, wg *sync.WaitGroup, mu *sync.Mutex, tea string) {
	defer wg.Done()
	mu.Lock()                 // 上锁：别人先别动 arr
	*arr = append(*arr, tea)  // 临界区：只有拿到锁的协程能改
	mu.Unlock()               // 解锁：下一位可以进来
}

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	arr := []string{"芋泥波波", "绿茶", "红茶"}

	wg.Add(3)
	go worker(&arr, &wg, &mu, "柠檬茶")
	go worker(&arr, &wg, &mu, "霸王茶姬")
	go worker(&arr, &wg, &mu, "爷爷不泡茶")
	wg.Wait()

	fmt.Println(arr)
}
```

多次运行，切片里都会有 **6 个名字**（原来 3 个 + 新加 3 个）。顺序可能每次不同（谁先抢到锁谁先加），但**不会少**：

```text
[芋泥波波 绿茶 红茶 柠檬茶 霸王茶姬 爷爷不泡茶]
```

或：

```text
[芋泥波波 绿茶 红茶 霸王茶姬 柠檬茶 爷爷不泡茶]
```

要点：

| 写法                   | 含义                                                  |
| ---------------------- | ----------------------------------------------------- |
| `mu.Lock()`            | 拿到锁；别人正在锁里时，这里会**卡住等待**            |
| `mu.Unlock()`          | 放开锁；漏写会让别的协程永远等                        |
| `*arr`                 | 传的是切片指针，改的才是 `main` 里那份                |
| `wg` / `mu` 也要传指针 | 和上一篇 WaitGroup 一样，传值等于改副本，外面感知不到 |

`defer wg.Done()` 放在函数开头，保证协程怎么退出都会报到；锁则建议**尽快 Unlock**，不要把无关的慢操作包在锁里。

### 不加锁会怎样：看起来像「被盖掉」

把 `Lock` / `Unlock` 拿掉，三个协程同时 `append` 同一份切片：

```go
func worker(arr *[]string, wg *sync.WaitGroup, tea string) {
	defer wg.Done()
	*arr = append(*arr, tea) // 没有锁，多人同时改
}
```

`append` 内部要读长度、扩容、写元素——几步叠在一起。多协程交错执行时，可能出现：

- 有的茶**没加进去**（少了几杯，像被覆盖）
- 偶发 **panic**（例如索引越界）
- 用 `go run -race` 会报 **data race**

所以入门结论很简单：**多人同时改同一份切片 / map，先加锁（或改成只用 channel 串行化）。**

## 读写锁：读多写少时更合适

`Mutex` 读写都互斥：有人读菜单时，别人也不能读。
若业务是「很多人都在看菜单，偶尔改一次」，可以用 **`sync.RWMutex`**：

- `RLock` / `RUnlock`：加**读锁**，多个读可以同时进行。
- `Lock` / `Unlock`：加**写锁**，写的时候别人既不能读也不能写。

下面继续用菜单：两个协程读、一个协程写。
写这边故意 `Sleep(1ms)`，先让读锁先拿到，方便体会「读可以并行、写要等读完」。

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

func read(arr *[]string, wg *sync.WaitGroup, mu *sync.RWMutex) {
	defer wg.Done()
	mu.RLock()
	fmt.Println(*arr)
	time.Sleep(20 * time.Millisecond) // 模拟慢慢读
	mu.RUnlock()
}

func write(arr *[]string, wg *sync.WaitGroup, mu *sync.RWMutex, tea string) {
	defer wg.Done()
	time.Sleep(1 * time.Millisecond) // 先让读锁触发
	mu.Lock()
	time.Sleep(10 * time.Millisecond) // 模拟写入耗时
	*arr = append(*arr, tea)
	mu.Unlock()
}

func main() {
	var wg sync.WaitGroup
	var rwmu sync.RWMutex
	menu := []string{"芋泥波波", "绿茶", "红茶"}

	wg.Add(3)
	go read(&menu, &wg, &rwmu)
	go read(&menu, &wg, &rwmu)
	go write(&menu, &wg, &rwmu, "爷爷不泡茶")
	wg.Wait()

	fmt.Println("最终菜单", menu)
}
```

可能的输出（两个读可以重叠，所以往往先看到两次「还没加料」的菜单；写在读锁释放后才能完成）：

```text
[芋泥波波 绿茶 红茶]
[芋泥波波 绿茶 红茶]
最终菜单 [芋泥波波 绿茶 红茶 爷爷不泡茶]
```

要点：

| 写法                       | 含义                                             |
| -------------------------- | ------------------------------------------------ |
| `mu.RLock()` / `RUnlock()` | 读锁；多个 `read` 可以同时拿着                   |
| `mu.Lock()` / `Unlock()`   | 写锁；要等读锁都放开，且期间别人读不了、也写不了 |
| `write` 里先 `Sleep(1ms)`  | 演示用：让读先执行；真实代码一般不需要这样凑时序 |

记住：

- 只读 → `RLock` / `RUnlock`
- 要改 → 还是 `Lock` / `Unlock`（写锁）
- 入门阶段：**不确定就用 `Mutex`**；确认读远多于写，再换成 `RWMutex`

## 和 channel 怎么选（入门版）

| 场景                                  | 建议                    |
| ------------------------------------- | ----------------------- |
| 协程算完，把结果交回                  | channel                 |
| 多个结果依次汇总、或流水线            | channel                 |
| 多个协程共改同一个切片 / map / 计数器 | `Mutex`（或 `RWMutex`） |
| 读很多、写很少的共享数据              | 优先考虑 `RWMutex`      |
| 只是等任务做完、不改共享数据          | 继续用 `WaitGroup` 即可 |

Go 社区常说「不要通过共享内存来通信，而要通过通信来共享内存」——意思是**优先 channel**。
但计数器、缓存、就地改结构体字段这类场景，**加锁往往更直接**。两种都会用，按场景选。

## 作业

1. **复现丢数据**：用本文「不加锁」版本跑十几次，观察切片长度是否总是 6；有条件的话用 `go run -race .` 看是否报 data race。
2. **加上 Mutex**：恢复 `Lock` / `Unlock`，确认每次长度都是 6，并打印出完整菜单。
3. **计数器**：多个协程给同一个 `var n int` 各加 1000 次。分别写「无锁」和「`Mutex` 保护 `n++`」两版，对比最终 `n` 是否等于期望值。
4. **读写锁小练习**：一个 `map[string]int` 表示库存。启动 3 个协程用 `RLock` 打印某商品库存；再启动 1 个协程用写锁把该商品库存 `+1`。用 `WaitGroup` 等全部结束。

# 17 GMP 协程调度模型

前面几篇讲了怎么用 `go` 开协程、用 channel / 锁协作。
这篇往下挖一层：**这些协程到底是谁在跑、怎么排到 CPU 上的**——也就是 Go 运行时的 **GMP 调度模型**。

入门先记住四件事：

1. **G**：Goroutine，一个待执行的任务。
2. **M**：Machine，操作系统内核线程，真正干活的人。
3. **P**：Processor，逻辑处理器，手里攥着一批 G，并绑定到某个 M 上跑。
4. Go 尽量在**用户态**切换 G，所以比直接开大量 OS 线程轻得多。

> 口诀：G 是任务，M 是工人，P 是工位。工人必须坐到工位上，才能领任务干活。

![GMP 三者关系](https://go-docs.pages.dev/post/gmp-overview.svg)

## 为什么要有 GMP：直接用 OS 线程太贵

如果并发全靠**操作系统原生线程**：

- 一个线程栈通常大约 **1MB**，开多了内存先顶不住。
- 线程切换要进**内核态**，上下文切换开销大。
- 想支撑十几万、上百万并发，用「一任务一线程」基本不现实。

Go 的做法是自己管一批轻量协程（G），再映射到少量 OS 线程（M）上跑：

- 切换多在**用户态**完成，不用每次都进内核。
- G 的初始栈大约只有 **2KB**，还带着**栈扩缩容**：用多少占多少，大任务再长大，用完可再缩。
- 所以同样硬件上，轻松扛住**十几万甚至百万级**并发，这就是 GMP 要解决的痛点。

## G、M、P 分别是什么

| 缩写  | 全称      | 可以怎么理解                                 |
| ----- | --------- | -------------------------------------------- |
| **G** | Goroutine | 协程 / 任务：`go func()` 出来的就是一个 G    |
| **M** | Machine   | OS 内核线程：真正执行机器码的人              |
| **P** | Processor | 逻辑处理器：管理一批 G，并决定「下一个跑谁」 |

几个关键关系：

- **M 必须绑定 P，才能执行 G**。没有 P，M 就算醒着也拿不到本地任务队列。
- **P 的数量默认等于 CPU 逻辑核数**（由 `GOMAXPROCS` 控制）。核有几个，默认就准备几个「工位」，尽量把多核吃满。
- G 可以远多于 M 和 P：任务很多没关系，轮流上工位执行就行。

可以脑补成奶茶店：

- **G** = 一杯待做的订单
- **P** = 一个操作台（台上最多排队一批订单）
- **M** = 一个店员
  店员必须站到某个操作台前，才能按台上的订单做茶。

## 旧模型的问题：没有 P，大家抢一把全局锁

更早的调度（可以记成 **GM 模型**）大致是：

- 只有一个**全局队列**装着所有待跑的 G
- 多个 M 直接去这个全局队列里领任务

多核 CPU 上一忙起来就会难受：

- 多个 M 同时抢同一条全局队列
- 队列本身要加锁保护
- 抢锁、放锁变成常态，**锁竞争很重**，调度本身反而拖慢业务

所以后来改进成现在的 **GMP**：加上 **P** 这一层，把「抢全局」改成「大部分时候各忙各的」。

![旧 GM 与新 GMP 对比](https://go-docs.pages.dev/post/gmp-old-vs-new.svg)

## 改进后的 GMP：每个 P 有自己的本地队列

现在的主流程可以记成：

1. 每个 **M 绑定一个 P**
2. 每个 **P 管一个本地队列**（容量常见是 **256** 个 G）
3. 新 G 优先放进当前 P 的本地队列
4. 本地队列满了，再往**全局队列**放（或把一部分挪到全局）

好处很直观：

- 多数时候，M 只从**自己绑的 P** 的本地队列取 G
- 本地取任务基本不用和别的核抢同一把大锁
- 多核可以真正并行调度，性能上去了

> 不是「完全没有锁」，而是把热点从「所有人抢一个全局队列」拆开了：本地为主，全局为辅。

## 本地空了、全局也空了：Work Stealing（工作窃取）

那如果某个 P：

- 本地队列已经跑完了
- 全局队列也空了

这个核就要闲着吗？不会。运行时会触发 **Work Stealing（工作窃取）**：

- 当前 P 去看别的 P 的本地队列
- 若别人还有积压，就**偷走大约一半**任务到自己这边
- 目的很简单：**别让 CPU 空转**，有活大家一起干

可以继续用奶茶店类比：自己台子上的单做完了，别的台子还堆着，就过去搬一半回来做——总比干站着强。

![Work Stealing 工作窃取](https://go-docs.pages.dev/post/gmp-work-stealing.svg)

## M 卡住了怎么办：Hand-off（交接）

还有一种常见情况：某个 M 正在跑 G，结果这个 G **阻塞**了（比如陷入会卡住线程的系统调用）。

这时如果还死抱着 P，那个「工位」就浪费了：核还在，却派不出新活。

所以会触发 **Hand-off（交接）**：

1. 当前阻塞的 **M 和 P 立刻解绑**
2. **P 马上去找空闲的 M** 接手
3. 若没有空闲 M，就**再创建一个 M**
4. 新的（或空闲的）M 绑上这个 P，继续跑队列里剩下的 G

这样多核不会因为「某个线程卡住」而整段闲置。等原来那个阻塞结束，M 还可以再找机会绑回某个 P 继续干活。

> 补充一句入门口径：网络 I/O 很多会被运行时的 **netpoller** 做成「看起来阻塞、其实不白白占死线程」；这里说的 Hand-off，重点记「线程级阻塞时，P 会交出去」即可。

![Hand-off 交接](https://go-docs.pages.dev/post/gmp-handoff.svg)

## 串起来看一眼完整链路

一次比较典型的调度路径：

1. `go xxx` → 产生一个 **G**
2. G 进当前 **P** 的本地队列（满了再进全局队列）
3. **M** 绑定 **P**，从本地队列取 G 执行
4. 本地空了 → 先看全局队列；还没有 → **Work Stealing** 偷别人一半
5. 执行中线程阻塞 → **Hand-off**：P 换人（换 M）继续跑

![GMP 完整调度链路](https://go-docs.pages.dev/post/gmp-flow.svg)

所以你写业务时只看到 `go`，底下其实是：

**很多轻量 G → 少量 P（约等于核数）→ 再映射到若干 M（OS 线程）**

这也是 Go 能「代码写起来像随便开协程，跑起来又相对省」的原因。

## 和前面几篇怎么对上

| 你在业务里写的         | 在 GMP 里大致对应                    |
| ---------------------- | ------------------------------------ |
| `go func() { ... }()`  | 创建一个 **G**                       |
| 同时跑很多协程也不慌   | G 很轻，初始栈小，用户态切换         |
| `GOMAXPROCS`           | 大致控制 **P** 的个数                |
| 协程里做很重的阻塞调用 | 可能让 **M** 卡住，触发 **Hand-off** |
| channel / 锁让出执行权 | 当前 G 可以停下，P/M 去跑别的 G      |

入门阶段不必把源码细节背下来。先抓住这句就够：

**G 是任务，P 是工位，M 是工人；本地队列减少抢锁，偷工避免空闲，交接避免阻塞拖死多核。**

# 18 文件与目录操作

日常写服务、写脚本，免不了和磁盘打交道：建文件、写内容、读出来、删掉，再进阶一点就是目录。
这些能力主要在标准库 **`os`**（以及一点 `io`）里。

入门先记住五件事：

1. 简单读写可用 `os.WriteFile` / `os.ReadFile`。
2. 更精细的控制用 **`os.OpenFile`**（flag + 权限）。
3. 文件用完要 **`Close`**，习惯写 `defer f.Close()`。
4. Unix 权限用八进制，如 `0644`、`0755`；**在 Windows 上这些权限位基本无效**。
5. 目录用 `Mkdir` / `MkdirAll`，删目录用 `Remove` / `RemoveAll`。

> 口诀：简单读写用 `ReadFile` / `WriteFile`；要追加、排他创建、自定义权限，再上 `OpenFile`。

## 创建文件

最直接：`os.Create`。文件不存在就创建；**已存在会清空成 0 字节**（相当于截断）。

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("note.txt")
	if err != nil {
		fmt.Println("创建失败:", err)
		return
	}
	defer f.Close()

	fmt.Println("创建成功:", f.Name())
}
```

要点：

| 写法              | 含义                               |
| ----------------- | ---------------------------------- |
| `os.Create(name)` | 创建或截断文件，返回 `*os.File`    |
| `defer f.Close()` | 函数结束时关文件，避免句柄泄漏     |
| `err != nil`      | 磁盘满、没权限、路径非法时都会报错 |

> `Create` 底层大致等于：`OpenFile(name, O_RDWR|O_CREATE|O_TRUNC, 0666)`。

## 写入内容

### 方式一：拿到 `*File` 再写

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("note.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	n, err := f.WriteString("你好，小满\n")
	if err != nil {
		panic(err)
	}
	fmt.Println("写入字节数:", n)

	_, err = f.Write([]byte("第二行：奶茶菜单\n"))
	if err != nil {
		panic(err)
	}
}
```

| 写法             | 含义             |
| ---------------- | ---------------- |
| `WriteString(s)` | 写字符串         |
| `Write(b)`       | 写字节切片       |
| 返回值 `n`       | 实际写入的字节数 |

### 方式二：一行写完（更常用）

```go
err := os.WriteFile("note.txt", []byte("你好，小满\n"), 0644)
if err != nil {
	panic(err)
}
```

`WriteFile` 会：创建/截断 → 写入 → 关闭。适合「整份内容一次写完」。

## 读取详情

### 一次读完整文件

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	data, err := os.ReadFile("note.txt")
	if err != nil {
		fmt.Println("读取失败:", err)
		return
	}
	fmt.Println(string(data))
}
```

### 打开后再读（适合大文件、边读边处理）

```go
package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	f, err := os.Open("note.txt") // 只读打开；不存在会报错
	if err != nil {
		panic(err)
	}
	defer f.Close()

	buf := make([]byte, 64)
	for {
		n, err := f.Read(buf)
		if n > 0 {
			fmt.Print(string(buf[:n]))
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
	}
}
```

### 看文件元信息：`Stat`

```go
info, err := os.Stat("note.txt")
if err != nil {
	panic(err)
}
fmt.Println("名字:", info.Name())
fmt.Println("大小:", info.Size(), "字节")
fmt.Println("是目录吗:", info.IsDir())
fmt.Println("权限:", info.Mode())
fmt.Println("修改时间:", info.ModTime())
```

| API           | 用途                           |
| ------------- | ------------------------------ |
| `os.ReadFile` | 小文件一次性读完               |
| `os.Open`     | 只读打开，再 `Read`            |
| `os.Stat`     | 大小、权限、是否目录、修改时间 |

## 删除文件

```go
err := os.Remove("note.txt")
if err != nil {
	fmt.Println("删除失败:", err)
	return
}
fmt.Println("已删除")
```

| API                  | 行为                                   |
| -------------------- | -------------------------------------- |
| `os.Remove(path)`    | 删文件；也可删**空目录**               |
| `os.RemoveAll(path)` | 递归删除：文件、非空目录都能删（慎用） |

文件不存在时，`Remove` 会返回错误（可用 `os.IsNotExist(err)` 判断）。

------

## 高级：`os.OpenFile`

当「只创建 / 只读 / 只写」不够用时，上 **`OpenFile`**：

```go
func OpenFile(name string, flag int, perm FileMode) (*File, error)
```

三个参数：

1. **`name`**：路径
2. **`flag`**：怎么打开（可读？追加？不存在是否创建？）——常用常量按位或 `|` 组合
3. **`perm`**：新建文件时的权限（如 `0644`）——**主要对 Unix/Linux/macOS 有意义**

示例：追加写入（文件不存在就创建）：

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.OpenFile("log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	_, err = f.WriteString("又来了一单奶茶\n")
	if err != nil {
		panic(err)
	}
	fmt.Println("追加成功")
}
```

### flag 常用枚举（可 `|` 组合）

| 常量          | 含义                                                   |
| ------------- | ------------------------------------------------------ |
| `os.O_RDONLY` | 只读                                                   |
| `os.O_WRONLY` | 只写                                                   |
| `os.O_RDWR`   | 读写                                                   |
| `os.O_CREATE` | 不存在则创建                                           |
| `os.O_APPEND` | 写的时候追加到末尾                                     |
| `os.O_TRUNC`  | 打开时把文件截断为 0 字节                              |
| `os.O_EXCL`   | 与 `O_CREATE` 一起用：文件**已存在则失败**（排他创建） |
| `os.O_SYNC`   | 尽量同步写到磁盘（更安全，更慢）                       |

读写权限类（`O_RDONLY` / `O_WRONLY` / `O_RDWR`）通常**三选一**；其它用 `|` 叠上去：

```go
// 读写 + 不存在则创建 + 追加
os.O_RDWR | os.O_CREATE | os.O_APPEND

// 只写 + 不存在则创建 + 每次打开清空（覆盖写）
os.O_WRONLY | os.O_CREATE | os.O_TRUNC

// 创建但已存在就报错（防止误覆盖）
os.O_WRONLY | os.O_CREATE | os.O_EXCL
```

和快捷 API 的对应关系（帮助记忆）：

| 快捷写法          | 大致等价                                        |
| ----------------- | ----------------------------------------------- |
| `os.Open(name)`   | `OpenFile(name, O_RDONLY, 0)`                   |
| `os.Create(name)` | `OpenFile(name, O_RDWR|O_CREATE|O_TRUNC, 0666)` |

### 权限 `perm`：`0777` / `0644` 到底什么意思

`perm` 是 **Unix 文件模式**，Go 里常写成**八进制**字面量：以 **`0` 开头**，如 `0644`、`0755`、`0777`。

#### 先记三个基本位：`r` `w` `x`

每一位权限用一个数字相加（不是 4、3、2）：

| 符号  | 含义                                   | 数值  |
| ----- | -------------------------------------- | ----- |
| **r** | read 读                                | **4** |
| **w** | write 写                               | **2** |
| **x** | execute 执行（对目录表示「能否进入」） | **1** |

所以一个人的权限可以是：

| 组合  | 计算        | 常见读法   |
| ----- | ----------- | ---------- |
| `---` | 0           | 什么都不能 |
| `--x` | 1           | 只能执行   |
| `-w-` | 2           | 只能写     |
| `-wx` | 2+1=3       | 写 + 执行  |
| `r--` | 4           | 只能读     |
| `r-x` | 4+1=5       | 读 + 执行  |
| `rw-` | 4+2=6       | 读 + 写    |
| `rwx` | 4+2+1=**7** | 全部可以   |

#### 再记三段人：所有者、组、其他人

一个完整权限是 **三组** `rwx`，从左到右：

```text
0  7  7  7
│  │  │  └─ 其他人（other）
│  │  └──── 所属组（group）
│  └─────── 所有者（owner / user）
└────────── 八进制前缀（Go 字面量写法）
```

| 位置           | 是谁                               |
| -------------- | ---------------------------------- |
| 左数第一位数字 | **文件所有者**（创建者账号）       |
| 中间一位       | **所属组**（同一组的用户）         |
| 右边一位       | **其他人**（既不是主人也不是同组） |

例子：

| 权限   | 拆解                                           | 含义                                                |
| ------ | ---------------------------------------------- | --------------------------------------------------- |
| `0644` | owner=`rw-`(6)，group=`r--`(4)，other=`r--`(4) | 自己可读写；别人只读。**文件很常见**                |
| `0755` | owner=`rwx`(7)，group=`r-x`(5)，other=`r-x`(5) | 自己全能；别人可读可进。**目录 / 可执行文件很常见** |
| `0777` | 三段都是 `rwx`(7)                              | **所有人都能读、写、执行**——方便但危险，生产少用    |
| `0600` | 只有 owner=`rw-`                               | 只有自己能读写，适合密钥、token                     |

对照 Linux 上常见显示：

```text
-rw-r--r--   →  0644
drwxr-xr-x   →  0755
-rwxrwxrwx   →  0777
```

最左边的 `-` / `d` 表示普通文件还是目录，不算在三位数字里。

#### 重要：在 Windows 上这些权限基本无效

Go 在 Windows 上**仍然要你传 `perm`**，但系统**不会按 Unix 那套 owner/group/other 去落实**。

入门可以记：

- 写跨平台代码时，`0644`、`0755` 照常写，Linux/macOS 上生效。
- 到了 **Windows**，别指望 `0777`、`0644` 真的控制「谁能读谁能写」。
- Windows 更认自己的 ACL 权限模型；Go 的 `perm` 在这里往往被忽略或只产生很有限影响。

所以：
**权限位是 Unix 世界的规则；Windows 上别拿 `0777` 当安全配置。**

------

## 文件夹（目录）操作

文件会了，目录也是同一套 `os` 思路。

### 创建目录

```go
// 只创建一层；父目录不存在会失败
err := os.Mkdir("data", 0755)

// 递归创建：a/b/c 一层层建好
err = os.MkdirAll("data/logs/2026", 0755)
```

| API                       | 行为                         |
| ------------------------- | ---------------------------- |
| `os.Mkdir(path, perm)`    | 创建单层目录                 |
| `os.MkdirAll(path, perm)` | 递归创建；已存在则通常不报错 |

目录权限常用 **`0755`**：自己可进可改，别人可进可列，不能随便改。

### 读目录里有什么

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	entries, err := os.ReadDir(".")
	if err != nil {
		panic(err)
	}
	for _, e := range entries {
		kind := "文件"
		if e.IsDir() {
			kind = "目录"
		}
		fmt.Println(kind, e.Name())
	}
}
```

`ReadDir` 返回 `[]fs.DirEntry`，比老的 `Readdir` 更轻量，入门优先用它。

### 删除 / 重命名 / 判断存在

```go
// 删空目录（里面还有东西会失败）
os.Remove("data")

// 递归删整个目录树（危险，先确认路径）
os.RemoveAll("data")

// 重命名或移动
os.Rename("note.txt", "note.bak.txt")

// 判断路径是否存在
_, err := os.Stat("note.txt")
if os.IsNotExist(err) {
	fmt.Println("不存在")
} else if err != nil {
	fmt.Println("其它错误:", err)
} else {
	fmt.Println("存在")
}
```

| API                         | 用途               |
| --------------------------- | ------------------ |
| `os.Remove`                 | 删文件或**空**目录 |
| `os.RemoveAll`              | 递归删除           |
| `os.Rename`                 | 改名 / 移动        |
| `os.Stat` + `os.IsNotExist` | 判断在不在         |

### 小综合：建目录 → 写文件 → 读回来 → 清理

```go
package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	dir := "demo_data"
	if err := os.MkdirAll(dir, 0755); err != nil {
		panic(err)
	}

	path := filepath.Join(dir, "menu.txt")
	content := []byte("芋泥波波\n绿茶\n红茶\n")

	if err := os.WriteFile(path, content, 0644); err != nil {
		panic(err)
	}

	// 追加一行
	f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	_, _ = f.WriteString("爷爷不泡茶\n")
	f.Close()

	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	fmt.Println("文件内容:\n" + string(data))

	// 演示完清理
	_ = os.RemoveAll(dir)
}
```

`path/filepath.Join` 会按当前系统拼路径（Windows 用 `\`，Linux 用 `/`），比手写 `"a/" + "b"` 更稳。

## 速查表

| 你想做的事        | 优先用                                            |
| ----------------- | ------------------------------------------------- |
| 整文件写入        | `os.WriteFile`                                    |
| 整文件读取        | `os.ReadFile`                                     |
| 追加日志          | `OpenFile(..., O_CREATE|O_WRONLY|O_APPEND, 0644)` |
| 防止覆盖已有文件  | `O_CREATE|O_EXCL`                                 |
| 看大小 / 是否目录 | `os.Stat`                                         |
| 删一个文件        | `os.Remove`                                       |
| 建多层目录        | `os.MkdirAll`                                     |
| 列目录            | `os.ReadDir`                                      |
| 删整棵目录树      | `os.RemoveAll`（慎用）                            |

入门阶段先把 **`WriteFile` / `ReadFile` / `Remove` / `MkdirAll` / `OpenFile+APPEND`** 练熟；权限先记住文件 `0644`、目录 `0755`，并清楚 **Windows 上 Unix 权限位基本当摆设**。

## 作业

1. 用 `WriteFile` 写一个 `hello.txt`，再用 `ReadFile` 打印出来。
2. 用 `OpenFile` + `O_APPEND` 往同一文件追加 3 行，确认不会覆盖旧内容。
3. 用 `O_CREATE|O_EXCL` 创建文件两次：第二次应失败。
4. 手算并说明：`0755`、`0644`、`0777` 对 owner / group / other 分别意味着什么。
5. `MkdirAll("a/b/c", 0755)`，在 `c` 下写文件，`ReadDir` 列出，最后 `RemoveAll("a")` 清理。

# 19 JSON 序列化与反序列化

写接口、存配置、跟前端或别的服务交换数据，几乎都会碰到 **JSON**。
Go 标准库 **`encoding/json`** 负责两件事：

- **序列化（Marshal）**：把 Go 的值变成 JSON 字节（`[]byte`）
- **反序列化（Unmarshal）**：把 JSON 字节解析回 Go 的值

入门先记住五件事：

1. 序列化用 `json.Marshal`，反序列化用 `json.Unmarshal`。
2. 结构体字段必须**导出**（首字母大写），否则 JSON 里看不到。
3. 用 ``json:"字段名"`` 标签控制 JSON 里的名字。
4. `omitempty` 可以在「空值」时省略该字段。
5. 反序列化时第二个参数必须是**指针**。

> 口诀：出站 `Marshal`，进站 `Unmarshal`；字段要大写，名字靠 tag。

## 序列化：Marshal

把结构体、map、切片等变成 JSON：

```go
package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string
	Age  int
	City string
}

func main() {
	p := Person{Name: "小满", Age: 18, City: "杭州"}

	data, err := json.Marshal(p)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(data))
	// {"Name":"小满","Age":18,"City":"杭州"}
}
```

| 写法              | 含义                            |
| ----------------- | ------------------------------- |
| `json.Marshal(v)` | 返回 `[]byte` 和 `error`        |
| `string(data)`    | 把字节转成可读字符串            |
| 字段名            | 默认用结构体字段名（如 `Name`） |

想格式化（缩进、好看一点）用 `MarshalIndent`：

```go
data, err := json.MarshalIndent(p, "", "  ")
if err != nil {
	panic(err)
}
fmt.Println(string(data))
```

输出类似：

```json
{
  "Name": "小满",
  "Age": 18,
  "City": "杭州"
}
```

## 反序列化：Unmarshal

把 JSON 填进结构体。**第二个参数必须是指针**：

```go
package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string
	Age  int
	City string
}

func main() {
	raw := []byte(`{"Name":"小满","Age":18,"City":"杭州"}`)

	var p Person
	err := json.Unmarshal(raw, &p) // 注意 &p
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", p)
	// {Name:小满 Age:18 City:杭州}
}
```

| 写法                       | 含义                 |
| -------------------------- | -------------------- |
| `json.Unmarshal(data, &v)` | 把 JSON 写入 `v`     |
| 必须传指针                 | 传值无法修改外部变量 |
| JSON 多了字段              | 默认忽略，不报错     |
| JSON 少了字段              | 对应字段保持零值     |

> 忘了写 `&` 会报错：`json: Unmarshal(non-pointer ...)`。

## 结构体标签（json tag）

接口里常见小写字段名（`name`、`age`），而 Go 导出字段必须大写。用 tag 对齐两边：

```go
package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	City string `json:"city"`
}

func main() {
	p := Person{Name: "小满", Age: 18, City: "杭州"}

	data, _ := json.Marshal(p)
	fmt.Println(string(data))
	// {"name":"小满","age":18,"city":"杭州"}

	var p2 Person
	_ = json.Unmarshal([]byte(`{"name":"奶茶","age":20,"city":"上海"}`), &p2)
	fmt.Printf("%+v\n", p2)
	// {Name:奶茶 Age:20 City:上海}
}
```

常用 tag 选项：

| tag                       | 作用                       |
| ------------------------- | -------------------------- |
| ``json:"name"``           | JSON 字段名用 `name`       |
| ``json:"name,omitempty"`` | 空值时不输出该字段         |
| ``json:"-"``              | 序列化/反序列化都忽略      |
| ``json:",omitempty"``     | 保留字段名，只加 omitempty |

### omitempty

「空」的含义：`0`、`""`、`false`、`nil`、空切片/map 等零值。

```go
package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name   string   `json:"name"`
	Age    int      `json:"age,omitempty"`
	Tags   []string `json:"tags,omitempty"`
	Secret string   `json:"-"` // 永远不出现在 JSON 里
}

func main() {
	u := User{Name: "小满", Secret: "不要对外暴露"}
	data, _ := json.Marshal(u)
	fmt.Println(string(data))
	// {"name":"小满"}
	// Age=0、Tags=nil 被 omitempty 省略；Secret 被 - 忽略
}
```

## 嵌套结构体

JSON 里经常有对象套对象，对应 Go 的嵌套结构体：

```go
package main

import (
	"encoding/json"
	"fmt"
)

type Address struct {
	City    string `json:"city"`
	Street  string `json:"street"`
}

type Order struct {
	ID      int     `json:"id"`
	Product string  `json:"product"`
	Addr    Address `json:"address"`
}

func main() {
	o := Order{
		ID:      1001,
		Product: "芋泥波波",
		Addr:    Address{City: "杭州", Street: "西湖路"},
	}

	data, _ := json.Marshal(o)
	fmt.Println(string(data))
	// {"id":1001,"product":"芋泥波波","address":{"city":"杭州","street":"西湖路"}}

	raw := []byte(`{"id":1002,"product":"绿茶","address":{"city":"上海","street":"南京路"}}`)
	var o2 Order
	_ = json.Unmarshal(raw, &o2)
	fmt.Printf("%+v\n", o2)
}
```

## map 与切片

不一定非要结构体，map 和切片也能直接编解码：

```go
package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// map → JSON
	m := map[string]any{
		"name": "小满",
		"age":  18,
		"ok":   true,
	}
	data, _ := json.Marshal(m)
	fmt.Println(string(data))

	// JSON → map
	var m2 map[string]any
	_ = json.Unmarshal(data, &m2)
	fmt.Println(m2["name"], m2["age"])
	// 注意：JSON 数字默认解成 float64

	// 切片
	list := []string{"芋泥波波", "绿茶", "红茶"}
	b, _ := json.Marshal(list)
	fmt.Println(string(b)) // ["芋泥波波","绿茶","红茶"]
}
```

> `map[string]any` 很灵活，但类型要自己断言。业务数据优先用结构体，类型更安全。

数字反序列化进 `any` 时是 **`float64`**，不是 `int`：

```go
var m map[string]any
_ = json.Unmarshal([]byte(`{"age":18}`), &m)
age := m["age"].(float64) // 对
// age := m["age"].(int)  // panic：类型不对
```

## Encoder / Decoder（流式）

跟文件、网络连接打交道时，不必先读成整段 `[]byte`，可以用流式接口：

```go
package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	// Encoder：往 Writer 写 JSON
	enc := json.NewEncoder(os.Stdout)
	_ = enc.Encode(Person{Name: "小满", Age: 18})
	// 输出一行 JSON，末尾自带换行

	// Decoder：从 Reader 读 JSON
	r := strings.NewReader(`{"name":"奶茶","age":20}`)
	dec := json.NewDecoder(r)

	var p Person
	_ = dec.Decode(&p)
	fmt.Printf("%+v\n", p)
}
```

| API                         | 场景                        |
| --------------------------- | --------------------------- |
| `Marshal` / `Unmarshal`     | 内存里一整段 JSON           |
| `NewEncoder` / `NewDecoder` | 文件、HTTP Body、长连接等流 |

## 常见坑

### 1. 字段未导出

```go
type Bad struct {
	name string // 小写：Marshal 后是 {}
}
```

必须写成 `Name`，再用 tag 映射成 `name`。

### 2. 时间字段

`time.Time` 默认按 RFC3339 字符串编解码：

```go
type Event struct {
	Name string    `json:"name"`
	At   time.Time `json:"at"`
}
```

JSON 侧一般是 `"2026-10-15T23:52:00+08:00"` 这种格式。

### 3. 指针字段与 omitempty

想用「没传」和「传了零值」区分时，常用指针：

```go
type Patch struct {
	Age *int `json:"age,omitempty"`
}
// Age == nil → 没传；Age 指向 0 → 明确传了 0
```

### 4. 严格校验多余字段

默认会忽略 JSON 里多出来的键。若要拒绝未知字段：

```go
dec := json.NewDecoder(r)
dec.DisallowUnknownFields()
err := dec.Decode(&p)
```

## 速查表

| 你想做的事      | 优先用                          |
| --------------- | ------------------------------- |
| 值 → JSON 字节  | `json.Marshal`                  |
| 美化输出        | `json.MarshalIndent`            |
| JSON 字节 → 值  | `json.Unmarshal(data, &v)`      |
| 改 JSON 字段名  | ``json:"xxx"``                  |
| 空值不输出      | ``json:"xxx,omitempty"``        |
| 永久隐藏字段    | ``json:"-"``                    |
| 写到文件/响应流 | `json.NewEncoder(w).Encode(v)`  |
| 从请求体读      | `json.NewDecoder(r).Decode(&v)` |

入门阶段先把 **`Marshal` / `Unmarshal` + 结构体 tag** 练熟；需要读写 HTTP Body 或文件时再上 `Encoder` / `Decoder`。

## 作业

1. 定义 `User`（`name`、`age`、`email`），`Marshal` 成 JSON 并打印。
2. 把一段 JSON 字符串 `Unmarshal` 进同一个结构体，打印各字段。
3. 给 `email` 加 `omitempty`，不赋值时确认输出里没有该字段。
4. 用 ``json:"-"`` 隐藏一个 `Password` 字段，确认序列化结果不含密码。
5. 写嵌套结构体 `Product` + `Price`，完成一次往返编解码。

# 20 自定义类型

基本类型（`int`、`string`…）能用，但有时不够「说话」。比如用户名和地址都是字符串，用户 ID 和订单 ID 都是整数——混着用很容易传错参。

Go 允许用 `type` 基于已有类型定义**新类型**：

```go
type 新类型名 底层类型
```

新类型和底层类型「长得像」，但在编译器眼里是**不同的类型**。这一点带来两个很实用的能力：

1. **可以挂方法**——像结构体一样，给自己的类型加行为
2. **类型安全**——同底层类型的不同语义，编译期就能拦住传参错误

------

## 1. 自定义类型可以挂方法

定义一个基于 `string` 的 `Name`，再给它挂两个方法：加前缀、拼随机后缀。

```go
package main

import (
	"fmt"
	"math/rand/v2"
	"strconv"
)

// 基于 string 定义新类型 Name
type Name string

// 给 Name 挂方法：加前缀
func (n Name) PrefixName() Name {
	return "哈夫克-" + n
}

// 给 Name 挂方法：拼一个 1~100 的随机数字后缀
func (n Name) RandomName() Name {
	ran := rand.IntN(100) + 1 // IntN(100) 是 [0,100)，再 +1 变成 [1,100]
	return n + Name(strconv.Itoa(ran))
}

func main() {
	var name Name = "张三"

	fmt.Println(name.PrefixName()) // 哈夫克-张三
	fmt.Println(name.RandomName()) // 张三42（数字每次不同）
}
```

几点说明：

- `func (n Name) Xxx()` 里的 `(n Name)` 叫**接收者**，表示这个方法挂在 `Name` 上，用 `变量.方法()` 调用
- 自定义类型和底层类型不能直接相加：`"哈夫克-" + n` 可以，是因为字符串字面量和 `Name` 的底层都是 `string`，Go 会自动处理；而数字要先 `strconv.Itoa`，再显式转成 `Name(...)`
- 普通 `string` **没有**这两个方法；只有 `Name` 才能调，这正是「挂方法」的意义

> 方法也能挂在基于 `int`、`float64` 等类型上，写法一样。结构体挂方法更常见，本质相同。

------

## 2. 规避同类型传参错误

实际业务里，用户 ID、订单 ID 往往都是 `int`。如果函数写成：

```go
func CreateOrder(uid int, oid int) { ... }
```

调用时把两个参数写反，编译器**完全不管**——两个都是 `int`，语法没错，逻辑却炸了。

用自定义类型把语义拆开：

```go
package main

import "fmt"

type Uid int // 用户 ID
type Oid int // 订单 ID

// 参数用自定义类型，语义一目了然
func CreateOrder(uid Uid, oid Oid) (Uid, Oid) {
	fmt.Println("创建订单：用户", uid, "订单", oid)
	return uid, oid
}

func main() {
	uid := Uid(10)
	oid := Oid(10086)

	// 正确：类型对得上
	CreateOrder(uid, oid)

	// 错误示范（解开注释会编译失败）：
	// CreateOrder(oid, uid)
	// 报错类似：cannot use oid (variable of type Oid) as Uid value in argument to CreateOrder

	// 就算底层都是 int，也不能直接把 int 传进去：
	// var n int = 10
	// CreateOrder(n, oid) // cannot use n (variable of type int) as Uid value
}
```

对比一下：

| 写法                                 | 传反参数时                               |
| ------------------------------------ | ---------------------------------------- |
| `func CreateOrder(uid int, oid int)` | 编译通过，运行时逻辑错误                 |
| `func CreateOrder(uid Uid, oid Oid)` | **编译直接报错**，问题在写代码时就被拦住 |

自定义类型和底层类型、以及其他自定义类型之间，需要**显式转换**，例如 `Uid(10)`、`int(uid)`。多写几个括号，换来的是少踩坑——很值。

------

## 小结

| 能力     | 做法                                       | 好处                                 |
| -------- | ------------------------------------------ | ------------------------------------ |
| 挂方法   | `type Name string` + `func (n Name) Xxx()` | 给类型绑定专属行为                   |
| 类型安全 | `type Uid int` / `type Oid int`            | 同底层类型不同语义，编译期防传参错误 |

记住一句话：**自定义类型 ≈ 底层类型的「带标签版本」**——能当同类数据用，但名字不同、方法不同、传参也不能乱混。

# 21 反射

上一篇讲了**自定义类型**。类型在编译期就定死了——那运行时怎么「看」一个值到底是什么类型、能不能改、有没有某个方法？

答案是标准库 `reflect`：**在运行时检查类型、读写值、调用方法**。框架、ORM、JSON 编解码底层都大量用到它。

> 反射灵活，但代码更绕、也更慢。日常业务能用接口 / 泛型解决的，优先别上反射。

下面用同一个自定义类型 `Age`，分三个案例：

```go
type Age int
```

------

## 1. 用反射做判断

先拿到类型信息，再看「叫什么名」「底层是什么」。

```go
package main

import (
	"fmt"
	"reflect"
)

type Age int

func main() {
	var age Age = 99
	t := reflect.TypeOf(age)

	fmt.Println(t.Name()) // Age —— 类型名（自定义类型的名字）
	fmt.Println(t.Kind()) // int —— 底层种类（拿来做判断）

	if t.Kind() == reflect.Int {
		fmt.Println("他绝对是一个 int 类型")
	}
}
```

| API                 | 作用                               | 本例结果                |
| ------------------- | ---------------------------------- | ----------------------- |
| `reflect.TypeOf(x)` | 拿到 `x` 的类型描述                | `Age` 的 `reflect.Type` |
| `t.Name()`          | 类型名称                           | `"Age"`                 |
| `t.Kind()`          | 底层种类（int / string / struct…） | `reflect.Int`           |

`Name` 和 `Kind` 别混：

- `Name()`：你起的名字（`Age`）
- `Kind()`：真正的底层形态（`int`）

做分支判断时，几乎总是比 `Kind()`，例如 `reflect.Int`、`reflect.String`、`reflect.Struct`。

------

## 2. 通过反射修改值 + 类型断言

想改原变量，必须传**指针**，再用 `.Elem()` 解引用，得到可写的 `reflect.Value`。

```go
package main

import (
	"fmt"
	"reflect"
)

type Age int

func main() {
	var age Age = 99

	// 必须传 &age，再 Elem()，否则 CanSet() 为 false，改不了
	val := reflect.ValueOf(&age).Elem()

	// val.CanSet()  是否允许修改
	// val.IsValid() 是否是有效的 reflect.Value
	if val.CanSet() && val.IsValid() && val.Kind() == reflect.Int {
		val.SetInt(10086) // 通过反射写入
	}

	// 类型断言：Interface() 先变成 any，再 .(Age) 还原成具体类型
	v := val.Interface().(Age)

	fmt.Println(val) // 10086（reflect.Value，打印时会显示内部值）
	fmt.Println(age) // 10086（原变量已被改掉）
	fmt.Println(v)   // 10086（断言回来的 Age）
}
```

要点：

1. **`reflect.ValueOf(age)` 得到的是副本**，`CanSet()` 为 `false`，`SetInt` 会 panic
2. **`reflect.ValueOf(&age).Elem()`** 才指向原变量，可以 `SetInt` / `SetString` / `Set`
3. 改之前先检查：`CanSet()`、`IsValid()`、`Kind()` 对不对
4. **类型断言**：`val.Interface().(Age)` 把反射世界里的值拿回静态类型，之后就能当普通 `Age` 用

更稳妥的断言写法（带 `ok`）：

```go
v, ok := val.Interface().(Age)
if !ok {
	fmt.Println("断言失败")
	return
}
fmt.Println(v)
```

------

## 3. 通过反射调用方法

给 `Age` 挂一个方法，再用方法名动态调用。

```go
package main

import (
	"fmt"
	"reflect"
)

type Age int

func (a Age) Say(params string, n int) {
	fmt.Println(params, n)
}

func main() {
	var age Age = 99
	val := reflect.ValueOf(age)

	m := val.MethodByName("Say") // 按名字找方法（必须导出，首字母大写）
	if m.IsValid() {
		m.Call([]reflect.Value{
			reflect.ValueOf("我是字符串"),
			reflect.ValueOf(123123),
		})
		// 输出：我是字符串 123123
	}
}
```

流程拆开看：

| 步骤 | 代码                           | 含义                                        |
| ---- | ------------------------------ | ------------------------------------------- |
| 1    | `reflect.ValueOf(age)`         | 拿到值的反射包装                            |
| 2    | `MethodByName("Say")`          | 按名字找方法；找不到则 `IsValid()` 为 false |
| 3    | `m.Call([]reflect.Value{...})` | 传入参数列表（每个参数都要先 `ValueOf`）    |

注意：

- 方法名必须**导出**（大写开头），否则 `MethodByName` 找不到
- `Call` 的参数个数、类型要和方法签名一致，否则 panic
- 有返回值时，`Call` 会返回 `[]reflect.Value`，再按需 `Interface()` 取出来

------

## 小结

| 案例   | 核心 API                                            | 用途                         |
| ------ | --------------------------------------------------- | ---------------------------- |
| 判断   | `TypeOf` / `Name` / `Kind`                          | 运行时识别类型               |
| 改值   | `ValueOf(&x).Elem()` + `SetInt` + `Interface().(T)` | 修改原变量，再断言回具体类型 |
| 调方法 | `MethodByName` + `Call`                             | 按名字动态调用               |

三板斧记牢：

1. **看类型**用 `TypeOf`，**动数据**用 `ValueOf`
2. **要改值就传指针再 `Elem()`**
3. **反射进出都要装箱/拆箱**：进去 `ValueOf`，出来 `Interface()` + 类型断言

# 22 Go Modules

前面写的都是「单文件 `package main`」。真正做项目时，代码要拆成多个包、依赖要版本可控——这就靠 **Go Modules**。

本篇按动手顺序来：

1. `go mod init` 建模块
2. `go mod tidy` 整理依赖
3. 自定义包和导入路径
4. 大写导出 / 小写私有
5. `cmd`、`internal`、`pkg` 简单目录
6. 引入一个第三方依赖演示

------

## 1. go mod init：创建模块

在空目录里初始化：

```bash
mkdir hello-mod
cd hello-mod
go mod init example.com/hello
```

会生成 `go.mod`：

```go
module example.com/hello

go 1.26.4
```

| 字段                       | 含义                           |
| -------------------------- | ------------------------------ |
| `module example.com/hello` | **模块路径**，也是包的导入前缀 |
| `go 1.22`                  | 本模块使用的 Go 语言版本       |

模块路径习惯写成「能当导入前缀」的名字，常见两种：

- 学习演示：`example.com/hello`、`myapp`
- 真仓库：`github.com/你的用户名/仓库名`

> `go.mod` 管的是**整个模块**（一个仓库通常一个），不是每个子目录一个。

------

## 2. go mod tidy：整理依赖

写完代码、增删 `import` 之后执行：

```bash
go mod tidy
```

它会：

- 把代码里**真正用到**的依赖写进 `go.mod`，并生成/更新 `go.sum`（校验和）
- 删掉**已经不用**的依赖

记住分工：

| 命令          | 什么时候用                   |
| ------------- | ---------------------------- |
| `go mod init` | 项目一开始，只跑一次         |
| `go mod tidy` | 改过 import / 依赖后，随时跑 |

`go.sum` 不要手改，提交到 Git 即可，用来保证大家下到同一份依赖。

------

## 3. 自定义包和导入路径

目录名 ≈ 包名；导入路径 = **模块路径 + 子目录**。

```text
hello-mod/
├── go.mod                 # module example.com/hello
├── main.go                # package main
└── greet/
    └── greet.go           # package greet
```

`greet/greet.go`：

```go
package greet

import "fmt"

func Hello(name string) {
	fmt.Println("你好,", name)
}
```

根目录 `main.go`：

```go
package main

import "example.com/hello/greet" // 模块路径 + 子目录

func main() {
	greet.Hello("小满")
}
```

运行：

```bash
go run .
# 你好, 小满
```

要点：

- 导入写的是**路径**，不是文件名：`example.com/hello/greet`
- 调用时用**包名**：`greet.Hello(...)`（一般等于最后一级目录名）
- 同一个目录里所有 `.go` 必须是同一个 `package` 名

------

## 4. 大写导出 / 小写私有

Go **没有** `public` / `private` 关键字，靠首字母大小写：

| 写法                        | 可见范围                   |
| --------------------------- | -------------------------- |
| `Hello`、`User`、`MaxRetry` | **导出**：别的包可以访问   |
| `hello`、`user`、`maxRetry` | **未导出**：仅本包内部可用 |

```go
package greet

import "fmt"

// 导出：外部可调用
func Hello(name string) {
	fmt.Println(prefix(), name)
}

// 未导出：外部写 greet.prefix() 会编译失败
func prefix() string {
	return "你好,"
}
```

外面只能：

```go
greet.Hello("小满") // OK
// greet.prefix()  // 报错：prefix 未导出
```

结构体字段同样规则：`Name` 导出，`age` 仅包内可见（JSON 标签、跨包赋值都会受影响）。

------

## 5. cmd、internal、pkg 简单目录

小项目一个 `greet` 包就够。稍大一点，社区常见这种骨架：

```text
hello-mod/
├── go.mod
├── cmd/
│   └── hello/
│       └── main.go      # 可执行入口：go run ./cmd/hello
├── internal/
│   └── auth/
│       └── auth.go      # 仅本模块能导入（编译器强制）
└── pkg/
    └── greeter/
        └── greeter.go   # 可被外部模块导入的库代码
```

| 目录           | 放什么                   | 谁能 import                                                  |
| -------------- | ------------------------ | ------------------------------------------------------------ |
| `cmd/<名字>/`  | `package main`，程序入口 | 不作为库被导入，用来 `go build` / `go run`                   |
| `internal/...` | 模块内部实现             | **只有本模块**（及 `internal` 的父树）能导入；外人导入会编译失败 |
| `pkg/...`      | 打算给别人用的库         | 任意模块都可导入                                             |

`cmd/hello/main.go` 示例：

```go
package main

import (
	"example.com/hello/internal/auth"
	"example.com/hello/pkg/greeter"
)

func main() {
	auth.Check()
	greeter.Hello("小满")
}
```

`internal/auth/auth.go`：

```go
package auth

import "fmt"

func Check() {
	fmt.Println("auth ok")
}
```

`pkg/greeter/greeter.go`：

```go
package greeter

import "fmt"

func Hello(name string) {
	fmt.Println("你好,", name)
}
```

说明：

- `internal` 是语言级约定：别的仓库 `import "example.com/hello/internal/auth"` 会直接报错
- `pkg` 不是强制的，很多项目把可导出库直接放在模块根下；放 `pkg` 只是「这里是对外 API」更清晰
- 入门项目不必硬套三层；**有多个可执行文件再用 `cmd`，有不想被外人用的代码再放 `internal`**

------

## 6. 引入第三方依赖演示

### 先配置国内代理（重要）

默认的官方代理 `proxy.golang.org` 在国内经常超时或拉不下来。拉第三方库之前，建议先把模块代理改成国内镜像。

推荐使用七牛云维护的 **[https://goproxy.cn](https://goproxy.cn/)**：

```bash
# 永久生效（推荐）
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn,direct
```

验证是否设置成功：

```bash
go env GOPROXY
# 应输出：https://goproxy.cn,direct
```

说明：

| 写法                 | 含义                                |
| -------------------- | ----------------------------------- |
| `https://goproxy.cn` | 优先走国内代理下载模块              |
| `,direct`            | 代理没有时，再直连源站（如 GitHub） |

也可以临时设置（只对当前终端有效）：

```bash
# Windows PowerShell
$env:GOPROXY = "https://goproxy.cn,direct"

# macOS / Linux
export GOPROXY=https://goproxy.cn,direct
```

其他可选镜像（备用）：

| 代理地址                             | 说明                         |
| ------------------------------------ | ---------------------------- |
| `https://goproxy.cn,direct`          | **首选**，七牛云，国内最常用 |
| `https://proxy.golang.com.cn,direct` | 官方域名的国内镜像站         |
| `https://goproxy.io,direct`          | 另一家公共代理               |

公司私有仓库不要走公共代理时，可再加：

```bash
go env -w GOPRIVATE=git.公司域名.com
```

### 演示：安装并使用第三方库

以常用的颜色输出库 [`github.com/fatih/color`](https://github.com/fatih/color) 为例。

在 `main.go`（或 `cmd/hello/main.go`）里先写上 import 和用法：

```go
package main

import (
	"github.com/fatih/color"
)

func main() {
	color.Green("go mod 依赖引入成功")
	color.Red("这是红色警告")
}
```

然后拉依赖并整理：

```bash
go get github.com/fatih/color
go mod tidy
```

`go.mod` 里会出现类似：

```go
module example.com/hello

go 1.22

require github.com/fatih/color v1.18.0
```

再运行：

```bash
go run .
# 终端里会看到绿色 / 红色文字
```

常用依赖命令：

| 命令               | 作用                                         |
| ------------------ | -------------------------------------------- |
| `go get 路径@版本` | 添加或升级某个依赖，如 `@v1.18.0`、`@latest` |
| `go mod tidy`      | 按实际 import 增删依赖                       |
| `go list -m all`   | 查看当前模块解析到的依赖树                   |

换依赖时：改代码里的 `import` → `go get`（可选）→ `go mod tidy`。

------

## 小结

| 主题                       | 记住什么                                                     |
| -------------------------- | ------------------------------------------------------------ |
| `go mod init`              | 定模块路径，生成 `go.mod`                                    |
| `go mod tidy`              | 按代码同步依赖与 `go.sum`                                    |
| 导入路径                   | `模块路径 + 子目录`                                          |
| 导出规则                   | 大写对外，小写包内                                           |
| `cmd` / `internal` / `pkg` | 入口 / 私有实现 / 对外库                                     |
| 第三方库                   | 先配 `GOPROXY=https://goproxy.cn,direct`，再 `go get` + import + `tidy` |

最小心智模型：

> **一个 `go.mod` = 一个模块；一个目录 = 一个包；导入靠路径，可见性靠大小写。**



# 23 http

Go 自带 `net/http`，不用额外框架，就能写 HTTP 服务。底层每个请求一个 goroutine，默认并发很强。

本篇按动手顺序来：

1. `http.NewServeMux` + `HandleFunc` 注册路由（重点：第一个参数 = 请求方式 + 路径）
2. 启动服务
3. `w.Write` 返回字符串（以及 `[]byte` 转换）
4. 返回 JSON：`json.NewEncoder(w).Encode(...)`
5. 带查询参数的用户列表搜索

------

## 1. 创建路由：NewServeMux + HandleFunc

`ServeMux` 是路由器：根据**请求方式**和**路径**，把请求交给对应处理函数。

```go
mux := http.NewServeMux()
mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
	// 处理 /hello
})
```

| API                                 | 作用                                                |
| ----------------------------------- | --------------------------------------------------- |
| `http.NewServeMux()`                | 新建一个路由器                                      |
| `mux.HandleFunc(pattern, 处理函数)` | 用第一个参数 `pattern` 匹配请求，命中则执行处理函数 |

处理函数签名固定：

```go
func(w http.ResponseWriter, r *http.Request)
```

- `w`：用来**写响应**（状态码、Header、Body）
- `r`：用来**读请求**（方法、路径、Query、Body）

> 也可以用包级函数 `http.HandleFunc(...)`，它挂在默认的 `DefaultServeMux` 上。自己 `NewServeMux()` 更清晰，推荐项目里用独立 mux。

### HandleFunc 的第一个参数：pattern

签名可以看成：

```go
mux.HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request))
```

**第一个参数 `pattern` 不是随便写的字符串**，它描述「什么样的请求才进这个处理函数」。

Go 1.22 起，`pattern` 的常见写法是：

```text
[请求方式 ]路径
```

中间有一个空格：前面是 HTTP 方法（可省略），后面是路径。

#### 只写路径（不限制方法）

```go
mux.HandleFunc("/hello", handler)
```

含义：只要路径是 `/hello`，**不管**是 GET、POST 还是 DELETE，都会进这个函数。

适合入门演示；正式接口更建议把方法写上，避免误匹配。

#### 请求方式 + 路径（推荐）

```go
mux.HandleFunc("GET /hello", handler)   // 只匹配 GET /hello
mux.HandleFunc("POST /hello", handler)  // 只匹配 POST /hello
mux.HandleFunc("PUT /user", handler)
mux.HandleFunc("DELETE /user", handler)
```

| pattern            | 匹配什么                                                   |
| ------------------ | ---------------------------------------------------------- |
| `"/hello"`         | 任意方法 + 路径 `/hello`                                   |
| `"GET /hello"`     | 仅 GET（以及 HEAD）+ `/hello`                              |
| `"POST /hello"`    | 仅 POST + `/hello`                                         |
| `"GET /user/{id}"` | GET，路径形如 `/user/1`，可用 `r.PathValue("id")` 取 `"1"` |

常见请求方式：

| 方法            | 一般用途                                   |
| --------------- | ------------------------------------------ |
| `GET`           | 查数据（浏览器地址栏、`<a>` 默认就是 GET） |
| `POST`          | 新建 / 提交表单                            |
| `PUT` / `PATCH` | 更新                                       |
| `DELETE`        | 删除                                       |

同一个路径、不同方法，可以挂不同处理函数：

```go
mux.HandleFunc("GET /user", listUser)    // 查列表
mux.HandleFunc("POST /user", createUser) // 新建
```

若只注册了 `GET /user`，却用 POST 访问，ServeMux 会自动回 **405 Method Not Allowed**（方法不对），不用你自己先判断。

#### 路径里的通配（可选）

```go
mux.HandleFunc("GET /user/{id}", func(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id") // /user/1 → "1"
	fmt.Fprint(w, "用户 id = ", id)
})
```

注意和 Query 的区别：

| 写法     | 例子              | 怎么取                                       |
| -------- | ----------------- | -------------------------------------------- |
| 路径参数 | `/user/1`         | `r.PathValue("id")`（pattern 里要有 `{id}`） |
| 查询参数 | `/user/list?id=1` | `r.URL.Query().Get("id")`                    |

后面搜索案例用的是 **Query**；路径参数知道有这回事即可。

#### 一眼对照

```go
mux := http.NewServeMux()

// 第一个参数 = 请求方式 + 空格 + 路径
mux.HandleFunc("GET /hello", func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello get"))
})
mux.HandleFunc("POST /hello", func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello post"))
})
```

浏览器打开 `http://localhost:8080/hello` 走的是 **GET**；用 Postman / curl 发 POST 才会进第二个处理函数。

> 需要 Go **1.22+**。`go.mod` 里 `go 1.22`（或更高）才支持 `"GET /path"` 这种写法；更老的版本把 `"GET /path"` 会理解错，路由对不上。

------

## 2. 创建并启动服务

把 mux 交给 `http.ListenAndServe`：

```go
package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "hello")
	})

	fmt.Println("服务启动：http://localhost:8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		panic(err)
	}
}
```

- `":8080"`：监听本机 8080 端口
- 第二个参数传 `mux`；若传 `nil`，则用默认的 `DefaultServeMux`
- 浏览器访问：`http://localhost:8080/hello`

------

## 3. 返回字符串：Write 与 []byte 转换

`ResponseWriter.Write` 的签名是：

```go
Write([]byte) (int, error)
```

它只收 **`[]byte`**，不直接收 `string`。所以要先转换：

```go
mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
	msg := "你好，小满"
	w.Write([]byte(msg)) // string → []byte
})
```

`[]byte(字符串)` 是类型转换（有人叫它「语法糖」）：把字符串变成字节切片，内存里就是那段 UTF-8 字节。

反过来：

```go
b := []byte("abc")
s := string(b) // []byte → string
```

更省事的写法（内部一样会写到 `w`）：

```go
fmt.Fprint(w, "你好，小满")
io.WriteString(w, "你好，小满")
```

入门先记住：`Write` 要 `[]byte`，字符串用 `[]byte("...")` 包一下。

------

## 4. 返回 JSON

接口常见做法：设 `Content-Type`，再用 `json.NewEncoder(w).Encode(...)` 直接往响应里写 JSON。

注意大小写（你草稿里的写法对照）：

| 你写的                       | 正确写法                         |
| ---------------------------- | -------------------------------- |
| `json.newEncode(w).encode()` | `json.NewEncoder(w).Encode(...)` |
| `type User stauct`           | `type User struct`               |

完整例子：

```go
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		u := User{Id: 1, Name: "小满"}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		err := json.NewEncoder(w).Encode(u)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	fmt.Println("服务启动：http://localhost:8080")
	http.ListenAndServe(":8080", mux)
}
```

访问 `http://localhost:8080/user`，响应大致是：

```json
{"id":1,"name":"小满"}
```

拆开看：

1. ``json:"id"``：JSON 字段名用小写 `id`（字段本身仍要大写才能导出）
2. `w.Header().Set("Content-Type", ...)`：告诉客户端这是 JSON
3. `json.NewEncoder(w)`：创建一个往 `w` 写的编码器
4. `.Encode(u)`：序列化并写入响应（末尾会自动加换行）

也可以先 `json.Marshal` 再 `w.Write`，效果类似；`NewEncoder(w).Encode` 少一次中间变量，写接口更常见。

------

## 5. 案例：按 id 搜索用户列表

目标地址：

```text
http://localhost:8080/user/list?id=1
```

你草稿里的几处要改：

| 你写的                  | 问题                                                   | 正确写法                  |
| ----------------------- | ------------------------------------------------------ | ------------------------- |
| `r.URL.Query.Get("id")` | `Query` 是**方法**，要加 `()`                          | `r.URL.Query().Get("id")` |
| `item.Id == targetId`   | `Get` 得到的是 **string**，`Id` 是 **int**，不能直接比 | 先 `strconv.Atoi`         |
| `stauct`                | 拼写错误                                               | `struct`                  |

完整可运行示例：

```go
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

var userList = []User{
	{Id: 1, Name: "小满"},
	{Id: 2, Name: "张三"},
	{Id: 3, Name: "李四"},
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/user/list", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")

		idStr := r.URL.Query().Get("id") // ?id=1 → "1"

		// 没传 id：返回全部
		if idStr == "" {
			json.NewEncoder(w).Encode(userList)
			return
		}

		targetId, err := strconv.Atoi(idStr) // string → int
		if err != nil {
			http.Error(w, "id 必须是数字", http.StatusBadRequest)
			return
		}

		result := make([]User, 0)
		for _, item := range userList {
			if item.Id == targetId {
				result = append(result, item)
			}
		}

		json.NewEncoder(w).Encode(result)
	})

	fmt.Println("服务启动：http://localhost:8080")
	http.ListenAndServe(":8080", mux)
}
```

试一下：

| 请求                | 结果                       |
| ------------------- | -------------------------- |
| `/user/list`        | 全部用户                   |
| `/user/list?id=1`   | `[{"id":1,"name":"小满"}]` |
| `/user/list?id=99`  | `[]`（没有匹配）           |
| `/user/list?id=abc` | 400，提示 id 必须是数字    |

`r.URL.Query()` 返回 `url.Values`（本质是 `map[string][]string`），`.Get("id")` 取第一个值；没有该参数时返回空字符串 `""`。

------

## 小结

| 主题       | 记住什么                                                     |
| ---------- | ------------------------------------------------------------ |
| 路由       | `NewServeMux` + `HandleFunc`；第一个参数是 `"[METHOD ]路径"`，如 `"GET /hello"` |
| 启动       | `ListenAndServe(":8080", mux)`                               |
| 写文本     | `w.Write([]byte("..."))`                                     |
| 写 JSON    | `Header` 设 Content-Type + `json.NewEncoder(w).Encode(v)`    |
| 结构体标签 | `Id int `json:"id"``，字段要导出                             |
| Query      | `r.URL.Query().Get("id")`，和 int 比较前先 `Atoi`            |

最小流程：

> **mux 注册路径 → ListenAndServe 启动 → 处理函数里读 `r`、写 `w`。**

# 24 Gin

Gin 是 Go 里常用的 HTTP 框架。路由按请求方法分开注册，取参数、回 JSON 都比 `net/http` 少写几行。它底层仍是标准库，上一篇里「读请求、写响应」的概念这里照样用。

本篇按动手顺序来：

1. 把 Go 升到最新稳定版（旧版和 SSE 依赖有兼容问题），或按版本另装一份
2. 安装 Gin
3. 安装热更新工具 Air，并看懂那条 `go install`
4. 起一个最小服务
5. 传参：路径参数、Query、表单、JSON

中间件、路由分组这些先不讲。

------

## 1. Go 版本：建议升到最新

先看本机版本：

```bash
go version
```

Gin 会依赖 [`github.com/gin-contrib/sse`](https://github.com/gin-contrib/sse)（Server-Sent Events，服务端按事件往客户端推数据）。这里有两层问题：

- 新版 `sse` 提高了 Go 语言版本要求。Go 太旧时，`go get` 最新 Gin 会在拉依赖阶段失败。
- 旧版 `sse` 跟较新的 `net/http` 在刷新响应（Flush）上行为不一致，真正用到推流时容易出兼容问题。

当前最新稳定版是 **Go 1.27.4**。直接覆盖安装最省事：打开 https://go.dev/dl/ 下载对应系统的安装包。Windows 也可以：

```powershell
winget upgrade GoLang.Go
```

装完再执行 `go version`，应看到 `go1.27.4`。

### 不想替换系统里的 go：按版本另装

机器上已经有一个旧的 `go`，又不想动它，可以再装一份指定版本，两条命令：

```bash
go install golang.org/dl/go1.27.4@latest
go1.27.4 download
```

| 命令                                       | 做了什么                                                     |
| ------------------------------------------ | ------------------------------------------------------------ |
| `go install golang.org/dl/go1.27.4@latest` | 只装一个名叫 `go1.27.4` 的小命令，放到 `GOPATH/bin`，这时还没有完整工具链 |
| `go1.27.4 download`                        | 用上面那个命令，把 Go 1.27.4 本体下载下来                    |

之后在这个项目里把 `go` 换成 `go1.27.4`，例如 `go1.27.4 mod init`、`go1.27.4 get ...`、`go1.27.4 run .`。原来的 `go version` 不变。

------

## 2. 安装 Gin

在模块目录里执行（还没有模块就先 `go mod init example.com/gin-demo`）：

```bash
go get github.com/gin-gonic/gin@latest
```

`@latest` 拉当前最新稳定版。写作时是 **v1.12.0**，`go.mod` 里要求 **Go 1.25+**，所以前面才先升级。

代码里加上导入，再整理一次依赖：

```go
import "github.com/gin-gonic/gin"
go mod tidy
```

国内拉模块超时的话，先配代理（Modules 那篇写过）：

```bash
go env -w GOPROXY=https://goproxy.cn,direct
```

用了上一节的并存版本时，这几条里的 `go` 同样换成 `go1.27.4`。

------

## 3. 热更新：Air

开发时每次改完都要自己停进程、再 `go run`，比较烦。Air 就是干这个的：盯着项目里的源文件，一保存就重新编译并重启进程。

安装命令：

```bash
go install github.com/air-verse/air@latest
```

把它拆开看：

| 片段                       | 含义                                                         |
| -------------------------- | ------------------------------------------------------------ |
| `go install`               | 编译这个程序，把可执行文件装进 `GOPATH/bin`。它**不会**写进当前项目的 `go.mod`，因为 Air 是你本机的开发工具，不是业务依赖 |
| `github.com/air-verse/air` | 模块路径。早期仓库在 `github.com/cosmtrek/air`，已经迁到 `air-verse`，安装用新路径 |
| `@latest`                  | 装最新版。Air 本身要求 Go 1.25+，所以同样用刚升级的那套 Go   |

装好后进入项目目录：

```bash
air
```

它会按默认配置监视 `.go` 文件。想改监视规则再执行 `air init`，会在当前目录生成 `.air.toml`。

记住它的边界：Air 只在**你自己电脑上写代码**时用。上线还是正常 `go build` 出一个二进制，不会在生产环境里靠它热重启。

Windows 上如果提示找不到 `air`，把 `go env GOPATH` 下面的 `bin` 加进 PATH（常见是 `C:\Users\你的用户名\go\bin`），新开一个终端再执行。并存安装时，用 `go1.27.4 install github.com/air-verse/air@latest`，装出来的命令同样在那个 `bin` 里。

------

## 4. 最小服务

```go
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "你好")
	})

	if err := router.Run(); err != nil {
		panic(err)
	}
}
```

| API                                 | 作用                                              |
| ----------------------------------- | ------------------------------------------------- |
| `gin.Default()`                     | 创建一个引擎，自带访问日志和 panic 恢复           |
| `router.GET(路径, 处理函数)`        | 只接收 GET。POST 就写成 `router.POST`             |
| `ctx.String(状态码, 格式, 参数...)` | 往响应里写一段文本，格式和 `fmt.Sprintf` 一样     |
| `router.Run()`                      | 监听 `:8080`。写成 `router.Run(":8080")` 是一回事 |

`*gin.Context` 同时负责读这次请求、写这次响应，对应标准库里的 `r` 和 `w`。

`go run .` 或 `air` 启动后，浏览器打开 `http://localhost:8080`，能看到「你好」。端口被占用时 `Run` 会返回错误，这里直接 `panic` 让终端把原因打出来。

------

## 5. 传参

接口里数据通常从四个地方来：

| 方式     | 长什么样                                | 怎么取                           |
| -------- | --------------------------------------- | -------------------------------- |
| 路径参数 | `/user/12`                              | `ctx.Param("id")`                |
| Query    | `/user?id=1&name=小满`                  | `ctx.DefaultQuery` / `ctx.Query` |
| 表单     | POST，正文是 `id=1&name=小满`           | `ctx.DefaultPostForm`            |
| JSON     | POST，正文是 `{"name":"小满","age":18}` | `ctx.ShouldBindJSON`             |

下面每种先看一小段，最后合成一个能跑的 `main.go`。

### 路径参数

路径里用 `:名字` 占一位：

```go
router.GET("/user/:id", func(ctx *gin.Context) {
	id := ctx.Param("id") // /user/12 → "12"
	ctx.String(http.StatusOK, "id=%s", id)
})
```

`Param` 拿到的永远是字符串。请求是 `/user/12`，`id` 就是 `"12"`，不会自动变成数字。

### Query

问号后面的键值对，用在「筛选、搜索」这种可选条件上：

```go
router.GET("/user", func(ctx *gin.Context) {
	id := ctx.DefaultQuery("id", "none")
	name := ctx.DefaultQuery("name", "none")
	ctx.String(http.StatusOK, "id=%s name=%s", id, name)
})
```

`DefaultQuery` 的第二个参数是**没带这个键**时的默认值。`/user` 什么都不带，两个都是 `none`；`/user?id=1&name=小满` 则取到实际的值。

只要键出现了，哪怕值是空的（`?name=`），也不会再用默认值。想要「空字符串就是没有」，用 `ctx.Query("name")`，缺省时它返回 `""`。

`/user/:id` 和 `/user` 可以同时注册：一个要路径上多一段，一个就是 `/user` 本身，Gin 能分开。

### 表单

POST 的正文是表单（`application/x-www-form-urlencoded` 或 `multipart/form-data`）时，用 `DefaultPostForm`。第二个参数同样是「没带这个字段」时的默认值：

```go
router.POST("/user/formdata", func(ctx *gin.Context) {
	id := ctx.DefaultPostForm("id", "none")
	name := ctx.DefaultPostForm("name", "none")
	ctx.String(http.StatusOK, "id=%s name=%s", id, name)
})
```

它读的是 **Body**，不是 URL 上的 `?id=1`。用 Query 那套函数去读表单，会一直拿到默认值。

### JSON

定义一个结构体，字段要大写才能被填上。``json`` 决定 JSON 里的名字，``binding`` 决定要不要校验：

```go
type User struct {
	Name string `json:"name" binding:"required"`
	Age  int    `json:"age"`
}
```

标签是 `binding`，少写一个字母就不会校验，请求照样能进来。

```go
router.POST("/user/create", func(ctx *gin.Context) {
	var user User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"code":    1,
		"data":    user,
		"success": true,
	})
})
```

`ShouldBindJSON` 做两件事：按 `json` 标签把正文填进结构体，再按 `binding` 标签检查。`Name` 标了 `required`，缺 `name` 或是空字符串都会返回错误。`Age` 没标，不传就是 `0`。

失败时要 `return`。否则函数还会继续往下走，同一次响应里先写 400、再写 201。

`gin.H` 是 `map[string]any` 的别名，用来随手组一个 JSON 对象。`ctx.JSON` 会设置 `Content-Type: application/json` 并序列化。这里成功用 `201 Created`，参数不合法用 `400 Bad Request`。

请求需要：

- 方法：`POST`
- 头：`Content-Type: application/json`
- 正文：`{"name":"小满","age":18}`

### 合在一起

```go
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Name string `json:"name" binding:"required"`
	Age  int    `json:"age"`
}

func main() {
	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "你好")
	})

	// 路径参数：/user/12
	router.GET("/user/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		ctx.String(http.StatusOK, "id=%s", id)
	})

	// Query：/user?id=1&name=小满
	router.GET("/user", func(ctx *gin.Context) {
		id := ctx.DefaultQuery("id", "none")
		name := ctx.DefaultQuery("name", "none")
		ctx.String(http.StatusOK, "id=%s name=%s", id, name)
	})

	// 表单正文：id=1&name=小满
	router.POST("/user/formdata", func(ctx *gin.Context) {
		id := ctx.DefaultPostForm("id", "none")
		name := ctx.DefaultPostForm("name", "none")
		ctx.String(http.StatusOK, "id=%s name=%s", id, name)
	})

	// JSON 正文：{"name":"小满","age":18}
	router.POST("/user/create", func(ctx *gin.Context) {
		var user User
		if err := ctx.ShouldBindJSON(&user); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusCreated, gin.H{
			"code":    1,
			"data":    user,
			"success": true,
		})
	})

	if err := router.Run(); err != nil {
		panic(err)
	}
}
```

用 `curl` 对一下（服务已在 8080）：

```bash
curl http://localhost:8080/
curl http://localhost:8080/user/12
curl "http://localhost:8080/user?id=1&name=小满"
curl -X POST http://localhost:8080/user/formdata -d "id=1&name=小满"
curl -X POST http://localhost:8080/user/create \
  -H "Content-Type: application/json" \
  -d "{\"name\":\"小满\",\"age\":18}"
```

| 请求                                                 | 结果                       |
| ---------------------------------------------------- | -------------------------- |
| `GET /`                                              | `你好`                     |
| `GET /user/12`                                       | `id=12`                    |
| `GET /user?id=1&name=小满`                           | `id=1 name=小满`           |
| `GET /user`                                          | `id=none name=none`        |
| `POST /user/formdata`，正文 `id=1&name=小满`         | `id=1 name=小满`           |
| `POST /user/create`，正文 `{"name":"小满","age":18}` | `201`，`data` 里是这个用户 |
| `POST /user/create`，正文 `{}`                       | `400`，提示 `name` 为必填  |

------

## 小结

| 主题     | 记住什么                                                     |
| -------- | ------------------------------------------------------------ |
| Go 版本  | 升到最新稳定版（目前 1.27.4）。旧 Go 和新版 `gin-contrib/sse` 对不上；不想覆盖安装就 `go install golang.org/dl/go1.27.4@latest` 再 `download` |
| 装 Gin   | `go get github.com/gin-gonic/gin@latest`，然后 `go mod tidy` |
| Air      | `go install github.com/air-verse/air@latest`，项目目录执行 `air`。它监视文件并重启进程，只用于本地开发 |
| 启动     | `gin.Default()` 注册路由，`router.Run()` 监听 `:8080`        |
| 路径参数 | `/user/:id`，用 `ctx.Param("id")`，结果是字符串              |
| Query    | `ctx.DefaultQuery("id", "none")`，第二个参数是没带这个键时的默认值 |
| 表单     | `ctx.DefaultPostForm`，读的是 POST 正文，不是 URL            |
| JSON     | 结构体加 `json` / `binding` 标签，`ShouldBindJSON` 绑定；失败要 `return`，再用 `ctx.JSON` 写响应 |

# 25 Gin 上传文件

上传走的是 `multipart/form-data`：浏览器把文件拆成一段一段放进请求正文，服务端按**表单字段名**把文件取出来，再写到磁盘。

本篇按动手顺序来：

1. 单文件上传
2. 校验大小和后缀
3. 多文件上传

上一篇的 `DefaultPostForm` 只能读普通表单字段，读不到文件。文件用 `FormFile` 或 `MultipartForm`。

------

## 1. 单文件上传

前端字段名是什么，服务端就按什么取。下面约定字段名是 `file`：

```html
<form action="/upload" method="post" enctype="multipart/form-data">
  <input type="file" name="file">
  <button type="submit">上传</button>
</form>
```

`enctype="multipart/form-data"` 必须写上，否则正文里没有文件。

```go
router.POST("/upload", func(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// etc/static/1.png → 1.png，只留文件名
	dst := filepath.Join("./files/", filepath.Base(file.Filename))
	if err := ctx.SaveUploadedFile(file, dst); err != nil {
		ctx.String(http.StatusInternalServerError, "保存失败")
		return
	}
	ctx.String(http.StatusOK, "上传成功")
})
```

| API                                 | 作用                                                         |
| ----------------------------------- | ------------------------------------------------------------ |
| `ctx.FormFile("file")`              | 按字段名取出**一个**文件。返回的是 `*multipart.FileHeader`，里面有原始文件名和大小 |
| `file.Filename`                     | 客户端带来的文件名，可能带目录，例如 `etc/static/1.png`      |
| `filepath.Base(file.Filename)`      | 只留最后一段：`etc/static/1.png` → `1.png`                   |
| `filepath.Join("./files/", 文件名)` | 拼出落盘路径 `files/1.png`。分隔符按当前系统处理             |
| `ctx.SaveUploadedFile(file, dst)`   | 把文件写到 `dst`。父目录不存在时会自己建                     |

`Base` 要加上。客户端可以带上 `../` 或一整段路径，直接拿 `file.Filename` 去拼接，文件会写到 `./files/` 外面。`Base` 丢掉目录，只留下 `1.png` 这种名字，再和 `./files/` 拼在一起。

`SaveUploadedFile` 的错误要判断。保存失败时如果继续往下写「上传成功」，前端看到的结果是错的。

用 `curl` 试（`-F` 会自动带上 `multipart/form-data`）：

```bash
curl -F "file=@1.png" http://localhost:8080/upload
```

`file=` 对应 `FormFile("file")`，`@1.png` 表示把本地文件 `1.png` 传上去。成功后项目下会出现 `files/1.png`。

------

## 2. 校验大小和后缀

文件取出来之后、写盘之前做两道检查：太大的丢掉，后缀不在白名单里的丢掉。

### 大小：`5 << 20` 就是 5MB

`<<` 是左移。二进制往左挪 n 位，数值乘上 2 的 n 次方。内存里文件大小的单位是**字节**，所以用位移写出 KB、MB、GB：

| 写法      | 等于                    | 含义 |
| --------- | ----------------------- | ---- |
| `1 << 10` | 1024 字节               | 1KB  |
| `1 << 20` | 1024 × 1024 字节        | 1MB  |
| `1 << 30` | 1024 × 1024 × 1024 字节 | 1GB  |
| `5 << 20` | 5 × 1MB                 | 5MB  |

左移的位数决定单位：`10` 是 KB，`20` 是 MB，`30` 是 GB。前面的数字是「几个」这种单位。`5 << 20` 就是 5 个 MB。

```go
const MaxSize = 5 << 20 // 5MB

if file.Size > MaxSize {
	ctx.String(http.StatusRequestEntityTooLarge, "文件过大")
	return
}
```

`file.Size` 是字节数。超过 5MB 回 **413**（`StatusRequestEntityTooLarge`），然后 `return`，后面的保存不再执行。

### 后缀：只允许 png 和 jpg

用 map 当白名单，查找是 O(1)，不在里面的键读出来是 `false`：

```go
var allowExt = map[string]bool{
	".png": true,
	".jpg": true,
}

ext := strings.ToLower(filepath.Ext(file.Filename))
if !allowExt[ext] {
	ctx.String(http.StatusBadRequest, "该后缀暂时不支持")
	return
}
```

`filepath.Ext` 带点，`"1.png"` 得到 `".png"`，所以白名单里要写成 `".png"`，不能写成 `"png"`。

后缀大小写不固定，Windows 上经常是 `.PNG`、`.JPG`。`Ext` 不会帮你转小写，比较前用 `strings.ToLower`，否则合法图片会被拒。

两道检查都放在 `SaveUploadedFile` 前面：

```go
router.POST("/upload", func(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if file.Size > MaxSize {
		ctx.String(http.StatusRequestEntityTooLarge, "文件过大")
		return
	}
	ext := strings.ToLower(filepath.Ext(file.Filename))
	if !allowExt[ext] {
		ctx.String(http.StatusBadRequest, "该后缀暂时不支持")
		return
	}

	dst := filepath.Join("./files/", filepath.Base(file.Filename))
	if err := ctx.SaveUploadedFile(file, dst); err != nil {
		ctx.String(http.StatusInternalServerError, "保存失败")
		return
	}
	ctx.String(http.StatusOK, "上传成功")
})
```

| 请求              | 结果                                    |
| ----------------- | --------------------------------------- |
| `1.png`，小于 5MB | `files/1.png`，响应「上传成功」         |
| 6MB 的 png        | `413`，「文件过大」，磁盘上没有这个文件 |
| `1.gif`           | `400`，「该后缀暂时不支持」             |
| `photo.JPG`       | 转小写后是 `.jpg`，允许保存             |

------

## 3. 多文件上传

一次选好几个文件时，`FormFile` 只会拿到该字段的第一个。整包解析用 `MultipartForm`，再按字段名取出文件切片。

前端字段名改成 `files`，并加上 `multiple`：

```html
<input type="file" name="files" multiple>
form, err := ctx.MultipartForm()
if err != nil {
	ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	return
}

files := form.File["files"]
```

`form.File` 的类型是 `map[string][]*multipart.FileHeader`。键是字段名 `"files"`，值是这个字段上的全部文件。字段名和前端 `name` 不一致时，取出来是 `nil`，循环一次都不会进。

大小、后缀和单文件是同一套判断，放进循环里，每个文件单独看。有一个不合格就 `return`，不再保存后面的：

```go
router.POST("/upload", func(ctx *gin.Context) {
	form, err := ctx.MultipartForm()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	files := form.File["files"]
	if len(files) == 0 {
		ctx.String(http.StatusBadRequest, "没有文件")
		return
	}

	for _, file := range files {
		if file.Size > MaxSize {
			ctx.String(http.StatusRequestEntityTooLarge, "文件过大")
			return
		}
		ext := strings.ToLower(filepath.Ext(file.Filename))
		if !allowExt[ext] {
			ctx.String(http.StatusBadRequest, "该后缀暂时不支持")
			return
		}

		dst := filepath.Join("./files/", filepath.Base(file.Filename))
		if err := ctx.SaveUploadedFile(file, dst); err != nil {
			ctx.String(http.StatusInternalServerError, "保存失败")
			return
		}
	}
	ctx.String(http.StatusOK, "上传成功")
})
```

循环结束后再回复「上传成功」。写在循环里面的话，第一个文件保存完就会结束响应。

两个文件一起传：

```bash
curl -F "files=@1.png" -F "files=@2.jpg" http://localhost:8080/upload
```

两次 `-F` 的字段名都是 `files`，对应 `form.File["files"]`。

### 合在一起

```go
package main

import (
	"net/http"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

// 1 << 10 是 1KB，1 << 20 是 1MB，1 << 30 是 1GB
const MaxSize = 5 << 20 // 5MB

var allowExt = map[string]bool{
	".png": true,
	".jpg": true,
}

func main() {
	router := gin.Default()

	router.POST("/upload", func(ctx *gin.Context) {
		form, err := ctx.MultipartForm()
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		files := form.File["files"]
		if len(files) == 0 {
			ctx.String(http.StatusBadRequest, "没有文件")
			return
		}

		for _, file := range files {
			if file.Size > MaxSize {
				ctx.String(http.StatusRequestEntityTooLarge, "文件过大")
				return
			}
			ext := strings.ToLower(filepath.Ext(file.Filename))
			if !allowExt[ext] {
				ctx.String(http.StatusBadRequest, "该后缀暂时不支持")
				return
			}

			dst := filepath.Join("./files/", filepath.Base(file.Filename))
			if err := ctx.SaveUploadedFile(file, dst); err != nil {
				ctx.String(http.StatusInternalServerError, "保存失败")
				return
			}
		}
		ctx.String(http.StatusOK, "上传成功")
	})

	if err := router.Run(); err != nil {
		panic(err)
	}
}
```

只传一个文件时，用同一个接口即可，字段名仍然是 `files`：

```bash
curl -F "files=@1.png" http://localhost:8080/upload
```

单文件那节的 `FormFile("file")` 适合确定只有一个文件的接口；多个文件用 `MultipartForm`。

------

## 小结

| 主题   | 记住什么                                                     |
| ------ | ------------------------------------------------------------ |
| 单文件 | `ctx.FormFile("file")`，字段名和前端 `name` 一致             |
| 落盘   | `filepath.Join("./files/", filepath.Base(file.Filename))`，`Base` 去掉客户端带来的目录 |
| 保存   | `ctx.SaveUploadedFile(file, dst)`，错误要 `return`           |
| 大小   | `5 << 20` 是 5MB。`10` / `20` / `30` 分别对应 KB / MB / GB。超限回 413 |
| 后缀   | `filepath.Ext` 的结果带点，白名单写成 `".png"`。比较前 `strings.ToLower` |
| 多文件 | `ctx.MultipartForm()`，用 `form.File["files"]` 取出切片，循环里做同一套校验再保存 |

# 26 Gin 中间件

什么是中间件?

中间件属于，[真正的业务代码]，之前，以及之后一段的通用逻辑，对请求做统一的加工，拦截，包装等，基本所有的框架都有这个思想。

例如：

- 跨域 CORS
- 登录鉴权、权限校验
- 请求日志、耗时统计
- panic 捕获 & 统一异常返回
- 请求限流、ip 黑名单
- 统一加请求 ID（traceId）

### 注册中间件

中间件，是通过`Use`方法去进行注册的，接受一个函数，并且注入参数`gin.Context`上下文，

> 需注意：中间件是放在顶部的，而不能放在路由下面，否则不生效

```go
package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Cors(ctx *gin.Context) {
	//请求之前做点什么
	fmt.Println("==================before")
	ctx.Next()
    //请求之后做点什么
	fmt.Println("==================after")
}

func main() {
	router := gin.New()
	router.Use(Cors)

	router.GET("/ping", func(ctx *gin.Context) {
		fmt.Println("result")
		ctx.String(http.StatusOK, "请求成功:%v", ctx.Request.URL)
	})

    //router.Use(Cors) 如放在路由下面则不生效

	router.Run()
}
```

输出结果

```txt
==================before
result
==================after
```

### Next & Abort

#### Next函数

请求进来时，Gin 把中间件和路由处理函数排成一条链，从头开始调用。遇到 Next()，当前函数先停住，去执行链上还没跑的部分：后面的中间件，以及最后的 GET /ping。这些全部 return 之后，Next() 才返回，after 才会打印

```go
func cors(c *gin.Context) {
    // 👉【前置逻辑】先跑这里
    fmt.Println("before next")

    c.Next() // ✅ 关键点：放行，去跑后面的链条，这里代码会卡住，等待后面全部跑完

    // 👉【后置逻辑】等后面所有中间件、handler 跑完之后，才回来执行这里
    fmt.Println("after next")
}
```

#### Abort

`Abort` 把这条链标记为已中止。还没开始的后续中间件，以及最后的业务函数，都不会再被调用。`/ping` 里的 `result` 不会打印。

它不会让当前函数停在这一行。`Abort()` 下面的语句照常往下走。想跳过当前函数剩下的代码，要自己 `return`。`return` 只离开当前函数，拦住后面链条的是 `Abort`。

已经在外层执行、并且调用过 `Next()` 的中间件也不受影响。内层中止之后，外层的 `Next()` 会返回，写在它下面的后置逻辑仍会执行。

`Abort` 本身不写响应。只调用它时，客户端拿到的是默认的 **200** 和空正文。拦截时先把状态码和内容写出去，或者直接用 `AbortWithStatusJSON`（内部就是先 `Abort`，再 `JSON`）。

```go
func logCost(ctx *gin.Context) {
	fmt.Println("Log before")
	ctx.Next()
	fmt.Println("Log after")
}

func auth(ctx *gin.Context) {
	fmt.Println("before Abort")

	if ctx.GetHeader("token") == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "未登录",
		})
		fmt.Println("after Abort")
		return
	}

	ctx.Next()
}
```

注册顺序是先 `logCost` 再 `auth`。没带 `token` 访问 `/ping` 时：

```txt
Log before
before Abort
after Abort
Log after
```

响应是 `401`，正文 `{"error":"未登录"}`。`Log after` 仍会打印，因为 `logCost` 已经停在自己的 `Next()` 上，等的就是内层返回。`Abort` 拦的是还没轮到的函数。

### 多个中间件执行例子

`Use` 可以连续注册。先注册的包在外面，后注册的在里面。同一次请求先按注册顺序走进去，业务函数返回后，再按相反的顺序走出来。

下面先挂 `Cors`，再挂 `Cors2`。两个路由都写在 `Use` 后面，所以 `/ping` 和 `/test` 都会经过这两个中间件。下面这份输出是访问 `/ping` 的：只有它在处理函数里打印了 `result`。

```go
package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Cors(ctx *gin.Context) {
	//请求之前做点什么
	fmt.Println("==================before")
	ctx.Next()
	fmt.Println("==================after")
}
func Cors2(ctx *gin.Context) {
	//请求之前做点什么
	fmt.Println("==================before2")
	ctx.Next()
	fmt.Println("==================after2")
}

func main() {
	router := gin.New()

	router.Use(Cors)
	router.Use(Cors2)

	router.GET("/ping", func(ctx *gin.Context) {
		fmt.Println("result")
		ctx.String(http.StatusOK, "请求成功:%v", ctx.Request.URL)
	})

	router.GET("/test", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "请求成功:%v", ctx.Request.URL)
	})

	router.Run()
}
```

访问 `GET /ping` 的输出：

```txt
==================before
==================before2
result
==================after2
==================after
```

套在一起就是：

```txt
Cors before
  Cors2 before
    result
  Cors2 after
Cors after
```

`Cors` 先打印 `before`，调用 `Next()` 后停住，把控制权交给后注册的 `Cors2`。`Cors2` 同样先打印 `before2`，再 `Next()` 进入 `/ping`。`/ping` 返回之后，才轮到 `after2`；接着 `Cors` 的 `Next()` 返回，最后打印 `after`。

谁先 `Use`，谁的前置最早执行，谁的后置最晚执行。把 `Use(Cors)` 和 `Use(Cors2)` 对调，`before` / `after` 的先后也会对调。

访问 `/test` 时，两层中间件的顺序相同，中间没有 `result` 这一行。

### 案例：跨域 CORS

浏览器默认只让页面读**同源**接口的响应。源看三样：协议、主机、端口，三样全相同才算同一个源。

| 页面地址                | 接口地址                     | 是否同源                                        |
| ----------------------- | ---------------------------- | ----------------------------------------------- |
| `http://127.0.0.1:5500` | `http://127.0.0.1:5500/ping` | 同源                                            |
| `http://127.0.0.1:5500` | `http://127.0.0.1:8080/ping` | 跨域，端口不同                                  |
| `http://localhost:8080` | `http://127.0.0.1:8080`      | 跨域，`localhost` 和 `127.0.0.1` 是两台「主机」 |
| `https://example.com`   | `http://example.com`         | 跨域，协议不同                                  |

上面这段脚本如果跑在 `http://127.0.0.1:5500` 的页面里，去请求 `http://localhost:8080/ping`，就是跨域。服务端不给 CORS 响应头时，接口其实可能已经执行了，但浏览器会把结果拦下，前端 `fetch` 直接报错。

```html
<script>
    fetch("http://localhost:8080/ping", {
        method: "DELETE",
        headers: {
            "Content-Type": "application/json; charset=utf-8"
        }
    })
</script>
```

#### 请求方法

`method` 就是这次请求的方法，和 Gin 里 `router.GET`、`router.DELETE` 是同一组名字：

| 方法            | 一般用途                                           |
| --------------- | -------------------------------------------------- |
| `GET`           | 查询                                               |
| `POST`          | 新建、提交                                         |
| `PUT` / `PATCH` | 更新                                               |
| `DELETE`        | 删除                                               |
| `OPTIONS`       | 浏览器用来问「这个跨域请求允不允许」，不是业务接口 |

这里写的是 `DELETE`，所以服务端要有 `router.DELETE("/ping", ...)`。写成 `router.GET` 时，真正的删除请求对不上，会得到 404。

#### 预检请求

有的跨域请求浏览器会直接发，有的会先发一次 **预检**。预检是浏览器自动补的 `OPTIONS`，页面里的 `fetch` 不用自己写。

这次会预检，因为同时占了两条：

- 方法是 `DELETE`（不是 `GET` / `HEAD` / `POST`）
- 请求头 `Content-Type` 是 `application/json`（不是表单或纯文本）

浏览器实际会发两次：

1. `OPTIONS /ping`，头里带上 `Origin`（页面的源，例如 `http://127.0.0.1:5500`）、`Access-Control-Request-Method: DELETE`、`Access-Control-Request-Headers: content-type`。问服务器：这个源、这个方法、这个头，能不能发。
2. 预检通过之后，才发真正的 `DELETE /ping`。

预检没有对应的业务路由。中间件不接住 `OPTIONS` 的话，Gin 回 404，浏览器认为不允许跨域，第二次 `DELETE` 也不会发，终端里看不到 `result`。

#### 三个响应头

手写的中间件在进业务之前，把允许范围写到响应头上。真正的 `DELETE` 也要带上这些头，所以写在 `OPTIONS` 判断的前面，两条请求都能带上。

| 响应头                         | 例子里的值                  | 作用                                                         |
| ------------------------------ | --------------------------- | ------------------------------------------------------------ |
| `Access-Control-Allow-Origin`  | `*`                         | 允许哪个源读取响应。`*` 表示任意源。页面要带 Cookie 时不能用 `*`，得写成具体源 |
| `Access-Control-Allow-Methods` | `GET,POST,DELETE,PUT,PATCH` | 预检时告诉浏览器，这些方法可以跨域调用。要包含这次的 `DELETE` |
| `Access-Control-Allow-Headers` | `Content-Type`              | 预检时告诉浏览器，请求可以带哪些头。前端写了 `Content-Type`，这里就要允许它 |

```go
func Cors(ctx *gin.Context) {
	fmt.Println("==================before")
	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.Header("Access-Control-Allow-Methods", "GET,POST,DELETE,PUT,PATCH")
	ctx.Header("Access-Control-Allow-Headers", "Content-Type")

	// 预检不是业务请求，回 204 后结束，不再进入 /ping
	if ctx.Request.Method == "OPTIONS" {
		ctx.AbortWithStatus(http.StatusNoContent)
		return
	}

	ctx.Next()
	fmt.Println("==================after")
}
```

`AbortWithStatus(204)` 把预检标成成功（没有正文），并中止后面的链条。这里要 `return`：否则还会继续执行 `Next()`。`Abort` 之后 `Next()` 不会进 `/ping`，但当前函数剩下的代码仍会走，和上一节一样。

完整例子：

```go
package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Cors(ctx *gin.Context) {
	fmt.Println("==================before")
	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.Header("Access-Control-Allow-Methods", "GET,POST,DELETE,PUT,PATCH")
	ctx.Header("Access-Control-Allow-Headers", "Content-Type")

	if ctx.Request.Method == "OPTIONS" {
		ctx.AbortWithStatus(http.StatusNoContent)
		return
	}

	ctx.Next()
	fmt.Println("==================after")
}

func main() {
	router := gin.New()
	router.Use(Cors)

	router.DELETE("/ping", func(ctx *gin.Context) {
		fmt.Println("result")
		ctx.String(http.StatusOK, "请求成功:%v", ctx.Request.URL)
	})

	router.Run()
}
```

用另一个源的页面点这次 `fetch`，顺序是：

| 顺序 | 请求            | 终端                        | 响应                                                         |
| ---- | --------------- | --------------------------- | ------------------------------------------------------------ |
| 1    | `OPTIONS /ping` | 只有 `before`               | `204`，带上面三个头                                          |
| 2    | `DELETE /ping`  | `before`、`result`、`after` | `200`，正文是「请求成功」，响应里同样有 `Access-Control-Allow-Origin` |

第二次如果缺 `Allow-Origin`，浏览器还是不会把正文交给 `fetch`。

#### 项目里用现成的

上面这份是为了看清预检和三个头，项目里不用自己维护。Gin 官方配套的是 [`github.com/gin-contrib/cors`](https://github.com/gin-contrib/cors)，它会处理 `OPTIONS`、补响应头。

```bash
go get github.com/gin-contrib/cors
go mod tidy
```

和手写 `*` 同样宽松的写法：

```go
import "github.com/gin-contrib/cors"

func main() {
	router := gin.New()
	router.Use(cors.Default())

	router.DELETE("/ping", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "请求成功:%v", ctx.Request.URL)
	})

	router.Run()
}
```

`cors.Default()` 允许任意源，效果接近 `Access-Control-Allow-Origin: *`，也不带 Cookie。要收成某一个前端地址时，改用配置。`AllowOrigins` 写页面的源，方法、请求头和手写的那三行是同一件事：

```go
config := cors.DefaultConfig()
config.AllowOrigins = []string{"http://127.0.0.1:5500"}
config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}
config.AllowHeaders = []string{"Content-Type"}

router.Use(cors.New(config))
```

`DefaultConfig()` 里其实已经带了这些方法和 `Content-Type`。源必须自己填，漏了的话 `cors.New` 会直接报配置冲突。页面地址是 `localhost` 还是 `127.0.0.1`，要和 `AllowOrigins` 里的字符串一致。