package angrypurpletiger

import (
	"crypto/md5"
	"fmt"
)

/// Sum returns the animal-based hash digest of the data.
func Sum(data []byte) string {
	d := md5.Sum(data)
	c := compress(d[:])

	return fmt.Sprintf("%s-%s-%s", Adjectives[c[0]], Colors[c[1]], Animals[c[2]])
}

func compress(data []byte) [3]byte {
	if len(data) < 3 {
		panic(fmt.Sprintf("data length cannot be less than 3, got %d", len(data)))
	}

	size := len(data) / 3
	var segments [3]byte
	for i := range segments {
		// Define the indexes of data to be sliced for this segment.
		from, to := i*size, i*size+size

		// Ensure all the data is included in the final segment.
		if i == 2 {
			to += len(data) % size
		}

		// Build the segment.
		var segment byte
		for _, b := range data[from:to] {
			segment ^= b
		}
		segments[i] = segment
	}
	return segments
}
