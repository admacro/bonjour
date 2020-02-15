package main

import (
	"fmt"

	// package biz cannot be used unless all files in biz have the same package name
	//   found packages biz (info.go) and main (main.go) in /Users/james/prog/bonjour/biz
	// "github.com/admacro/bonjour/biz"

	"github.com/admacro/bonjour/biz/p1"
	"github.com/admacro/bonjour/biz/p2"
)

func main() {
	fmt.Println(p1.Info())
	fmt.Println(p2.Info())
}
