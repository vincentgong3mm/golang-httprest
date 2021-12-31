package mongowrap

import (
	"fmt"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	acc := LoadSetting()

	fmt.Println(acc)
}
