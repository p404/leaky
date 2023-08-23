# Leaky
## Overview

This Go program serves as a resource leak simulator, capable of intentionally causing memory leaks and file descriptor leaks for educational or research purposes. The program provides command-line options to control the type of resource to leak.

## Installation

1. Clone this repository.
2. Navigate to the directory containing the source code.
3. Run `go build` to compile the program.

## Usage

Run the compiled executable with the desired flags to simulate the corresponding type of resource leak:

- To leak file descriptors:

  ```bash
  ./leaky --fds
  ```

- To leak memory:

  ```bash
  ./leaky --memory
  ```

- To leak both file descriptors and memory:

  ```bash
  ./leaky --fds --memory
  ```

## CLI Options

- `--fds`: Activate file descriptor leak.
- `--memory`: Activate memory leak.