package mercarigo

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

type payload struct {
	Iat int    `json:"iat"`
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

func intToByte(target int) []byte {
	result := make([]byte, 8)
	binary.BigEndian.PutUint64(result, uint64(target))
	return result
}

func intToBase64URL(target int) string {
	return byteToBase64URL(intToByte(target))
}

func byteToBase64URL(target []byte) string {
	return strings.TrimRight(base64.StdEncoding.EncodeToString(target), "=")
}

func dPoPGenerator(uuid_ string, method string, url_ string) string { //因为有 url和uuid 包了
	private_key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		fmt.Println("Error at mercarigo//dPoP.go//dPoPGenerator//ecdsa.GenerateKey():\n", err)
		os.Exit(60)
	}

	pl := payload{int(time.Now().Unix()), uuid_, url_, strings.ToUpper(method)}
	pkjwk := pkey_jwk{"P-256", "EC", byteToBase64URL(private_key.PublicKey.X.Bytes()), byteToBase64URL(private_key.Y.Bytes())}
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

	tmp := append(headerString, "."...)
	tmp = append(tmp, payloadString...)
	data_unsigned := byteToBase64URL(tmp)

	h := sha256.New()
	h.Write([]byte(data_unsigned))
	hValue := h.Sum(nil)

	//signature, err := private_key.Sign(rand.Reader, data_unsigned, crypto.SHA256)
	r, s, err := ecdsa.Sign(rand.Reader, private_key, hValue)

	if err != nil {
		fmt.Println("Error at mercarigo//dPoP.go//dPoPGenerator//ecdsa.Sign():\n", err)
		os.Exit(63)
	}

	signatured := r.Bytes()
	signatured = append(signatured, s.Bytes()...)

	signaturedString := byteToBase64URL(signatured)

	result := fmt.Sprintf("%s.%s", data_unsigned, signaturedString)
	return result
}
