package validator

type challengeWatcher struct {
	// Will keep track of ancestor histories for honest
	// branches per challenge.
	// Will scan for all previous events, and poll for new ones.
	// Will scan for level zero edges being confirmed and track
	// their claim id in this struct.
}

func newChallengeWatcher() *challengeWatcher {
	return &challengeWatcher{}
}
