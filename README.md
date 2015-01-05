# go-spider

## Example

```go
spider := NewSpider()
spider.Pipe(Retry{})

NewSpider()
    .Pipe(a)
    .Pipe(b)
    .IfUri('xxx', p)
    .IfUri('xxx', bb)
    
```