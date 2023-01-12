package utils

import (
	"math/rand"
	"unsafe"
)

func RevertStringBytes(s string) {
	buf := *(*[]byte)(unsafe.Pointer(&struct {
		string
		Cap int
	}{s, len(s)}))
	for i := 0; i < len(buf)/2; i++ {
		buf[i], buf[len(buf)-i-1] = buf[len(buf)-i-1], buf[i]
	}
}

func RevertString(s string) string {
	buf := []rune(s)
	for i := 0; i < len(buf)/2; i++ {
		buf[i], buf[len(buf)-i-1] = buf[len(buf)-i-1], buf[i]
	}
	return string(buf)
}

func IntSlicesAreEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func Abs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}

func Min(a int, b ...int) int {
	min := a
	for _, num := range b {
		if num < min {
			min = num
		}
	}
	return min
}

func Max(a int, b ...int) int {
	max := a
	for _, num := range b {
		if num > max {
			max = num
		}
	}
	return max
}

func GenRandomSlice(length, limit int) []int {
	s := make([]int, length)
	for i := range s {
		s[i] = rand.Intn(limit)
	}
	return s
}

func IsLetterOrDigit(b byte) bool {
	if b >= 'a' && b <= 'z' ||
		b >= 'A' && b <= 'Z' ||
		b >= '0' && b <= '9' {
		return true
	}
	return false
}

func ToLowerCase(b byte) byte {
	if b >= 'A' && b <= 'Z' {
		b += 32
	}
	return b
}
