# fs-scratch-0

This project is a console shell terminal program written in golang. It writes a message or messages to stdout then exits.

## Usage

Requires Go 1.21 or later.

Build the project:

```sh
make build
```

or directly:

```sh
go build -o fs-scratch-0 .
```

This produces the `fs-scratch-0` binary.

Run it to print the default message:

```sh
./fs-scratch-0
```

```text
Hello World!
```

Pass the `-message` flag to print a custom message:

```sh
./fs-scratch-0 -message "your message"
```
