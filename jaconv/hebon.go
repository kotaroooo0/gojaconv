package jaconv

import (
	"strings"
	"unicode/utf8"

	"golang.org/x/exp/utf8string"
)

func ToHebon(kana string) string {
	isOmitted := map[string]bool{
		"aa": true, "ee": true, "ii": false, // i は連続しても省略しない
		"oo": true, "ou": true, "uu": true,
	}

	var hebon string
	var lastHebon string

	i := 0
	for {
		ch := charHebonByIndex(kana, i)
		if ch.Char == "っ" {
			// "っち"
			nextCh := charHebonByIndex(kana, i+1)
			if nextCh.Hebon != "" {
				ch.Hebon = "t"
			}
		} else if ch.Char == "ん" {
			// B,M,P の前の "ん" は "M" とする。
			nextCh := charHebonByIndex(kana, i+1)
			if nextCh.Hebon != "" && strings.Index("bmp", nextCh.Hebon[0:1]) != -1 {
				ch.Hebon = "m"
			} else {
				ch.Hebon = "n"
			}
		} else if ch.Char == "ー" {
			// 長音は無視
			ch.Hebon = ""
		}

		if ch.Hebon != "" {
			// 変換できる文字の場合
			if lastHebon != "" {
				// 連続する母音の除去
				joinedHebon := lastHebon + ch.Hebon
				if len(joinedHebon) > 2 {
					joinedHebon = joinedHebon[len(joinedHebon)-2:]
				}
				if isOmitted[joinedHebon] {
					ch.Hebon = ""
				}
			}
			hebon += ch.Hebon
		} else {
			if ch.Char != "ー" {
				// 変換できない文字の場合
				hebon += ch.Char
			}
		}

		lastHebon = ch.Hebon
		i += utf8.RuneCountInString(ch.Char)
		if i >= utf8.RuneCountInString(kana) {
			break
		}
	}

	return hebon
}

type CharHebon struct {
	Char  string
	Hebon string
}

func charHebonByIndex(kana string, index int) CharHebon {
	hebonMap := map[string]string{
		"あ": "a", "い": "i", "う": "u", "え": "e", "お": "o",
		"か": "ka", "き": "ki", "く": "ku", "け": "ke", "こ": "ko",
		"さ": "sa", "し": "shi", "す": "su", "せ": "se", "そ": "so",
		"た": "ta", "ち": "chi", "つ": "tsu", "て": "te", "と": "to",
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
		"しゃ": "sha", "しゅ": "shu", "しょ": "sho",
		"ちゃ": "cha", "ちゅ": "chu", "ちょ": "cho", "ちぇ": "che",
		"にゃ": "nya", "にゅ": "nyu", "にょ": "nyo",
		"ひゃ": "hya", "ひゅ": "hyu", "ひょ": "hyo",
		"みゃ": "mya", "みゅ": "myu", "みょ": "myo",
		"りゃ": "rya", "りゅ": "ryu", "りょ": "ryo",
		"ぎゃ": "gya", "ぎゅ": "gyu", "ぎょ": "gyo",
		"じゃ": "ja", "じゅ": "ju", "じょ": "jo",
		"びゃ": "bya", "びゅ": "byu", "びょ": "byo",
		"ぴゃ": "pya", "ぴゅ": "pyu", "ぴょ": "pyo",
	}

	var hebon string
	var char string
	utfstr := utf8string.NewString(kana)
	// 2文字ヒットするとき
	if index+1 < utf8.RuneCountInString(kana) {
		char = utfstr.Slice(index, index+2)
		hebon = hebonMap[char]
	}
	// 2文字はヒットしないが1文字はヒットするとき
	if hebon == "" && index < utfstr.RuneCount() {
		char = utfstr.Slice(index, index+1)
		hebon = hebonMap[char]
	}
	return CharHebon{Char: char, Hebon: hebon}
}
