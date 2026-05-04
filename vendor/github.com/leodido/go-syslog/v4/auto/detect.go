package auto

// detect inspects the bytes after the PRI closing bracket to determine
// whether the input is RFC 5424 or RFC 3164. It returns the detected format.
//
// The function examines at most ~6 bytes and runs in O(1) time.
// Decision table:
//   - Letter, space, '*'     → RFC 3164 (month name, leading space, Cisco star)
//   - '0'                    → RFC 3164 (zero-padded or Cisco counter)
//   - 4+ digits              → RFC 3164 (year prefix / long Cisco counter)
//   - 1-3 digits + space     → RFC 5424 (VERSION SP)
//   - 1-3 digits + ':'       → RFC 3164 (Cisco counter)
//   - Default                → RFC 5424
func detect(input []byte) Format {
	// Find the closing '>' of PRI.
	pos := 0
	for pos < len(input) && input[pos] != '>' {
		pos++
	}
	if pos >= len(input) {
		// No '>' found — default to RFC 5424 (will fail with PRI error).
		return FormatRFC5424
	}

	// Move past '>'.
	pos++
	if pos >= len(input) {
		// Nothing after '>' — default to RFC 5424 (will fail with version error).
		return FormatRFC5424
	}

	b := input[pos]

	// Letter, space, or '*' → RFC 3164.
	if (b >= 'A' && b <= 'Z') || (b >= 'a' && b <= 'z') || b == ' ' || b == '*' {
		return FormatRFC3164
	}

	// '0' → RFC 3164 (RFC 5424 version cannot start with 0).
	if b == '0' {
		return FormatRFC3164
	}

	// Digit 1-9: scan forward to distinguish VERSION (5424) from Cisco counter (3164).
	if b >= '1' && b <= '9' {
		digitCount := 1
		pos++
		for pos < len(input) && input[pos] >= '0' && input[pos] <= '9' {
			digitCount++
			pos++
		}

		// VERSION is 1-3 digits. 4+ digits → Cisco message counter → RFC 3164.
		if digitCount > 3 {
			return FormatRFC3164
		}

		// Check the byte after the digit run.
		if pos < len(input) {
			if input[pos] == ' ' {
				return FormatRFC5424
			}
			if input[pos] == ':' {
				return FormatRFC3164
			}
		}

		// Ambiguous or truncated — default to RFC 5424.
		return FormatRFC5424
	}

	// Any other byte — default to RFC 5424 (stricter grammar gives clearer errors).
	return FormatRFC5424
}
