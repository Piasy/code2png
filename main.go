package main

import (
	_ "github.com/Piasy/code2png/routers"

	"github.com/astaxie/beego"
)

func main() {
	beego.BConfig.WebConfig.DirectoryIndex = true
	beego.BConfig.WebConfig.StaticDir["/images"] = "images"
	beego.Run()
}
