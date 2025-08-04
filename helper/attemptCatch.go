package helper

import (
	"math/rand"
	"time"
)

var rng = rand.New(rand.NewSource(time.Now().UnixNano()))

func AttemptCatch(baseExp int) bool {
	chance := calculateCatchChance(baseExp)
	return rng.Float64() < chance
}

func calculateCatchChance(baseExp int) float64 {
	minExp := 50.0
	maxExp := 300.0

	if float64(baseExp) < minExp {
		baseExp = int(minExp)
	}
	if float64(baseExp) > maxExp {
		baseExp = int(maxExp)
	}

	normalized := (maxExp - float64(baseExp)) / (maxExp - minExp)

	minChance := 0.1
	maxChance := 0.9
	return minChance + normalized*(maxChance-minChance)
}
