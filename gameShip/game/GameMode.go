package game

// Mode denotes various game modes available
type Mode int

const (
	// EndlessMode - game played with no time limit to end
	EndlessMode Mode = iota
	// TimerMode - game played with fixed turn timings
	TimerMode
)
