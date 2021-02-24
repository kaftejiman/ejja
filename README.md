# ejja

A modular Golang source code level control flow obfuscator, made in Golang.

Runs the selected module on the target project.


*Please report any bugs you encounter.*

## Architecture

![architecture](https://raw.githubusercontent.com/kaftejiman/ejja/main/assets/architecture.png)

## Quick Start

Usage:

```
λ ejja run --project "C:\Users\kaftejiman\sampleGolangProject" --module "analyser"

[+] Running analyser..

=== Object Summary ===

basic objects        : 2
array objects        : 2
slice objects        : 13
struct objects       : 5
pointer objects      : 2
tuple objects        : 1
signature objects    : 29
interface objects    : 3
map objects          : 0
chan objects         : 1

λ ejja run --project "C:\Users\kaftejiman\sampleGolangProject" --module "flattener" --function "main"

[+] Running flattener..
[+] Found function main in C:\Users\kaftejiman\sampleGolangProject\main.go, flattening..

[+] Done.

```

![quickstart](https://raw.githubusercontent.com/kaftejiman/ejja/main/assets/quickstart.gif)

## Available Modules

### Analyser

Runs an analysis on the target project's codebase, returns summary of object analysis.

* Name: analyser
* Usage: `ejja --project="example/project" --module="analyser"`

### Flattener

Flattens the target function's control flow graph. [More](http://ac.inf.elte.hu/Vol_030_2009/003.pdf)

* Name: flattener
* Usage: `ejja --project="example/project" --module="flattener" --function="main"`
  
*before* *after*

*example cfgs from r2*


## Install

Install directions

## How to contribute your own module

Each module should export two required methods:
 * `Manifest()` -- Module manifestation with a unique name and description.
 * `Run()` -- The entry point of the module.

You can find template module in samples folder.

## Known issues

## Release Notes

**[CHANGELOG](https://github.com/kaftejiman/ejja/blob/main/CHANGELOG.md)**

