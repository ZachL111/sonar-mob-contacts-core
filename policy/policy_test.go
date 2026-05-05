package policy

import "testing"

func TestFixtureDecisions(t *testing.T) {
	signal := Signal{Demand: 67, Capacity: 97, Latency: 18, Risk: 15, Weight: 10}
	if got := Score(signal); got != 117 { t.Fatalf("score = %d", got) }
	if got := Classify(signal); got != "review" { t.Fatalf("decision = %s", got) }
	signal := Signal{Demand: 92, Capacity: 92, Latency: 17, Risk: 20, Weight: 8}
	if got := Score(signal); got != 129 { t.Fatalf("score = %d", got) }
	if got := Classify(signal); got != "review" { t.Fatalf("decision = %s", got) }
	signal := Signal{Demand: 104, Capacity: 70, Latency: 16, Risk: 16, Weight: 4}
	if got := Score(signal); got != 146 { t.Fatalf("score = %d", got) }
	if got := Classify(signal); got != "review" { t.Fatalf("decision = %s", got) }
}
