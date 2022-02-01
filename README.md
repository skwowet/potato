# Potato

Potato is a golang module that is intended for testing & publishing modules on [pkg.go.dev](https://pkg.go.dev/) 

## Installation

Use the package manager [go get ](https://pkg.go.dev/cmd/go) to install foobar.

```bash
go get github.com/yeganathan18/potato
```

## Usage

```go
import (
	"fmt"
	"log"

	"github.com/yeganathan18/potato"
)

message, err := potato.Hello('John')

// If an error was returned, print it to the console
if err != nil {
    log.Fatal(err)
}

// returns 'Hi, John. Welcome!'
fmt.Println(message)
```

## Any Questions
If you have any questions or comments, please open an issue or open a pull request on this repository.

Feel free to use this is as a reference for creating your own modules and publish it on [pkg.go.dev](https://pkg.go.dev/).

## License
[MIT](LICENSE)