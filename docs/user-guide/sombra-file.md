---
title: Sombra File
---

## What is the Sombra File?

The `sombra.yaml` file defines how templates are applied to a project. It lives in the root of a **target repository** — the repo you are creating or updating using a Sombra template.

This file tells `sombra` which template(s) to use, which variables to pass, and which version (branch or tag) to pull.

---

## File Location

Place `sombra.yaml` at the root of your generated project:

```

my-app/
├── sombra.yaml
├── go.mod
├── README.md
└── ...

````

---

## Configuration Reference


### `templates`

List of one or more templates to apply. Each template includes a Git repo name and optional variables.

```yaml
templates:
  - name: cool-org/playground-django-api-template
    vars:
      project: My Awesome Project
      author: Jane Doe
      email: jane@example.com
      entity: Settings
```

#### Fields:

* `name`: The GitHub path to the template repo
* `vars`: Key-value pairs that are injected into the template

---

## Full Example

```yaml
branch: main

templates:
  - name: sombrahq/playground-django-api-template
    vars:
      project: Internal API
      author: Dev Team
      email: dev@example.com
      entity: Config
```

---

Need to apply this file? Use:

```bash
sombra local init
```

Or to update a project:

```bash
sombra local update
```

For more, check out the [CLI Commands](commands.md) or [Installation Guide](installation.md).

