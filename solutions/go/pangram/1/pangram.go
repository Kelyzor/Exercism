package pangram

import "strings"


func IsPangram(input string) bool {
	input = strings.ToLower(input)
    
    alphabet := map[rune]int{}
    
    for i := 'a'; i <= 'z'; i++ {
        alphabet[i] = 0
    }
    
    for _, r := range input {
        _, e := alphabet[r]
        if e {
            alphabet[r]++
        } else {
            continue
        }
    }

	for i := 'a'; i <= 'z'; i++ {
        if alphabet[i] == 0 {
            return false
        }
    }
    
    return true
}
