package challengemanager

type Mode uint8

const (
	// WatchTowerMode mode is the default mode for the challenge manager.
	// It will not trigger a challenge creation, but will agree if it agrees with assertions and log errors if it disagrees.
	WatchTowerMode Mode = iota
	// DefensiveMode mode will not post assertion, but will post and open challenges if it disagrees with any assertions.
	DefensiveMode
	// ResolveMode mode will not post assertion, but will stake on challenge leaves and support confirmation for vertices that it agrees with.
	ResolveMode
	// MakeMode mode will perform everything, ranging from posting assertions to staking to challenging and confirming.
	MakeMode
)
