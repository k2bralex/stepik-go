package internal

import (
	"testing"
)

func TestPositiveInt(t *testing.T) {
	token, err := ReverseNumber(1212)
	if err != nil {
		t.Error(err.Error())
	}
	if token != 2121 {
		t.Errorf("invalid token: %d", 2121)
	}
}

func TestNegativeInt(t *testing.T) {
	token, err := ReverseNumber(-1212)
	if err != nil {
		t.Error(err.Error())
	}
	if token != -2121 {
		t.Errorf("invalid token: %d", -2121)
	}
}

func TestZeroValue(t *testing.T) {
	token, err := ReverseNumber(0)
	if err != nil {
		t.Error(err.Error())
	}
	if token != 0 {
		t.Errorf("invalid token: %d", 0)
	}
}

func TestString(t *testing.T) {
	token, err := ReverseNumber("1212")
	if err != nil {
		t.Error(err.Error())
	}
	if token != "2121" {
		t.Errorf("invalid token: %s", "2121")
	}
}

func TestInvalidType(t *testing.T) {
	_, err := ReverseNumber(12.22)
	if err == nil {
		t.Error(err.Error())
	}
}
