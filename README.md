# xk6-sign

This is a k6 extends to support 数字签名, 目前支持如下签名:
* sha256withrsa


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