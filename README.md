# Webref - Go types representing Web IDL standards

This library exposes data from Web IDL specifications as Go types.

Data is sourced from [github.com/w3c/webref](https://github.com/w3c/webref)

> [!IMPORTANT]
>
> This tool is not complete. Not all IDL information are exposed. If you find
> that you need something not exposed, please file an issue. Or even better,
> make a PR.

## Completeness

This is not a complete representation of all data, but strives to be the place
to look.

Features are added as 

## Packages

This package is divided into subpackages:

- `html` contains mapping from HTML element tag names to IDL interface name.
- `idl` contains the Web IDL specifications.

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
