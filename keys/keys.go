package keys

import (
	"crypto/rand"
	"crypto/sha256"
)

func GenerateKey() (*PrivateKey, error) {
	sk0, pk0, err := createRow()
	if err != nil {
		return nil, err
	}

	sk1, pk1, err := createRow()
	if err != nil {
		return nil, err
	}

	return &PrivateKey{
		key: &key{
			Row0: sk0,
			Row1: sk1,
		},
		publicKey: &PublicKey{
			Row0: pk0,
			Row1: pk1,
		},
	}, nil
}

type key struct {
	Row0 [256]*Int256
	Row1 [256]*Int256
}

func pickCorrespondingRowElement(inputKey *key, bits []byte) [256]*Int256 {

	sgn := [256]*Int256{}
	for i, b := range bits {
		if b == 0 {
			sgn[i] = inputKey.Row0[i]
		} else if b == 1 {
			sgn[i] = inputKey.Row1[i]
		} else {
			panic("bit is not 0 or 1")
		}
	}
	return sgn
}

func createRow() ([256]*Int256, [256]*Int256, error) {
	row := [256]*Int256{}
	image := [256]*Int256{}

	bs := make([]byte, 32)
	s256 := sha256.New()
	for i := 0; i < 256; i++ {
		n, err := rand.Read(bs)
		if n != len(bs) {
			return [256]*Int256{}, [256]*Int256{}, nil
		}

		if err != nil {
			return [256]*Int256{}, [256]*Int256{}, err
		}
		i256, err := CreateInt256(bs)
		if err != nil {
			return [256]*Int256{}, [256]*Int256{}, err
		}

		s256.Write(i256.Bytes())
		hash := s256.Sum(nil)
		s256.Reset()

		row[i] = i256
		j, err := CreateInt256(hash)
		if err != nil {
			return [256]*Int256{}, [256]*Int256{}, err
		}

		image[i] = j
	}

	return row, image, nil
}
