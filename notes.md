# Go Foundation

## Magesh Kuppan
- tkmagesh77@gmail.com

## Schedule
| What | When |
|-----|----|
| Commence | 9:00 AM |
| Tea Break | 10:30 AM (20 mins) |
| Lunch Break | 12:30 PM (1 hour) |
| Tea Break | 3:00 PM (20 mins) |
| Wind up | 4:30 PM |

## Software Prerequisites
- Go Tools (https://go.dev/dl)
- Visual Studio Code (http://code.visualstudio.com)
    - Go Extension (https://marketplace.visualstudio.com/items?itemName=golang.Go)

## Repository
- https://github.com/tkmagesh/cisco-gofoundation-dec-2025

## Methodology
- No powerpoint
- Code & Discussion

## Why Go?
- Simplicity
    - ONLY 25 keywords
    - No access modifiers (public, private, protected)
    - No reference types (everything is a value)
    - No class (only structs)
    - No inheritance (only composition)
    - No pointer arithmatic
    - No exceptions (only errors)
    - No try..catch..finally construct
- Performance
    - Equivalent to C++
    - Close to hardware
        - Compile to machine code
        - Tooling support for cross-compilation
- Support for Concurrency
    - Concurrency problems are typically solved using OS Threads
        - OS Threads are costly (~2MB of memory)
    - Go offers a cheaper & efficient alternatives called "Goroutines"
        - Goroutines (2KB of memory)
        - Built in scheduler
    - Language support for concurrency
        - "go" `keyword`
        - "chan" `data type`
        - "<-" `operator`
        - "range" `construct`
        - "select-case" `construct`
    - Standard Library support
        - `sync` package
        - `sync/atomic` package