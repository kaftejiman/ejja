# ejja

A modular Golang source code level control flow obfuscator, made in Golang.

Runs the selected module on the target project.


*Please report any bugs you encounter.*

## Architecture

![architecture](https://raw.githubusercontent.com/kaftejiman/ejja/main/assets/architecture.png)

## Quick Start

Usage:

```golang
λ go run main.go run --project "samples" --module "flattener" --functions "main"
[+] Running flattener..
[+] Found function `main` in `main`..

a := []int{2, 212, 3001, 14, 501, 7800, 9932, 33, 45, 45, 45, 91, 99, 37, 102, 102, 104, 106, 109, 106}

var c10hsl3m9cj2651d78mg string
c10hsl3m9cj2651d78mg = "c10hsl3m9cj2651d78n0"
for c10hsl3m9cj2651d78mg != "c10hsl3m9cj2651d78ng" {
        switch(c10hsl3m9cj2651d78mg){
        case "c10hsl3m9cj2651d78n0":
                 fmt.Println(sort(a))
                 c10hsl3m9cj2651d78mg = "c10hsl3m9cj2651d78o0"
                 break
        case "c10hsl3m9cj2651d78o0":
        if (1 > 2) {
                 c10hsl3m9cj2651d78mg = "c10hsl3m9cj2651d78p0"
        }else{
                 c10hsl3m9cj2651d78mg = "c10hsl3m9cj2651d78pg"
        }
        break
        case "c10hsl3m9cj2651d78p0":
                 fmt.Println("no")
                 c10hsl3m9cj2651d78mg = "c10hsl3m9cj2651d78og"
                 break
                case "c10hsl3m9cj2651d78pg":
                         fmt.Println("yes")
                         c10hsl3m9cj2651d78mg = "c10hsl3m9cj2651d78og"
                         break
        case "c10hsl3m9cj2651d78og":
                 fmt.Println(fibonacci(30))
                 c10hsl3m9cj2651d78mg = "c10hsl3m9cj2651d78ng"
                 break

        }

}

[+] Done.

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

```

![quickstart](https://raw.githubusercontent.com/kaftejiman/ejja/main/assets/quickstart.gif)

## Available Modules

### Flattener

Flattens the target function's control flow graph. [More](http://ac.inf.elte.hu/Vol_030_2009/003.pdf)

* Name: flattener
* Usage: `ejja --project="example/project" --module="flattener" --function="main"`
  
*before* *after*

*example cfgs from r2*

### Analyser

Runs an analysis on the target project's codebase, returns summary of object analysis.

* Name: analyser
* Usage: `ejja --project="example/project" --module="analyser"`

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

