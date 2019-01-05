package liba

import (
	"fmt"

	"github.com/BurntSushi/toml"
	"github.com/pkg/errors"
)

func MethodA() {
	fmt.Println("[called] liba.MethodA()")
}

func MethodB(success bool) error {
	fmt.Println("[called] liba.MethodB()")

	if !success {
		return errors.New("MethodB Error")
	}

	enc := toml.Encoder{}
	fmt.Println(enc)

	return nil
}
