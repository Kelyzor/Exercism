package luhn

import (
    "strings"
)

func Valid(id string) bool {
    
    id = strings.ReplaceAll(id, " ", "")

    if len(id) <= 1 {
        return false
    }

	sum := 0
    flag := 0

	for i := len(id) - 1; i >= 0; i-- {
        n := int(id[i]) - '0'
        if n > 9 {
            return false
        }
        if flag == 1 {
            n *= 2
            if n > 9 {
                sum += n - 9
            } else {
                sum += n
            }
            flag = 0
        } else {
			sum += n
            flag = 1
        }
    }
    
    return sum % 10 == 0
}
