package test

import (
	"testing"
	"udmey/module06/internal/util"
)

func BenchmarkContainsDuplicate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		util.ContainsDuplicate([]int{1, 2, 3, 4, 5, 6, 9, 9})
		util.ContainsDuplicate([]int{1, 2, 3, 4, 5, 6, 9})
	}
}
