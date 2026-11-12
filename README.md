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

Pass the `--colour` flag to print the message in a colour. The named
colours `red`, `blue` and `green` are supported, as are six-digit HTML
colour hashes such as `#FF0000` (case-insensitive):

```sh
./fs-scratch-0 --colour "#FF0000"
```

## Testing

Run the test suite:

```sh
go test ./...
```

Run the vet checks:

```sh
go vet ./...
```

Check formatting (expected to print nothing):

```sh
gofmt -l .
```
