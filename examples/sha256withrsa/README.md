This is a example show how to use sha256withrsa

## Run
1. Make a file with a name like `test.js` then add this code:

```JavaScript
import sign from 'k6/x/sign'

const privateKey = open('./private.pem')

export default function () {
    const clearText = 'Hello World!'
    console.log(sign.sha256WithRsa(privateKey, clearText))
}
```

2. Run the test with ./k6 run test.js.
```shell
./k6 run test.js
```

It should output the following:

```shell
INFO[0000] dIfTth2navNGkjE4Mbxc6vFICPWDzWteku6oX2bAdZw/ft1u6TKolQz0tHL329DBS7wo6cfSv/y9uLHOARPNiDeLW1ToQyC4dLELqameMUDhDswYlOMLPUoe24vyvLCLw44KxiDF7+6eP6Lmc8Q1viNrcnc8jztAGmf48wF4ZsMd0OyM3qWc/qpnou48AHtcFq27LXTFw9NzPTwRGX+gDsMSxxnjMBEPJDACZt4qjZYh2xYMaAYbknIAmaNOOtEUhZmmOYEsEhP7hLHpPH3Zx89q5Hh6YPWAmfPgdmcjwKG4vwaaJqaFGz3XICp8EDcvjAeXJ1bjMLMVreDgXt28cQ==  source=console
```

## Verify By Openssl
1. Make a file with a name like `sign.base64` then add output:
```base64
dIfTth2navNGkjE4Mbxc6vFICPWDzWteku6oX2bAdZw/ft1u6TKolQz0tHL329DBS7wo6cfSv/y9uLHOARPNiDeLW1ToQyC4dLELqameMUDhDswYlOMLPUoe24vyvLCLw44KxiDF7+6eP6Lmc8Q1viNrcnc8jztAGmf48wF4ZsMd0OyM3qWc/qpnou48AHtcFq27LXTFw9NzPTwRGX+gDsMSxxnjMBEPJDACZt4qjZYh2xYMaAYbknIAmaNOOtEUhZmmOYEsEhP7hLHpPH3Zx89q5Hh6YPWAmfPgdmcjwKG4vwaaJqaFGz3XICp8EDcvjAeXJ1bjMLMVreDgXt28cQ==
```

2. Decode base64 and output to a file like `sign.bin`
```shell
openssl base64 -d -A -in sign.base64 -out sign.bin
```

3. Verify By Openssl
```shell
openssl dgst -sha256 -verify public.pem -signature sign.bin sign.txt
```
It should output the following:
```shell
Verified OK
```
