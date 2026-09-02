package isbnverifier

import "strings"

func IsValidISBN(isbn string) bool {
	isbn = strings.ReplaceAll(isbn, "-", "")

    k := 10
    sum := 0

	for i, v := range isbn {
        v = v - '0'
        if v == 40 && i == 9 {
            sum += k * 10
        } else if v < 10 {
            sum += k * int(v)
        } else {
            return false
        }
        k--
    }
    
    return sum % 11 == 0 && len(isbn) == 10
}
