package main

import (
	"crypto/rsa"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"messenger_server/pkg/rest"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
)

func main() {

	rest.Router.Run(":8080")

}

type JWKSet struct {
	Keys []json.RawMessage `json:"keys"`
}

// JWK represents a JSON Web Key (JWK).
type JWK struct {
	Kid string `json:"kid"`
	Kty string `json:"kty"`
	Alg string `json:"alg"`
	N   string `json:"n"`
	E   string `json:"e"`
	Use string `json:"use"`
}

func validateToken() {
	jwkSet, err := getJWKSet("")

	if err != nil {
		fmt.Printf("jwk get error : %s", err)
	}

	var tokenString string = "eyJhbGciOiJSUzI1NiIsImtpZCI6IjkxNDEzY2Y0ZmEwY2I5MmEzYzNmNWEwNTQ1MDkxMzJjNDc2NjA5MzciLCJ0eXAiOiJKV1QifQ.eyJpc3MiOiJodHRwczovL2FjY291bnRzLmdvb2dsZS5jb20iLCJhenAiOiI4MTM5NTQwMjk0NzItb3VtcmRoMHQwc25nMWx1YzdnaGhhZWF2YmFpMzM0cDIuYXBwcy5nb29nbGV1c2VyY29udGVudC5jb20iLCJhdWQiOiI4MTM5NTQwMjk0NzItMDA1NWw3NG5lbmZhNjRja2t2bjM2MzRhN21xYTFkYW8uYXBwcy5nb29nbGV1c2VyY29udGVudC5jb20iLCJzdWIiOiIxMTM3OTAxNjAwMDc0NzcxMzA1NzMiLCJlbWFpbCI6ImJ1Y2tndXNhdWRAZ21haWwuY29tIiwiZW1haWxfdmVyaWZpZWQiOnRydWUsIm5hbWUiOiLqsJXtmITrqoUiLCJwaWN0dXJlIjoiaHR0cHM6Ly9saDMuZ29vZ2xldXNlcmNvbnRlbnQuY29tL2EvQUNnOG9jTHd6d3daM0M4WTRZN3RYWllpT1lJMHlEcWhtMXZPYjNzMkVieUNSR0tyR3EwPXM5Ni1jIiwiZ2l2ZW5fbmFtZSI6Iu2YhOuqhSIsImZhbWlseV9uYW1lIjoi6rCVIiwibG9jYWxlIjoia28iLCJpYXQiOjE3MDQxOTY0NzMsImV4cCI6MTcwNDIwMDA3M30.fGlxkoHmwkCS_c8avaJN5AnVsRNAlHo38I_W2upT5X0a4A-DXSdgE3km05MgIaigAt_YpFVGyu2098zChIkXzE5OOxq66qdbVih60G7fWo-q99sWnJlyrRvCCVIb45TncB5gM3wunRtsHA7w5EgE7xwIYYCAkw1GZrVQT8oH8OsVBRv8PraRr510pXLNXZjic_8OaV0HIa2UPlBEKH-6ybu5qwvFKueT1nuW7WgC8mKg4fcFJyjY8GRloY1HV1br0lPNMfk81sXg7e8BAUn9ODzSD2VjlnLOb5LctVwajSd6gwb-hpzmSMER5M8ZaauFeYMmKLhOWgdHGP_LQ2NM-A"

	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		kid, ok := t.Header["kid"].(string)

		if !ok {
			return nil, fmt.Errorf("Key Id (kid) not found")
		}

		var selectKey *JWK

		for _, key := range jwkSet.Keys {
			var jwk JWK
			if err := json.Unmarshal(key, &jwk); err != nil {
				return nil, err
			}

			if jwk.Kid == kid {
				selectKey = &jwk
				break
			}
		}

		if selectKey == nil {
			return nil, fmt.Errorf("Key not Found")
		}

		dn, err := base64.RawURLEncoding.DecodeString(selectKey.N)

		if err != nil {
			return nil, err
		}
		de, err := base64.RawStdEncoding.DecodeString(selectKey.E)

		if err != nil {
			return nil, err
		}

		var publicKey *rsa.PublicKey = &rsa.PublicKey{
			N: new(big.Int).SetBytes(dn),
			E: int(new(big.Int).SetBytes(de).Int64()),
		}

		if err != nil {
			return nil, err
		}

		return publicKey, nil
	})

	if err != nil {
		fmt.Printf("error token vaildate : %s\n", err)
		return
	}

	fmt.Println(token)
}

func getJWKSet(url string) (*JWKSet, error) {
	response, err := http.Get("https://www.googleapis.com/oauth2/v3/certs")
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var jwkSet JWKSet
	if err := json.Unmarshal(body, &jwkSet); err != nil {
		return nil, err
	}

	return &jwkSet, nil
}
