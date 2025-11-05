package main

import (
	"bufio"
	"fmt"
	"gsh/internal"
	"os"
)

// 示例假设是 Unix 系统。
// 它们可能无法在 Windows 上运行
// 也不会在 Go Playground 中运行 go.dev 和 internal.go.dev 使用
func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("please enter:")
	for {
		fmt.Print(">")
		// 读取键盘输入
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		internal.ExecInp(input)
	}
}
