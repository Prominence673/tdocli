## License

MIT

![Go](https://img.shields.io/badge/Go-1.25-blue)
![License](https://img.shields.io/badge/License-MIT-green)
![Platform](https://img.shields.io/badge/Platform-Windows%20%7C%20Linux%20%7C%20macOS-orange)

# TDOCLI

A lightweight and fast task manager written in Go.

TDOCLI is a cross-platform command-line application for managing tasks with JSON persistence.

## Features

- ✅ Add tasks
- ✅ Edit tasks
- ✅ Remove tasks
- ✅ Mark tasks as completed
- ✅ Undo completed tasks
- ✅ Export task lists
- ✅ Use custom task files
- ✅ Colored terminal output
- ✅ Cross-platform binaries

## Installation

### Linux / macOS

```bash
curl -sSL https://raw.githubusercontent.com/Prominence673/tdocli/main/install.sh | bash
```

### Go

```bash
go install github.com/Prominence673/tdocli@latest
```

## Usage

![Demo](demo.gif)

```bash
tasker add "Learn Go"
```

```bash
tasker list
```

```bash
tasker done 1
```

```bash
tasker edit 1 "Learn Go deeply"
```

```bash
tasker remove 1
```

```bash
tasker export backup
```

Using another task file:

```bash
tasker --file work.json list
```

## Built With

- Go
- Cobra
- Lip Gloss
