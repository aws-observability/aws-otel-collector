# Contributing to Timberjack

Thank you for considering contributing to **Timberjack**! Please take a moment to read through these guidelines.

## Guidelines

- All contributions must follow the [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0) format.
- Ensure all tests pass before pushing or opening a pull request.
- Document new functionality clearly in code comments and README if applicable.
- Avoid unnecessary commits or unrelated changes.

## Commit Message Format

We follow the [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/) standard:

```
<type>[optional scope]: <description>
```
### Examples

- `feat: add Rotate feature`
- `fix: correct rotation timestamp logic`
- `docs: update README with usage example`
- `test: cover edge case for rotation overlap`

### Allowed Types

| Type      | Use For                                                                 |
|-----------|-------------------------------------------------------------------------|
| `feat`    | A new feature                                               |
| `fix`     | A bug fix                                             |
| `chore`   | Build process, CI, tooling                        |
| `docs`    | Changes to documentation (README, comments)                             |
| `test`    | Adding or modifying tests                                               |
| `refactor`| Code refactoring without behavior change               |
| `style`   | Code style changes (whitespace, etc)                                      |
| `perf`    | Performance improvements                                                |
| `ci`      | Changes to GitHub Actions, CI/CD workflows                              |


## Scopes (Optional)

Use a **scope** in parentheses after the type to clarify what part of the code is affected. Examples:

- `fix(logger): correct timestamp handling`
- `feat(rotation): add weekday scheduling`
- `docs(readme): document Rotate usage`

### Breaking Changes

If your commit introduces a **breaking change** — something that would require users to modify their code — add a `!` immediately after the type:

- `feat!: remove support for deprecated config field`

This will mark the commit as a major version bump when release automation is used.

### Other Guidelines

- Run all tests before pushing: `go test ./...`
- Keep PRs focused and self-contained
- Use clear, descriptive commit messages
- If you add a new feature, consider adding an example in the README.

## Code Coverage Requirements

To maintain high-quality tests, all pull requests must meet the following **code coverage** rules:

- Overall coverage must be **at least 85%**
- Coverage **must not decrease by more than 2.5%**

### How to Check Locally

Before opening a PR, run this to verify your test coverage:

```bash
go test -covermode=atomic -coverprofile=covprofile ./...
```