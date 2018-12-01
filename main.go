package main

import (
	"flag"
	"fmt"
	"os"
	"github.com/frankpf/undbg/undbg"
)

func main() {
	flag.Parse()
	input := flag.Arg(0)

	if input == "" {
		fmt.Println("Usage: undbg <target>")
		os.Exit(1)
	}

	undbg.Start(input)
}
