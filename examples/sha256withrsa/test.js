import sign from 'k6/x/sign'

const privateKey = open('./private.pem')

export default function () {
    const clearText = 'Hello World!'
    console.log(sign.sha256WithRsa(privateKey, clearText))
}