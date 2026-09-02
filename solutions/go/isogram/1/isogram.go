package isogram

import "strings"

func IsIsogram(word string) bool {
    word = strings.ToLower(word)
    word = strings.ReplaceAll(word, " ", "")
    word = strings.ReplaceAll(word, "-", "")
	chars := map[rune]int{}
    
	for _, v := range word {
        _, ex := chars[v]
        if ex {
            return false
        } else {
            chars[v] = 1
        }
    }

    return true
}
