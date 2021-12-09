package main

import (
	"flag"
	"log"
)

func main() {
	var name string

	// 用标准库 flag 的 StringVar 方法实现了对命令行参数 name 的解析和绑定
	// 第二个参数为 标示位名称  第三个参数为默认值
	flag.StringVar(&name, "name", "默认Go Name", "帮助信息")
	flag.StringVar(&name, "n", "Go Name", "帮助信息")
	flag.Parse()
	log.Printf("name: %s", name)

	//go run main.go -name=测试 -n="233232"
	//	2021/12/09 11:12:53 name: 233232  (最终输出了-n的内容,flag包不支持长短选项)
}
