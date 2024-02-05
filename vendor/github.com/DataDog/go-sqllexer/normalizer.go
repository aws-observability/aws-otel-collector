package sqllexer

import (
	"strings"
)

type normalizerConfig struct {
	// CollectTables specifies whether the normalizer should also extract the table names that a query addresses
	CollectTables bool

	// CollectCommands specifies whether the normalizer should extract and return commands as SQL metadata
	CollectCommands bool

	// CollectComments specifies whether the normalizer should extract and return comments as SQL metadata
	CollectComments bool

	// CollectProcedure specifies whether the normalizer should extract and return procedure name as SQL metadata
	CollectProcedure bool

	// KeepSQLAlias specifies whether SQL aliases ("AS") should be truncated.
	KeepSQLAlias bool

	// UppercaseKeywords specifies whether SQL keywords should be uppercased.
	UppercaseKeywords bool

	// RemoveSpaceBetweenParentheses specifies whether spaces should be kept between parentheses.
	// Spaces are inserted between parentheses by default. but this can be disabled by setting this to true.
	RemoveSpaceBetweenParentheses bool
}

type normalizerOption func(*normalizerConfig)

func WithCollectTables(collectTables bool) normalizerOption {
	return func(c *normalizerConfig) {
		c.CollectTables = collectTables
	}
}

func WithCollectCommands(collectCommands bool) normalizerOption {
	return func(c *normalizerConfig) {
		c.CollectCommands = collectCommands
	}
}

func WithCollectComments(collectComments bool) normalizerOption {
	return func(c *normalizerConfig) {
		c.CollectComments = collectComments
	}
}

func WithKeepSQLAlias(keepSQLAlias bool) normalizerOption {
	return func(c *normalizerConfig) {
		c.KeepSQLAlias = keepSQLAlias
	}
}

func WithUppercaseKeywords(uppercaseKeywords bool) normalizerOption {
	return func(c *normalizerConfig) {
		c.UppercaseKeywords = uppercaseKeywords
	}
}

func WithCollectProcedures(collectProcedure bool) normalizerOption {
	return func(c *normalizerConfig) {
		c.CollectProcedure = collectProcedure
	}
}

func WithRemoveSpaceBetweenParentheses(removeSpaceBetweenParentheses bool) normalizerOption {
	return func(c *normalizerConfig) {
		c.RemoveSpaceBetweenParentheses = removeSpaceBetweenParentheses
	}
}

type StatementMetadata struct {
	Size       int
	Tables     []string
	Comments   []string
	Commands   []string
	Procedures []string
}

type groupablePlaceholder struct {
	groupable bool
}

type Normalizer struct {
	config *normalizerConfig
}

func NewNormalizer(opts ...normalizerOption) *Normalizer {
	normalizer := Normalizer{
		config: &normalizerConfig{},
	}

	for _, opt := range opts {
		opt(normalizer.config)
	}

	return &normalizer
}

// Normalize takes an input SQL string and returns a normalized SQL string, a StatementMetadata struct, and an error.
// The normalizer collapses input SQL into compact format, groups obfuscated values into single placeholder,
// and collects metadata such as table names, comments, and commands.
func (n *Normalizer) Normalize(input string, lexerOpts ...lexerOption) (normalizedSQL string, statementMetadata *StatementMetadata, err error) {
	lexer := New(
		input,
		lexerOpts...,
	)

	var normalizedSQLBuilder strings.Builder

	statementMetadata = &StatementMetadata{
		Tables:     []string{},
		Comments:   []string{},
		Commands:   []string{},
		Procedures: []string{},
	}

	var lastToken Token // The last token that is not whitespace or comment
	var groupablePlaceholder groupablePlaceholder

	for {
		token := lexer.Scan()
		if token.Type == EOF {
			break
		}
		n.collectMetadata(&token, &lastToken, statementMetadata)
		n.normalizeSQL(&token, &lastToken, &normalizedSQLBuilder, &groupablePlaceholder, lexerOpts...)
	}

	normalizedSQL = normalizedSQLBuilder.String()

	// Dedupe collected metadata
	dedupeStatementMetadata(statementMetadata)

	return strings.TrimSpace(strings.TrimSuffix(normalizedSQL, ";")), statementMetadata, nil
}

func (n *Normalizer) collectMetadata(token *Token, lastToken *Token, statementMetadata *StatementMetadata) {
	if n.config.CollectComments && (token.Type == COMMENT || token.Type == MULTILINE_COMMENT) {
		// Collect comments
		statementMetadata.Comments = append(statementMetadata.Comments, token.Value)
	} else if token.Type == IDENT || token.Type == QUOTED_IDENT {
		tokenVal := token.Value
		if token.Type == QUOTED_IDENT {
			// remove all open and close quotes
			tokenVal = trimQuotes(tokenVal, tokenVal[0:1], tokenVal[len(tokenVal)-1:])
		}
		if n.config.CollectCommands && isCommand(strings.ToUpper(tokenVal)) {
			// Collect commands
			statementMetadata.Commands = append(statementMetadata.Commands, strings.ToUpper(tokenVal))
		} else if n.config.CollectTables && isTableIndicator(strings.ToUpper(lastToken.Value)) {
			// Collect table names
			statementMetadata.Tables = append(statementMetadata.Tables, tokenVal)
		} else if n.config.CollectProcedure && isProcedure(lastToken) {
			// Collect procedure names
			statementMetadata.Procedures = append(statementMetadata.Procedures, tokenVal)
		}
	}
}

func (n *Normalizer) normalizeSQL(token *Token, lastToken *Token, normalizedSQLBuilder *strings.Builder, groupablePlaceholder *groupablePlaceholder, lexerOpts ...lexerOption) {
	if token.Type != WS && token.Type != COMMENT && token.Type != MULTILINE_COMMENT {
		if token.Type == DOLLAR_QUOTED_FUNCTION && token.Value != StringPlaceholder {
			// if the token is a dollar quoted function and it is not obfuscated,
			// we need to recusively normalize the content of the dollar quoted function
			quotedFunc := token.Value[6 : len(token.Value)-6] // remove the $func$ prefix and suffix
			normalizedQuotedFunc, _, err := n.Normalize(quotedFunc, lexerOpts...)
			if err == nil {
				// replace the content of the dollar quoted function with the normalized content
				// if there is an error, we just keep the original content
				var normalizedDollarQuotedFunc strings.Builder
				normalizedDollarQuotedFunc.WriteString("$func$")
				normalizedDollarQuotedFunc.WriteString(normalizedQuotedFunc)
				normalizedDollarQuotedFunc.WriteString("$func$")
				token.Value = normalizedDollarQuotedFunc.String()
			}
		}

		if !n.config.KeepSQLAlias {
			// discard SQL alias
			if strings.ToUpper(token.Value) == "AS" {
				// if current token is AS, then continue to next token
				// because without seeing the next token, we cannot
				// determine if the current token is an alias or not
				*lastToken = *token
				return
			}

			if strings.ToUpper(lastToken.Value) == "AS" {
				if token.Type == IDENT && !isSQLKeyword(token) {
					// if the last token is AS and the current token is IDENT,
					// then the current token is an alias, so we discard it
					*lastToken = *token
					return
				} else {
					// if the last token is AS and the current token is not IDENT,
					// this could be a CTE like WITH ... AS (...),
					// so we do not discard the current token
					n.appendWhitespace(lastToken, token, normalizedSQLBuilder)
					n.writeToken(lastToken, normalizedSQLBuilder)
				}
			}
		}

		// group consecutive obfuscated values into single placeholder
		if n.isObfuscatedValueGroupable(token, lastToken, groupablePlaceholder) {
			// return the token but not write it to the normalizedSQLBuilder
			*lastToken = *token
			return
		}

		// determine if we should add a whitespace
		n.appendWhitespace(lastToken, token, normalizedSQLBuilder)
		n.writeToken(token, normalizedSQLBuilder)

		*lastToken = *token
	}
}

func (n *Normalizer) writeToken(token *Token, normalizedSQLBuilder *strings.Builder) {
	if n.config.UppercaseKeywords && isSQLKeyword(token) {
		normalizedSQLBuilder.WriteString(strings.ToUpper(token.Value))
	} else {
		normalizedSQLBuilder.WriteString(token.Value)
	}
}

func (n *Normalizer) isObfuscatedValueGroupable(token *Token, lastToken *Token, groupablePlaceholder *groupablePlaceholder) bool {
	if token.Value == NumberPlaceholder || token.Value == StringPlaceholder {
		if lastToken.Value == "(" || lastToken.Value == "[" {
			// if the last token is "(" or "[", and the current token is a placeholder,
			// we know it's the start of groupable placeholders
			// we don't return here because we still need to write the first placeholder
			groupablePlaceholder.groupable = true
		} else if lastToken.Value == "," && groupablePlaceholder.groupable {
			return true
		}
	}

	if (lastToken.Value == NumberPlaceholder || lastToken.Value == StringPlaceholder) && token.Value == "," && groupablePlaceholder.groupable {
		return true
	}

	if groupablePlaceholder.groupable && (token.Value == ")" || token.Value == "]") {
		// end of groupable placeholders
		groupablePlaceholder.groupable = false
	}

	return false
}

func (n *Normalizer) appendWhitespace(lastToken *Token, token *Token, normalizedSQLBuilder *strings.Builder) {
	// do not add a space between parentheses if RemoveSpaceBetweenParentheses is true
	if n.config.RemoveSpaceBetweenParentheses && (lastToken.Value == "(" || lastToken.Value == "[") {
		return
	}

	if n.config.RemoveSpaceBetweenParentheses && (token.Value == ")" || token.Value == "]") {
		return
	}

	switch token.Value {
	case ",":
	case "=":
		if lastToken.Value == ":" {
			// do not add a space before an equals if a colon was
			// present before it.
			break
		}
		fallthrough
	default:
		normalizedSQLBuilder.WriteString(" ")
	}
}

func dedupeCollectedMetadata(metadata []string) (dedupedMetadata []string, size int) {
	// Dedupe collected metadata
	// e.g. [SELECT, JOIN, SELECT, JOIN] -> [SELECT, JOIN]
	dedupedMetadata = []string{}
	var metadataSeen = make(map[string]struct{})
	for _, m := range metadata {
		if _, seen := metadataSeen[m]; !seen {
			metadataSeen[m] = struct{}{}
			dedupedMetadata = append(dedupedMetadata, m)
			size += len(m)
		}
	}
	return dedupedMetadata, size
}

func dedupeStatementMetadata(info *StatementMetadata) {
	var tablesSize, commentsSize, commandsSize, procedureSize int
	info.Tables, tablesSize = dedupeCollectedMetadata(info.Tables)
	info.Comments, commentsSize = dedupeCollectedMetadata(info.Comments)
	info.Commands, commandsSize = dedupeCollectedMetadata(info.Commands)
	info.Procedures, procedureSize = dedupeCollectedMetadata(info.Procedures)
	info.Size += tablesSize + commentsSize + commandsSize + procedureSize
}
