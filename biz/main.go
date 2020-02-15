// package main

package biz // make biz package importable in allinfo.go temporarily

import (
	"fmt"

	// Multiple package declaration not allowed in the same directory.
	// Here's the error message if `"github.com/admacro/bonjour/biz"` is uncommented:
	//   found packages biz (info.go) and main (main.go) in /Users/james/prog/bonjour/biz
	// When Go tries to import package `"github.com/admacro/bonjour/biz"`, it found not only
	// files that have package declared as `biz`, but also `main` (this file). Thus the complain.
	// "github.com/admacro/bonjour/biz"

	// importing packages in the subdirectories
	"github.com/admacro/bonjour/biz/p1"
	"github.com/admacro/bonjour/biz/p2"
)

func main() {
	// once there is a file with a different package name as the directory, the package cannot
	// be imported anyhow, thus cannot be used.
	// fmt.Println(biz.Info())

	fmt.Println(p1.Info())
	fmt.Println(p2.Info())
}
