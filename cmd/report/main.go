package main

import (
	"fmt"

	"github.com/ZachL111/sonar-mob-contacts-core/policy"
)

func main() {
	signal := policy.Signal{Demand: 80, Capacity: 95, Latency: 12, Risk: 8, Weight: 7}
	fmt.Printf("score=%d decision=%s
", policy.Score(signal), policy.Classify(signal))
}
