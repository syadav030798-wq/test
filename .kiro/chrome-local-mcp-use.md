Use when agent needs to open and inspect the browser running localhost.

1. Start debug chrome

```bash
/Applications/Google\ Chrome.app/Contents/MacOS/Google\ Chrome \
    --remote-debugging-port=9222 \
    --user-data-dir=/tmp/chrome-debug

 Keep this chrome window running. MCP attaches to port 9222; it does not use normal chrome profile


2. Prompt away. LLM will decide to use mcp based on prompt or explicitly mention it in prompts.