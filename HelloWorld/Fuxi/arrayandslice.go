package main

func main() {
	//// 显式指定长度
	//var arr1 [3]int = [3]int{1, 2, 3}
	//
	//// 使用 ... 让编译器计算长度
	//arr2 := [...]int{1, 2, 3, 4}
	//
	//// 指定索引初始化
	//arr3 := [5]int{1: 10, 3: 30}
	//fmt.Println(arr1, arr2, arr3)
	//

	//// 直接声明
	//var s1 []int
	//
	//// 从数组创建
	//arr := [5]int{1, 2, 3, 4, 5}
	//s2 := arr[1:3] // [2, 3]
	//
	//// 使用 make 创建
	//s3 := make([]int, 3, 5) // 长度3，容量5
	//
	//// 字面量创建
	//s4 := []int{1, 2, 3}
	//fmt.Println(s1, s2, arr, s3, s4)
	//

	//// 遍历
	//for i, v := range arr {
	//	fmt.Println(i, v)
	//}
	//
	//// 长度
	//length := len(arr)
	//fmt.Println(length)
	//

	//// 添加元素
	//s = append(s, 4)
	//
	//// 复制切片
	//s2 := make([]int, len(s))
	//copy(s2, s)
	//
	//// 截取子切片
	//sub := s[1:3]
	//
	//// 容量扩展
	//newS := make([]int, len(s), cap(s)*2)
	//copy(newS, s)
	//fmt.Println(s2, sub, newS)
	/*一等公民的定义：

	在编程语言中，如果函数是一等公民，意味着函数可以：

	赋值给变量

	作为参数传递给其他函数

	作为其他函数的返回值

	存储在数据结构中
	*/

	/*基本类型的零值
	类型	        零值	           说明
	bool	    false	      布尔型
	int	        0	          所有整数类型（int8/16/32/64, uint等）
	float32/64	0.0	          浮点数
	string	    ""（空字符串）  字符串
	byte	    0	         （uint8 的别名）
	rune	    0	         （int32 的别名，表示Unicode码点）
	复合类型的零值
	类型	        零值	             说明
	数组	     所有元素为零值       如 [3]int → [0, 0, 0]
	结构体	 所有字段为零值	   递归初始化每个字段
	指针	      nil	           未指向任何内存地址
	切片	      nil	           长度=0，容量=0，底层数组=nil
	Map	      nil	           无法直接使用（需make初始化）
	Channel	  nil	           未初始化的通道
	函数	      nil	           未初始化的函数引用
	接口	      nil	           类型和值均为nil
	*/

	/*所以，if 的条件表达式可以是：

	布尔变量（isReady）

	比较运算（x == y、a > b）

	逻辑运算（a && b、!condition）

	返回 bool 的函数调用（strings.HasPrefix(s, "go")）

	带初始化的 if，但最终条件仍必须是 bool

	但 不能是数字、字符串、指针（除非显式比较）等非 bool 类型。
	*/

	/*
				不能用作 map 键的类型：
			切片（slice）

			因为切片是引用类型，不支持 == 比较。

			映射（map）

			同样不支持 == 比较。

			函数（func）

			函数不能比较。

			包含不可比较字段的结构体

		Go 的 map 键必须是可比较的类型，即支持 == 和 != 操作。通常使用 string、int、float64
		等基本类型作为键，而 slice、map 和 func 不能作为键。
	*/
}
