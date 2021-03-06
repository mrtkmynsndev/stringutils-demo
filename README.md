![Version](https://img.shields.io/badge/version-0.0.4-orange.svg)
![Test](https://github.com/mrtkmynsndev/stringutils-demo/actions/workflows/go.yml/badge.svg)
![Go](https://img.shields.io/github/go-mod/go-version/mrtkmynsndev/stringutils-demo)
![MIT License](https://img.shields.io/github/license/mrtkmynsndev/stringutils-demo)
[![Documentation](https://godoc.org/github.com/mrtkmynsndev/stringutils-demo?status.svg)](https://pkg.go.dev/github.com/mrtkmynsndev/stringutils-demo)
[![Go Report Card](https://goreportcard.com/badge/github.com/mrtkmynsndev/stringutils-demo)](https://goreportcard.com/report/github.com/mrtkmynsndev/stringutils-demo)

# stringutils

A basic golang package for demonstration purpose. Package currently contains 
only one function:

```go
func Reverse(s string) string
```

---

## Installation

Go to your project root, where `go.mod` file exists, than grab the library via:

```bash
go get github.com/mrtkmynsndev/stringutils-demo@latest
```

---

## Usage

```go
package main

import (
	"fmt"

	"github.com/mrtkmynsndev/stringutils-demo"
)

func main(){
	reversed := stringutils.Reverse("mert")
	fmt.Println(reversed) // trem
}
```

---

## Makefile

```bash
make help
```

Commands usage:

```bash
make <command>

commands:

test    run tests
bench   run benchmark tests
doc     run godoc server at 3000 unless PORT env-var is set
```

- `make test`: Runs tests
- `make bench`: Runs benchmark tests
- `make doc`: Runs godoc server on port 3000. Use `PORT` environment variable
  for different port -> `PORT=4000 make doc`

---

## Contributor(s)

* [Mert Kimyonşen](https://github.com/mrtkmynsndev) - Creator, maintainer

---

## Contribute

All PR’s are welcome!

1. `fork` (https://github.com/mrtkmynsndev/stringutils-demo/fork)
1. Create your `branch` (`git checkout -b my-feature`)
1. `commit` yours (`git commit -am 'add some functionality'`)
1. `push` your `branch` (`git push origin my-feature`)
1. Than create a new **Pull Request**!

---

## License

This project is licensed under MIT

---

This project is intended to be a safe, welcoming space for collaboration, and
contributors are expected to adhere to the [code of conduct][coc].

[coc]: https://github.com/mrtkmynsndev/stringutils-demo/blob/main/CODE_OF_CONDUCT.md