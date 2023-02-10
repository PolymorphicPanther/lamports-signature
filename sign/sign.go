package sign

import (
	"LamportsSignature/bits"
	"LamportsSignature/keys"
	"crypto/sha256"
)

type Signature [256]*keys.Int256

func Sign(pk *keys.PrivateKey, m string) Signature {
	s256 := sha256.New()
	s256.Write(([]byte)(m))
	h := s256.Sum(nil)

	bs := bits.Get(h)

	sgn := pk.PickCorrespondingRowElement(bs)

	return sgn
}
