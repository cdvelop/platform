# Website Component Structure

This directory contains the website rebuilt using TinyWasm components.

## Structure

- `components/`: Contains individual UI components.
    - Each component has:
        - `component.go`: Logic and rendering.
        - `ssr.go`: CSS embedding (build tag `!wasm`).
        - `front.go`: WASM-specific logic (build tag `wasm`).
- `models/`: Data models adapted for use with `tinywasm/form`.
- `main.go`: WASM entry point.
- `index.html`: HTML shell.

## Build

To build the WASM binary:

```bash
GOOS=js GOARCH=wasm go build -o main.wasm main.go
```

Note: CSS is embedded in `ssr.go` files and requires a backend process (like `tinywasm/site`) to extract and serve it, or manual extraction.
