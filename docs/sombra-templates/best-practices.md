---
title: Best Practices
---

## Best Practices for Creating Sombra Templates

Creating effective templates with Sombra means writing definitions that are reusable, clear, and easy to maintain over time.  
Below are best practices we recommend following when designing your `.sombra/default.yaml`.

---

### 1. Keep Definitions Simple

While Sombra supports Go templating and regex, complex logic increases the risk of errors and reduces maintainability.  
Favor simple mappings and clear structure over clever or deeply nested patterns.

---

### 2. Separate Template Logic from Production Code

Never modify production code just to turn it into a Sombra template.  
Use abstract patterns and search/replace mappings instead — this keeps your codebase clean and reusable.

---

### 3. Use Abstract Patterns for Global Rules

Use `abstract: true` patterns to define shared mappings that apply to multiple files:

```yaml
- pattern: "*"
  abstract: true
  default:
    my-app-template: "{{ .project }}"
````

This centralizes your replacements and avoids repetition.

---

### 4. Add Comments for Context

Document your mappings so others (or future-you) understand why replacements exist.

```yaml
default:
  my-app-template: "{{ .project }}"  # Replace project name in path/content
```

---

### 5. Favor Explicit File Scoping

Instead of using `*`, prefer more specific glob patterns when possible:

```yaml
- pattern: config/**/*.yaml
- pattern: README.md
```

This improves performance and reduces accidental replacements.

---

### 6. Use Semantic Versioning in Template Repos

Tag your template repositories with semantic version tags (`v1.0.0`, `v1.1.0`, etc.).
This enables consumers to pin and upgrade templates predictably via:

```yaml
branch: v1.0.0
```

---

### 7. Test Template Behavior in a Sandbox

Use a throwaway repo or directory to validate your template results.
Run:

```bash
sombra local init your/template-repo
```

Verify filenames, paths, and content transformations behave as expected before releasing.

---

### 8. Commit `.sombra/` Separately

Treat template definition files like infrastructure. Commit changes separately from feature code to make diffs easier to review.

---

These practices help keep templates readable, reusable, and safe to apply at scale.

For hands-on examples, visit [Start a Template](start-a-template.md).
