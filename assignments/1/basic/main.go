package main

import (
	"fmt"

	"github.com/otus-golang/assignments/1/basic/hellonow"
	"github.com/otus-golang/tools"
)

func main() {
	time, err := hellonow.GetTime()
	tools.CheckAndPanicError(err)

	fmt.Printf("The most precision time is: %s\n", time)
}
