package resistorcolor

import "fmt"

// Colors returns the list of all colors.
func Colors() []string {
	return []string{"black", "brown", "red", "orange", "yellow", "green", "blue", "violet", "grey", "white"}
}

// ColorCode returns the resistance value of the given color.
func ColorCode(color string) int {
    colors := Colors()
	for i := 1; i < len(colors); i++ {
        fmt.Sprint(colors)
        if colors[i] == color {
            return i
        }
    }
    return 0
}
