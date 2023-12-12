package challengemanager

// Checks if there are active, honest validators involved onchain.
func (m *Manager) checkHonestValidatorActivity() {
	// - Identify the honest party by its history commitment in a challenge
	// - Make sure it is the unrivaled and making txs. Check if its edges are being bisected or if subchallenges are being opened on its edges
	// - Show the time between challenge moves
	// - **If an honest edge is rivaled and is not being bisected for some X period of time, the honest validator might be down**
	// How can we do this?

	// First, identify the challenge
}
