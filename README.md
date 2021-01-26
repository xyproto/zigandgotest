# Zig and Go test

## Requirements

* `zig`, `go` and `make`.
* x86_64 Linux.

## Building

Just type `make` to build on Arch Linux.

## Running

Just type `make run`. This fails.

`go build` followed by `make run` works.

## What happens?

It builds, but when running I get:

`./zigandgotest: error while loading shared libraries: /usr/lib/libc.so: invalid ELF header`

Why? I don't know, yet.
