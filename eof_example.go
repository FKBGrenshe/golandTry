package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	// 示例1: 从字符串读取，模拟文件结束
	fmt.Println("=== 示例1: 从字符串读取 ===")
	reader := bufio.NewReader(strings.NewReader("Hello\nWorld"))
	readUntilEOF(reader)

	// 示例2: 从标准输入读取（需要用户输入）
	fmt.Println("\n=== 示例2: 从标准输入读取 ===")
	fmt.Println("请输入一些文本，然后按Ctrl+D(Unix)或Ctrl+Z(Windows)结束:")
	stdinReader := bufio.NewReader(os.Stdin)
	readUntilEOF(stdinReader)
}

// readUntilEOF 演示EOF错误处理
func readUntilEOF(in *bufio.Reader) {
	lineCount := 0
	for {
		line, err := in.ReadString('\n')
		if err == io.EOF {
			// 文件结束 - 正常退出循环
			fmt.Printf("文件结束，共读取了 %d 行\n", lineCount)
			break
		}
		if err != nil {
			// 其他错误 - 需要处理
			fmt.Printf("读取失败: %v\n", err)
			break
		}
		// 正常读取到数据
		lineCount++
		fmt.Printf("第%d行: %s", lineCount, line)
	}
}

// 另一个示例：读取固定字节数
func readFixedBytesExample() {
	fmt.Println("\n=== 示例3: 读取固定字节数 ===")

	data := "Hello, 世界!"
	reader := strings.NewReader(data)

	// 尝试读取比实际数据更多的字节
	buf := make([]byte, 20)
	n, err := reader.Read(buf)

	if err == io.EOF {
		fmt.Printf("读取到文件结束，读取了 %d 字节: %s\n", n, string(buf[:n]))
	} else if err != nil {
		fmt.Printf("读取错误: %v\n", err)
	} else {
		fmt.Printf("成功读取 %d 字节: %s\n", n, string(buf[:n]))
	}
}
