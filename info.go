package main

import (
	"fmt"

	"github.com/admacro/bonjour/biz"
	"github.com/admacro/bonjour/biz/p1"
	"github.com/admacro/bonjour/biz/p2"
)

func main() {
	fmt.Println(biz.Info())
	fmt.Println(p1.Info())
	fmt.Println(p2.Info())

	fmt.Println(biz.AllInfo())
}
