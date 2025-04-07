# gguf-to-ollama
Importing GGUF to Ollama made easy

## Features
- Import GGUF files to Ollama
- Auto-generate Modelfile
- Automatic STOP detection

## STOP supported tags
```
[INST] ... [/INST]
[SYSTEM_PROMPT] ... [/SYSTEM_PROMPT]
[AVAILABLE_TOOLS] ... [/AVAILABLE_TOOLS]
[TOOL_RESULTS] ... [/TOOL_RESULTS]
[TOOL_CALL] ... [/TOOL_CALL]
<s> ... </s>
<start_of_turn> ... <end_of_turn>
```

## Usage
```bash
gguf_to_ollama <gguf-file> [<name>]
```