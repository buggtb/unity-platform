package main

import (
	_ "unity-platform/routers"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	beego.Run()
}

