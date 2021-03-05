# ejja

A modular Go source code level control flow obfuscator, made in Go.

Runs the selected module on the target project.


## Demonstration

![quickstart](https://raw.githubusercontent.com/kaftejiman/ejja/main/assets/quickstart.gif)

## Quick Start

Usage:

```golang
λ ejja run --project "samples" --module "flattener" --functions "main"
[+] Running flattener..
[+] Found function `main` in `main`..
[+] Emitting body of the transformed function..

    a := []int{2, 212, 3001, 14, 501, 7800, 9932, 33, 45, 45, 45, 91, 99, 37, 102, 102, 104, 106, 109, 106}

    var c10hsl3m9cj2651d78mg string
    c10hsl3m9cj2651d78mg = "c10hsl3m9cj2651d78n0"
    for c10hsl3m9cj2651d78mg != "c10hsl3m9cj2651d78ng" {
            switch(c10hsl3m9cj2651d78mg){
            case "c10hsl3m9cj2651d78n0":
                    fmt.Println(Sort(a))
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

[+] Done.
```

## Available Modules

### Flattener

Flattens the target function's control flow graph.

Implementation of the flattener algorithm in ["OBFUSCATING C++ PROGRAMS VIA CONTROL FLOW FLATTENING" by T. L ́aszl ́o and ́A. Kiss](http://ac.inf.elte.hu/Vol_030_2009/003.pdf).

* Name: flattener
* Usage: `ejja --project="example/project" --module="flattener" --function="main"`
* Description: *The idea behind control flow flattening is to transform the structure of the source code in such a way that the targets of branches cannot be easily determined by static analysis, thus hindering the comprehension of the program.*  

#### Before/After source code level
  
![before/after source code level](assets/before_after.png)

#### Before/After binary level (IDA 7.0)

![Before/After binary level (IDA 7.0)](assets/ida_comparison.png)



### Analyser

Displays object metrics about the target project codebase, returns summary of object analysis.

* Name: analyser
* Usage: `ejja --project="example/project" --module="analyser"`
* Description: *Runs an analysis on the target project's codebase, returns summary of object analysis.*

## Install


## How to contribute your own module

Each module should export two required methods:
 * `Manifest()` -- Module manifestation with a unique name and description.
 * `Run()` -- The entry point of the module.

You can use helper functions found in utils. They provide basic ast operations.

You can find a sample module in samples folder, move the sample module to `modules` folder for actually running.

## Known issues

Flattening module is lacking a lot as of now, I will be supporting more statements asap, feel free to PR.

Supported statements:
* ExprStatements
* IfStatements
* ForStatements
## Release Notes

**[CHANGELOG](https://github.com/kaftejiman/ejja/blob/main/CHANGELOG.md)**

