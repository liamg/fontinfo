# FontInfo

FontInfo is a Go package to list available fonts on a Linux system.

- No CGO required
- Doesn't wrap `fontconfig` or other utilities
- Pure Go
- No external dependencies
- Provides `family` and `style` for each font
- Supports TTF and OTF
- Fast (typically parses 1k fonts in ~100ms)

## Example

```go
package main

import (
	"fmt"

	"github.com/liamg/fontinfo"
)

func main() {

	fonts, err := fontinfo.List()
	if err != nil {
		panic(err)
	}

	for _, font := range fonts {
		fmt.Printf("Family=%s Style=%s Path=%s\n", font.Family, font.Style, font.Path)
	}
}

```