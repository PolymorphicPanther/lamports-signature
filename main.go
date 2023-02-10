package main

import (
	"LamportsSignature/keys"
	"LamportsSignature/sign"
	"LamportsSignature/verify"
	"fmt"
)

func main() {

	k, err := keys.GenerateKey()

	if err != nil {
		return
	}

	m := "Pip is an orphan, about seven years old, who lives with his hot-tempered older sister and her kindly blacksmith husband Joe Gargery on the coastal marshes of Kent."

	sgn := sign.Sign(k, m)

	valid := verify.Verify(k.GetPublicKey(), m, sgn)
	fmt.Println("Signature is valid: ", valid)
}
