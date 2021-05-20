package main

import (
	_ "github.com/antbiz/antapi/internal/router"
	"github.com/gogf/gf/frame/g"
)

func main() {
	g.Server().Run()
}
