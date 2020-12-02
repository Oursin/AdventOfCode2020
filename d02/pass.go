package main

import (
	"AoC2020/lib"
	"log"
	"strings"
)

type pass struct {
	l string
	r []int
	pass string
}

func getPass(line string) pass {
	parts := strings.Split(line, " ")
	parts[1] = strings.TrimRight(parts[1], ":")
	r := strings.Split(parts[0], "-")
	i, err := lib.ConvToInt(r)
	if err != nil {
		log.Fatal(err)
	}
	return pass{
		l:    parts[1],
		r:    i,
		pass: parts[2],
	}
}

func (p pass) check1() bool {
	c := 0
	for i := 0; i < len(p.pass); i++ {
		if p.pass[i] == p.l[0] {
			c += 1
		}
	}
	if c >= p.r[0] && c <= p.r[1] {
		return true
	}
	return false
}

func (p pass) check2() bool {
	c := 0
	if p.pass[p.r[0] -1] == p.l[0] {
		c++
	}
	if p.pass[p.r[1] -1] == p.l[0] {
		c++
	}
	return c == 1
}