package main

import (
	"fmt"
	"github.com/admacro/bonjour/salut"
	"math/rand"
	"time"
)

func main() {
	for {
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		s := salut.Salut(rand.Intn(3))
		if s == "" {
			fmt.Println("Bonjour !")
		} else {
			fmt.Printf("Bonjour %s !\n", s)
		}
	}
}
