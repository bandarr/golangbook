package popcount

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint64) int {

	var blarg int

	for i := 0; i < 64; i++ {
		blarg += int(pc[byte(x>>(i*8))])
	}

	return blarg
}