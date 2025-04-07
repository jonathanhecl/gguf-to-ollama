# gguf-to-ollama
Importing GGUF to Ollama made easy

## Features
- Import GGUF files to Ollama
- Auto-generate Modelfile
- Automatic STOP detection

## Requirements
- [Ollama](https://ollama.com/)

## STOP tags supported
```
[INST] ... [/INST]
[SYSTEM_PROMPT] ... [/SYSTEM_PROMPT]
[AVAILABLE_TOOLS] ... [/AVAILABLE_TOOLS]
[TOOL_RESULTS] ... [/TOOL_RESULTS]
[TOOL_CALL] ... [/TOOL_CALL]
<<SYS>> ... <</SYS>>
<s> ... </s>
<start_of_turn> ... <end_of_turn>
<|start_header_id|> ... <|end_header_id|>
<｜begin▁of▁sentence｜> ... <｜begin▁of▁sentence｜>
<｜User｜>
<｜Assistant｜>
<|eot_id|>
<|im_sep|>
<|reserved_special_token
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
	Context size fixed (optional, default by ollama: 2048)
```

## Modelfile

If modelfile already exists, it will not be overwritten.

You can edit the modelfile manually. [Documentation](https://github.com/ollama/ollama/blob/main/docs/modelfile.md)

## License

[MIT](LICENSE)

## Links

- [Releases](https://github.com/jonathanhecl/gguf-to-ollama/releases/)