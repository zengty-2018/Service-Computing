# 服务计算第六次作业——修改、改进 RxGo 包
## 任务
请在 pmlpml/RxGo 基础上，  
~~1. 修改、改进它的实现~~  
2. 或添加一组新的操作，如 filtering  

该库的基本组成：  
rxgo.go 给出了基础类型、抽象定义、框架实现、Debug工具等  
generators.go 给出了 sourceOperater 的通用实现和具体函数实现  
transforms.go 给出了 transOperater 的通用实现和具体函数实现

## 使用案例
main.go代码：
```
package main

import (
	"fmt"
	"time"

	rxgo "github.com/pmlpml/rxgo"
)

func main() {
	fmt.Println("Debounce:")
	rxgo.Just(1, 2, 3, 4, 5, 6).Map(func(x int) int {
		switch x {
		case 1:
			time.Sleep(0 * time.Millisecond)
		case 2:
			time.Sleep(11 * time.Millisecond)
		case 3:
			time.Sleep(15 * time.Millisecond)
		case 4:
			time.Sleep(5 * time.Millisecond)
		case 5:
			time.Sleep(12 * time.Millisecond)
		case 6:
			time.Sleep(2 * time.Millisecond)
		}
		return x
	}).Debounce(10 * time.Millisecond).Subscribe(func(x int) {
		fmt.Print(x)
	})
	fmt.Println()

	fmt.Println("Distinct:")
	rxgo.Just(1, 2, 1, 3, 1, 4, 2, 4).Distinct().Subscribe(func(x int) {
		fmt.Print(x)
	})
	fmt.Println()

	fmt.Println("ElementAt:")
	for i := 0; i < 6; i++ {
		rxgo.Just(1, 2, 3, 4, 5, 6).ElementAt(i).Subscribe(func(x int) {
			if x%2 == 1 {
				fmt.Printf("%d:%d\n", i, x)
			}
		})
	}
	fmt.Println()

	fmt.Println("First:")
	rxgo.Just(11, 22, 33, 44).First().Subscribe(func(x int) {
		fmt.Print(x)
	})
	fmt.Println()

	fmt.Println("IgnoreElements:")
	rxgo.Just(1, 2, 3, 4, 5, 6).IgnoreElements().Subscribe(func(x int) {
		fmt.Print(x)
	})
	fmt.Println()

	fmt.Println("Last:")
	rxgo.Just(1, 2, 3, 4, 5, 6).Last().Subscribe(func(x int) {
		fmt.Print(x)
	})
	fmt.Println()

	fmt.Println("Sample:")
	rxgo.Just(1, 2, 3, 4, 5, 6).Map(func(x int) int {
		switch x {
		case 1:
			time.Sleep(0 * time.Millisecond)
		case 2:
			time.Sleep(10 * time.Millisecond)
		case 3:
			time.Sleep(5 * time.Millisecond)
		case 4:
			time.Sleep(20 * time.Millisecond)
		case 5:
			time.Sleep(20 * time.Millisecond)
		case 6:
			time.Sleep(50 * time.Millisecond)
		}
		return x
	}).Sample(25 * time.Millisecond).Subscribe(func(x int) {
		fmt.Print(x)
	})
	fmt.Println()

	fmt.Println("Skip:")
	rxgo.Just(1, 2, 3, 4, 5, 6).Skip(2).Subscribe(func(x int) {
		fmt.Print(x)
	})
	fmt.Println()

	fmt.Println("SkipLast")
	rxgo.Just(1, 2, 3, 4, 5, 6).SkipLast(2).Subscribe(func(x int) {
		fmt.Print(x)
	})
	fmt.Println()

	fmt.Println("Take:")
	rxgo.Just(1, 2, 3, 4, 5, 6).Take(2).Subscribe(func(x int) {
		fmt.Print(x)
	})
	fmt.Println()

	fmt.Println("TakeLast:")
	rxgo.Just(1, 2, 3, 4, 5, 6).TakeLast(2).Subscribe(func(x int) {
		fmt.Print(x)
	})
	fmt.Println()
}
```
运行结果：  
[pic1](img/1.jpg)

## 设计说明
详情查看[specification](https://gitee.com/zengty18342007/Service-Computing/blob/master/RxGO/specification.md)