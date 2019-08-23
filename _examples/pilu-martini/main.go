package main

import (
	"github.com/lordtatty/martini"
)

func main() {
	m := martini.Classic()
	m.Get("/", func() string {
		return "Hello world - Martini"
	})
	m.Run()
}
