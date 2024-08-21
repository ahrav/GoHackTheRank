package gohacktherank

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSockMerchant(t *testing.T) {
	tests := []struct {
		name     string
		n        int32
		ar       []int32
		expected int32
	}{
		{"AllPairs", 6, []int32{1, 2, 1, 2, 1, 2}, 3},
		{"NoPairs", 5, []int32{1, 2, 3, 4, 5}, 0},
		{"SomePairs", 7, []int32{1, 2, 1, 2, 1, 3, 2}, 2},
		{"EmptyArray", 0, []int32{}, 0},
		{"SingleElement", 1, []int32{1}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := SockMerchant(tt.n, tt.ar)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func BenchmarkSockMerchant(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		SockMerchant(7, []int32{1, 2, 1, 2, 1, 3, 2})
	}
}

func TestPageCount(t *testing.T) {
	tests := []struct {
		name     string
		n        int32
		p        int32
		expected int32
	}{
		{"FirstPage", 5, 1, 0},
		{"LastPage", 5, 5, 0},
		{"MiddlePage", 5, 3, 1},
		{"EvenPages", 6, 2, 1},
		{"OddPages", 7, 4, 1},
		{"SinglePageBook", 1, 1, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := PageCount(tt.n, tt.p)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func BenchmarkPageCount(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		PageCount(7, 4)
	}
}

func TestTowerBreakers(t *testing.T) {
	tests := []struct {
		name     string
		n        int32
		m        int32
		expected int32
	}{
		{"SingleTowerSingleHeight", 1, 1, 2},
		{"SingleTowerMultipleHeight", 1, 2, 1},
		{"MultipleTowersSingleHeight", 2, 1, 2},
		{"MultipleTowersEvenCount", 4, 3, 2},
		{"MultipleTowersOddCount", 3, 3, 1},
		{"LargeNumberOfTowers", 1000000, 1000000, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := TowerBreakers(tt.n, tt.m)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func BenchmarkTowerBreakers(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		TowerBreakers(1000000, 1000000)
	}
}

func TestCaesarCipher(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		k        int32
		expected string
	}{
		{"NoShift", "abc", 0, "abc"},
		{"ShiftByOne", "abc", 1, "bcd"},
		{"ShiftByTwentySix", "abc", 26, "abc"},
		{"ShiftWithWrapAround", "xyz", 3, "abc"},
		{"MixedCase", "AbC", 2, "CdE"},
		{"NonAlphabeticCharacters", "a!b.c", 1, "b!c.d"},
		{"LargeShift", "abc", 52, "abc"},
		{"NegativeShift", "abc", -1, "zab"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := CaesarCipher(tt.s, tt.k)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func BenchmarkCaesarCipher(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		CaesarCipher("The quick brown fox jumps over the lazy dog", 5)
	}
}
