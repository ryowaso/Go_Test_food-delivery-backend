package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	for {
		fmt.Println("\n选择操作:")
		fmt.Println("1. 其他进制 → 十进制")
		fmt.Println("2. 十进制 → 其他进制")
		fmt.Println("3. 二进制 ↔ 八进制/十六进制")
		fmt.Println("4. 退出")
		fmt.Print("请输入选项: ")

		var choice int
		_, err := fmt.Scan(&choice)
		if err != nil {
			fmt.Println("输入错误，请重新输入")
			continue
		}

		switch choice {
		case 1:
			otherToDecimal()
		case 2:
			decimalToOther()
		case 3:
			binaryToOctalOrHex()
		case 4:
			fmt.Println("退出程序")
			return
		default:
			fmt.Println("无效选项")
		}
	}
}

// 其他进制转十进制
func otherToDecimal() {
	fmt.Print("输入数字: ")
	var numStr string
	fmt.Scan(&numStr)

	fmt.Print("输入当前进制 (2, 8, 16): ")
	var base int
	fmt.Scan(&base)

	// 处理带有前缀的情况 (0x, 0b)
	if strings.HasPrefix(numStr, "0x") || strings.HasPrefix(numStr, "0X") {
		numStr = numStr[2:]
		base = 16
	} else if strings.HasPrefix(numStr, "0b") || strings.HasPrefix(numStr, "0B") {
		numStr = numStr[2:]
		base = 2
	}

	numStr = strings.ToUpper(numStr) // 统一转为大写处理十六进制

	decimal, err := strconv.ParseInt(numStr, base, 64)
	if err != nil {
		fmt.Println("转换错误:", err)
		return
	}

	fmt.Printf("%s (基数%d) = %d (十进制)\n", numStr, base, decimal)
}

// 十进制转其他进制
func decimalToOther() {
	fmt.Print("输入十进制数字: ")
	var decimal int
	fmt.Scan(&decimal)

	fmt.Print("输入目标进制 (2, 8, 16): ")
	var base int
	fmt.Scan(&base)

	if base == 16 {
		fmt.Printf("%d (十进制) = 0x%X (十六进制)\n", decimal, decimal)
	} else if base == 8 {
		fmt.Printf("%d (十进制) = 0%o (八进制)\n", decimal, decimal)
	} else if base == 2 {
		fmt.Printf("%d (十进制) = %s (二进制)\n", decimal, strconv.FormatInt(int64(decimal), 2))
	} else {
		fmt.Println("不支持的目标进制")
	}
}

// 二进制与八进制/十六进制互转
func binaryToOctalOrHex() {
	fmt.Print("输入二进制数字: ")
	var binaryStr string
	fmt.Scan(&binaryStr)

	// 去除可能的0b前缀
	if strings.HasPrefix(binaryStr, "0b") || strings.HasPrefix(binaryStr, "0B") {
		binaryStr = binaryStr[2:]
	}

	// 验证是否为有效二进制
	for _, ch := range binaryStr {
		if ch != '0' && ch != '1' {
			fmt.Println("无效的二进制数字")
			return
		}
	}

	// 补零使长度为3的倍数(八进制)或4的倍数(十六进制)
	paddedBinary := binaryStr
	for len(paddedBinary)%3 != 0 {
		paddedBinary = "0" + paddedBinary
	}

	// 二进制转八进制
	var octalBuilder strings.Builder
	for i := 0; i < len(paddedBinary); i += 3 {
		chunk := paddedBinary[i : i+3]
		val, _ := strconv.ParseInt(chunk, 2, 64)
		octalBuilder.WriteString(fmt.Sprintf("%d", val))
	}
	fmt.Printf("八进制: 0%s\n", octalBuilder.String())

	// 重新补零为4的倍数
	paddedBinary = binaryStr
	for len(paddedBinary)%4 != 0 {
		paddedBinary = "0" + paddedBinary
	}

	// 二进制转十六进制
	var hexBuilder strings.Builder
	for i := 0; i < len(paddedBinary); i += 4 {
		chunk := paddedBinary[i : i+4]
		val, _ := strconv.ParseInt(chunk, 2, 64)
		hexBuilder.WriteString(fmt.Sprintf("%X", val))
	}
	fmt.Printf("十六进制: 0x%s\n", hexBuilder.String())
}
