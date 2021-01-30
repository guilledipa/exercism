package raindrops

import "fmt"

func Convert(number int) string {
	var droplet string
	if number%3 == 0 {
		droplet += "Pling"
	}
	if number%5 == 0 {
		droplet += "Plang"
	}
	if number%7 == 0 {
		droplet += "Plong"
	}
	// Has not matched any of the above
	if droplet == "" {
		droplet = fmt.Sprint(number)
	}
	return droplet
}
