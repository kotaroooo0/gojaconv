package jaconv

import (
	"unicode/utf8"

	"golang.org/x/exp/utf8string"
)

func ToNippon(kana string) string {
	isOmitted := map[string]bool{
		"aa": true, "ee": true, "ii": false, // i は連続しても省略しない
		"oo": true, "ou": true, "uu": true,
	}

	var nippon string
	var lastNippon string

	i := 0
	for {
		ch := charNipponByIndex(kana, i)
		if ch.Char == "っ" {
			// "っち"
			nextCh := charNipponByIndex(kana, i+1)
			if nextCh.Nippon != "" {
				ch.Nippon = "t"
			}
		} else if ch.Char == "ん" {
			// 後続の文字に無関係にnを用いる
			ch.Nippon = "n"
		} else if ch.Char == "ー" {
			// 長音は無視
			ch.Nippon = ""
		}

		if ch.Nippon != "" {
			// 変換できる文字の場合
			if lastNippon != "" {
				// 連続する母音の除去
				joinedNippon := lastNippon + ch.Nippon
				if len(joinedNippon) > 2 {
					joinedNippon = joinedNippon[len(joinedNippon)-2:]
				}
				if isOmitted[joinedNippon] {
					ch.Nippon = ""
				}
			}
			nippon += ch.Nippon
		} else {
			if ch.Char != "ー" {
				// 変換できない文字の場合
				nippon += ch.Char
			}
		}

		lastNippon = ch.Nippon
		i += utf8.RuneCountInString(ch.Char)
		if i >= utf8.RuneCountInString(kana) {
			break
		}
	}

	return nippon
}

type CharNippon struct {
	Char   string
	Nippon string
}

func charNipponByIndex(kana string, index int) CharNippon {
	nipponMap := map[string]string{
		"あ": "a", "い": "i", "う": "u", "え": "e", "お": "o",
		"か": "ka", "き": "ki", "く": "ku", "け": "ke", "こ": "ko",
		"さ": "sa", "し": "si", "す": "su", "せ": "se", "そ": "so",
		"た": "ta", "ち": "ti", "つ": "tu", "て": "te", "と": "to",
		"な": "na", "に": "ni", "ぬ": "nu", "ね": "ne", "の": "no",
		"は": "ha", "ひ": "hi", "ふ": "fu", "へ": "he", "ほ": "ho",
		"ま": "ma", "み": "mi", "む": "mu", "め": "me", "も": "mo",
		"や": "ya", "ゆ": "yu", "よ": "yo",
		"ら": "ra", "り": "ri", "る": "ru", "れ": "re", "ろ": "ro",
		"わ": "wa", "ゐ": "i", "ゑ": "e", "を": "o",
		"ぁ": "a", "ぃ": "i", "ぅ": "u", "ぇ": "e", "ぉ": "o",
		"が": "ga", "ぎ": "gi", "ぐ": "gu", "げ": "ge", "ご": "go",
		"ざ": "za", "じ": "ji", "ず": "zu", "ぜ": "ze", "ぞ": "zo",
		"だ": "da", "ぢ": "ji", "づ": "zu", "で": "de", "ど": "do",
		"ば": "ba", "び": "bi", "ぶ": "bu", "べ": "be", "ぼ": "bo",
		"ぱ": "pa", "ぴ": "pi", "ぷ": "pu", "ぺ": "pe", "ぽ": "po",
		"きゃ": "kya", "きゅ": "kyu", "きょ": "kyo",
		"しゃ": "sya", "しゅ": "syu", "しょ": "syo",
		"ちゃ": "tya", "ちゅ": "tyu", "ちょ": "tyo", "ちぇ": "tye",
		"にゃ": "nya", "にゅ": "nyu", "にょ": "nyo",
		"ひゃ": "hya", "ひゅ": "hyu", "ひょ": "hyo",
		"みゃ": "mya", "みゅ": "myu", "みょ": "myo",
		"りゃ": "rya", "りゅ": "ryu", "りょ": "ryo",
		"ぎゃ": "gya", "ぎゅ": "gyu", "ぎょ": "gyo",
		"じゃ": "zya", "じゅ": "zyu", "じょ": "zyo",
		"びゃ": "bya", "びゅ": "byu", "びょ": "byo",
		"ぴゃ": "pya", "ぴゅ": "pyu", "ぴょ": "pyo",
	}

	var nippon string
	var char string
	utfstr := utf8string.NewString(kana)
	// 2文字ヒットするとき
	if index+1 < utf8.RuneCountInString(kana) {
		char = utfstr.Slice(index, index+2)
		nippon = nipponMap[char]
	}
	// 2文字はヒットしないが1文字はヒットするとき
	if nippon == "" && index < utfstr.RuneCount() {
		char = utfstr.Slice(index, index+1)
		nippon = nipponMap[char]
	}
	return CharNippon{Char: char, Nippon: nippon}
}
