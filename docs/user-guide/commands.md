---
title: CLI Commands
---

## Overview

The `sombra` CLI provides commands to generate and update projects from templates, and to create templates from existing projects.

Run `sombra --help` at any time to view global help.

---


## 🔧 `local` Commands

Used to apply and update templates in your working project.

### `sombra local init`

Generate a new project from a remote Git template.

```bash
sombra local init TEMPLATE
```

#### Positional:

* `TEMPLATE`: Git repo URL of the template

#### Example:

```bash
sombra local init github.com/sombrahq/playground-django-api-template
```

---

### `sombra local update`

Update your current project using the source template.

```bash
sombra local update [--tag TAG] [--method METHOD] TEMPLATE
```

#### Positional:

* `TEMPLATE`: Git repo URL of the template

#### Options:

* `--tag`: Specific git tag or version to use
* `--method`: `copy` (default) or `diff` for smarter merging
* `--help, -h`: Show help

#### Example:

```bash
sombra local update --tag v1.2.0 --method diff github.com/org/template-repo
```

---

## 🧪 `template` Commands

Used to turn existing codebases into reusable templates.

### `sombra template init`

Initialize a `.sombra/default.yaml` template from an existing project.

```bash
sombra template init [--exclude PATTERN] [--only PATTERN] [DIR]
````

#### Positional:

* `DIR`: Path to the project directory to convert (default: current dir)

#### Options:

* `--exclude, -e`: Glob to exclude files (e.g. `"*.pyc"`)
* `--only, -o`: Glob to include files
* `--help, -h`: Show help

#### Example:

```bash
sombra template init --exclude "README.md" ./my-project
```

---

For detailed usage, see the [Sombra File](sombra-file.md) or [Template Guide](../sombra-templates/index.md).
