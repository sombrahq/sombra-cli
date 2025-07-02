---
title: Sombra Templates
---

## Overview

Sombra Templates let you turn real production code into reusable, versioned project generators — without modifying the original source.

This section explains how to author, structure, and publish templates that can be used by others with the `sombra` CLI.

---

## What Is a Sombra Template?

A **Sombra template** is just a Git repository that contains:

- Real project files (application, service, etc.)
- A `.sombra/default.yaml` file defining transformation rules

When used with a `sombra.yaml`, the CLI can generate new projects by replacing values in filenames, content, and directory paths — all using Go templates.

---

## Who Is This For?

This guide is for developers, consultants, or teams who want to:

- Share boilerplate across projects or clients
- Enforce a consistent project structure and tooling setup
- Reuse proven, production-quality code without duplication

---

## Get Started

Ready to build a template?

- 🏗 [Start a Template](start-a-template.md) — Fork a real project and convert it into a Sombra template
- 🧠 [Concepts](concepts.md) — Understand how templates, patterns, and mappings work
- 📌 [Best Practices](best-practices.md) — Keep your templates clean and maintainable

To use a template instead, head to the [User Guide](../user-guide/index.md).
