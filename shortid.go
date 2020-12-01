package shortid

import (
	"math"
	"strings"
)

const (
	// Alphabet is fully customizable and may contain any number of
	// characters.  By default, digits and lower-case letters are used, with
	// some removed to avoid confusion between characters like o, O and 0.  The
	// default alphabet is shuffled and has a prime number of characters to further
	// improve the results of the algorithm.
	Alphabet = "JedR8LNFY2j6MrhkBSADUyfP5amuH9xQCX4VqbgpsGtnW7vc3TwKE"
	// BlockSize specifies how many bits will be shuffled.  The lower BLOCK_SIZE
	// bits are reversed.  Any bits higher than BLOCK_SIZE will remain as is.
	// BLOCK_SIZE of 0 will leave all bits unaffected and the algorithm will simply
	// be converting your integer to a different base.
	BlockSize = 24

	maxLen = 5
)

// DefaultEncoding default
var DefaultEncoding = NewEncoding(Alphabet, 0, BlockSize, maxLen)

// Encoding encode
type Encoding struct {
	alphabet  string
	padding   byte
	mask      int
	blockSize int
	maxLen    int
}

// NewEncoding new
func NewEncoding(alphabet string, padding byte, blockSize, maxLen int) *Encoding {
	enc := &Encoding{
		alphabet:  alphabet,
		padding:   padding,
		blockSize: blockSize,
		maxLen:    maxLen,
		mask:      (1 << blockSize) - 1,
	}
	if padding == 0 {
		padding = alphabet[0]
	}
	return enc
}

// String encodes int to string
func (enc *Encoding) String(x int) string {
	buf := make([]byte, enc.maxLen)
	enc.Encode(x, buf)
	return string(buf)
}

// Encode encode
func (enc *Encoding) Encode(x int, s []byte) {
	enc.enbase(enc.encode(x), s)
}

func (enc *Encoding) encode(x int) int {
	return (x & ^enc.mask) | enc._encode(x&enc.mask)
}

func (enc *Encoding) _encode(x int) int {
	r, bs := 0, enc.blockSize-1
	for i := bs; i > -1; i-- {
		if x&(1<<i) != 0 {
			r |= (1 << (bs - i))
		}
	}
	return r
}

func (enc *Encoding) enbase(x int, s []byte) {
	n, m := len(enc.alphabet), len(s)
	i := m - 1
	for x >= n {
		s[i] = enc.alphabet[x%n]
		x /= n
		i--
	}
	s[i] = enc.alphabet[x]
	// add padding
	for i = i - 1; i > -1; i-- {
		s[i] = enc.alphabet[0]
	}
}

// Int decodes string to int
func (enc *Encoding) Int(s string) int {
	return enc.Decode([]byte(s))
}

// Decode decodes
func (enc *Encoding) Decode(b []byte) int {
	return enc.decode(enc.debase(b))
}

func (enc *Encoding) decode(x int) int {
	return (x & ^enc.mask) | enc._decode(x&enc.mask)
}

func (enc *Encoding) _decode(x int) int {
	r, bs := 0, enc.blockSize-1
	for i := bs; i > -1; i-- {
		if x&(1<<(bs-i)) != 0 {
			r |= 1 << i
		}
	}
	return r
}

func (enc *Encoding) debase(b []byte) int {
	n, m, r := len(enc.alphabet), len(b), 0
	for i := 0; i < m; i++ {
		r += strings.IndexByte(enc.alphabet, b[m-i-1]) * int(math.Pow(float64(n), float64(i)))
	}
	return r
}

// String encodes int to string
func String(x int) string {
	return DefaultEncoding.String(x)
}

// Int decodes string to int
func Int(s string) int {
	return DefaultEncoding.Decode([]byte(s))
}
