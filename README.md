# Sombra CLI

**Sombra** is an open-source command-line tool that helps you automate project scaffolding by turning production-ready repositories into reusable, version-controlled templates.

Built for developers, consultants, and teams who want to:

* ⚡ Quickly start new projects with consistent setup
* 🧱 Reuse real, tested code without changing it
* 🔄 Keep projects up to date with shared boilerplate

---

## ✨ Features

* ✅ Use any Git repository as a template source
* ⚙️ Define flexible rules using Go templates + YAML
* ♻️ Reuse code without modifying production files
* 🔍 Match and transform paths, filenames, and content
* 🏷 Semantic versioning with Git tags

---

## 📦 Install

### Option 1: via Go

```bash
go install github.com/sombrahq/sombra-cli@latest
```

### Option 2: Prebuilt Binaries

Download from [GitHub Releases](https://github.com/sombrahq/sombra-cli/releases)

### Option 3: Build from source

```bash
git clone https://github.com/sombrahq/sombra-cli.git
cd sombra-cli
make build WHAT=sombra
```

See [Installation Guide](https://sombra-cli.sombrahq.com/user-guide/installation/) for more details.

---

## 🚀 Quick Start

### 1. Create a Template

Convert a production repo into a template:

```bash
sombra template init ./my-app
```

This creates `.sombra/default.yaml`.

### 2. Apply a Template

In a new repo, create a `sombra.yaml`:

```yaml
branch: main
templates:
  - name: https://github.com/your-org/your-template
    vars:
      project: New API
```

Then run:

```bash
sombra local init
```

### 3. Update a Project

```bash
sombra local update --tag v1.0.0 --method copy
```

---

## 📖 Documentation

Full docs available at 👉 [sombra-cli.sombrahq.com](https://sombra-cli.sombrahq.com)

Key topics:

* [Installation](https://sombra-cli.sombrahq.com/user-guide/installation/)
* [Creating Templates](https://sombra-cli.sombrahq.com/sombra-templates/start-a-template/)
* [sombra.yaml Config](https://sombra-cli.sombrahq.com/user-guide/sombra-file/)
* [Command Reference](https://sombra-cli.sombrahq.com/user-guide/commands/)

---

## 🤝 Contributing

Issues and PRs welcome! Start with the [Contact page](https://sombra-cli.sombrahq.com/contact/) or open an [Issue](https://github.com/sombrahq/sombra-cli/issues).

MIT licensed. Made with ❤️ by [@yunier](https://www.linkedin.com/in/yunier-rojas-garc%C3%ADa/)
