package main

import (
	"fmt"
	"math"
	"strconv"
	"time"
)

func main() {
	var u uint8 = 255
	fmt.Println(u, u+1, u*u) // "255 0 1"

	var i int8 = 127
	fmt.Println(i, i+1, i*i) // "127 -128 1"

	var x uint8 = 1<<1 | 1<<5
	var y uint8 = 1<<1 | 1<<2

	fmt.Printf("%08b\n", x) // "00100010", the set {1, 5}
	fmt.Printf("%08b\n", y) // "00000110", the set {1, 2}

	fmt.Printf("%08b\n", x&y)  // "00000010", the intersection {1}
	fmt.Printf("%08b\n", x|y)  // "00100110", the union {1, 2, 5}
	fmt.Printf("%08b\n", x^y)  // "00100100", the symmetric difference {2, 5}
	fmt.Printf("%08b\n", x&^y) // "00100000", the difference {5}

	for i := uint(0); i < 8; i++ {
		if x&(1<<i) != 0 { // membership test
			fmt.Println(i) // "1", "5"
		}
	}

	fmt.Printf("%08b\n", x<<1) // "01000100", the set {2, 6}
	fmt.Printf("%08b\n", x>>1) // "00010001", the set {0, 4}

	f := 3.14
	j := int(f)
	fmt.Println(f, j)

	f = 1.99
	fmt.Println(int(f))

	o := 0666
	fmt.Printf("%d %[1]o %#[1]o\n", o) // "438 666 0666"
	z := int64(0xdeadbeef)
	fmt.Printf("%d %[1]x %#[1]x %#[1]X\n", z)

	ascii := 'a'
	unicode := '国'
	newline := '\n'
	fmt.Printf("%d %[1]c %[1]q\n", ascii)   // "97 a 'a'"
	fmt.Printf("%d %[1]c %[1]q\n", unicode) // "22269 国 '国'"
	fmt.Printf("%d %[1]q\n", newline)       // "10 '\n'"

	var f1 float32 = 167777216
	fmt.Println(f1 == f1+1)

	for x := 0; x < 8; x++ {
		fmt.Printf("x = %d e^x = %8.3f\n", x, math.Exp(float64(x)))
	}

	var z1 float64
	fmt.Println(z1, -z1, 1/z1, -1/z1, z1/z1) // "0 -0 +Inf -Inf NaN"

	nan := math.NaN()
	fmt.Println(nan == nan, nan < nan, nan > nan) // "false false false"

	var x2 complex128 = complex(1, 2) // 1+2i
	var y2 complex128 = complex(3, 4) // 3+4i
	fmt.Println(x2 * y2)              // "(-5+10i)"
	fmt.Println(real(x2 * y2))        // "-5"
	fmt.Println(imag(x2 * y2))        // "10"

	s := "hello world"
	fmt.Println(len(s))
	fmt.Println(s[0], s[7])
	fmt.Println(s[0:5])
	fmt.Println(s[:5])
	fmt.Println(s[7:])
	fmt.Println(s[:])
	fmt.Println("goodbye" + s[5:])

	for i, r := range "Hello, 世界" {
		fmt.Printf("%d\t%q\t%d\n", i, r, r)
	}

	// "program" in Japanese katakana
	s2 := "プログラム"
	fmt.Printf("% x\n", s2) // "e3 83 97 e3 83 ad e3 82 b0 e3 83 a9 e3 83 a0"
	r := []rune(s)
	fmt.Printf("%x\n", r) // "[30d7 30ed 30b0 30e9 30e0]"

	x3 := 123
	y3 := fmt.Sprintf("%d", x3)
	fmt.Println(y3, strconv.Itoa(x3)) // "123 123"

	fmt.Println(strconv.FormatInt(int64(x3), 2)) // "1111011"

	const noDelay time.Duration = 0
	const timeout = 5 * time.Minute
	fmt.Printf("%T %[1]v\n", noDelay)     // "time.Duration 0"
	fmt.Printf("%T %[1]v\n", timeout)     // "time.Duration 5m0s"
	fmt.Printf("%T %[1]v\n", time.Minute) // "time.Duration 1m0s"

	const (
		_   = 1 << (10 * iota)
		KiB // 1024
		MiB // 1048576
		GiB // 1073741824
		TiB // 1099511627776             (exceeds 1 << 32)
		PiB // 1125899906842624
		EiB // 1152921504606846976
		ZiB // 1180591620717411303424    (exceeds 1 << 64)
		YiB // 1208925819614629174706176
	)

	var f2 float64 = 212
	fmt.Println((f2 - 32) * 5 / 9)     // "100"; (f - 32) * 5 is a float64
	fmt.Println(5 / 9 * (f2 - 32))     // "0";   5/9 is an untyped integer, 0
	fmt.Println(5.0 / 9.0 * (f2 - 32)) // "100"; 5.0/9.0 is an untyped float

	fmt.Printf("%T\n", 0)      // "int"
	fmt.Printf("%T\n", 0.0)    // "float64"
	fmt.Printf("%T\n", 0i)     // "complex128"
	fmt.Printf("%T\n", '\000') // "int32" (rune)

}
