package cli

import (
	"fmt"
	"gsh/internal"
	"os"
	"os/user"
	"strings"

	"github.com/chzyer/readline"
)

func RunShell(rl *readline.Instance) {
	var history []string
	for {
		// 获取当前目录和用户名
		dir, _ := os.Getwd()
		//baseDir := filepath.Base(dir)
		usr, _ := user.Current()

		prompt := fmt.Sprintf("\033[1;34m%s\033[0m:\033[1;32m%s\033[0m$ ", usr.Username, dir)
		rl.SetPrompt(prompt)

		// 读取输入
		s, err := rl.Readline()
		if err != nil {
			if err == readline.ErrInterrupt {
				fmt.Println("\n输入 Ctrl-C，退出 shell")
				break
			}
			fmt.Println("读取输入失败:", err)
			break
		}

		s = strings.TrimSpace(s)
		if s == "" {
			continue
		}

		history = append(history, s)
		rl.SaveHistory(s)

		// 执行命令
		err = internal.ExecInp(s)
		if err != nil {
			if err.Error() == "exit" { // internal.ExecInp 返回 exit 信号
				break
			}
			fmt.Println(err)
		}
	}
}
