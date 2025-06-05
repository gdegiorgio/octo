# Contributing to Octo

Thanks for your interest in contributing to **Octo**, a cross-platform Homebrew alternative written in Go! 🎉

Please take a moment to review this guide. Following these standards helps us maintain a high-quality, consistent codebase and smooth review process.

---

## 📦 Project Setup

### 1. **Fork, then clone the repository**

```bash
   git clone https://github.com/YOUR_USERNAME/octo.git
   cd octo
```

### 2. **Install dependencies**

```bash
    make install
```

### 1. **Run tests**

```bash
    make test
```

### 1. **Lint Code**

```bash
    make lint
```



## ✅ Commit Guidelines

  ### All commits must be signed

  We require GPG or SSH signed commits.

  ### Use conventional commits

  This helps with automated changelogs and semantic versioning:

  ```
    feat(cli): add install command for Linux
    fix(resolver): correct path resolution on Windows
    chore: update Go modules
  ```

  Types allowed:

    ```
        feat: New feature
        fix: Bug fix
        chore: Tooling or maintenance
        docs: Documentation only
        refactor: Code change that neither fixes a bug nor adds a feature
        test: Adding or improving tests
        ci: CI/CD related changes
        perf: Performance improvement
    ```

## 📦 Semantic Versioning

We follow Semantic Versioning:

  - MAJOR version when you make incompatible API changes
  - MINOR version when you add functionality in a backward-compatible manner
  - PATCH version when you make backward-compatible bug fixes

Release tags and changelogs will be automatically generated based on commits.

## 📂 Pull Requests

  Create a new branch for each contribution:

  ```bash
    git checkout -b feat/some-feature
  ```

  Ensure all tests pass and code is linted before submitting a PR.
  Describe why the change is needed and what it does.
  PR titles should use the same conventional commit format.

## 🤝 Code of Conduct

Please be respectful and inclusive. See [CODE_OF_CONDUCT](CODE_OF_CONDUCT.md) for details.

Thank you for contributing! ✨
