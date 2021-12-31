package main

import (
	"fmt"
	"testing"

	"github.com/vincentgong3mm/golanghttprest/seaside/mongowrap"
)

func TestMain(t *testing.T) {
	fmt.Println("test hello world")

	acc := mongowrap.Connect()
	fmt.Println(acc)
}

func TestLogPrint(t *testing.T) {
	NewSlog()
	sl.Info.Println("info log.")
	sl.Error.Println("error log.")
}
