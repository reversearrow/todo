package auth

import (
	"fmt"
	"gopkg.in/square/go-jose.v2"
	"gopkg.in/square/go-jose.v2/jwt"
	"time"
)

type JWT struct {
	sessions []*jwt.JSONWebToken
}

func (j *JWT) Create(){
	key := []byte("secret")
	signingKey := jose.SigningKey{
		Algorithm: jose.HS512,
		Key: key,
	}

	signatureOption := (&jose.SignerOptions{}).WithType("JWT")

	signature, err := jose.NewSigner(signingKey, signatureOption)

	if err != nil{
		fmt.Println(err)
	}

	claims := jwt.Claims{
		IssuedAt: jwt.NewNumericDate(time.Now()),
		Expiry: jwt.NewNumericDate(time.Now().Add(time.Minute * 10)),
	}

	raw, err := jwt.Signed(signature).Claims(claims).CompactSerialize()

	if err != nil {
		panic(err)
	}

	fmt.Println(raw)
}