# Product Requirements Document (PRD) - go-reloaded

## 1. Problem Statement
We need to build a command-line interface (CLI) tool in Go that acts as a text completion, editing, and auto-correction system. It takes a raw text file as input, applies a series of complex formatting and transformation rules (such as number conversion, casing, punctuation spacing, and grammar correction), and writes the polished result to an output file.

## 2. User / Use Case
- **Primary User:** Students, writers, or auditors needing to normalize text files according to a strict set of style rules.
- **Use Case:** A user has a draft text file (`sample.txt`) containing markers like `(hex)`, `(up)`, and messy punctuation. They run the tool to generate a clean, readable version (`result.txt`).

## 3. CLI Contract
- **Command:** `go run . <inputFile> <outputFile>`
- **Inputs:**
  - `<inputFile>`: Path to an existing file containing the text to process.
  - `<outputFile>`: Path where the transformed text will be saved.
- **Output:**
  - The processed text is written directly to `<outputFile>`.
- **Error Handling:**
  - If arguments are missing: Print usage message.
  - If input file is missing/unreadable: Print error and exit.
  - If output file cannot be created: Print error and exit.

## 4. Functional Requirements (Rules)

### 4.1 Marker Conversions (Numbers)
- **(hex):** Convert the word immediately preceding this marker from hexadecimal to decimal.
  - *Example:* `1E (hex) files` -> `30 files`
- **(bin):** Convert the word immediately preceding this marker from binary to decimal.
  - *Example:* `It has been 10 (bin) years` -> `It has been 2 years`

### 4.2 Case Transformations
- **(up):** Convert the preceding word to UPPERCASE.
  - *Example:* `Ready, set, go (up) !` -> `Ready, set, GO!`
- **(low):** Convert the preceding word to lowercase.
  - *Example:* `I should stop SHOUTING (low)` -> `I should stop shouting`
- **(cap):** Convert the preceding word to Capitalized (first letter upper, rest lower).
  - *Example:* `Welcome to the Brooklyn bridge (cap)` -> `Welcome to the Brooklyn Bridge`
- **(up, N), (low, N), (cap, N):** Apply the transformation to the *previous N words*.
  - *Example:* `This is so exciting (up, 2)` -> `This is SO EXCITING`

### 4.3 Punctuation Formatting
- **Standard Punctuation (`.`, `,`, `!`, `?`, `:`, `;`):**
  - Must adhere to the previous word (no space before).
  - Must have a space after (unless it's the end of the text).
  - *Example:* `I was sitting over there ,and then BAMM !!` -> `I was sitting over there, and then BAMM!!`
- **Punctuation Groups:**
  - Groups like `...` or `!?` must remain together and are treated as a single punctuation unit regarding spacing.
  - *Example:* `I was thinking ... You were right` -> `I was thinking... You were right`

### 4.4 Quotes (`'`)
- Single quotes `'` always appear in pairs.
- The spaces *inside* the quotes must be removed (i.e., stick the quotes to the words inside).
- *Example:* `As Elton John said: ' I am the most well-known homosexual in the world '` -> `As Elton John said: 'I am the most well-known homosexual in the world'`

### 4.5 Grammar (`a` -> `an`)
- Replace `a` with `an` if the following word starts with a vowel (`a`, `e`, `i`, `o`, `u`) or `h`.
- *Example:* `There it was. A amazing rock!` -> `There it was. An amazing rock!`

## 5. Non-Goals
- **No external libraries:** The project must use Standard Go packages only (e.g., `os`, `strings`, `strconv`).
- **No advanced NLP:** The tool does not need to understand complex grammar rules outside of the specific `a/an` logic provided.

## 6. Acceptance Criteria & Tests

### 6.1 Audit Cases (Golden Tests)
These cases must pass exactly as specified in the subject.

- **Case 1: General Casing**
  - *Proves:* Correct handling of `(cap)`, `(up)`, `(low)` markers and their numeric variants (e.g., `(cap, 6)`), plus punctuation spacing.
  - *Input:* `it (cap) was the best of times, it was the worst of times (up) , it was the age of wisdom, it was the age of foolishness (cap, 6) , it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, IT WAS THE (low, 3) winter of despair.`
  - *Expected:* `It was the best of times, it was the worst of TIMES, it was the age of wisdom, It Was The Age Of Foolishness, it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, it was the winter of despair.`

- **Case 2: Hex and Bin Conversion**
  - *Proves:* `(hex)` and `(bin)` markers successfully convert the preceding word to decimal.
  - *Input:* `Simply add 42 (hex) and 10 (bin) and you will see the result is 68.`
  - *Expected:* `Simply add 66 and 2 and you will see the result is 68.`

- **Case 3: Grammar (a -> an)**
  - *Proves:* The particle `a` transforms to `an` when followed by a vowel or `h`.
  - *Input:* `There is no greater agony than bearing a untold story inside you.`
  - *Expected:* `There is no greater agony than bearing an untold story inside you.`

- **Case 4: Punctuation Groups**
  - *Proves:* Grouped punctuation marks (like `...`) are treated as a single unit and spacing is applied correctly (attached to previous word, space after).
  - *Input:* `Punctuation tests are ... kinda boring ,what do you think ?`
  - *Expected:* `Punctuation tests are... kinda boring, what do you think?`

### 6.2 Additional Edge Cases (Golden Tests)
These additional tests cover scenarios not fully exercised by the audit cases.

- **Extra 1: Quotes with Punctuation**
  - *Proves:* Single quotes `'` remove internal spaces, and `(cap)` works correctly inside/near quotes and punctuation.
  - *Input:* `He said: ' hello , world (cap) ' .`
  - *Expected:* `He said: 'Hello, World'.`

- **Extra 2: Order of Operations**
  - *Proves:* Markers are processed in a logical order (left-to-right) and affect the immediately preceding state.
  - *Input:* `test (up) (low)`
  - *Expected:* `test`

- **Extra 3: Case Insensitivity for 'an'**
  - *Proves:* The `a` -> `an` rule applies even if the following word is capitalized (e.g., `Honest`).
  - *Input:* `It was a Honest mistake.`
  - *Expected:* `It was an Honest mistake.`

## 7. Implementation Approach

### Architecture Comparison

We considered two main architectural approaches for this project:

| Feature | Option A: Pipeline | Option B: Finite State Machine (FSM) |
| :--- | :--- | :--- |
| **Concept** | "Car Wash": Data flows through independent stages (Tokenize -> Markers -> Punctuation -> Quotes -> Output). | "Conveyor Belt": Single pass where a scanner reacts to each token based on current state (e.g., "InsideQuotes"). |
| **Pros** | - Easy to test each stage in isolation.<br>- Separation of concerns (one function per rule).<br>- Simple to understand for teams. | - Efficient single-pass processing.<br>- Handles context-heavy rules (quotes, punctuation spacing) naturally.<br>- Low memory overhead. |
| **Cons** | - Hard to handle "global" context (e.g., quotes affecting spacing across stages).<br>- Multiple passes over data can be inefficient.<br>- "Chain reaction" bugs (stage 1 error breaks stage 4). | - Can become complex ("spaghetti code") if states aren't managed well.<br>- Harder to debug a single specific rule in isolation. |

### Decision: Finite State Machine (FSM)

We have chosen the **Finite State Machine (FSM)** approach.

#### Rationale
The complexity of punctuation spacing and quote handling (where an opening quote changes the rules for spacing until a closing quote is found) makes FSM a superior choice. A pipeline might struggle to know if a quote is "open" or "closed" without passing complex state objects between stages. An FSM allows us to maintain flags like `inQuote` or `lastWasPunctuation` easily while iterating through the tokens.

#### High-Level Design (ASCII Diagram)
*(For detailed logic and state transitions, see `architecture.md`)*

```
[Start] -> [Scan Token]
              |
              v
      Is it a Marker? --Yes--> [Apply Transformation to Buffer] -> [Back to Scan]
              | No
              v
      Is it Punctuation? --Yes--> [Adjust Spacing in Buffer] -> [Append Punct] -> [Back to Scan]
              | No
              v
      Is it Quote? --Yes--> [Toggle Quote State] -> [Adjust Spacing] -> [Append Quote] -> [Back to Scan]
              | No
              v
        [Append Word]
              |
              v
      [Check 'a' -> 'an' Candidate] -> [Back to Scan]
```

## 8. Milestones (Implementation Plan)

1.  **Project Skeleton & CLI:** Setup `main.go` to accept arguments, read file, and write file. Basic test setup.
2.  **Tokenizer & FSM Loop:** Implement the core loop that breaks input into tokens (words/symbols) and iterates through them.
3.  **Marker Logic:** Implement `hex`, `bin`, `up`, `low`, `cap` (single and multi-word).
4.  **Punctuation & Quote State:** Implement the spacing logic and quote state tracking.
5.  **Grammar & Polish:** Implement `a` -> `an` logic, final code cleanup, and full test pass.
