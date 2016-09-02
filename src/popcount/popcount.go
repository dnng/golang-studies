// gopl.io/ch2/popcount
// Package popcount provides the number of bits whose value is 1 in an uint64
// value
package popcount

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func PopCountLoop(x uint64) int {
	var count int
	for i := 0; i < 8; i++ {
		count += int(byte(x>>(i*8)))
	}
	return count
}

func PopCountShift(x uint64) {
	var count int
	while x != 0 {
		x = x&(x-1)
		count++
	}
	return count
}

