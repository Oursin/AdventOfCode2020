package main

import (
	"AoC2020/lib"
	"fmt"
	"log"
)

func main() {
	content, err := lib.ReadFileLines("input")
	if err != nil {
		log.Fatal(err)
	}
	c := 0
	for _, l := range content {
		pass := getPass(l)
		fmt.Println(pass, pass.check1())
		if pass.check1() { // pass.check2 for part 2
			c++
		}
	}
	fmt.Println(c)
}
