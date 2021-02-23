package jaconv

import (
	"fmt"
	"testing"
)

func TestKatakanaToHiragana(t *testing.T) {
	cases := []struct {
		katakana string
		expected string
	}{
		{katakana: "", expected: ""},
		{katakana: "アカサタナ", expected: "あかさたな"},
		{katakana: "ワボン", expected: "わぼん"},
	}
	for _, tt := range cases {
		act := KatakanaToHiragana(tt.katakana)
		if act != tt.expected {
			t.Error(fmt.Sprintf("%s is not %s", act, tt.expected))
		}
	}
}
