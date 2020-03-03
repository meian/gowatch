# gowatch

![Go](https://github.com/meian/gowatch/workflows/Go/badge.svg)
[![GitHub license](https://img.shields.io/github/license/meian/gowatch?color=orange)](https://github.com/meian/gowatch/blob/master/LICENSE)
[![codecov](https://codecov.io/gh/meian/gowatch/branch/master/graph/badge.svg)](https://codecov.io/gh/meian/gowatch)

Watch file change and run test.  
Current support is only test.

## Feature

- watch files in specific directory, and run test for **only change target** when detect changes.
  - when change test source(`foo_test.go`), run test only in this file.
  - when change no test source(`bar.go`), run test only in companion test file(`bar_test.go`). 
- watch single directory or recursive.

## Install

On go 1.12 or later with go modules, you can install with `go get`

```bash
go get github.com/meian/gowatch
```

## Usage

### test

Watch directory change and trigger test.

```bash
# watch current directory(no recursive)
$ gowatch test

# watch foo directory and subdirectories
$ gowatch test -r foo
```

All options shows with `gowatch test -h`
