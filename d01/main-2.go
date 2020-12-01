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
			for k := j + 1; k < len(data); k++ {
				if data[i] + data[j] + data[k] == 2020 {
					fmt.Println(i, j, k, data[i], data[j], data[k], data[i] + data[j] + data[k], data[i] * data[j] * data[k])
					return
				}
			}
		}
	}
}
