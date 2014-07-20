package main

import (
	"github.com/dustin/randbo"
)

const (
	ReadableText = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
)

var (
	UsedTokens   = make(map[string]struct{})
	CryptoReader = randbo.New()
)

func GenerateRandomString(n int) string {
	for {
		bytes := make([]byte, n)
		CryptoReader.Read(bytes)
		for k, v := range bytes {
			bytes[k] = ReadableText[v%byte(len(ReadableText))]
		}

		str := string(bytes)
		if _, ok := UsedTokens[str]; !ok {
			UsedTokens[str] = struct{}{}
			return str
		}
	}
}
