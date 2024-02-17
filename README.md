
# Go LPM Trie and Go LPM Hash Map Library

This is a standalone library that implements the `longest prefix match` algorithm
for use with `Subnets`, `CIDRs` and `IPs`. It supports both `IPv4` and `IPv6`
prefix notations.


## Usage

The [examples](/examples) folder contains usage examples.

As pointed out in my [Web Log article about LPM Tries](https://cookie.engineer/weblog/articles/you-dont-need-lpm-tries.html)
it is heavily recommended to use the `HashMap` wherever possible, as it's the
most efficient way to represent a `longest prefix match` respecting data
structure.

```go
import "github.com/cookiengineer/lpm"
import "fmt"

hashmap := lpm.NewHashMap()

hashmap.Insert("192.168.0.0/24")
hashmap.Insert("192.169.128.0/17")
hashmap.Insert("192.169.0.0/16")
hashmap.Insert("192.169.0.0/24")

fmt.Println("parent of 192.168.0.123 is:", hashmap.Search("192.168.0.123/32"))
fmt.Println("parent of 192.169.0.123 is:", hashmap.Search("192.169.0.123/32"))

fmt.Println("")
fmt.Println("hashmap:")
hashmap.Print()
```


## License

AGPL3

