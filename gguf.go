package main

import (
	"fmt"
	"os"
	"strings"
)

var GGUF_STOP_TAGS []string = []string{
	"[INST]", "[/INST]",
	"[SYSTEM_PROMPT]", "[/SYSTEM_PROMPT]",
	"[AVAILABLE_TOOLS]", "[/AVAILABLE_TOOLS]",
	"[TOOL_RESULTS]", "[/TOOL_RESULTS]",
	"[TOOL_CALL]", "[/TOOL_CALL]",
	"<<SYS>>", "<</SYS>>",
	"<s>", "</s>",
	"<start_of_turn>", "<end_of_turn>",
	"<|start_header_id|>", "<|end_header_id|>",
	"<｜begin▁of▁sentence｜>", "<｜begin▁of▁sentence｜>",
	"<｜User｜>", "<｜Assistant｜>",
	"<|eot_id|>", "<|im_sep|>",
	"<|reserved_special_token",
}

func GetGGUFStops(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("error opening file:", err)
		return nil, err
	}
	defer file.Close()

	stopTags := []string{}

	buf := make([]byte, 32*1024) // 32KB
	n, err := file.Read(buf)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	for _, tag := range GGUF_STOP_TAGS {
		if strings.Contains(string(buf[:n]), tag) {
			stopTags = append(stopTags, tag)
		}
	}

	return stopTags, nil
}
