package anagram

import "strings"

func GetAlphabet() map[rune]int {
    alphabet := map[rune]int{}
    for i := 'a'; i <= 'z'; i++ {
        alphabet[i] = 0
    }
    return alphabet
}

func GetChars (word string) map[rune]int {
    word = strings.ToLower(word)
	alphabet := GetAlphabet()

    for _, r := range word {
        alphabet[r]++
    }

    return alphabet
}

func IsEqual(subject, candidate map[rune]int) bool {       
    for i := range subject {
        if subject[i] != candidate[i] {
            return false
        }
    }
    return true
}

func Detect(subject string, candidates []string) []string {
	result := []string{}
    
    subjectChars := GetChars(subject)
    for _, w := range candidates {
        candidateChars := GetChars(w)
        if IsEqual(subjectChars, candidateChars) && strings.ToLower(subject) != strings.ToLower(w) {
            result = append(result, w)
        }
    }
    
    return result
}
