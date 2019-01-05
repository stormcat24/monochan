package main

import (
	"fmt"
	"log"

	"github.com/stormcat24/monochan/pkg/lib/liba"
)

func main() {
	fmt.Println("*** gazelle main ***")
	liba.MethodA()
	if err := liba.MethodB(false); err != nil {
		log.Fatal(err)
	}
}
