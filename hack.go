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

func PageCount(n, p int32) int32 {
	front := p / 2
	back := (n / 2) - front
	if front < back {
		return front
	}
	return back
}

func TowerBreakers(n, m int32) int32 {
	if m == 1 {
		return 2
	}

	if n%2 == 0 {
		return 2
	}
	return 1
}
