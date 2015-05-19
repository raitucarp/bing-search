# Bing web search for Go language

Get useful information with Bing search engine


# Usage
Run this in your terminal
```
go get github.com/raitucarp/bing-search
```
Then import it in your Go code:
```
import "github.com/raitucarp/bing-search"
```


# Example
```
package main

import (
    "github.com/raitucarp/bing-search"
    "fmt"
)

func main() {
    options := Options{
        Query: "how to create cookies?",
        Count: 10,
        Tor: false,
    }
    results, ok := search.WebSearch(options)
    if !ok {
        fmt.Println("not ok")
    }
    
    // results will be []Item
    fmt.Println(results)
    
    // example of urlinfo
    info, _ := URLinfo("http://techcrunch.com/2015/05/12/the-ultimate-interface-is-your-brain/", false)
    fmt.Println(info)
}
```
# Usage
type Options is required when doing WebSearch

## Options
```
type Options struct {
    Query string
    Count int
    Tor bool
}
```

## WebSearch(options *bing.Options) (results []Item, ok bool)

Get search results with bing web search

## URLInfo(url string, withTor bool) (info Item, ok bool)

Get single url Info

# TODO
- Image Search
- Build Advance Query
- Refactoring code
- Documentation
- Fix some bugs

# License

The MIT License (MIT)

Copyright (c) 2015 Ribhararnus Pracutiar

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
