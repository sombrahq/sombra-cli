---
title: Start a Template
---

## Overview

This guide walks you through creating a Sombra template from an existing codebase — no changes to the source code required.

Use this when you want to:

- Reuse a production-ready project as a starting point
- Share best practices and config across multiple teams or services
- Turn internal tooling into reusable templates

---

## Example: Fork the Hextra Starter Template

We'll use [imfing/hextra-starter-template](https://github.com/imfing/hextra-starter-template) as the source project.

### Step 1: Fork the Project

Fork the repository to your own GitHub org/account.

> ✅ Tip: You can use private repos for internal templates.

---

### Step 2: Create the `.sombra/` Directory

In the root of your forked project, create a directory named `.sombra` and add a `default.yaml` inside it:

```

.
└── .sombra/
└── default.yaml

````

This file defines how the template will transform content.

---

### Step 3: Define Your Template

Here’s a minimal `default.yaml` example:

```yaml
vars:
  - repository
  - title

patterns:
  - pattern: "*"
    abstract: true
    default:
      imfing/hextra-starter-template: "{{ .repository }}"
      My Site: "{{ .title }}"
      Hextra Starter Template: "{{ .title }}"

  - pattern: LICENSE
    path:
      "/": "vendors"
    name:
      LICENSE: hextra-starter-template.LICENSE
    verbatim: true
````

This configuration:

* Replaces the repo name and title
* Copies `LICENSE` into a `vendors/` folder without modification

---

### Step 4: Expand the Pattern List

You can match specific files or folders like so:

```yaml
patterns:
  - pattern: README.md
    content:
      https://imfing.github.io/hextra-starter-template/: "{{ .demo }}"

  - pattern: .devcontainer/*
  - pattern: .github/**/*

  - pattern: content/**/*.md
  - pattern: hugo.yaml
```

Use globs like `**/*` to match recursively.

---

### Step 5: Tag a Release

To make your template available for versioned use:

```bash
git tag v1.0.0
git push origin v1.0.0
```

Consumers can now use your template with:

```yaml
branch: v1.0.0
templates:
  - name: your-org/your-template
```

---

## Summary

To convert any repo into a Sombra template:

1. Fork the project
2. Add `.sombra/default.yaml`
3. Define transformation patterns
4. Tag a release
5. Use it in a target repo with `sombra.yaml`

Next: learn about [Template Concepts](concepts.md) to understand pattern structure in depth.

