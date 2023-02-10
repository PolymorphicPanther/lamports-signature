package keys

type PublicKey key

func (p *PublicKey) PickCorrespondingRowElement(bitArray []byte) [256]*Int256 {
	return pickCorrespondingRowElement((*key)(p), bitArray)
}
