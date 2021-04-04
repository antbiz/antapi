package main

import (
	_ "antapi/boot"
	_ "antapi/router"

	"github.com/gogf/gf/frame/g"
)

func main() {
	g.Server().Run()
}
