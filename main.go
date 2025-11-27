package main

import (
	"fmt"
	"os"

	"github.com/liyatai/gin/data"
	execute "github.com/liyatai/gin/exec"
	"github.com/liyatai/gin/writer"
)

func generator(arg string) {
	data.Main("avalon")
	writer.WriteStringToFile(data.Main(arg), "./main.go")
	writer.WriteStringToFile(data.Runner(), "./runner.conf")
	writer.WriteStringToFile(data.Service(), "./service/service.go")
	writer.WriteStringToFile(data.Db(arg), "./db/db.go")
	writer.WriteStringToFile(data.Cors(), "./cor/cor.go")
	writer.WriteStringToFile(data.Controller(arg), "./controller/controller.go")
	writer.WriteStringToFile(data.Config(), "./config/config.go")
	writer.WriteStringToFile(data.ConfigYml(), "./config/config.yml")

	// 执行命令
	execute.DoCommand("go", "mod", "init", arg)
	execute.DoCommand("go", "mod", "tidy")
}

func main() {
	// fmt.Println("传入：", os.Args[len(os.Args)-1])
	arg := os.Args[len(os.Args)-1]
	if arg == "-v" {
		fmt.Println("gin generator v1.0.2")
		fmt.Println("author:liyatai")
		fmt.Println("blog:https://blog.lyt11.cn")
	} else {
		generator(arg)
	}
}
