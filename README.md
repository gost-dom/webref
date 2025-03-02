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
- `events` information about the different event types being dispatched.

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

#### Legacy data structures

The first version of `idl` module consisted was a "reverse engineering" of the
JSON data. This data _should_ be contain all information in the source data. But
this exposes a less useful model.

A new interface was starter, better representing the Web IDL itself, using terms
like `Interface` instead of `Name` for types representing an interface
specification.

The new data model is not completed, i.e., not all source data is exposed. As a
solution to missing information, the data has references to the the legacy
objects representing the same data.

The legacy model is in the `idl/legacy` subpackage, and should not be used
unless the information is not present in the `idl` package itself.

If you depend on this information, please submit an issue, or even better, file
a PR, 

### Package `events`

A click event bubbles and is cancelable. A formdata also bubbles, but isn't
cancelable. That is the type of information present in the `events` package.

## Coding guidelines

This codebase is tested in general terms by inspecting properties on select
values, and comparing them against expectations. E.g., the `URL` interface from
the `url` specification should have a non-static `toJSON` operation, and a
static `parse` operation.

When support is added for a new type of information, new tests should be added
verifying the information on _real types in the IDL specification_, both
positive and negative tests, i.e., add tests for types that don't have this type
of information.

## How data is sourced

The folder `internal/specs/sources` is a git submodule pointing to the
[curated](https://github.com/w3c/webref/tree/curated) list of webref files.

A make target copies the used files to `internal/specs/curated`, and processes
them for smaller footprint, removing whitespace and unused attributes

### Updating sources

To update the sources, you must get the submodule, update to the latest branch,
and rerun the make script

If you didn't already, get the submodule `git submodule update --init`

If you already had the data, from the `internal/specs/sources` folder run

```sh
git fetch # if you have previously retrieved the data, get latest)
git checkout origin/curated
```

From the root of the project, run (note, this requires a unix-like shell)

```
make specs
```

Run the test suite, and commit the changes if successful, including the
`internal/specs/sources` containing the new commit hash.

### Including more types

If you need to add support for more "folders" in the sources, you need to:
- Add this to the list of processed folders in the Makefiles. Check comments for how
- Create the output folder, the target fails if the folder doesn't exist.
- Run the `specs` make target.

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
