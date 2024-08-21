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
