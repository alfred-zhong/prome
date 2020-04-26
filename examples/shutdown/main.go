package main

import (
	"fmt"
	"time"

	"github.com/alfred-zhong/prome"
)

func main() {
	client := prome.NewClient("test", "/foo")
	go func() {
		if err := client.ListenAndServe(":9000"); err != nil {
			panic(err)
		}
		fmt.Println("server shutdown")
	}()

	time.Sleep(5 * time.Second)
	if err := client.Close(); err != nil {
		panic(err)
	}
	time.Sleep(10 * time.Second)
}
