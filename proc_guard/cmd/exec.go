package cmd

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"proc_guard/global"
	"strings"
)

func IsProcessExist(appName string) (bool, string) {
	cmd := exec.Command("cmd", "/C", "tasklist")
	output, _ := cmd.Output()
	//fmt.Printf("fields: %v\n", output)
	n := strings.Index(string(output), "System")
	if n == -1 {
		fmt.Println("no find")
		os.Exit(1)
	}
	data := string(output)[n:]
	fields := strings.Fields(data)
	for _, v := range fields {
		if v == appName {
			return true, ""
		}
	}

	return false, appName
}

func LoadRunTxt() {
	//fmt.Println("cmd init")
	// 指定文件路径
	filePath := "run.txt"

	// 打开文件
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("打开文件错误:", err)
		global.Logger.Error("打开文件错误:" + err.Error())
		return
	}
	defer file.Close()

	// 使用 bufio 包创建一个 scanner
	scanner := bufio.NewScanner(file)

	// 逐行读取文件
	for scanner.Scan() {
		// 获取每一行的内容
		line := scanner.Text()
		if line == "" {
			continue
		}

		// 使用 # 号分割每一行
		parts := strings.Split(line, "#")
		fmt.Println("内容", parts[:])
		global.RunPath[parts[0]] = parts[1]
	}

	// 检查是否有错误导致扫描结束
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	for s, s2 := range global.RunPath {
		fmt.Println("守护进程名:", s, "进程路径:", s2)
	}
}
