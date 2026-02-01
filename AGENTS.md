# AI Agent Instructions - go-reloaded

This document defines how AI agents should work within the `go-reloaded` repository. Treat these instructions as strict requirements for all tasks.

## 1. Project Context
- **Goal:** Build a Go CLI tool for text formatting and auto-correction.
- **Status:** Week 1 (Planning) is complete. Week 2 (Implementation) is starting.
- **Reference:** Always check `PRD.md` for the single source of truth regarding rules and requirements, and `architecture.md` for algorithm design.

## 2. Technical Constraints
- **Standard Library Only:** Do NOT use any external packages (no `github.com/...` imports). Use only Go's standard library (e.g., `strings`, `strconv`, `unicode`, `os`).
- **Go Version:** Assume current stable Go version.
- **Build:** The project must run with `go run . <input> <output>`.

## 3. Workflow & Philosophy
- **TDD (Test-Driven Development):**
  1.  Read the requirement (e.g., "Implement hex marker").
  2.  Create/Update a test case in `main_test.go` (or equivalent) that fails.
  3.  Implement the minimal code to pass the test.
  4.  Refactor for clarity.
- **Architecture:**
  - Follow the **Finite State Machine (FSM)** approach decided in the PRD.
  - Do not implement a "Pipeline" or "Regex-heavy" replacement unless explicitly requested to refactor the architecture.
- **Commits:**
  - If asked to commit, write clear, imperative messages in the convential commit style (e.g., "feat: implement hex marker logic", "fix: correct spacing for quotes").

## 4. Code Style
- **Formatting:** Code must always be formatted with `gofmt`.
- **Naming:** Use idiomatic Go naming (CamelCase, short variable names where context is clear, e.g., `i` for index, `r` for reader).
- **Comments:** Comment *why*, not *what*. Explain complex FSM state transitions.

## 5. Safety & Quality
- **File Operations:** Never overwrite the input file unless explicitly told to use the same path for output.
- **Testing:** Ensure `go test ./...` passes before considering a task complete.
- **No Hallucinations:** If a rule in the PRD is ambiguous, ask the user for clarification rather than guessing.

## 6. Interaction
- **Brevity:** Be concise. Don't explain Go syntax unless asked.
- **Proactive:** If you see a missing edge case in the PRD or tests, suggest adding it.
