package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/admacro/bonjour/salut"
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
