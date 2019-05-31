package main

import (
	"flag"
	"fmt"
	"github.com/beevik/ntp"

	"github.com/otus-golang/tools"
)

func main() {
	ntpServer := flag.String("ntp", "0.beevik-ntp.pool.ntp.org", "ntp server string")

	flag.Parse()

	resp, err := ntp.Query(*ntpServer)
	tools.CheckAndPanicError(err)

	fmt.Printf("The most precision time is: %s\n", resp.Time)
}
