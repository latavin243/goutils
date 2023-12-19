# goutils

This projects provides many Golang tools for reference.  
Most packages supports generics, which needs Golang version >= 1.18.  
To use tools in this project, just `go get` the project and import the packages:

```bash
go get github.com/latavin243/goutils
```

## utils without dependency

- [x] set - a set implementation with generics
- [x] strcase - convert between camelCase, snake_case, TitleCase, etc
- [x] fnwrap - wrap a func with additional functions
- [x] iterop - operation funcs for iterables, e.g. map, flatmap, etc
- [x] number - number conversion, etc
- [x] roundrobin - round-robin balancing
- [x] stopwatch - a stopwatch to get time duration elapsed
- [x] ringbuffer - ring buffer with generics
- [x] ringhash - ring hash for consistent hashing
- [ ] bus - message bus for pub/sub, producer/consumer

## utils based on other packages

- [x] hashutil - hash functions and examples, e.g. md5, murmur3, etc
- [x] reflectutil - functions based on golang reflect
- [x] routinegroup - goroutine group, depends on ants, errgroup, etc
- [x] filereader - read file (e.g. config) from json, yaml, toml, etc
- [x] timeutil - functions based on golang time package
- [x] requtil - http request utils
- [x] sampling - a little wrap of downsampling package
- [x] fileutil - file tools, e.g. absolute path, size, md5, etc
- [x] excelutil - excel functions in practice based on qax-os/excelize

## todo

- [ ] reflectutil struct fold support more types: number, string, bool, slice, map, etc
