---
title: Template Concepts
---

## Overview

This page explains the internal structure and logic of **Sombra Templates**, including how files are transformed using Go templates, YAML definitions, and mappings.

A Sombra template is composed of:

- Real source files (code, configs, etc.)
- A `.sombra/default.yaml` file with transformation rules

The CLI uses this configuration to generate new projects with modified paths, filenames, and content based on input variables.

---

## Directory Structure

Sombra templates are Git repositories that contain a `.sombra` directory with a `default.yaml` file:

```

.
└── .sombra/
└── default.yaml

````

Only one definition file is currently supported.

---

## Go Templates + YAML

Sombra uses [Go’s templating engine](https://pkg.go.dev/text/template) to apply variable substitutions. It also includes [Sprig](https://masterminds.github.io/sprig/) functions for string, list, and math utilities.

All variables come from the target repo’s `sombra.yaml`.

---

## Template Structure

### `vars`

A list of variable names that are expected to be defined by the user when applying the template:

```yaml
vars:
  - project
  - author
````

---

### `patterns`

Each pattern block applies transformation rules to files matching a glob-style pattern.

```yaml
patterns:
  - pattern: "*"
    abstract: true
    default:
      project-template: "{{ .project | kebabcase }}"
```

---

### Pattern Matching Categories

Each pattern supports three transformation scopes:

* `path`: Folder and subdirectory names
* `name`: Filenames
* `content`: File contents

The `default` section applies to all three.

---

## Pattern Field Reference

| Field      | Description                                          |
| ---------- | ---------------------------------------------------- |
| `pattern`  | Required. Glob to match files                        |
| `abstract` | If true, this rule applies only as a base pattern    |
| `verbatim` | If true, disables all replacements for matched files |
| `default`  | General-purpose replacements                         |
| `path`     | Folder path replacements                             |
| `name`     | Filename replacements                                |
| `content`  | Content-specific replacements                        |
| `except`   | Files to exclude from the rule                       |

---

## How Replacements Are Applied

When a file matches multiple patterns:

1. All relevant mappings are collected
2. Mappings are sorted by key length (longest first)
3. Each key-value pair is applied sequentially to:

	* Path
	* Name
	* Content

This ensures deterministic and consistent replacements.

---

## Regex Support

Sombra supports [RE2-compatible](https://github.com/google/re2/wiki/Syntax) regular expressions for advanced text processing:

```yaml
"re:description = .*\n": ""
```

This removes lines that start with `description = ...`.

---

For a step-by-step guide, continue to [Start a Template](start-a-template.md).
