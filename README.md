# fossil

fossil is a small kinda-but-not-really code analyzer for Go. It walks a directory structure and builds a memory structure of functions within files within packages. It can be queried against to find all files in a package, all files that contain a specific function, and `all`.

`main.go` is moreso an example on how to use fossil (no idea where I got the name, was just feeling it) than anything.

fossil was built as part of a take-home project for a job interview with Orijtech.

## Getting Started

fossil was built using Go 1.17.2 and has been tested and targetted against the following OSes:
* x86_64 Windows 10/11
* x86_64 Linux
* x86_64 macOS
* ARM M1 macOS
* Windows 11 WSL2