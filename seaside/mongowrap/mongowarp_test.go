package mongowrap

import (
	"fmt"
	"log"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	acc := loadSetting()

	fmt.Println(acc)
}

func TestConnect(t *testing.T) {
	log.Println(Connect())
}

func TestSelectCollection(t *testing.T) {

}
