// bcc verification algorithm implementation.
//     BCC(Block Check Character/信息组校验码)，因校验码是将所有数据异或得出，故俗称异或校验。
// 具体算法是：将每一个字节的数据（一般是两个16进制的字符）进行异或后即得到校验码。 
// 例如16进制数据：01 A0 7C FF 02 
// 计算：01 xor A0 xor 7C xor FF xor 02 = 20 
// 校验码是：20 
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
