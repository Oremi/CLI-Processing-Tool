# CLI-Processing-Tool

[![Go Reference](https://pkg.go.dev/badge/github.com/Oremi/CLI-Processing-Tool.svg)](https://pkg.go.dev)
[![Go Version](https://img.shields.io/badge/go-1.20+-blue)](https://go.dev)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg)]()

## What the project does

`CLI-Processing-Tool` is a lightweight text processing CLI utility written in Go. It reads a plaintext input file, applies a stepwise transformation pipeline, and writes formatted output. The processor handles:

- binary/hex number conversion markers `(... (bin), ... (hex))`
- case conversions `(up)`, `(low)`, `(cap)` and multi-word forms like `(up, N)`
- punctuation spacing sanity (`. , ! ? : ;`)
- quote trimming (`' quoted text ' -> 'quoted text'`)
- article correction (`a` -> `an` before vowels and `h`)

## Why the project is useful

- Designed as a learning example of string parsing and slice manipulation in Go
- Useful for text pre-processing in data-cleanup workflows
- Demonstrates idiomatic testing with Go's `testing` package

## Key features

- Converts marked numeric tokens to decimal
- Normalizes uppercase/lowercase/capitalized words by marker
- Collapses repeated spaces and fixes punctuation spacing
- Auto-fixes article usage for vowels/h
- Opinionated quote handling to remove inner spacing

## How users can get started

### Prerequisites

- Go 1.20+

### Build and run

```bash
$ cd CLI-Processing-Tool
$ go run . sample.txt result.txt
```

### With `go install`

```bash
$ cd CLI-Processing-Tool
$ go install ./...
$ CLI-Processing-Tool input.txt output.txt
```

### Example

Input `sample.txt`:

```
I was thinking ... You were right
```

Run:

```bash
go run . sample.txt result.txt
```

Output `result.txt`:

```
I was thinking... You were right
```

### Run tests

```bash
$ go test ./...
```


---

## Supported Markers


| Marker | Description |
|------|-------------|
| `(bin)` | Convert binary to decimal |
| `(hex)` | Convert hexadecimal to decimal |
| `(up)` | Uppercase previous word |
| `(low)` | Lowercase previous word |
| `(cap)` | Capitalize previous word |
| `(up, N)` | Uppercase previous N words |

---

## Project Structure

```text

CLI-Processing-Tool/
│
├── main.go          # CLI entrypoint and processing pipeline
├── ├── processor/
│   └── processor.go # transformation functions
├── main_test.go     # unit tests
├── go.mod           # go module file
├── sample.txt       # example input
├── result.txt       # example output
├── CONTRIBUTING.md
├── LICENSE.md
└── README.md
```
### Learning Objectives

- String manipulation in Go
- File I/O operations
- CLI argument parsing
- Building transformation pipelines
- Writing unit tests with Go's testing package

## Where users can get help

- Raise an issue in this repository
- Add a `discussion` in GitHub Discussions
- Inspect tests in `main_test.go` for expected processor behavior

## Maintainer

This project is maintained by **[@Oremi](https://github.com/Oremi)**.

Contributions are welcome!  
Please submit a pull request or open an issue to discuss changes.

## Contributing Guidelines

For details on development setup, tests, style, and pull request expectations, see [CONTRIBUTING.md](CONTRIBUTING.md).

## About the Learning Context

This repository contains a project completed as part of the
[01Edu](https://github.com/01edu) software engineering curriculum.

[![01Edu GitHub](https://img.shields.io/badge/01Edu-GitHub-blue?logo=github)](https://github.com/01edu)

The program focuses on project-based learning, peer reviews,
and building practical programming skills through real-world tasks.
