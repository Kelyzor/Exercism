package atbashcipher

import "strings"

func Atbash(s string) string {
    s = strings.ToLower(s)
    s = strings.ReplaceAll(s, " ", "")
    s = strings.ReplaceAll(s, ",", "")
    s = strings.ReplaceAll(s, ".", "")
    
    result := ""
    count := 0
    
    for _, r := range s {
        if count == 5 {
            result += " "
            count = 0
        }
    
        count++
        
        if r >= 'a' && r <= 'z' {
            result += string('z' - r + 'a')
        } else {
            result += string(r)
        }
    }
    
    return result
}
