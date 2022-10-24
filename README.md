# wordninja

## Getting Started
```bash
go get github.com/goftic/wordninja
```

## Usage
```go
package main

import (
	"github.com/goftic/wordninja"
	"fmt"
)

func main()  {
	// only English characters
	eng := "thisisatest"
	fmt.Println(wordninja.SplitEnglish(eng))
	
	// multi characters
	mul := "this哈isa,test"
	fmt.Println(wordninja.Split(mul))
}
```
## reference
> 1. https://pypi.org/project/wordninja/
> 2. https://github.com/keredson/wordninja
> 3. https://stackoverflow.com/questions/8870261/how-to-split-text-without-spaces-into-list-of-words
> 4. https://github.com/willsmil/go-wordninja
