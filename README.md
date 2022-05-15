# hatebu
Crawl [hatebu](https://b.hatena.ne.jp/hotentry/it) to use something like daily chat bot

# Usage

write the code like this.

```
package main

import (
	"fmt"

	"github.com/maito1201/hatebu"
)

func main() {
	results, err := hatebu.ScrapeHotEntry()
	if err != nil {
		panic(err)
	}
	for _, v := range results {
		fmt.Printf("%s %s\n", v.Title, v.Href)
	}
}
```