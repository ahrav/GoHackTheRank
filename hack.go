package gohacktherank

func SockMerchant(n int32, ar []int32) int32 {
	pairs := int32(0)
	socks := make(map[int32]int32, n)
	for _, sock := range ar {
		socks[sock]++
		if socks[sock]%2 == 0 {
			pairs++
		}
	}

	return pairs
}
