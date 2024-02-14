import sign from 'k6/x/sign'
import encoding from 'k6/encoding';

const privateKey = open('./private.pem')
const clearText = open('./sign.txt')

export default function () {
    console.log(encoding.b64encode(sign.sha256WithRsa(privateKey, clearText)))
}