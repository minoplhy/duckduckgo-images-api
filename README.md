# DuckDuckGo-images-api

This project is Go adaptation of Python3 forked [duckduckgo-images-api](https://github.com/joeyism/duckduckgo-images-api) . I made this modules because this module will be need in my future project .

Contribution are always welcome

# How to use

Import

# Example Use Cases

#### Get Search Results
```go
func main() {
	hunsen := goduckgo.Search(goduckgo.Query{Keyword: "duck"})
    fmt.Print(hunsen.Results)
}

```
#### Get Search Result Image
```go
func main() {
	hunsen := goduckgo.Search(goduckgo.Query{Keyword: "duck"})
	for _, somtam := range hunsen.Results {
        // This Can be use with all hunsen.Results(or anything.Results depending on your goduckgo.Search)
        // e.g. Title or URL depending on what you want 
		fmt.Println(somtam.Image)
	}
}
```
#### Specific P and S
```go
func main() {
	hunsen := goduckgo.Search(goduckgo.Query{Keyword: "duck", P: "1", S: "200"})
    fmt.Print(hunsen.Results)
}

```
