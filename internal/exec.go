package internal

import (
	"errors"
	"os"
	"os/exec"
	"strings"
)

// ExecInp 执行输入的命令
func ExecInp(input string) error {
	// 删除换行符
	// strings.TrimSuffix 传入一个字符串，并删除后缀 suffix
	// https://pkg.go.dev/strings#TrimSuffix
	input = strings.TrimSuffix(input, "\n")

	// 当命令后带参数时
	// 需要切分单独读取参数
	args := strings.Split(input, " ")

	switch args[0] {
	case "cd":
		// 不能没有指定cd的路径
		if len(args) < 2 {
			return errors.New("路径不存在")
		}
		return os.Chdir(args[1])

	// 当输入 'exit' 退出程序
	// 等同于 ctrl + c
	case "exit":
		os.Exit(0)
	}

	// 创建命令
	cmd := exec.Command(args[0], args[1:]...)

	// 设置输出到控制台
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	// 执行命令
	return cmd.Run()
}
