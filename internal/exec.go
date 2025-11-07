package internal

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// ExecInp 执行输入的命令
func ExecInp(input string) error {
	input = strings.TrimSuffix(input, "\n")
	args := strings.Split(input, " ")

	switch args[0] {
	case "cd":
		if len(args) < 2 {
			return errors.New("路径不存在")
		}
		return os.Chdir(args[1])

	case "exit":
		os.Exit(0)
	}

	if args[0] == "pwd" {
		dir, err := os.Getwd()
		if err != nil {
			return err
		}
		fmt.Println(strings.TrimSpace(dir))
		return nil
	}

	// 执行其他系统命令
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	return cmd.Run()
}
