# Readability

## Code Formatting

- Is the code formatted using go fmt?
- Are lines kept under 80 characters when possible (except for imports, const declarations, type definitions)?
- Are tabs used for indentation (as per Go convention), not spaces?
- Are naming conventions consistent with Go standards? (e.g., CamelCase for exported names, camelCase for unexported)
- Are package names short, lowercase, and descriptive? (e.g., http, not net/http)
- Are file names lowercase with underscores preferred? (e.g., my_utility.go)

## Comments

- Are exported functions, types, constants, and variables documented with clear, concise comments?
- Are documentation comments placed immediately before the declaration?
- Are complex logic and design decisions clearly explained?
- Do comments focus on explaining "why" rather than "what"?
- Are documentation comments structured with blank lines between paragraphs?

## Packages & Imports

- Are imports grouped in order: standard library, external packages, internal packages?
- Are imported packages referenced using their original names unless aliasing is necessary?
- Are unused imports removed?
- Are package names descriptive and relevant to their content?
- Are vague names like util, helper, common avoided (except as part of a longer name)?
- Are package names chosen to avoid name collisions or shadowing at the call site?

## Functions

- Are functions short and designed with a single responsibility?
- Are errors returned explicitly?
- Is error handling done correctly?
- Do functions with multiple return values return an error as the last value?
- Is defer used appropriately for cleanup tasks?
- Are function/method names chosen with readability and usage context in mind?
- Is excessive repetition at the call site avoided?
- Are types or pointer status omitted from names where context is clear?
- Are package names avoided in function names?
- Are receiver names avoided in method names?
- Are variable names not repeated as parameter names?
- Are return value names not unnecessarily repeated?
- If similar function names are needed, is extra context added?
- Are functions returning a value named like nouns?
- Are verbs used for functions that perform actions?
- Are Get prefixes avoided in function/method names?
- Are type names added to function names only when necessary to distinguish?
- If there's a clear "primary" version of a function, is the type name omitted?

## Types

- Do struct field names follow Go naming conventions?
- Are struct fields grouped logically, considering padding?
- Are interfaces small, ideally with a single method?

## Variables

- Are variable names short and meaningful?
- Is variable scope minimized where possible?
- Are zero values relied upon instead of explicit initialization when appropriate?
- If shadowing is used, is it done intentionally and clearly?
- Are variable names that shadow standard library packages avoided (except in very limited scope)?

## Control Structures

- Are if conditions concise?
- Are for loops clear and readable?
- Is it understood that switch statements don't require break? (use fallthrough if needed)

## Error Handling

- Are errors properly checked and handled?
- Do error messages clearly describe the issue?
- Is panic used only for unrecoverable errors?
- Do libraries avoid panicking in most cases?
- Do functions used in tests avoid stopping the test unless it's unrecoverable?

## Concurrency

- Are goroutines and channels used safely and effectively?
- Is access to shared mutable state properly synchronized (e.g., using sync.Mutex)?
- Is code designed to avoid deadlocks?

## General

- Are tests written? (unit, integration, etc.)
- Do tests cover a wide range of cases?
- Is the purpose and rationale of the code clear to readers?
- Is code written as simply as possible to achieve its purpose?
- Does code maintain a high signal-to-noise ratio?
- Is code written to be easily maintainable?
- Does the code achieve its goals with the simplest solution in terms of behavior and performance?
- When multiple valid options exist, is the most idiomatic one chosen?
- Are core language constructs (channels, slices, maps, loops, structs) preferred over unnecessary abstractions?
- Are all source files gofmt compliant?
- Is camelCase or MixedCaps used for multi-word names, not snake_case?
- If a line feels too long, is the code refactored rather than simply split?
- Are names chosen to avoid repetition in usage context?
- Are these guidelines learned and followed until readability is second nature?
- Is support code for tests written consistently with Go style?
- Is the global state handled with care?
- In tests, is the naming of doubles context-aware and distinct from production types?
- If the original value is no longer needed, is it safe to discard it?
- Are test packages named by appending test to the original package name?
