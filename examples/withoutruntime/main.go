package main

import (
	"github.com/alfred-zhong/prome"
)

func main() {
	client := prome.NewClient("test", "/foo")
	client.EnableRuntime = false
	if err := client.ListenAndServe(":9000"); err != nil {
		panic(err)
	}
}
