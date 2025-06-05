# 🐙 Octo



> ⚠️ **Early Stage Project**
> Octo is under **active development** and not yet production-ready. Expect frequent changes, bugs, and missing features. Contributions are welcome as we build this together!


Octo is a cross-platform, fast, and minimal package manager inspired by [Homebrew](https://brew.sh), but built from the ground up in Go. It's designed for developers who want a universal, self-contained package management tool that just works across **macOS**, **Linux**, and **Windows**.


---

## ✨ Features

- 📦 Install and manage software packages easily
- 🧠 Smart dependency resolution
- 💡 Built-in support for GitHub releases and archives
- 🛠️ Written in Go for speed and portability
- 🌍 Cross-platform support (Linux, macOS, Windows)

---

## 🚀 Getting Started

### Install

> **Note**: Precompiled binaries coming soon. For now, build from source.

```bash
git clone https://github.com/YOUR_USERNAME/octo.git
cd octo
go build -o octo ./cmd/octo
./octo --help
```

## 🧪 Example Usage

```bash
octo install bat
octo upgrade bat
octo remove bat
octo list
```

## 🧙 Contributing

We welcome contributions of all kinds!

- All PRs must follow Semantic Commit Messages
- Read our [Contributing guide](CONTRIBUTING.md) before submitting a PR
