package test

import (
	"gsh/internal"
	"testing"
)

func TestExecInput(t *testing.T) {
	// 模拟用户输入的命令
	commands := []string{
		"pwd",
		"cd ..",
		"echo HelloWorld",
	}

	for _, cmd := range commands {
		err := internal.ExecInp(cmd)
		if err != nil {
			t.Errorf("执行命令失败: %s, err: %v", cmd, err)
		}
	}
}
