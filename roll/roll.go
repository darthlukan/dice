package roll

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func randomInt64(n int64) int64 {
	if n == 0 {
		n = 1
	}
	return rand.Int63n(n) + 1
}

func Roll(sides, qty int64) int64 {
	var val int64

	for i := int64(1); i <= qty; i++ {
		val += randomInt64(sides)
	}
	return val
}
