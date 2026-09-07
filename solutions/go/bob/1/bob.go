// Package bob should have a package comment that summarizes what it's about.
package bob

import "strings"

// Hey should have a comment documenting it.
func Hey(remark string) string {
	remark = strings.TrimRight(remark, " \t\n\r")
    
    if len(remark) == 0 {
        return "Fine. Be that way!"
    }
    
    isSilence := true
    hasCapital := false
    isYelling := true
    isQuestion := false
    
    if remark[len(remark) - 1] == '?' {
        isQuestion = true
    }
    
    for _, r := range remark {
        if r != ' ' {
            isSilence = false
        }

        if r >= 'A' && r <= 'Z' {
            hasCapital = true
        }

        if r >= 'a' && r <= 'z' {
            isYelling = false
        }
    }
    
    if isSilence {
        return "Fine. Be that way!"
    }
    
    if isQuestion {
        if isYelling && hasCapital {
            return "Calm down, I know what I'm doing!"
        }
        return "Sure."
    }

    if isYelling && hasCapital {
        return "Whoa, chill out!"
    }

	return "Whatever."
}
