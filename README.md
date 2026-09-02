# Go Reloaded

A standard-library Go CLI that cleans assignment-specific text markup into readable prose. It converts binary and hexadecimal values, applies case markers, normalizes punctuation and quotes, and corrects `a` to `an` where required.

## What it demonstrates

- Unicode-aware tokenization with preserved line breaks
- Look-behind transformations such as `(up, 3)`, `(cap)`, `(hex)`, and `(bin)`
- Stateful quote and punctuation reconstruction
- Separation between file I/O and the pure `Process` transformation
- Golden audit cases plus focused unit and CLI tests

## Example

Input:

```text
Simply add 42 (hex) and 10 (bin) and you will see the result is 68.
```

Output:

```text
Simply add 66 and 2 and you will see the result is 68.
```

## Run it

Go 1.25 or newer is required.

```bash
go run . input.txt output.txt
```

The command reads `input.txt`, processes its full contents, and writes the result to `output.txt`.

## Supported rules

| Marker or rule | Result |
| --- | --- |
| `(hex)` / `(bin)` | Converts the preceding hexadecimal or binary token to decimal |
| `(up)` / `(low)` / `(cap)` | Changes the case of the preceding token |
| `(up, N)` / `(low, N)` / `(cap, N)` | Changes the case of the preceding `N` tokens |
| `. , ! ? : ;` | Removes space before punctuation and preserves one separator afterwards |
| `' quoted text '` | Removes padding inside paired single quotes |
| `a` before a vowel or `h` | Rewrites the article as `an`, preserving initial capitalization |

Transformation markers are consumed from the output. Unrecognized parenthesized text is kept as ordinary content.

## Design

Processing is deliberately split into a small pipeline:

1. `tokenize` separates words, marker expressions, punctuation groups, quotes, and line breaks.
2. The marker pass transforms already-seen tokens and removes recognized commands.
3. The grammar pass applies the assignment's `a`/`an` rule.
4. The reconstruction pass emits normalized spacing while tracking whether a quote is open.

The core transformation is exposed as `Process(string) string`, so behavior can be tested without touching the filesystem. See [architecture.md](architecture.md) for the detailed data flow and [PRD.md](PRD.md) for the original assignment requirements.

## Verify it

```bash
go test ./...
go test -race ./...
go vet ./...
```

The test suite covers tokenization, transformations, CLI file handling, the official-style golden cases, and edge cases such as grouped punctuation and multi-word markers.

## Project layout

```text
.
├── main.go          # CLI and file I/O
├── fsm.go           # tokenizer and text-processing pipeline
├── *_test.go        # unit, CLI, and audit-style tests
├── PRD.md           # assignment requirements
└── architecture.md  # implemented design
```

## Author

Built by [Stefanos Kamprogiannis](https://github.com/skamprogiannis) during the Zone01 Athens program.

## License

This project is available under the [MIT License](LICENSE).
