package config

import (
	"fmt"
	"testing"
)

func Test_getConfig(t *testing.T) {
	c := GetConfig()
	if c == nil {
		fmt.Printf("not to get config for harbor\n")
		return
	}

	fmt.Printf("harbor config is %v\n", *c)

}
