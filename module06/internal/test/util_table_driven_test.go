package test

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"udmey/module06/internal/util"
)

var tests = map[string]struct {
	input []int
	want  bool
}{

	"nil input":              {input: nil, want: false},
	"one value input":        {input: []int{1}, want: false},
	"has duplicate input":    {input: []int{1, 2, 3, 4, 1}, want: true},
	"has no duplicate input": {input: []int{1, 2, 3, 4}, want: false},
}

func TestContainsDuplicate(t *testing.T) {
	for name, testCase := range tests {
		t.Run(name, func(t *testing.T) {
			res := util.ContainsDuplicate(testCase.input)
			assert.Equal(t, testCase.want, res)
		})
	}
}
