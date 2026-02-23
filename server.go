package main

import (
	"fmt"
	"net/http"
)

func main() {
	// 设置端口
	port := "8000"
	
	// 将当前目录作为根目录提供文件服务
	// 这等同于 python -m http.server 8000
	fs := http.FileServer(http.Dir("."))
	http.Handle("/", fs)

	// 控制台提示
	fmt.Printf("--------------------------------------------------\n")
	fmt.Printf("  AutoLiveMusic 服务已启动\n")
	fmt.Printf("  地址: http://localhost:%s\n", port)
	fmt.Printf("  请在直播软件中添加浏览器源，地址填写上面一行。\n")
	fmt.Printf("  请勿关闭此窗口，关闭将停止服务。\n")
	fmt.Printf("--------------------------------------------------\n")

	// 启动服务器
	err := http.ListenAndServe(":"+port, nil)
	
	// 如果启动失败（例如端口被占用），打印错误
	if err != nil {
		fmt.Printf("\n[错误] 启动失败: %v\n", err)
		fmt.Println("请检查端口 8000 是否被其他程序占用。")
		// 防止窗口一闪而过
		fmt.Scanln() 
	}
}
