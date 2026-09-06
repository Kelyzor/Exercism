package microblog

func Truncate(phrase string) string {
    const limit = 5
    count := 0
    output := ""
    for _, runeValue := range phrase {
        count++
        if count > limit {
            return output
        } else {
            output += string(runeValue)
        }
    }
    return output
}
