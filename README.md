# lib1n
This is a simple package, for a **very** simple key/value format.
With this format you can map 1 key to n values, provided those keys and values are strings.

### Example
Here is some example content:
```
foo=bar;baz
lorem=ipsum;dolor;sit;amet
hotel=trivago
this=Sparta!
goto=considered harmful
oop=no
csp=yes
```

### Operations
Concurrent operations:
* Concurrent decode -> GoDecode(dataset)
* Concurrent encode -> GoEncode(dataset)
* Concurrent find -> Find(dataset, pattern)
* Concurrent replace -> Replace(dataset, pattern, replacement)
* Concurrent clean -> Clean(dataset)

Procedural operations:
* Procedural map/filter -> MapFilter(dataset, function)
* Procedural reduce -> Reduce(dataset, function)
* Procedural decode -> Decode(buffer)
* Procedural encode -> Encode(dataset)
