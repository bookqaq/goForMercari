package mercarigo

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/asn1"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"math/big"
	"os"
	"strings"
	"time"
)

type ECDSASignature struct {
	R, S *big.Int
}

type payload struct {
	Iat int64  `json:"iat"`
	Jti string `json:"jti"`
	Htu string `json:"htu"`
	Htm string `json:"htm"`
}

type pkey_jwk struct {
	Crv string `json:"crv"`
	Kty string `json:"kty"`
	X   string `json:"x"`
	Y   string `json:"y"`
}

type pkey_header struct {
	Typ string   `json:"typ"`
	Alg string   `json:"alg"`
	Jwk pkey_jwk `json:"jwk"`
}

func byteToBase64URL(target []byte) string {
	return base64.RawURLEncoding.EncodeToString(target)
}

func dPoPGenerator(uuid_ string, method string, url_ string) string { //因为有 url和uuid 包了
	private_key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		fmt.Println("Error at mercarigo//dPoP.go//dPoPGenerator//ecdsa.GenerateKey():\n", err)
		os.Exit(60)
	}

	pl := payload{time.Now().Unix(), uuid_, url_, strings.ToUpper(method)}
	pkjwk := pkey_jwk{"P-256", "EC", byteToBase64URL(private_key.PublicKey.X.Bytes()), byteToBase64URL(private_key.PublicKey.Y.Bytes())}
	pkh := pkey_header{"dpop+jwt", "ES256", pkjwk}

	headerString, err := json.Marshal(pkh)
	if err != nil {
		fmt.Println("Error at mercarigo//dPoP.go//dPoPGenerator//json.Marshal(pkh):\n", err)
		os.Exit(61)
	}
	payloadString, err := json.Marshal(pl)
	if err != nil {
		fmt.Println("Error at mercarigo//dPoP.go//dPoPGenerator//json.Marshal(pl):\n", err)
		os.Exit(62)
	}

	data_unsigned := fmt.Sprintf("%s.%s", byteToBase64URL(headerString), byteToBase64URL(payloadString))

	hval := sha256.Sum256([]byte(data_unsigned))

	signature, err := ecdsa.SignASN1(rand.Reader, private_key, hval[:])
	if err != nil {
		fmt.Println(err)
		os.Exit(63)
	}
	sig := &ECDSASignature{}
	if _, err := asn1.Unmarshal(signature, sig); err != nil {
		fmt.Println(err)
		os.Exit(64)
	}

	signatured := append(sig.R.Bytes(), sig.S.Bytes()...)

	signaturedString := byteToBase64URL(signatured)

	result := fmt.Sprintf("%s.%s", data_unsigned, signaturedString)
	return result
}
