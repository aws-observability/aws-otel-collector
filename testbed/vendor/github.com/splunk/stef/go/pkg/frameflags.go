package pkg

type FrameFlags byte

const (
	// RestartDictionaries resets and restarts all dictionaries at frame beginning.
	RestartDictionaries FrameFlags = 1 << iota
	// RestartCompression resets and restarts the compression stream at frame beginning.
	RestartCompression
	// RestartCodecs resets the state of all encoders/decoders at frame beginning.
	RestartCodecs

	FrameFlagsMask = RestartDictionaries | RestartCompression | RestartCodecs
)
