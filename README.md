# go-reloaded

A CLI tool for text formatting, auto-correction, and transformation, written in Go.

## Overview
`go-reloaded` takes a text file as input, applies a set of formatting rules (hex/bin conversion, casing, punctuation spacing, grammar correction), and outputs the cleaned text to a new file.

## Usage

```bash
go run . <input_file> <output_file>
```

### Example

```bash
go run . sample.txt result.txt
```

## Testing

To run the golden tests (audit cases + extra scenarios):

```bash
go test -v ./...
```

## Features & Rules
- **Number Conversion:** Converts `(hex)` and `(bin)` numbers to decimal.
- **Casing:** Changes words to `(up)`, `(low)`, or `(cap)` (supports counts like `(up, 2)`).
- **Punctuation:** Fixes spacing around `. , ! ? : ;` and preserves groups like `...`.
- **Quotes:** Normalizes spacing inside single quotes `'`.
- **Grammar:** Changes `a` to `an` before vowels or `h`.

## Assumptions
- The input file is encoded in UTF-8.
- Markers like `(hex)` apply to the valid word immediately preceding them.
- If a marker command is invalid (e.g. `(bin)` on a non-binary word), the behavior is to keep the word as is (or as reasonable per the parser).
- "Word" is defined as a sequence of non-whitespace characters separated by spaces or punctuation boundaries.
