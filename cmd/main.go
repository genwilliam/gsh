package main

import (
	"fmt"
	cli "gsh/pkg"
	"os"

	"github.com/chzyer/readline"
)

func main() {
	rl, err := readline.New("")
	if err != nil {
		fmt.Println("初始化 readline 出错:", err)
		os.Exit(1)
	}
	defer rl.Close()

	cli.RunShell(rl)
}
