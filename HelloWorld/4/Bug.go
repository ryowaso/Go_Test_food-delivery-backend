package main

import "fmt"

// 全局调试模式变量
var debugMode bool

// 状态管理器：用于在函数执行前后自动管理状态
func manageDebugState(originalState bool, enableDebug func(), disableDebug func()) func() {
	// 修改状态为调试模式
	enableDebug()
	fmt.Println("调试模式已开启")

	// 返回一个匿名函数，用于恢复原始状态
	return func() {
		disableDebug()
		fmt.Printf("调试模式已恢复为: %t\n", originalState)
	}
}

func main() {
	// 原始调试模式状态
	originalDebug := debugMode

	// 开启调试模式的函数
	enableDebug := func() {
		debugMode = true
	}

	// 关闭调试模式的函数
	disableDebug := func() {
		debugMode = originalDebug
	}

	// 使用 defer 确保调试模式一定会被恢复
	defer manageDebugState(originalDebug, enableDebug, disableDebug)()

	// 模拟业务逻辑
	fmt.Println("当前调试模式:", debugMode)
	fmt.Println("执行业务逻辑...")
	// 执行一些需要调试模式的操作
}
