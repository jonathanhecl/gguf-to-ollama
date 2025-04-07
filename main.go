package main

import (
	"fmt"
	"os"
	"path"
)

var (
	version = "0.1.0"
)

func main() {
	fmt.Println("gguf-to-ollama v", version)

	if len(os.Args) < 2 {
		fmt.Println("Usage: gguf-to-ollama <gguf-file> [<name>]")
		return
	}

	name := ""
	ggufFile := os.Args[1]

	if len(os.Args) > 2 {
		name = os.Args[2]
	} else {
		name = ggufFile[:len(ggufFile)-len(path.Ext(ggufFile))]
	}

	fmt.Println("GGUF ", ggufFile)
	fmt.Println("Name ", name)
}
