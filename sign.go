package sign

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"log"
	"go.k6.io/k6/js/modules"
)

func init() {
    modules.Register("k6/x/sign", new(Sign))
}

type Sign struct {}


func (s *Sign) Sha256WithRsa(priKey, clearText string) []byte {
	block, _ := pem.Decode([]byte(priKey))
	privKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		log.Fatal("read private key err: " + err.Error())
	}

	hasher := sha256.New()
	hasher.Write([]byte(clearText))
	sum := hasher.Sum(nil)
	signature, err := rsa.SignPKCS1v15(rand.Reader, privKey.(*rsa.PrivateKey), crypto.SHA256, sum)
	if err != nil {
		log.Fatal("rsa sign err: " + err.Error())
	}
	return signature
}
