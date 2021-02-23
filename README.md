# ejja

A modular Golang source code level control flow obfuscator, made in Golang.

Runs the selected module on the target project.


*Please report any bugs you encounter.*

## Overview

![overview](https://raw.githubusercontent.com/kaftejiman/ejja/main/assets/overview.png)

## Architecture

![architecture](https://raw.githubusercontent.com/kaftejiman/ejja/main/assets/architecture.png)

## Quick Start

Usage:

```
λ ejja run --project "C:\Users\kaftejiman\obfuscators\ejja" --module "analyser"

                === Object Summary ===

        basic objects        : 2
        array objects        : 0
        slice objects        : 0
        struct objects       : 5
        pointer objects      : 0
        tuple objects        : 0
        signature objects    : 29
        interface objects    : 3
        map objects          : 0
        chan objects         : 0
```

![quickstart](https://raw.githubusercontent.com/kaftejiman/ejja/main/assets/quickstart.gif)

## Available Modules

### Analyser

![analyser](https://raw.githubusercontent.com/kaftejiman/jamal/main/assets/analyser.png)

### Flattener

*before* *after*
*example cfgs*

![flattener](https://raw.githubusercontent.com/kaftejiman/jamal/main/assets/flattener.png)

## How to contribute your own module

Each module should provide two required methods:
 * `manifest()` -- Module manifestation with a unique name and description.
 * `run()` -- The entry point of the module.

## Known issues
## Requirements

* [Henrylee2cn aster](https://github.com/henrylee2cn/aster)

## Release Notes

**[CHANGELOG](https://github.com/kaftejiman/ejja/blob/main/CHANGELOG.md)**

