package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// 输入一行整数数组,用空格分开
	fmt.Print("请输入一行整数(用空格分隔): ")
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	line = strings.TrimSpace(line) // 去除换行符和空格
	parts := strings.Split(line, " ")

	// 转换为整数数组
	intArr := make([]int, len(parts))
	for i, part := range parts {
		intArr[i], _ = strconv.Atoi(part)
	}

	fmt.Println("转换后的整数数组:", intArr)
}
