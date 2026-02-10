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

### 1. Number Conversions

- **(hex):** Converts the preceding hexadecimal number to decimal.
  - _Example:_ `1E (hex) files` -> `30 files`
- **(bin):** Converts the preceding binary number to decimal.
  - _Example:_ `It has been 10 (bin) years` -> `It has been 2 years`

### 2. Case Transformations

- **(up):** Converts the preceding word to UPPERCASE.
  - _Example:_ `Ready, set, go (up) !` -> `Ready, set, GO!`
- **(low):** Converts the preceding word to lowercase.
  - _Example:_ `I should stop SHOUTING (low)` -> `I should stop shouting`
- **(cap):** Capitalizes the preceding word (first letter upper, rest lower).
  - _Example:_ `Welcome to the Brooklyn bridge (cap)` -> `Welcome to the Brooklyn Bridge`

You can also specify a number of words to transform:

- **(up, N), (low, N), (cap, N):** Applies the rule to the previous N words.
  - _Example:_ `This is so exciting (up, 2)` -> `This is SO EXCITING`

### 3. Punctuation

- **Spacing:** Punctuation marks (`.`, `,`, `!`, `?`, `:`, `;`) are formatted to have no space before them and one space after them.
  - _Example:_ `I was sitting over there ,and then BAMM !!` -> `I was sitting over there, and then BAMM!!`
- **Groups:** Groups of punctuation (like `...` or `!?`) are treated as a single unit.
  - _Example:_ `I was thinking ... You were right` -> `I was thinking... You were right`

### 4. Quotes

- Single quotes `'` are formatted to remove spaces inside them (sticking to the text they enclose).
  - _Example:_ `As Elton John said: ' I am the most well-known homosexual in the world '` -> `As Elton John said: 'I am the most well-known homosexual in the world'`

### 5. Grammar

- **A -> An:** The indefinite article `a` is changed to `an` if the next word starts with a vowel (`a`, `e`, `i`, `o`, `u`) or `h`.
  - _Example:_ `There it was. A amazing rock!` -> `There it was. An amazing rock!`

## Assumptions

- The input file is encoded in UTF-8.
- Markers like `(hex)` apply to the valid word immediately preceding them.
- If a marker command is invalid (e.g. `(bin)` on a non-binary word), the behavior is to keep the word as is (or as reasonable per the parser).
- "Word" is defined as a sequence of non-whitespace characters separated by spaces or punctuation boundaries.
