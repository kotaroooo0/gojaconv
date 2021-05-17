package jaconv

import (
	"fmt"
	"testing"
)

func TestToHebon(t *testing.T) {
	cases := []struct {
		hiragana string
		hebon    string
	}{
		{hiragana: "はくば47", hebon: "hakuba47"},
		{hiragana: "みょうこうすぎのはら", hebon: "myokosuginohara"},
		{hiragana: "たかすすのーぱーく", hebon: "takasusunopaku"},
		{hiragana: "GALAゆざわ", hebon: "GALAyuzawa"},
		{hiragana: "ごんどら", hebon: "gondora"},
		{hiragana: "ろっち", hebon: "rotchi"},
		{hiragana: "こんび", hebon: "kombi"},
		{hiragana: "じゃりみち", hebon: "jarimichi"},
	}
	for _, tt := range cases {
		act := ToHebon(tt.hiragana)
		if act != tt.hebon {
			t.Error(fmt.Sprintf("%s is not %s", act, tt.hebon))
		}
	}
}

func TestCharHebonByIndex(t *testing.T) {
	cases := []struct {
		s     string
		index int
		ch    CharHebon
	}{
		{s: "はくば47", index: 2, ch: CharHebon{Char: "ば", Hebon: "ba"}},
		{s: "はくば47", index: 3, ch: CharHebon{Char: "4", Hebon: ""}},
		{s: "はくば47", index: 4, ch: CharHebon{Char: "7", Hebon: ""}},
		{s: "みょうこうすぎのはら", index: 0, ch: CharHebon{Char: "みょ", Hebon: "myo"}},
		{s: "みょうこうすぎのはら", index: 1, ch: CharHebon{Char: "ょ", Hebon: ""}},
		{s: "みょうこうすぎのはら", index: 2, ch: CharHebon{Char: "う", Hebon: "u"}},
		{s: "たかすすのーぱーく", index: 4, ch: CharHebon{Char: "の", Hebon: "no"}},
		{s: "たかすすのーぱーく", index: 5, ch: CharHebon{Char: "ー", Hebon: ""}},
		{s: "たかすすのーぱーく", index: 6, ch: CharHebon{Char: "ぱ", Hebon: "pa"}},
		{s: "GALAゆざわ", index: 0, ch: CharHebon{Char: "G", Hebon: ""}},
		{s: "GALAゆざわ", index: 1, ch: CharHebon{Char: "A", Hebon: ""}},
		{s: "GALAゆざわ", index: 4, ch: CharHebon{Char: "ゆ", Hebon: "yu"}},
	}

	for _, tt := range cases {
		act := charHebonByIndex(tt.s, tt.index)
		if act != tt.ch {
			t.Error(fmt.Sprintf("%s is not %s", act, tt.ch))
		}
	}
}
