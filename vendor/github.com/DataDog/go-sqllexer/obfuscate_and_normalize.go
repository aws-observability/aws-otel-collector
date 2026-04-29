package sqllexer

// ObfuscateAndNormalize takes an input SQL string and returns an normalized SQL string with metadata
// This function is a convenience function that combines the Obfuscator and Normalizer in one pass
func ObfuscateAndNormalize(input string, obfuscator *Obfuscator, normalizer *Normalizer, lexerOpts ...lexerOption) (normalizedSQL string, statementMetadata *StatementMetadata, err error) {
	obfuscate := func(token *Token, lastValueToken *LastValueToken) {
		obfuscator.ObfuscateTokenValue(token, lastValueToken, lexerOpts...)
	}
	return normalizer.normalize(input, obfuscate, lexerOpts...)
}
