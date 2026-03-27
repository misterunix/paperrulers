# Copilot Workspace Instructions for paperrulers

## Overview
This project is a Go CLI tool for generating printable PDF paper rulers and patterns. All logic is in the root package, with modular drawing code and a global configuration pattern.

## Build & Run
- **Build:** Run `./build.sh` to build the binary to `bin/paperrulers-linux-amd64`.
- **Run:** Execute the binary with CLI flags to generate PDFs. All output is written to the `pdf/` directory.
- **Dependencies:** Ensure Go modules are up to date (`go mod tidy`).
- **No automated tests:** Manual validation is required for changes.

## Key Files & Structure
- `main.go`: Entry point, CLI flag parsing, main logic
- `common.go`: Global config/state (`Opt` struct)
- `drawing.go`, `ladder.go`, `lines.go`: Drawing logic (PDF, ladder/blackletter, lines)
- `build.sh`: Build script
- `pdf/`: Output directory for generated PDFs
- `bin/`: Output directory for built binaries

## Conventions & Patterns
- **Global State:** All config/state is in the global `Opt` variable (see `common.go`).
- **CLI Flags:** All runtime options are set via CLI flags (see `main.go`).
- **Output Naming:** Output PDF filenames encode parameters (e.g., `dots-LETTER-L-1.500000.pdf`).
- **No sub-packages:** All code is in the root package.

## Pitfalls & Environment Notes
- The `pdf/` directory must exist before running; create it if missing.
- Only certain paper sizes/orientations are valid; invalid values cause exit.
- The build script targets Linux AMD64; adjust for other platforms as needed.
- No automated tests; changes must be manually validated.

## Example Prompts
- "Build the project and generate a PDF with A4 paper, landscape orientation, and 1.5mm spacing."
- "Add a new CLI flag for custom margin size."
- "Refactor drawing logic to support additional paper sizes."

## Next Steps / Customizations
- Consider adding test instructions or a test script for manual validation.
- For complex changes, create applyTo-based instructions for drawing logic vs. CLI/flag handling.
