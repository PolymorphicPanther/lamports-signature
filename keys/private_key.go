package keys

type PrivateKey struct {
	*key
	publicKey *PublicKey
}

func (p *PrivateKey) GetPublicKey() *PublicKey {
	return p.publicKey
}

func (p *PrivateKey) PickCorrespondingRowElement(bitArray []byte) [256]*Int256 {
	return pickCorrespondingRowElement(p.key, bitArray)
}
