package main

import (
	"bufio"
	"os"
)

func main() {
	file, err := os.Create("output.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// 创建带缓冲的写入器（默认缓冲区大小4096字节）
	writer := bufio.NewWriter(file)

	// 写入数据（先存到缓冲区）
	writer.WriteString("Hello, bufio!\n")
	writer.WriteString("这是带缓冲的写入\n")

	// 手动刷新缓冲区（确保数据写入文件）
	writer.Flush()
}
