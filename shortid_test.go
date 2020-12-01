package shortid

import (
	"fmt"
	"testing"
)

// Example example
func Example() {
	fmt.Println(String(12))
	fmt.Println(Int("JyNvD"))
	// Output:
	// JyNvD
	// 12
}

func TestEncoding(t *testing.T) {
	for i := 0; i < 200000; i += 37 {
		b := String(i)
		d := Int(b)
		if i != d {
			t.Fatalf("%s -> %d cannot be decoded", b, d)
		}
	}
}

func BenchmarkString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		String(i)
	}
}

func BenchmarkInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Int("eRASX")
	}
}
