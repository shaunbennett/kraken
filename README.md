# kraken

[![build](https://github.com/shaunbennett/kraken/workflows/build/badge.svg)](https://github.com/shaunbennett/kraken/actions?query=workflow%3Abuild)
[![Go Report Card](https://goreportcard.com/badge/github.com/shaunbennett/kraken)](https://goreportcard.com/report/github.com/shaunbennett/kraken)
[![GoDoc](https://godoc.org/github.com/shaunbennett/kraken?status.svg)](https://godoc.org/github.com/shaunbennett/kraken)

A CLI and Go Client for the [Kaken API](https://www.kraken.com/features/api).

Note that this project is a work in progress. It is not ready for public use yet.

## Usage

### CLI

There is still a lot of work to be done for the CLI, check back later for installation and usage instructions.

### Go Client

kraken can also be used as a Go library to interact with Kraken's API.

```shell
go get -u github.com/shaunbennett/kraken
```

Now you can use kraken in your application:

```go
import "github.com/shaunbennett/kraken"

api := kraken.New("API_KEY")
```

The Go client provides types functions for each API call in the Kraken API. If needed, you can also make API calls directly,
deserializing into your own type or raw json values.

## Examples

See the `examples` folder for small examples using both the CLI and library.

## License

This code is licensed under the [MIT License](https://github.com/shaunbennett/kraken/blob/main/LICENSE).
