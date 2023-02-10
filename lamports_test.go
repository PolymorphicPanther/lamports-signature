package main

import (
	"LamportsSignature/keys"
	"LamportsSignature/sign"
	"LamportsSignature/verify"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSignatureModification(t *testing.T) {
	key, err := keys.GenerateKey()

	if err != nil {
		return
	}
	msg := "He has no running water."
	testCases := []struct {
		name         string
		message      string
		key          *keys.PrivateKey
		getSignature func(privateKey *keys.PrivateKey, message string) sign.Signature
		isValid      bool
	}{
		{
			name:    "Correct signature",
			message: msg,
			key:     key,
			getSignature: func(privateKey *keys.PrivateKey, message string) sign.Signature {
				return sign.Sign(privateKey, message)
			},
			isValid: true,
		},
		{
			name:    "Modified signature",
			message: msg,
			key:     key,
			getSignature: func(privateKey *keys.PrivateKey, message string) sign.Signature {
				sgn := sign.Sign(privateKey, message)

				flippedVal := 1 & ^sgn[5].Bit(200)
				sgn[5].SetBit(sgn[5].Int, 200, flippedVal) // Flip bit 200
				return sgn
			},
			isValid: false,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			sgn := test.getSignature(test.key, test.message)
			isValid := verify.Verify(test.key.GetPublicKey(), test.message, sgn)

			assert.Equal(t, test.isValid, isValid)
		})
	}

}

func TestModifyingPublicKey(t *testing.T) {
	key, err := keys.GenerateKey()
	if err != nil {
		return
	}
	msg := "That is training for the day."

	testCases := []struct {
		name         string
		message      string
		key          *keys.PrivateKey
		getPublicKey func(privateKey *keys.PrivateKey) (*keys.PublicKey, error)
		isValid      bool
	}{
		{
			name:    "Correct public key",
			message: msg,
			key:     key,
			getPublicKey: func(privateKey *keys.PrivateKey) (*keys.PublicKey, error) {
				return privateKey.GetPublicKey(), nil
			},
			isValid: true,
		}, {
			name:    "Modified public key",
			message: msg,
			key:     key,
			getPublicKey: func(privateKey *keys.PrivateKey) (*keys.PublicKey, error) {
				modifiedPublicKey := *key.GetPublicKey()

				// Note: bits must be flipped in both rows, in case the message changes
				flippedVal := 1 & ^modifiedPublicKey.Row0[5].Bit(10)
				modifiedPublicKey.Row0[5].SetBit(modifiedPublicKey.Row0[5].Int, 10, flippedVal)

				flippedVal = 1 & ^modifiedPublicKey.Row1[5].Bit(10)
				modifiedPublicKey.Row1[5].SetBit(modifiedPublicKey.Row1[5].Int, 10, flippedVal)

				return &modifiedPublicKey, nil
			},
			isValid: false,
		}, {
			name:    "Wrong public key",
			message: msg,
			key:     key,
			getPublicKey: func(*keys.PrivateKey) (*keys.PublicKey, error) {
				key, err := keys.GenerateKey()
				if err != nil {
					return nil, err
				}
				return key.GetPublicKey(), nil
			},
			isValid: false,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			sgn := sign.Sign(test.key, test.message)
			pubKey, err := test.getPublicKey(test.key)
			assert.Nil(t, err, "public key")

			isValid := verify.Verify(pubKey, test.message, sgn)

			assert.Equal(t, test.isValid, isValid)
		})
	}

}

func TestModifyingMessage(t *testing.T) {

	testCases := []struct {
		name              string
		messageToSign     string
		messageToValidate string
		isValid           bool
	}{
		{
			name:              "Matching message",
			messageToSign:     "There is a man in a cave.",
			messageToValidate: "There is a man in a cave.",
			isValid:           true,
		},
		{
			name:              "Different message",
			messageToSign:     "There is a man in a cave.",
			messageToValidate: "There is a man in a cavE.",
			isValid:           false,
		}, {
			name:              "Different message 1",
			messageToSign:     "There is a man in a cave.",
			messageToValidate: "There is a man in a cave",
			isValid:           false,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			key, err := keys.GenerateKey()
			if err != nil {
				t.Error(err)
				return
			}

			sgn := sign.Sign(key, test.messageToSign)

			isValid := verify.Verify(key.GetPublicKey(), test.messageToValidate, sgn)

			assert.Equal(t, test.isValid, isValid)
		})
	}

}
