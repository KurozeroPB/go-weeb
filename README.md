# go-weeb

## go-weeb is here to make it simple getting images from rra.ram.moe

### Install
`go get github.com/KurozeroPB/go-weeb`

### Usage
Quick example:
```go
package main

import (
	"fmt"

	"github.com/KurozeroPB/go-weeb"
)

func main() {
	img, err := weeb.GetImage("pat")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(img)
}
```
