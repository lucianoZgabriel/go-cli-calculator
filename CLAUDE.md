# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## 🎓 Learning-Focused Approach

**IMPORTANT**: This is a learning project. The user wants to strengthen their Go skills through hands-on practice. 

### Guidelines for Claude:
- **DO NOT** modify files directly
- **DO** provide guidance, explanations, and step-by-step instructions
- **DO** explain the reasoning behind each decision and Go best practices
- **DO** help the user understand concepts before implementing
- **DO** track progress by updating checkmarks in `go_calculator_plan.md`
- **DO** encourage the user to implement solutions themselves

### Teaching Approach:
1. Explain the concept or pattern first
2. Show examples of Go idioms and best practices
3. Guide the user through implementation steps
4. Help them understand why certain approaches are preferred
5. Update progress tracking in the planning document

## Common Development Commands

```bash
# Run all tests
go test ./...

# Run tests with verbose output
go test -v ./...

# Run tests with coverage
go test -cover ./...

# Run specific package tests
go test ./calculator

# Build the project
go build

# Format code
go fmt ./...

# Check for issues
go vet ./...

# Tidy up go.mod
go mod tidy
```

## Code Architecture

This is a Go CLI calculator project organized into packages:

- **main.go**: Entry point (currently minimal - needs implementation)
- **calculator/**: Core calculator logic package
  - `Calculator` struct with methods: `Add`, `Subtract`, `Multiply`, `Divide`
  - Error handling for division by zero
  - Implements `Stringer` interface
- **cli/**: CLI parsing package (empty - needs implementation)

## Project Status

The project follows a phased development approach as outlined in `go_calculator_plan.md`:

- **Phase 1** (partially complete): Basic calculator operations are implemented with partial test coverage
- **Phase 2** (pending): CLI interface needs implementation in main.go
- **Phase 3** (planned): Advanced CLI with flag package
- **Phase 4** (planned): Scientific operations and interactive mode

## Testing Approach

The project uses table-driven tests following Go idioms. Currently only the `Add` function has comprehensive tests. When adding tests for other operations, follow the existing pattern in `calculator/calculator_test.go`.

## Key Implementation Notes

1. The calculator package uses a struct-based approach with pointer receivers
2. Error handling follows Go idioms - returning error as second value
3. Division by zero returns a descriptive error message
4. The main.go file needs to be implemented to create a working CLI binary