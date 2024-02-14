# xk6-sign

This is a k6 extension that supports some signature algorithms

## Use
You need to install xk6:
```bash
$ go install go.k6.io/xk6/cmd/xk6@latest
```

### Compile
```bash
$ xk6 build --with xk6-sign=.
```

now you can use `sign` module


## signature algorithms
* [sha256withRsa](./examples/sha256withrsa/README.md)