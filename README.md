# Webref - Go types representing Web IDL standards

This library exposes data from Web IDL specifications as Go types.

Data is sourced from [github.com/w3c/webref](https://github.com/w3c/webref). The
data in the release package contains data from the _curated_ list, i.e. types
that have reached a level of maturity, and can assumed to be supported by major
browsers.

Please reach out if you are interested in the bleeding edge specs.

> [!IMPORTANT]
>
> This tool is not complete. Not all IDL information are exposed. If you find
> that you need something not exposed, please file an issue. Or even better,
> make a PR.

## Completeness

This is not a complete representation of all data, but strives to be the place
to look.

Features are added as necessary from gost-dom, or user feedback.

## Packages

This package is divided into subpackages:

- `elements` contains mapping from HTML element tag names to IDL interface name.
- `idl` contains the Web IDL specifications.

### Package `elements`

This exposes element tag name to interface name mapping, i.e., which type is
used at runtime to handle a specific element.

E.g., an `<a>` element in the HTML will be handled by an `HTMLAnchorElement` in
code.

There is not a one-to-one mapping between tag names, and elements names. In
the html specs, multiple elements are handled by the generic `HTMLElement`.
Examples include, `<article>`, `<section>`, and `<nav>`. Then are all elements
for representation with no behaviour associated.

### Package `idl`

Contain specifications for the interfaces available at runtime; and made
accessible to client-side code.

```
spec, err := idl.Load("html")
anchor := spec.Interfaces["HTMLAnchorElement"]
for _, o := anchor.Attributes {
  // ...
}
for _, o := anchor.Operations {
  // ...
}
```

## Coding guidelines

This codebase is tested in general terms by inspecting properties on select
values, and comparing them against expectations. E.g., the `URL` interface from
the `url` specification should have a non-static `toJSON` operation, and a
static `parse` operation.

When a new feature is supported, add a test for a type that uses the feature,
showing.

### Historic code

The first version of the code for the idl module consisted of structures
"reverse engineered" from the JSON data, not knowing exactly what they
represented. This data _should_ be complete, but exposes a less useful model.

Later, I started reading the specs for Web IDL itself, and started a new set of
types that has a model reflecting the standard itself. This is not complete
however.

Eventually, the old model will be removed (or unexported), leaving only the new
model.

## About the compiled file size

The compiled library takes up about 6Mb on disk, which is mainly embedded data
files, which also takes up space in source code.[^1]

It is not a priotity to reduce this. This is a tool intended for design-time
use, not runtime.

However, a PR to reduce compiled file size would be welcome, though a more
complete model should be 

### Reducing file size

The file size (and source code) could be reduced by one of two approaces

- Parse IDL files instead of JSON
- Generate Go code from JSON files

#### Parse IDL files instead of JSON

The JSON is a very verbose representation of the source IDL. By parsing and
embedding IDL files, the source code, and compiled output would be significantly
smaller.

#### Auto-generate Go code

You could generate Go source files with global variables initialised to the
values representing the data.

---

[^1]: The JSON data from the webref submodule is >15Mb in size. This tool copies
    and strips useless fields and whitespace, reducing the JSON data to ~5Mb.
