package main

import "fmt"
import "./check"

func bcc_test() {
	b := []byte{0x01, 0xA0, 0x7C, 0xFF, 0x02}

	var bcc uint8
	bcc = check.Bcc(b)
	fmt.Printf("bcc, %x\n", bcc)
}

func cs_test() {
	b := []byte{0x01, 0xA0, 0x7C, 0xFF, 0x02}

	var cs uint8
	cs = check.Cs(b)
	fmt.Printf("cs, %x\n", cs)
}

func main() {
	bcc_test()
	cs_test()
}
