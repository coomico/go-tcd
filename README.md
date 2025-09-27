
# Go TCD

Go package to get Tax Court Decisions (ID)


```bash
go get github.com/coomico/go-tcd
```
## Example

```go
package main

import (
	"log"

	"github.com/coomico/go-tcd"
)

func main() {
	raw, err := tcd.New().FetchData()
	if err != nil {
		panic(err)
	}

	raw.GetFileBulk()
}
```
Another example [here](https://github.com/coomico/go-tcd/tree/main/_example).

## Reference

Scrap PP: https://github.com/aldofebriii/backend-stanic/tree/main/scrap-pp