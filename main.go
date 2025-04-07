package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

var (
	version = "0.1.0"
)

func main() {
	fmt.Println("gguf-to-ollama v", version)

	if len(os.Args) < 2 {
		fmt.Println("Usage: gguf-to-ollama <gguf-file> [<name>] [-context=65536]")
		return
	}

	ggufFile := os.Args[1]
	name := ""
	context := 0

	if len(os.Args) > 2 {
		for i := range os.Args {
			if strings.HasPrefix(os.Args[i], "-context=") {
				context, _ = strconv.Atoi(os.Args[i][len("-context="):])
			} else {
				name = os.Args[i]
			}
		}
	} else {
		name = filepath.Base(ggufFile)
		ext := filepath.Ext(name)
		name = name[:len(name)-len(ext)]
	}

	fmt.Println() // break

	fmt.Println("GGUF File : ", ggufFile)
	fmt.Println("Name : ", name)
	if context > 0 {
		fmt.Println("Context : ", context)
	}

	stops, err := GetGGUFStops(ggufFile)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("STOP detected : ", strings.Join(stops, ", "))

	modelFilePath := ggufFile + ".modelfile"

	if _, err := os.Stat(modelFilePath); err == nil {
		fmt.Println("Modelfile already exists:", modelFilePath)
	} else {
		modelfile := fmt.Sprintf("FROM \"%s\"\n", ggufFile)
		for _, stop := range stops {
			modelfile += fmt.Sprintf("PARAMETER stop \"%s\"\n", stop)
		}
		if context > 0 {
			modelfile += fmt.Sprintf("PARAMETER num_ctx %d\n", context)
		}
		fmt.Printf("Modelfile:\n----------\n%s\n----------\n", modelfile)

		if err := os.WriteFile(modelFilePath, []byte(modelfile), 0644); err != nil {
			fmt.Printf("Error saving modelfile: %v\n", err)
			return
		}
		fmt.Printf("Modelfile saved to: %s\n", modelFilePath)
	}

	fmt.Println() // break

	if _, err := os.Stat(modelFilePath); err == nil {
		fmt.Println("Creating model on Ollama...")
		fmt.Println("This may take a while...")

		cmd := exec.Command("ollama", "create", name, "-f", modelFilePath)
		output, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Printf("Error creating model: %v\n", err)
			fmt.Printf("Command output: %s\n", string(output))
			return
		}

		fmt.Printf("Model created successfully!\n")
		fmt.Printf("Output: %s\n", string(output))
	}
}
