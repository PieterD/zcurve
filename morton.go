package zcurve

func interleave(x, y uint32) uint64 {
	return sparse(x) | sparse(y)<<1
}

func sparse(dnum uint32) uint64 {
	num := uint64(dnum)
	num |= num << 16
	num &= 0x0000FFFF0000FFFF
	num |= num << 8
	num &= 0x00FF00FF00FF00FF
	num |= num << 4
	num &= 0x0F0F0F0F0F0F0F0F
	num |= num << 2
	num &= 0x3333333333333333
	num |= num << 1
	num &= 0x5555555555555555
	return num
}

func deinterleave(num uint64) (uint32, uint32) {
	return unsparse(num), unsparse(num >> 1)
}

func unsparse(num uint64) uint32 {
	num &= 0x5555555555555555
	num |= num >> 1
	num &= 0x3333333333333333
	num |= num >> 2
	num &= 0x0F0F0F0F0F0F0F0F
	num |= num >> 4
	num &= 0x00FF00FF00FF00FF
	num |= num >> 8
	num &= 0x0000FFFF0000FFFF
	num |= num >> 16
	num &= 0x00000000FFFFFFFF
	return uint32(num)
}
