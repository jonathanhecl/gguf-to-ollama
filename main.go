package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
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
		name = filepath.Base(ggufFile)
		ext := filepath.Ext(name)
		name = name[:len(name)-len(ext)]
	}

	fmt.Println() // break

	fmt.Println("GGUF ", ggufFile)
	fmt.Println("Name ", name)

	stops, err := GetGGUFStops(ggufFile)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("STOP detected: ", strings.Join(stops, ", "))

	modelfile := fmt.Sprintf(`FROM "%s"\n`, ggufFile)
	for _, stop := range stops {
		modelfile += fmt.Sprintf("PARAMETER stop %s\n", stop)
	}

}
