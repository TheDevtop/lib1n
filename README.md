# lib1n
This is a **very** simple package, for a **very** simple key/value format.
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
os=plan9
```

### Operations
The following operations are supported:

* Concurrent decode -> GoDecode(dataset)
* Concurrent find -> Find(dataset, pattern)
* Concurrent replace -> Replace(dataset, pattern, replacement)
* Procedural map -> Map(dataset, function)
* Procedural clean -> Clean(dataset)
* Procedural decode -> Decode(buffer)
* Procedural encode -> Encode(dataset)
