package main

import (
	"fmt"

	"github.com/vincentgong3mm/golang-httprest/seaside/mongowrap"
)

func main() {
	fmt.Println("hello world!")

	acc := mongowrap.LoadSetting()
	fmt.Println(acc)
}

