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

## Windows

```
powershell -c "irm https://raw.githubusercontent.com/gdegiorgio/octo/refs/heads/main/scripts/install.ps1 | iex"
```

## MacOs & Linux

```bash
curl https://raw.githubusercontent.com/gdegiorgio/octo/refs/heads/main/scripts/install.sh | bash
```


## 🧪 Usage

```bash
🐙 Install your packages everywhere

Usage:
  octo [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  install     Install a new package
  list        List all installed packages
  update      Update a package to latest version
  upgrade     Upgrade Octo CLI to latest version
  version     Show current Octo CLI version

Flags:
  -h, --help   help for octo
```

## 🧙 Contributing

We welcome contributions of all kinds!

- All PRs must follow Semantic Commit Messages
- Read our [Contributing guide](CONTRIBUTING.md) before submitting a PR
