package jaconv

import (
	"fmt"
	"testing"
)

func TestToNippon(t *testing.T) {
	cases := []struct {
		hiragana string
		nippon   string
	}{
		{hiragana: "はくば47", nippon: "hakuba47"},
		{hiragana: "みょうこうすぎのはら", nippon: "myokosuginohara"},
		{hiragana: "たかすすのーぱーく", nippon: "takasusunopaku"},
		{hiragana: "GALAゆざわ", nippon: "GALAyuzawa"},
		{hiragana: "ごんどら", nippon: "gondora"},
		{hiragana: "ろっち", nippon: "rotti"},
		{hiragana: "こんび", nippon: "kombi"},
		{hiragana: "じゃりみち", nippon: "zyarimiti"},
	}
	for _, tt := range cases {
		act := ToNippon(tt.hiragana)
		if act != tt.nippon {
			t.Error(fmt.Sprintf("%s is not %s", act, tt.nippon))
		}
	}
}

func TestCharNipponByIndex(t *testing.T) {
	cases := []struct {
		s     string
		index int
		ch    CharNippon
	}{
		{s: "はくば47", index: 2, ch: CharNippon{Char: "ば", Nippon: "ba"}},
		{s: "はくば47", index: 3, ch: CharNippon{Char: "4", Nippon: ""}},
		{s: "はくば47", index: 4, ch: CharNippon{Char: "7", Nippon: ""}},
		{s: "みょうこうすぎのはら", index: 0, ch: CharNippon{Char: "みょ", Nippon: "myo"}},
		{s: "みょうこうすぎのはら", index: 1, ch: CharNippon{Char: "ょ", Nippon: ""}},
		{s: "みょうこうすぎのはら", index: 2, ch: CharNippon{Char: "う", Nippon: "u"}},
		{s: "たかすすのーぱーく", index: 4, ch: CharNippon{Char: "の", Nippon: "no"}},
		{s: "たかすすのーぱーく", index: 5, ch: CharNippon{Char: "ー", Nippon: ""}},
		{s: "たかすすのーぱーく", index: 6, ch: CharNippon{Char: "ぱ", Nippon: "pa"}},
		{s: "GALAゆざわ", index: 0, ch: CharNippon{Char: "G", Nippon: ""}},
		{s: "GALAゆざわ", index: 1, ch: CharNippon{Char: "A", Nippon: ""}},
		{s: "GALAゆざわ", index: 4, ch: CharNippon{Char: "ゆ", Nippon: "yu"}},
	}

	for _, tt := range cases {
		act := charNipponByIndex(tt.s, tt.index)
		if act != tt.ch {
			t.Error(fmt.Sprintf("%s is not %s", act, tt.ch))
		}
	}
}
