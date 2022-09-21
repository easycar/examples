package main

import (
	"flag"
	"fmt"
	"github.com/easycar/examples/withoutclient/commands"
)

var (
	easyCarAddr = flag.String("addr", "localhost:8089", "the address to connect easycar server")
)

func main() {
	commands.MustLoad(*easyCarAddr)

	flag.Parse()

	if err := commands.RunDemo(); err != nil {
		fmt.Println(err)
	}
}
