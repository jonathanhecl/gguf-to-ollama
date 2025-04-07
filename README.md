# gguf-to-ollama
Importing GGUF to Ollama made easy

## Features
- Import GGUF files to Ollama
- Auto-generate Modelfile
- Automatic STOP detection

## Requirements
- Ollama

## STOP supported tags
```
[INST] ... [/INST]
[SYSTEM_PROMPT] ... [/SYSTEM_PROMPT]
[AVAILABLE_TOOLS] ... [/AVAILABLE_TOOLS]
[TOOL_RESULTS] ... [/TOOL_RESULTS]
[TOOL_CALL] ... [/TOOL_CALL]
<s> ... </s>
<start_of_turn> ... <end_of_turn>
<|im_start|> ... <|im_end|>
```

## Usage
```bash
gguf_to_ollama <gguf-file> [<name>] [-context=65536]
```

```
<gguf-file>
	GGUF file to import to Ollama

<name>
	Name of the model (optional)

-context
	Context size fixed (optional)
```


## License

[MIT](LICENSE)

## Links

- [Ollama](https://ollama.com/)
- [Releases](https://github.com/jonathanhecl/gguf-to-ollama/releases/)