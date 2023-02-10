package bits

func Get(bs []byte) []byte {
	bits := make([]byte, 256)
	for i, b := range bs {
		ba := getBitArray(b)
		copy(bits[i:((i+1)*8-1)], ba[:])
	}
	return bits
}

// LSB
func getBitArray(b byte) [8]byte {
	bits := [8]byte{}

	j := 1
	for i := 0; i < 8; i++ {
		if (int(b) & (j << i)) > 0 {
			bits[i] = 1
		}
	}

	return bits
}
