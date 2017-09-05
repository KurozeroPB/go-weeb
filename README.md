# go-weeb
__go-weeb is here to make it simple getting images from rra.ram.moe__

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
Here's a list of types you can use to get images from:
- cry
- cuddle
- hug
- kiss
- lewd
- lick
- nom
- nyan
- owo
- pat
- pout
- rem
- slap
- smug
- stare
- tickle
- triggered
- nsfw-gtn
- potato
- kermit
