package gohacktherank

import (
	"sort"
	"strings"
)

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

func CaesarCipher(s string, k int32) string {
	k = k % 26
	if k < 0 {
		k += 26
	}

	var sb strings.Builder
	for _, char := range s {
		switch {
		case char >= 'A' && char <= 'Z':
			sb.WriteRune('A' + (char-'A'+rune(k))%26)
		case char >= 'a' && char <= 'z':
			sb.WriteRune('a' + (char-'a'+rune(k))%26)
		default:
			sb.WriteRune(char)
		}
	}

	return sb.String()
}

func MaxMin(k int32, arr []int32) int32 {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	minV := arr[k-1] - arr[0]

	for i := k; i < int32(len(arr)); i++ {
		if m := arr[i] - arr[i-k+1]; m < minV {
			minV = m
		}
	}

	return minV
}
