package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// func TestSquare1(t *testing.T) {
// 	var rst int = square(9)
// 	// rst := square(9)
// 	if rst != 81 {
// 		t.Errorf("error %d", rst)
// 	}
// }

// func TestSqure2(t *testing.T) {
// 	rst := square(3)
// 	if rst != 9 {
// 		t.Errorf("square(3) should be 9 but aquare(3) returns %d", rst)
// // 	}
// // }

// // func TestSquare1(t *testing.T) {
// // 	assert := assert.New(t)
// // 	assert.Equal(81, square(9), "square(9) should be 81")
// // }

// // func TestSquare2(t *testing.T) {
// // 	assert := assert.New(t)
// // 	assert.Equal(9, square(3), "sqaure(3) should be 9")
// // }

// func TestFibonacci1(t *testing.T) {
// 	assert := assert.New(t)

// 	assert.Equal(0, fibonacci1(-1), "fibonacci1(-1) should be 0")
// 	assert.Equal(0, fibonacci1(0), "fibonacci1(0) should be 0")
// 	assert.Equal(1, fibonacci1(1), "fibonacci1(1) should be 0")
// 	assert.Equal(2, fibonacci1(3), "fibonacci1(3) should be 0")
// 	assert.Equal(233, fibonacci1(13), "fibonacci1(13) should be 0")

// }

// func TestFibonacci2(t *testing.T) {
// 	assert := assert.New(t)

// 	assert.Equal(0, fibonacci2(-1), "fibonacci1(-1) should be 0")
// 	assert.Equal(0, fibonacci2(0), "fibonacci1(0) should be 0")
// 	assert.Equal(1, fibonacci2(1), "fibonacci1(1) should be 0")
// 	assert.Equal(2, fibonacci2(3), "fibonacci1(3) should be 0")
// 	assert.Equal(233, fibonacci2(13), "fibonacci1(13) should be 0")
// }

// func BenchmarkFibonacci1(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		fibonacci1(13)
// 	}
// }

// func BenchmarkFibonacci2(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		fibonacci2(13)
// 	}
// }

func TestAtoi(t *testing.T) {
	assert := assert.New(t)

	rst, err := Atoi("123")
	assert.Nil(err)
	assert.Equal(123, rst)

	rst, err = Atoi("  123")
	assert.Nil(err)
	assert.Equal(123, rst)

	rst, err = Atoi("  -123")
	assert.Nil(err)
	assert.Equal(-123, rst)

	rst, err = Atoi("  -123  ")
	assert.Nil(err)
	assert.Equal(-123, rst)

	rst, err = Atoi("  -123a")
	assert.Nil(err)
	assert.Equal(0, rst)

	rst, err = Atoi("  -123 4")
	assert.NotNil(err)
	assert.Equal(0, rst)
}
