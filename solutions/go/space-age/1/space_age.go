package spaceage

type Planet string

func Age(seconds float64, planet Planet) float64 {
    years := seconds / 31557600
	switch planet {
        case "Mercury":
        	return years / 0.2408467
        case "Venus":
        	return years / 0.61519726
        case "Earth":
        	return years
        case "Mars":
        	return years / 1.8808158
        case "Jupiter":
        	return years / 11.862615
        case "Saturn":
        	return years / 29.447498
        case "Uranus":
        	return years / 84.016846
        case "Neptune":
        	return years / 164.79132
    }
    return -1
}
