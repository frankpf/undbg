# undbg

undbg is an experimental debugger with support for reverse execution.

## Installing

For now, the only way to install is to build from source. Before installing, make sure you have the Go toolchain configured.

    git clone https://github.com/frankpf/undbg
    cd undbg
    go build src/main.go -o undbg

## Usage

    undbg <target program>

Currently, the only implemented command is `step` (aliased to `s`).

There is a sample program available in `samples/hello.asm`. You can use it to test the undbg. To build the executable, run the following commands:

    nasm -felf64 samples/hello.asm -o samples/hello.o
    ld samples/hello.o -o samples/hello
    undbg ./samples/hello

Example undbg session:

![example undbg session](./docs/undbg_session.png)
