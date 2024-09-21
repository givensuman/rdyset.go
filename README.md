# rdyset.go

Weirdly, the Go standard library does not have sets. I think sets are useful! I also think they're super neat.
This is a library that tries its best to fill that void.

## Usage
```bash
go get https://github.com/givensuman/rdyset.go
```

```go
import (
  "https://github.com/givensuman/rdyset.go"
)

func main() {
  A := Set(1, 2, 3)
}
```

## Should I use this?

Probably not. Check out [golang-set](https://github.com/deckarep/golang-set) instead.
