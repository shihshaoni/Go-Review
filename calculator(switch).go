// 小專案：命令行計算器
// 功能要求：

// 用戶輸入兩個數字。
// 用戶選擇運算類型（加法、減法、乘法、除法）。
// 使用 switch 語句根據運算符執行相應的計算。
// 輸出結果到終端。

// 說明：
// 使用 fmt.Scanln() 來接受用戶輸入。
// 透過 switch 語句來判斷用戶選擇的運算符號。
// 加入除法時防止除以 0 的處理邏輯。
package main

import "fmt"

func main() {
	var a, b float64
	var operator string
	fmt.Println("Type in a:")
	fmt.Scanln(&a)
	fmt.Println("Type in b (not 0):")
	fmt.Scanln(&b)
	fmt.Println("Type in operator(+,-,*,/ only):")
	fmt.Scanln(&operator)

	switch operator {
	case "+":
		fmt.Println(a + b)
	case "-":
		fmt.Println(a - b)
	case "*":
		fmt.Println(a * b)
	case "/":
		if b == 0 {
			fmt.Println("Type in b again:")
			fmt.Scanln(&b)
		}
		fmt.Println(a / b)
	}

}
