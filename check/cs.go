// cs verification algorithm implementation.
package check

func Cs(data []byte) uint8 {
	var (
		cs     uint8
		length int = len(data)
		index  int
	)

	for index = 0; index < length; index++ {
		cs += data[index]
	}

	return cs
}
