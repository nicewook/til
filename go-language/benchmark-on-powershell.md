# Benchmark on PowerShell of Windows

Generally, you can benchmark like this
```
$ go test -bench=.
```

But, on PowerShell, you should use quote
```
$ go test -bench="."
```
