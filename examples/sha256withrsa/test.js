import sign from 'k6/x/sign'

const privateKey = open('./private.pem')
const clearText = open('./sign.txt')

export default function () {
    console.log(sign.sha256WithRsa(privateKey, clearText))
}