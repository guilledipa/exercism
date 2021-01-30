package hamming

import "errors"

// Distance calculates the distance between two equal length DNA strands.
func Distance(a, b string) (int, error) {
	// Different length strands. Bail out.
	if len(a) != len(b) {
		return 0, errors.New("strands have different length, bailing out.")
	}
	var distance int
	// same strands
	if a == b {
		return distance, nil
	}
	// Calculate distance
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			distance++
		}
	}
	return distance, nil
}
