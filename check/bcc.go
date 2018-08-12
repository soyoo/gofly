// bcc verification algorithm implementation.
package check

func Bcc(data []byte) uint8 {
	var (
		bcc    uint8
		length int = len(data)
		index  int
	)

	for index = 0; index < length; index++ {
		bcc ^= data[index]
	}

	return bcc
}
