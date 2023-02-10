package verify

import (
	"LamportsSignature/bits"
	"LamportsSignature/keys"
	"LamportsSignature/sign"
	"crypto/sha256"
	"reflect"
)

func Verify(publicKey *keys.PublicKey, message string, signature sign.Signature) bool {
	s256 := sha256.New()
	s256.Write([]byte(message))
	hash := s256.Sum(nil)

	bs := bits.Get(hash)

	pubKeyBits := publicKey.PickCorrespondingRowElement(bs)

	sgnHash, _ := calculateRowHash(signature)

	return reflect.DeepEqual(sgnHash, pubKeyBits)
}

func calculateRowHash(input [256]*keys.Int256) ([256]*keys.Int256, error) {
	image := [256]*keys.Int256{}
	s256 := sha256.New()
	for i := 0; i < 256; i++ {
		s256.Write(input[i].Bytes())
		hash := s256.Sum(nil)
		s256.Reset()

		j, err := keys.CreateInt256(hash)
		if err != nil {
			return [256]*keys.Int256{}, err
		}

		image[i] = j
	}

	return image, nil
}
