package main

import (
	"fmt"
	"os"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lifei6671/godoc/routers"
	_ "github.com/astaxie/beego/session/redis"
	_ "github.com/astaxie/beego/session/memcache"
	_ "github.com/astaxie/beego/session/mysql"
	"github.com/astaxie/beego"
	"github.com/lifei6671/godoc/commands"
	"github.com/lifei6671/godoc/controllers"
)

var (
	VERSION    string
	BUILD_TIME string
	GO_VERSION string
)


func main() {

	fmt.Printf("MinDoc version => %s\nbuild time => %s\nstart directory => %s\n%s\n", VERSION, BUILD_TIME, os.Args[0],GO_VERSION)

	commands.RegisterDataBase()
	commands.RegisterModel()
	commands.RegisterLogger()
	commands.RegisterCommand()
	commands.RegisterFunction()

	beego.SetStaticPath("uploads","uploads")

	beego.ErrorController(&controllers.ErrorController{})

	beego.Run()
}

