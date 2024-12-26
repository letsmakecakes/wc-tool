# WC - Word Count

A clone of the Unix `wc` utility written in Go, implementing the core functionality of counting bytes, lines, words, and characters in text files.

## Features

- Count bytes (-c)
- Count lines (-l)
- Count words (-w)
- Count characters (-m)
- Default mode (equivalent to -c -l -w)
- Support for file input and standard input
- Unicode support

## Building

```bash
go build -o wc ./cmd/wc
```

## Usage

```bash
# Count bytes
./wc -c test.txt

# Count lines
./wc -l test.txt

# Count words
./wc -w test.txt

# Count characters
./wc -m test.txt

# Default mode (lines, words, bytes)
./wc test.txt

# Read from standard input
cat test.txt | ./wc -l
```

## Testing

```bash
go test ./...
```

## Project Structure

```
wc/
├── cmd/
│   └── wc/
│       └── main.go       # Entry point
├── internal/
│   └── counter/
│       ├── counter.go    # Core counting logic
│       └── counter_test.go
├── pkg/
│   └── utils/
│       └── utils.go      # Utility functions
├── go.mod
└── README.md
```

## License

MIT License