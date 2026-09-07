package bottlesong

import (
    "fmt";
	"strings"
)

func NtoS(n int) string {
    switch n {
        case 10:
        	return "Ten"
        case 9:
        	return "Nine"
        case 8:
        	return "Eight"
        case 7:
        	return "Seven"
        case 6:
        	return "Six"
        case 5:
        	return "Five"
        case 4:
        	return "Four"
        case 3:
        	return "Three"
        case 2:
        	return "Two"
        case 1:
        	return "One"
        default:
        	return "No"
    }
}

func OnTheWall(bottles int) string {
    if bottles == 1 {
        return fmt.Sprintf("One green bottle hanging on the wall,")
    }
    bottlesS := NtoS(bottles)
    return fmt.Sprintf("%s green bottles hanging on the wall,", bottlesS)
}

func Drop() string {
    return fmt.Sprintf("And if one green bottle should accidentally fall,")
} 

func Left(bottles int) string {
    if bottles == 1 {
        return fmt.Sprintf("There'll be one green bottle hanging on the wall.")
    }
    bottlesS := strings.ToLower(NtoS(bottles))
    return fmt.Sprintf("There'll be %s green bottles hanging on the wall.", bottlesS)
} 

func Recite(startBottles, takeDown int) []string {
    output := []string{}
	for takeDown > 0 {
        output = append(output, OnTheWall(startBottles), OnTheWall(startBottles), Drop())
        startBottles--
        output = append(output, Left(startBottles))
        takeDown--
        if takeDown > 0 {
            output = append(output, "")
        }
    }
    return output
}
