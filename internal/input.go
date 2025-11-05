package internal

import (
	"bufio"
	"os"
)

// ReadInput 封装读取逻辑（目前未用）
func ReadInput() (*bufio.Reader, error) {
	return bufio.NewReader(os.Stdin), nil
}
