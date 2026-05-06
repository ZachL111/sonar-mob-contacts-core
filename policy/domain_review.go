package policy

type DomainReview struct {
	Signal     int
	Slack      int
	Drag       int
	Confidence int
}

func DomainReviewScore(item DomainReview) int {
	return item.Signal*2 + item.Slack + item.Confidence - item.Drag*3
}

func DomainReviewLane(item DomainReview) string {
	score := DomainReviewScore(item)
	if score >= 140 {
		return "ship"
	}
	if score >= 105 {
		return "watch"
	}
	return "hold"
}
