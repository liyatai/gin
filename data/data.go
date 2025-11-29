package data

import (
	"bytes"
	"text/template"
)

func Main(Package string) string {

	model := `
package main

import (
	"{{.package}}/config"
	"{{.package}}/controller"
	"{{.package}}/cor"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(cor.Cors())
	controller.GetRouter(r)
	r.Run(":" + config.GetConfig("server.port"))
}	
	`
	data := map[string]interface{}{
		"package": Package,
	}
	tmpl, err := template.New("userInfo").Parse(model)
	if err != nil {
		panic(err)
	}
	var resultBuf bytes.Buffer

	err = tmpl.Execute(&resultBuf, data)
	if err != nil {
		panic(err)
	}
	// 5. 将缓冲区内容转为字符串变量
	resultStr := resultBuf.String()
	// fmt.Println(resultStr)
	return resultStr
}

func Runner() string {

	model := `
root:              .                # 应用根目录
tmp_path:          ./tmp            # 临时文件目录
build_name:        runner-build     # 构建产物名称
build_log:         runner-build-errors.log  # 构建错误日志文件
valid_ext:         .go, .tpl, .tmpl, .html, .yml  # 监听的文件扩展名
no_rebuild_ext:    .tpl, .tmpl, .html       # 无需重新构建的文件扩展名（仅重启应用）
ignored:           assets, tmp      # 忽略的目录
build_delay:       200              # 构建延迟（毫秒），避免频繁触发
colors:            1                # 是否启用彩色日志（1=启用，0=禁用）
log_color_main:    cyan             # 主日志颜色（青色）
log_color_build:   yellow           # 构建日志颜色（黄色）
log_color_runner:  green            # 运行器日志颜色（绿色）
log_color_watcher: magenta          # 监听器日志颜色（品红色）
log_color_app:                      # 应用日志颜色（留空使用默认颜色）
	`

	return model
}

func Service() string {

	model := `
package service

import "github.com/gin-gonic/gin"

func Hello(c *gin.Context) {
	c.String(200, "欢迎使用 ★gin★ 框架开发您的 🌐web🌐 程序！✨")
}

	`

	return model
}

func Db(Package string) string {

	model := `
package db

import (
	"fmt"
	"{{.package}}/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func init() {
	dsn := config.GetConfig("mysql.user") + ":" + config.GetConfig("mysql.password") + "@tcp(" + config.GetConfig("mysql.url") + ":" + config.GetConfig("mysql.port") + ")/" + config.GetConfig("mysql.name") + "?charset=utf8mb4&parseTime=True&loc=Local"
	// dsn := config.GetConfig("sqlite")
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
}

	`
	data := map[string]interface{}{
		"package": Package,
	}
	tmpl, err := template.New("userInfo").Parse(model)
	if err != nil {
		panic(err)
	}
	var resultBuf bytes.Buffer

	err = tmpl.Execute(&resultBuf, data)
	if err != nil {
		panic(err)
	}
	// 5. 将缓冲区内容转为字符串变量
	resultStr := resultBuf.String()
	// fmt.Println(resultStr)
	return resultStr
}

func Cors() string {

	model := `
package cor

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 处理跨域请求,支持options访问
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Header("Access-Control-Allow-Headers", "*")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	}
}
	`
	return model
}

func Controller(Package string) string {

	model := `
package controller

import (
	"fmt"
	"ginTemplate/service"

	"github.com/gin-gonic/gin"
)

// PrintAllRoutes 打印所有已注册的路由
func PrintAllRoutes(engine *gin.Engine) {
	fmt.Println("======= 所有注册的路由 =======")
	// 遍历所有HTTP方法的路由树
	for _, methodTree := range engine.Routes() {
		// methodTree 包含 Method（GET/POST）、Path（路由路径）、Handler（处理器名称）
		fmt.Printf("方法: %-6s 路径: %-20s 处理器: %s\n",
			methodTree.Method, methodTree.Path, methodTree.Handler)
	}
	fmt.Println("==============================")
}

// 业务层
func GetRouter(r *gin.Engine) {

	r.GET("/", service.Hello)

	PrintAllRoutes(r)

}

	`
	data := map[string]interface{}{
		"package": Package,
	}
	tmpl, err := template.New("userInfo").Parse(model)
	if err != nil {
		panic(err)
	}
	var resultBuf bytes.Buffer

	err = tmpl.Execute(&resultBuf, data)
	if err != nil {
		panic(err)
	}
	// 5. 将缓冲区内容转为字符串变量
	resultStr := resultBuf.String()
	// fmt.Println(resultStr)
	return resultStr
}

func Config() string {

	model := `
package config

import (
	"bytes"
	_ "embed"
	"fmt"

	"github.com/spf13/viper"
)

//go:embed config.yml
var configYAML []byte

func GetConfig(str string) string {
	viper.SetConfigType("yaml")
	if err := viper.ReadConfig(bytes.NewBuffer(configYAML)); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("找不到配置文件..")
		} else {
			fmt.Println("配置文件出错..")
		}
	}

	return viper.GetString(str)
}


	`

	return model
}

func ConfigYml() string {

	model := `mysql:
    url: 数据库地址
    port: 2323
    user: root
    password: rpt
    name: 数据库名字
server:
    port: 8080
sqlite: 文件地址`

	return model
}
