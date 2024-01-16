# WC - Word Count Command in Golang

A simple implementation of the `wc` (word count) command in Golang. This project provides basic functionality to count lines, words, characters, and bytes in a file.

## Table of Contents

- [About the Project](#about-the-project)
- [Usage](#usage)
- [Options](#options)
- [Getting Started](#getting-started)

## About the Project

The `wc` command is a utility in Unix and Unix-like operating systems that provides line, word, and byte count options. This Golang implementation offers a subset of the `wc` functionality.

## Usage

To use this `wc` implementation, run the following command:

```bash
go run main.go -<option> <filename>
```

Replace `<option>` with one of the following:

- `-c` : Byte count
- `-l` : Line count
- `-w` : Word count
- `-m` : Character count
- (No option) : Display line, word, byte count

Replace `<filename>` with the path to the file you want to analyze.

Example:

```bash
go run main.go -l sample.txt
```

## Options

- `-c` : Display the byte count.
- `-l` : Display the line count.
- `-w` : Display the word count.
- `-m` : Display the character count.
- (No option) : Display all counts (lines, words, bytes).

### Sample Output Cases

Assuming your `main.go` is set up to handle command-line arguments and call the appropriate `wc` function:

#### Byte Count (`-c`)

```bash
go run main.go -c example.txt
# Output: 123 example.txt
```

#### Line Count (`-l`)

```bash
go run main.go -l example.txt
# Output: 10 example.txt
```

#### Word Count (`-w`)

```bash
go run main.go -w example.txt
# Output: 35 example.txt
```

#### Character Count (`-m`)

```bash
go run main.go -m example.txt
# Output: 157 example.txt
```

#### All Counts (No Option)

```bash
go run main.go example.txt
# Output: 10 35 123 example.txt
```

## Getting Started

1. Clone the repository:

```bash
git clone https://github.com/yourusername/your-wc-project.git
cd your-wc-project
```

2. Run the example:

```bash
go run main.go -c example.txt
```
