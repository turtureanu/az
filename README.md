# az

Simple *fast* file finder made in [Go](https://go.dev/).

## Usage

```man
Search for files matching a pattern

Usage:
  az [pattern] [path] [flags]

Flags:
  -d, --detail   results return filename, path, size, date modified, and file type
  -H, --hidden   search hidden files and directories
  -j, --json     output results in JSON format
  -h, --help     help for az
```

RegEx support for patterns:

- \*
- ?
- \\ escaping
- character range
- character class
