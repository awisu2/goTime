package goTime

import (
	"testing"
)

func TestStrToTime(t *testing.T) {
	_, err := StrToTime("1234/12/23T12:34:56+09:00")
	if err != nil {
		t.Fatalf("convert err. %v", err)
	}

}