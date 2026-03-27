# paperrulers (Temporary README)

**Note:** This README is temporary. The program is in active development and both features and CLI usage are subject to change.

## Overview

paperrulers is a Go CLI tool for generating printable PDF paper rulers and patterns. It supports various paper sizes, orientations, and custom spacing.

## Building

```
./build.sh
```

- Output binary: `bin/paperrulers-linux-amd64`
- Requires Go modules (run `go mod tidy` if needed)

## Running (CLI Example)

```
# Example usage (flags/options may change soon):
./bin/paperrulers-linux-amd64 -paper A4 -orient L -spacing 1.5
```

- Output PDFs are written to the `pdf/` directory.
- All options are set via CLI flags (see `main.go` for current list).

## Key Files

- `main.go` — CLI parsing, main logic
- `common.go` — Global config/state
- `drawing.go`, `ladder.go`, `lines.go` — Drawing logic
- `build.sh` — Build script

## Notes

- The `pdf/` directory must exist before running.
- No automated tests; validate output manually.
- Only certain paper sizes/orientations are valid.

---
This file will be updated as the program evolves. Please check back for the latest usage and options.
