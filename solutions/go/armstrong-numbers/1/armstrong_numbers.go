package armstrongnumbers

import "math"

func IsNumber(n int) bool {
	result := 0.0
    power := 0
	k := n

    for k != 0 {
        k /= 10
        power++
    }

    k = n

    for k != 0 {
        result += math.Pow(float64(k % 10), float64(power))
        k /= 10
    }

    return result == float64(n)
}
