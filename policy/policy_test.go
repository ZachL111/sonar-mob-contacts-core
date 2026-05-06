package policy

import "testing"

func TestFixtureDecisions(t *testing.T) {
	tests := []struct {
		name         string
		signal       Signal
		wantScore    int
		wantDecision string
	}{
		{name: "case_1", signal: Signal{Demand: 67, Capacity: 97, Latency: 18, Risk: 15, Weight: 10}, wantScore: 117, wantDecision: "review"},
		{name: "case_2", signal: Signal{Demand: 92, Capacity: 92, Latency: 17, Risk: 20, Weight: 8}, wantScore: 129, wantDecision: "review"},
		{name: "case_3", signal: Signal{Demand: 104, Capacity: 70, Latency: 16, Risk: 16, Weight: 4}, wantScore: 146, wantDecision: "review"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := Score(tc.signal); got != tc.wantScore {
				t.Fatalf("score = %d, want %d", got, tc.wantScore)
			}
			if got := Classify(tc.signal); got != tc.wantDecision {
				t.Fatalf("decision = %s, want %s", got, tc.wantDecision)
			}
		})
	}
}
