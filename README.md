# Go Profiling Example

This repository contains a simple Go program demonstrating how to perform CPU and memory profiling using the `runtime/pprof` package.

## Overview

The program performs the following actions:
1. Starts CPU profiling and writes the profile to a file named `cpu.prof`.
2. Runs a simple test function (`testUser`) that simulates some work by sleeping for 10 seconds.
3. Stops CPU profiling.
4. Optionally writes a memory profile to a file named `mem.prof`.

## Prerequisites

- Go 1.16 or higher
- Basic knowledge of Go programming

## Getting Started

### Clone the Repository

```bash
git clone https://github.com/balasl342/go-pprof-examplee.git
```

### Running the Program

To run the program, use the following command:

```bash
go run main.go
```

This will generate two files:
- `cpu.prof`: Contains the CPU profiling data.
- `mem.prof`: Contains the memory profiling data.

### Analyzing the Profiles

You can analyze the generated profiles using the `go tool pprof` command.

#### CPU Profile

To analyze the CPU profile, run:

```bash
go tool pprof cpu.prof
```

This will open an interactive shell where you can explore the CPU usage data. Some useful commands within the pprof shell include:

- `top`: Shows the top functions by CPU usage.
- `list <function>`: Displays source code annotated with CPU usage for the specified function.
- `web`: Opens a graphical view of the profile in your web browser.

#### Memory Profile

To analyze the memory profile, run:

```bash
go tool pprof mem.prof
```

Similar to the CPU profile analysis, this will open the pprof shell for exploring memory usage data.

### Cleaning Up

To remove the generated profiling files:

```bash
rm cpu.prof mem.prof
```

## Code Structure

```plaintext
.
├── main.go        # Main Go file with profiling logic
└── README.md      # This README file
```

### main.go

The `main.go` file contains the following key sections:

- **CPU Profiling**: Starts and stops CPU profiling, saving the data to `cpu.prof`.
- **testUser Function**: Simulates some work by sleeping for 10 seconds.
- **Memory Profiling**: Writes a heap profile to `mem.prof` after the CPU profiling is completed.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- [Go Documentation on pprof](https://pkg.go.dev/runtime/pprof) for detailed information on profiling in Go.
- [Go tool pprof](https://pkg.go.dev/cmd/pprof) for profile analysis and visualization.
