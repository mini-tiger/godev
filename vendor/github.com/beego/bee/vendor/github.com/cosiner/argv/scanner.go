package argv

import "unicode"

// Scanner is a cmdline string scanner.
//
// It split cmdline string to tokens: space, string, pipe, reverse quote string.
type Scanner struct {
	env map[string]string

	text      []rune
	rpos      int
	dollarBuf []rune
}

// NewScanner create a scanner and init it's internal states.
func NewScanner(text []rune, env map[string]string) *Scanner {
	return &Scanner{
		text: text,
		env:  env,
	}
}

func (s *Scanner) envs() map[string]string {
	return s.env
}

const _RuneEOF = 0

func (s *Scanner) nextRune() rune {
	if s.rpos >= len(s.text) {
		return _RuneEOF
	}

	r := s.text[s.rpos]
	s.rpos++
	return r
}

func (s *Scanner) unreadRune(r rune) {
	if r != _RuneEOF {
		s.rpos--
	}
}

func (s *Scanner) isEscapeChars(r rune) (rune, bool) {
	switch r {
	case 'a':
		return '\a', true
	case 'b':
		return '\b', true
	case 'f':
		return '\f', true
	case 'n':
		return '\n', true
	case 'r':
		return '\r', true
	case 't':
		return '\t', true
	case 'v':
		return '\v', true
	case '\\':
		return '\\', true
	case '$':
		return '$', true
	}
	return r, false
}

func (s *Scanner) endEnv(r rune) bool {
	if r == '_' || (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') {
		return false
	}
	return true
}

// TokenType is the type of tokens recognized by the scanner.
type TokenType uint32

// Token is generated by the scanner with a type and value.
type Token struct {
	Type  TokenType
	Value []rune
}

const (
	// TokString for string, single quoted string and double quoted string
	TokString TokenType = iota + 1
	// TokPipe is the '|' character
	TokPipe
	// TokReversequote is reverse quoted string
	TokReversequote
	// TokSpace represent space character sequence
	TokSpace
	// TokEOF means the input end.
	TokEOF
)

func (s *Scanner) getEnv(name string) string {
	return s.env[name]
}

func (s *Scanner) specialVar(r rune) (string, bool) {
	switch r {
	case '0', '*', '#', '@', '?', '$':
		v, has := s.env[string(r)]
		return v, has
	default:
		return "", false
	}
}

func (s *Scanner) checkDollarStart(tok *Token, r rune, from, switchTo uint8) uint8 {
	state := from
	nr := s.nextRune()
	if val, has := s.specialVar(nr); has {
		if val != "" {
			tok.Value = append(tok.Value, []rune(val)...)
		}
	} else if s.endEnv(nr) {
		tok.Value = append(tok.Value, r)
		s.unreadRune(nr)
	} else {
		state = switchTo
		s.dollarBuf = append(s.dollarBuf[:0], nr)
	}
	return state
}

func (s *Scanner) checkDollarEnd(tok *Token, r rune, from, switchTo uint8) uint8 {
	var state = from
	if s.endEnv(r) {
		tok.Value = append(tok.Value, []rune(s.getEnv(string(s.dollarBuf)))...)
		state = switchTo
		s.unreadRune(r)
	} else {
		s.dollarBuf = append(s.dollarBuf, r)
	}
	return state
}

// Next return next token, if it reach the end, TOK_EOF will be returned.
//
// Error is returned for invalid syntax such as unpaired quotes.
func (s *Scanner) Next() (Token, error) {
	const (
		Initial = iota + 1
		Space
		ReverseQuote
		String
		StringDollar
		StringQuoteSingle
		StringQuoteDouble
		StringQuoteDoubleDollar
	)

	var (
		tok Token

		state uint8 = Initial
	)
	s.dollarBuf = s.dollarBuf[:0]
	for {
		r := s.nextRune()
		switch state {
		case Initial:
			switch {
			case r == _RuneEOF:
				tok.Type = TokEOF
				return tok, nil
			case r == '|':
				tok.Type = TokPipe
				return tok, nil
			case r == '`':
				state = ReverseQuote
			case unicode.IsSpace(r):
				state = Space
				s.unreadRune(r)
			default:
				state = String
				s.unreadRune(r)
			}
		case Space:
			if r == _RuneEOF || !unicode.IsSpace(r) {
				s.unreadRune(r)
				tok.Type = TokSpace
				return tok, nil
			}
		case ReverseQuote:
			switch r {
			case _RuneEOF:
				return tok, ErrInvalidSyntax
			case '`':
				tok.Type = TokReversequote
				return tok, nil
			default:
				tok.Value = append(tok.Value, r)
			}
		case String:
			switch {
			case r == _RuneEOF || r == '|' || r == '`' || unicode.IsSpace(r):
				tok.Type = TokString
				s.unreadRune(r)
				return tok, nil
			case r == '\'':
				state = StringQuoteSingle
			case r == '"':
				state = StringQuoteDouble
			case r == '\\':
				nr := s.nextRune()
				if nr == _RuneEOF {
					return tok, ErrInvalidSyntax
				}
				tok.Value = append(tok.Value, nr)
			case r == '$':
				state = s.checkDollarStart(&tok, r, state, StringDollar)
			default:
				tok.Value = append(tok.Value, r)
			}
		case StringDollar:
			state = s.checkDollarEnd(&tok, r, state, String)
		case StringQuoteSingle:
			switch r {
			case _RuneEOF:
				return tok, ErrInvalidSyntax
			case '\'':
				state = String
			case '\\':
				nr := s.nextRune()
				if escape, ok := s.isEscapeChars(nr); ok {
					tok.Value = append(tok.Value, escape)
				} else {
					tok.Value = append(tok.Value, r)
					s.unreadRune(nr)
				}
			default:
				tok.Value = append(tok.Value, r)
			}
		case StringQuoteDouble:
			switch r {
			case _RuneEOF:
				return tok, ErrInvalidSyntax
			case '"':
				state = String
			case '\\':
				nr := s.nextRune()
				if nr == _RuneEOF {
					return tok, ErrInvalidSyntax
				}
				if escape, ok := s.isEscapeChars(nr); ok {
					tok.Value = append(tok.Value, escape)
				} else {
					tok.Value = append(tok.Value, r)
					s.unreadRune(nr)
				}
			case '$':
				state = s.checkDollarStart(&tok, r, state, StringQuoteDoubleDollar)
			default:
				tok.Value = append(tok.Value, r)
			}
		case StringQuoteDoubleDollar:
			state = s.checkDollarEnd(&tok, r, state, StringQuoteDouble)
		}
	}
}

// Scan is a utility function help split input text as tokens.
func Scan(text []rune, env map[string]string) ([]Token, error) {
	s := NewScanner(text, env)
	var tokens []Token
	for {
		tok, err := s.Next()
		if err != nil {
			return nil, err
		}
		tokens = append(tokens, tok)
		if tok.Type == TokEOF {
			break
		}
	}
	return tokens, nil
}
