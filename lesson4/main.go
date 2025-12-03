package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"time"
)

func main() {
	//LV1
	file, err := os.Create("iotest.txt")
	if err != nil {
		fmt.Println("创建文件失败", err)
		return
	}
	file, err1 := os.OpenFile("iotest.txt", os.O_RDWR|os.O_APPEND, 0666)
	if err1 != nil {
		fmt.Println("打开文件失败", err1)
		return
	}
	defer file.Close()
	//不带缓冲I/O写
	starttime1 := time.Now()
	for i := 0; i < 10000; i++ {
		file.Write([]byte("无缓冲带写入内容\n"))
	}
	cost1 := time.Since(starttime1)
	//带缓冲带I/O写

	writer := bufio.NewWriter(file)
	defer writer.Flush()
	starttime2 := time.Now()
	for i := 0; i < 10000; i++ {
		_, err3 := writer.WriteString("带缓冲带的写入内容\n")
		if err3 != nil {
			fmt.Println("写入缓冲区失败", err)
			return
		}
	}
	cost2 := time.Since(starttime2)
	fmt.Printf("cost1: %v, cost2: %v\n", cost1, cost2)

	//LV2
	// 打开一个日志文件，如果文件不存在则创建，追加写入
	file2, err := os.OpenFile("example.txt", os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("打开文件失败：%v", err)
	}
	defer file2.Close()
	// 创建一个带时间戳的写入器
	multiWriter := io.MultiWriter(file2, os.Stdout)
	logWriter := newTimestampWriter(multiWriter)
	// 模拟用户操作并记录日志
	fmt.Fprintln(logWriter, "用户登录")
	time.Sleep(2 * time.Second)
	fmt.Fprintln(logWriter, "用户执行操作A")
	time.Sleep(1 * time.Second)
	fmt.Fprintln(logWriter, "用户执行操作B")
}

// timestampWriter 是一个实现 io.Writer 接口的结构体，它在写入数据前添加时间和时间戳

type timestampWriter struct {
	logFile io.Writer
}

// 传入一个io.writer,file实现了io.writer接口

func newTimestampWriter(w io.Writer) *timestampWriter {
	return &timestampWriter{logFile: w}
}
func (tw *timestampWriter) Write(p []byte) (n int, err error) {
	// 添加时间戳和时间
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	content := string(p)
	fullcontent := fmt.Sprintf("%s %s\n", timestamp, content)
	// 输出到文件
	return tw.logFile.Write([]byte(fullcontent))
}
