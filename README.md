<img src="https://user-images.githubusercontent.com/2049665/29219793-b4dcb942-7e7e-11e7-8785-761b0e784e04.png" width=100px />

# Word Ninja

Slice your munged together words! Seriously, Take anything, `'imateapot'` for example, would become `[im, a, teapot]`. Useful for humanizing stuff (like database tables when people don't like underscores).

## Usage
```go
package main

import (
	"github.com/goftic/wordninja"
	"fmt"
)

func main()  {
	// only English characters
	eng := "derekanderson"
	fmt.Println(wordninja.SplitEnglish(eng))
	
	// multi characters
	mul := "this哈isa,test"
	fmt.Println(wordninja.Split(mul))
}
```

## How to Install
```bash
go get github.com/goftic/wordninja
```

## reference
> 1. https://pypi.org/project/wordninja/
> 2. https://github.com/keredson/wordninja
> 3. https://stackoverflow.com/questions/8870261/how-to-split-text-without-spaces-into-list-of-words
> 4. https://github.com/willsmil/go-wordninja
