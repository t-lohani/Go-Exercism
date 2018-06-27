package space

import "math"

type Planet = string

func Age(secs float64, planet string) float64 {
	secsInYear := float64(31557600)
	earthYears := secs / secsInYear

	m := make(map[string]float64)

	m["Earth"] = 1.0
	m["Mercury"] = 0.2408467
	m["Venus"] = 0.61519726
	m["Mars"] = 1.8808158
	m["Jupiter"] = 11.862615
	m["Saturn"] = 29.447498
	m["Uranus"] = 84.016846
	m["Neptune"] = 164.79132

	planetYears := earthYears / m[planet]
	return math.Round(planetYears*100) / 100
}
