package kmeans2

import (
	"math/rand"
	"time"
)

var (
	// seeded indicates if the random number generator has been seeded.
	seeded bool
	// seed is the seed used to seed the random number generator.
	seed int64
)

// seedRNG the random number generator.
func seedRNG() {
	if !seeded {
		seed = int64(time.Now().Nanosecond())
		rand.Seed(seed)
		seeded = true
	}
}
