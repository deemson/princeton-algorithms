package hash

func String(value string) uint32 {
	hash := int32(7)
	for _, character := range value {
		hash = 31*hash + character
	}
	return uint32(hash)
}
