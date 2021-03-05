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
[+] Found function `main` in `test.go` ..

[+] Emitting transformed function..

func main(){
        a := []int{2, 212, 3001, 14, 501, 7800, 9932, 33, 45, 45, 45, 91, 99, 37, 102, 102, 104, 106, 109, 106}
        var c11apfbm9cj2dj36j3jg string
        c11apfbm9cj2dj36j3jg = "c11apfbm9cj2dj36j3k0"
        for c11apfbm9cj2dj36j3jg != "c11apfbm9cj2dj36j3kg" {
                switch(c11apfbm9cj2dj36j3jg){
                case "c11apfbm9cj2dj36j3k0":
                        fmt.Println(sort(a))
                        c11apfbm9cj2dj36j3jg = "c11apfbm9cj2dj36j3l0"
                        break
                case "c11apfbm9cj2dj36j3l0":
                        if (1 > 2) {
                                c11apfbm9cj2dj36j3jg = "c11apfbm9cj2dj36j3m0"
                        }else{
                                c11apfbm9cj2dj36j3jg = "c11apfbm9cj2dj36j3mg"
                        }
                        break
                case "c11apfbm9cj2dj36j3mg":
                        fmt.Println("yes")
                        c11apfbm9cj2dj36j3jg = "c11apfbm9cj2dj36j3lg"
                        break
                case "c11apfbm9cj2dj36j3lg":
                        i := 0
                        c11apfbm9cj2dj36j3jg = "c11apfbm9cj2dj36j3o0"
                        break
                case "c11apfbm9cj2dj36j3o0":
                        if i < 5 {
                                c11apfbm9cj2dj36j3jg = "c11apfbm9cj2dj36j3ng"
                        }else{
                                c11apfbm9cj2dj36j3jg = "c11apfbm9cj2dj36j3n0"
                        }
                        break
                case "c11apfbm9cj2dj36j3og":
                        i++
                        c11apfbm9cj2dj36j3jg = "c11apfbm9cj2dj36j3o0"
                        break
                case "c11apfbm9cj2dj36j3n0":
                        fmt.Println(fibonacci(30))
                        c11apfbm9cj2dj36j3jg = "c11apfbm9cj2dj36j3kg"
                        break
                }
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
Issues with nested statements, will be fixed asap.

Supported statements:
* ExprStatements
* IfStatements
* ForStatements
* ReturnStatements
* AssignmentStatements
## Release Notes

**[CHANGELOG](https://github.com/kaftejiman/ejja/blob/main/CHANGELOG.md)**

