# az

Simple *fast* file finder made in [Go](https://go.dev/).

## Usage

```man
Usage:
  az [pattern] [path] [flags]

Flags:
  -d, --detail   results return filename, path, size, date modified, and file type
  -h, --help     help for az
  -H, --hidden   search hidden files and directories
  -j, --json     output results in JSON format
```

RegEx support for patterns:

- *
- ?
- \\ escaping
- character range
- character class
