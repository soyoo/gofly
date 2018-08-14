// LRC verification algorithm implementation
//     纵向冗余校验（Longitudinal Redundancy Check，简称：LRC）是通信中常用的一种校验形式，也称LRC校验或纵向校验。
// 它是一种从纵向通道上的特定比特串产生校验比特的错误检测方法。在行列格式中（如磁带），LRC经常是与VRC一起使用，
// 这样就会为每个字符校验码。在工业领域Modbus协议Ascii模式采用该算法。
// 具体算法如下：
// 1、对需要校验的数据（2n个字符）两两组成一个16进制的数值求和。
// 2、将求和结果与256求模。
// 3、用256减去所得模值得到校验结果（另一种方法：将模值按位取反然后加1）。
// 例如16进制数据：01 A0 7C FF 02
// （16进制计算）　　求和：01 + A0 + 7C + FF + 02 = 21E　　取模：21E % 100 = 1E　　计算：100 - 1E = E2
// （10进制计算）　　求和：01 + 160 + 124 + 255 + 02 = 542　　取模：542 % 256 = 30　　计算：256 - 30 = 226
package check

func Lrc(data []byte) uint8 {
	var (
		sum    uint16
		length int = len(data)
		index  int
	)

	for index = 0; index < length; index++ {
		sum += (uint16)(data[index])
	}

	return (uint8)(256 - (sum % 256))
}
