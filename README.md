# Some `go` module named `gen` 

## Getting Started

`go get -u github.com/alexDtorres/gen`


## Example

```
package main

import "github.com/alexDtorres/gen"

func main() {
  pkg := gen.Install()
  pkg.Apply()
}
```


## Install the CLI Tool

`go install github.com/alexDtorres/gen/app@latest`


## CLI Tool Usage

`app gen <your_app_name>`


