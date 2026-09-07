// Package acronym should have a package comment that summarizes what it's about.
package acronym

import "strings"

// Abbreviate should have a comment documenting it.
func Abbreviate(s string) string {
    s = strings.ToUpper(s)
    
    output := string(s[0])
    
    for i := 1; i < len(s) - 1; i++ {
        if s[i] == ' ' && s[i + 1] != '-' && s[i + 1] != '_' || 
        	s[i] == '-' && s[i + 1] != ' ' || s[i] == '_' && s[i + 1] != ' ' {
            output += string(s[i + 1])
        }
    }

	return output
}
