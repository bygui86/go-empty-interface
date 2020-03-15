
# Go empty interface

Project to explore Golang empty interface concept

## Definition

Internally, an interface value is two words (a.k.a. sets of bytes):
- One word points to the value underlying type.
- One word points to the data.

---

## Run

### with Type
```shell
go test ./... -bench=BenchmarkWithType -benchmem -run='^$' -count=3
```

### with Empty interface to Type
```shell
go test ./... -bench=BenchmarkWithEmptyInterfaceToType -benchmem -run='^$' -count=3
```

### with Empty interface to Pointer
```shell
go test ./... -bench=BenchmarkWithEmptyInterfaceToPointer -benchmem -run='^$' -count=3
```

BenchmarkWithEmptyInterfaceToType takes more nanoseconds (~50) for the double conversion type to empty interface and then to the type back than copying the structure.
The time will increase if the number of fields in the structure increases.

`TIP:` a good solution would be to use pointer and convert back to this same struct pointer. See `BenchmarkWithEmptyInterfaceToPointer` and `emptyInterfaceToPointer` methods

Regarding the basics type like int or string, the performances are slightly different:

```
BenchmarkWithTypeInt-8              2000000000          1.42 ns/op
BenchmarkWithEmptyInterfaceInt-8    1000000000          2.02 ns/op
BenchmarkWithTypeString-8           1000000000          2.19 ns/op
BenchmarkWithEmptyInterfaceString-8  50000000           30.7 ns/op
```

Well used and with parsimony, in most of the case the empty interface should have a real impact on the performances of your application.

---

## Links

- https://tour.golang.org/methods/14
- https://flaviocopes.com/go-empty-interface/
- https://golang.org/doc/faq#convert_slice_of_interface
- [medium blanchon.vincent](https://medium.com/@blanchon.vincent/go-understand-the-empty-interface-2d9fc1e5ec72)
- https://dev.to/flrnd/understanding-the-empty-interface-in-go-4652
