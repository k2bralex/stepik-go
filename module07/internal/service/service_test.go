package service

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var tests = map[string]struct {
	input interface{}
	want  map[string]interface{}
}{
	"struct Alex": {input: struct {
		name string
		age  int
	}{name: "Alex", age: 22},
		want: map[string]interface{}{"name": "Alex", "age": 22}},
}

func TestName(t *testing.T) {
	for name, testCase := range tests {
		t.Run(name, func(t *testing.T) {
			res := StructToMap(testCase.input)
			assert.Equal(t, testCase.want, res)
		})
	}
}
