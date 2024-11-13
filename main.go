package main

import (
	"flag"
	"tiramisu/backend"
)

func main() {
	flag.Parse()
	backend.ConfigInit()
	backend.Main()
}
