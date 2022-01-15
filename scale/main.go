package main

import (
	"fmt"
	"strconv"
)

func main() {
	scale := scanlnScale()
	input := scanlnInput(scale)
	scalnValArr := []int{1: 2, 2: 8, 3: 10, 4: 16}
	scaleVal := scalnValArr[scale]
	convertScale(input, scaleVal)
}

/**
 *@note 选择输出的进制
 */
func scanlnScale() int {
	// 进制可选数组
	scalnArr := []int{1, 2, 3, 4}
	var scale int
	fmt.Println(`
请选择输入的内容进制：
1：2进制
2：8进制
3：10进制
4：16进制
	`)
	fmt.Scanln(&scale)

	if isContain(scalnArr, scale) == false {
		return scanlnScale()
	}
	return scale
}

/**
 *@note 输入需转换的内容
 */
func scanlnInput(scale int) string {
	var input string
	scalnArr := []string{"2进制", "8进制", "10进制", "16进制"}
	fmt.Printf("请输入要转换的内容（%v）：", scalnArr[scale-1])
	fmt.Scanln(&input)
	if input == "" {
		return scanlnInput(scale)
	}
	return input
}

/**
 *@note 判断数组中是否包含某个值
 */
func isContain(items []int, item int) bool {
	for _, eachItem := range items {
		if eachItem == item {
			return true
		}
	}
	return false
}

/**
 *@note 进制转换
 *@param string input 待转换的进制值
 *@param int scale 待转换的进制
 */
func convertScale(input string, scale int) {
	//指定进制（scale）转换成10进制
	i, _ := strconv.ParseInt(input, scale, 64)

	//10进制转换成其他进制
	var v int64 = i
	s2 := strconv.FormatInt(v, 2) //10 yo 2
	fmt.Printf("2进制： %v\n", s2)

	s8 := strconv.FormatInt(v, 8) //10 yo 8
	fmt.Printf("8进制： %v\n", s8)

	s10 := strconv.FormatInt(v, 10) //10 yo 10
	fmt.Printf("10进制：%v\n", s10)

	s16 := strconv.FormatInt(v, 16) //10 yo 16
	fmt.Printf("16进制：%v\n", s16)

	//继续执行
	var conti string
	fmt.Println("按下回车继续执行...")
	fmt.Scanln(&conti)
	if conti == "" {
		main()
	}
}
