package main

import "fmt"
import "./check"

func bcc_test() {
	b := []byte{0x01, 0xA0, 0x7C, 0xFF, 0x02}

	var bcc uint8
	bcc = check.Bcc(b)
	// fmt.Printf("bcc, %x\n", bcc)
	if bcc == 0x20 {
		fmt.Println("bcc test pass.")
	} else {
		fmt.Println("bcc test fail.")
	}
}

func cs_test() {
	b := []byte{0x01, 0xA0, 0x7C, 0xFF, 0x02}

	var cs uint8
	cs = check.Cs(b)
	// fmt.Printf("cs, %x\n", cs)
	if cs == 0x1e {
		fmt.Println("cs test pass.")
	} else {
		fmt.Println("cs test fail.")
	}
}

func main() {
	bcc_test()
	cs_test()
}
