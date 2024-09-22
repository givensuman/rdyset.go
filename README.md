# rdyset.go

Weirdly, the Go standard library does not have sets. I think sets are useful! I also think they're super neat.
This is a "good enough for government work" implementation.

## Usage

```bash
go get https://github.com/givensuman/rdyset.go
```

```go
import (
  "https://github.com/givensuman/rdyset.go"
)

func main() {
  A := Set(1, 2, 3) // { 1, 2, 3 }
  A.Add(3) // { 1, 2, 3 }
  A.Remove(2) // { 1, 3 }
  B := A.Union(Set(5, 7)) // { 1, 3, 5, 7 }
}
```

## Should I use this?

Probably not. Check out [golang-set](https://github.com/deckarep/golang-set) instead.
