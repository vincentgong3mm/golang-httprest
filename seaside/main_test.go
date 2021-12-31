package main

import (
	"fmt"
	"testing"

	"github.com/vincentgong3mm/golang-httprest/seaside/mongowrap"
)

func TestMain(t *testing.T) {
	fmt.Println("test hello world")

	acc := mongowrap.LoadSetting()
	fmt.Println(acc)

}
