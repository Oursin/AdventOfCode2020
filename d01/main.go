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
	data, err := lib.ConvToInt(content)
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < len(data); i++ {
		for j := i + 1; j < len(data); j++ {
			if data[i] + data[j] == 2020 {
				fmt.Println(i, j, data[i], data[j], data[i] + data[j], data[i] * data[j])
			}
		}
	}
}
